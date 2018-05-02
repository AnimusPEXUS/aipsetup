package aipsetup

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/checksums"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

type Packager struct {
	site *BuildingSiteCtl
}

func NewPackager(site *BuildingSiteCtl) *Packager {
	ret := new(Packager)
	ret.site = site
	return ret
}

func (self *Packager) Run(log *logger.Logger) error {
	for _, i := range [](func(log *logger.Logger) error){
		self.DestDirCheckCorrectness,
		self.DestDirFileList,
		self.DestDirChecksum,
		self.CompressPatchesAndDestDir,
		self.CompressFilesInListsDir,
		self.UpdateTimestamp,
		self.MakeChecksums,
		self.CompressLogs,
		self.Pack,
	} {
		// log.Info(fmt.Sprintf("Starting \"%v\" pack target", i))
		err := i(log)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self Packager) DestDirCheckCorrectness(log *logger.Logger) error {

	log.Info("Checking paths correctness")

	var allowed_in_root = []string{"multihost", "etc", "var", "boot"}
	// others not allowed in root

	var allowed_in_host = []string{
		basictypes.DIRNAME_BIN,
		basictypes.DIRNAME_SBIN,
		"opt",
		basictypes.DIRNAME_LIB,
		basictypes.DIRNAME_LIB64,
		basictypes.DIRNAME_SHARE,
		basictypes.DIRNAME_INCLUDE,
		"libexec",
		"multiarch",
		"src",

		"x86_64-pc-linux-gnu", // TODO: compiler dir's should be checcked smarter
		"i686-pc-linux-gnu",
	}

	var allowed_in_host_arch = []string{
		basictypes.DIRNAME_BIN,
		basictypes.DIRNAME_SBIN,
		"opt",
		basictypes.DIRNAME_LIB,
		basictypes.DIRNAME_LIB64,
		basictypes.DIRNAME_SHARE,
		basictypes.DIRNAME_INCLUDE,
		"libexec",
		"src",
		"x86_64-pc-linux-gnu", // TODO: compiler dir's should be checcked smarter
		"i686-pc-linux-gnu",
	}

	calc := self.site.GetBuildingSiteValuesCalculator()

	info, err := self.site.ReadInfo()
	if err != nil {
		return err
	}

	dest_dir := self.site.GetDIR_DESTDIR()

	{
		files, err := ioutil.ReadDir(dest_dir)
		if err != nil {
			return err
		}

	loop:
		for _, i := range files {

			for _, j := range allowed_in_root {
				if i.Name() == j {
					continue loop
				}
			}
			log.Error(fmt.Sprintf("file %s not allowed in root directory", i.Name()))
			return errors.New("found not allowed files in destdir root")
		}
	}

	dest_dir_host, err := calc.CalculateDstHostDir()
	if err != nil {
		return err
	}

	{
		files, err := ioutil.ReadDir(dest_dir_host)
		if err != nil {
			return err
		}

	loop2:
		for _, i := range files {

			for _, j := range allowed_in_host {
				if i.Name() == j {
					continue loop2
				}
			}

			log.Error(fmt.Sprintf("file %s not allowed in host directory", i.Name()))
			return errors.New("found not allowed files in host dir")
		}
	}

	if info.HostArch != info.Host {

		dest_dir_host_arch, err := calc.CalculateDstHostArchDir()
		if err != nil {
			return err
		}

		{
			files, err := ioutil.ReadDir(dest_dir_host_arch)
			if err != nil {
				return err
			}

		loop3:
			for _, i := range files {

				for _, j := range allowed_in_host_arch {
					if i.Name() == j {
						continue loop3
					}
				}

				log.Error(fmt.Sprintf("file %s not allowed in host's arch directory", i.Name()))
				return errors.New("found not allowed files in host's arch dir")
			}
		}

	}

	log.Info("   no problems found")

	return nil
}

func (self Packager) DestDirFileList(log *logger.Logger) error {

	log.Info("Creating file list")

	ddir := self.site.GetDIR_DESTDIR()
	ldir := self.site.GetDIR_LISTS()

	outfile, err := os.Create(path.Join(ldir, "DESTDIR.lst"))
	if err != nil {
		return err
	}
	defer outfile.Close()

	err = filetools.Walk(
		ddir,
		func(
			dir string,
			dirs []os.FileInfo,
			files []os.FileInfo,
		) error {
			drill, err := filepath.Rel(ddir, dir)
			if err != nil {
				return err
			}

			for _, i := range files {
				drillj := path.Join("/", drill, i.Name())

				outfile.WriteString(fmt.Sprintln(drillj))

			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (self Packager) DestDirChecksum(log *logger.Logger) error {

	log.Info("Calculating DESTDIR files' checksums")

	ddir := self.site.GetDIR_DESTDIR()
	ldir := self.site.GetDIR_LISTS()

	outfile, err := os.Create(path.Join(ldir, "DESTDIR.sha512"))
	if err != nil {
		return err
	}
	defer outfile.Close()

	err = filetools.Walk(
		ddir,
		func(
			dir string,
			dirs []os.FileInfo,
			files []os.FileInfo,
		) error {
			drill, err := filepath.Rel(ddir, dir)
			if err != nil {
				return err
			}

			for _, i := range files {

				if i.Mode()&os.ModeSymlink != 0 {
					continue
				}

				drilljn := path.Join(dir, i.Name())
				drillj := path.Join("/", drill, i.Name())

				f, err := os.Open(drilljn)
				if err != nil {
					return err
				}
				defer f.Close()

				h := sha512.New()

				buff := make([]byte, 2*(1024^2))
				for {
					s, err := f.Read(buff)
					if err != nil {
						if err == io.EOF {
							break
						} else {
							return err
						}
					}
					h.Write(buff[:s])
				}
				f.Close()

				outfile.WriteString(
					fmt.Sprintf(
						"%s *%s\n",
						hex.EncodeToString(h.Sum(nil)),
						drillj,
					),
				)

			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (self Packager) CompressPatchesAndDestDir(log *logger.Logger) error {
	log.Info(
		fmt.Sprintf(
			"Compressing  %s and %s",
			basictypes.DIR_PATCHES,
			basictypes.DIR_DESTDIR,
		),
	)
	return self._CompressPatchesDestDirAndLogs(
		log,
		[]string{basictypes.DIR_PATCHES, basictypes.DIR_DESTDIR},
	)
}
func (self Packager) CompressLogs(log *logger.Logger) error {
	log.Info(
		fmt.Sprintf(
			"Compressing %s",
			basictypes.DIR_BUILD_LOGS,
		),
	)
	return self._CompressPatchesDestDirAndLogs(
		log,
		[]string{basictypes.DIR_BUILD_LOGS},
	)
}

func (self Packager) _CompressPatchesDestDirAndLogs(log *logger.Logger, subject []string) error {
	for _, i := range subject {
		log.Info(fmt.Sprintf("  %s", i))
		dirname := path.Join(self.site.path, i)
		filename := fmt.Sprintf("%s.tar.xz", dirname)

		filename_o, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer filename_o.Close()

		xz_c := exec.Command("xz", "-v9", "-c", "-")
		xz_c.Stderr = os.Stderr
		xz_c.Stdout = filename_o

		tar_c := exec.Command(
			"tar",
			"-cv",
			"--sort=name",
			"--owner=root",
			"--group=root",
			"--mode=755",
			".",
		)
		tar_c.Dir = dirname

		if i != basictypes.DIR_BUILD_LOGS {
			tar_c.Stderr = log.StdoutLbl()
		} else {
			fmt.Println("Logs dir should be compressed without logging. sorry.")
		}

		if t, err := tar_c.StdoutPipe(); err == nil {
			xz_c.Stdin = t
		} else {
			return err
		}

		tar_c.Start()
		xz_c.Start()

		tar_err := tar_c.Wait()
		xz_err := xz_c.Wait()

		if tar_err != nil {
			return tar_err
		}

		if xz_err != nil {
			return xz_err
		}

		filename_o.Close()

		log.Info(fmt.Sprintf("    %s Done", i))

	}
	return nil
}

func (self Packager) CompressFilesInListsDir(log *logger.Logger) error {
	log.Info("Compressing files in lists dir")

	ldir := self.site.GetDIR_LISTS()

	for _, i := range []string{"DESTDIR.lst", "DESTDIR.sha512"} {

		log.Info(fmt.Sprintf("  %s", i))

		infile := path.Join(ldir, i)
		outfile := fmt.Sprintf("%s.xz", infile)

		f, err := os.Open(infile)
		if err != nil {
			return err
		}
		defer f.Close()

		f2, err := os.Create(outfile)
		if err != nil {
			return err
		}
		defer f2.Close()

		c := exec.Command("xz", "-9c", "-")
		c.Stdin = f
		c.Stdout = f2

		err = c.Run()
		if err != nil {
			return err
		}

		f.Close()
		f2.Close()

		log.Info(fmt.Sprintf("    %s Done", i))

	}
	return nil
}

func (self Packager) UpdateTimestamp(log *logger.Logger) error {
	info, err := self.site.ReadInfo()
	if err != nil {
		return err
	}

	t := time.Now().UTC()

	t_s := fmt.Sprintf(
		"%04d%02d%02d.%02d%02d%02d.%07d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1000,
	)

	info.PackageTimeStamp = t_s

	err = self.site.WriteInfo(info)
	if err != nil {
		return err
	}

	return nil
}

func (self Packager) _ListItemsToPack(include_checksum, include_build_logs bool) ([]string, error) {
	ret := make([]string, 0)
	pth := self.site.path
	ret = append(ret, path.Join(pth, basictypes.DIR_DESTDIR+".tar.xz"))
	ret = append(ret, path.Join(pth, basictypes.DIR_PATCHES+".tar.xz"))

	if include_build_logs {
		ret = append(ret, path.Join(pth, basictypes.DIR_BUILD_LOGS+".tar.xz"))
	}

	ret = append(ret, path.Join(pth, basictypes.PACKAGE_INFO_FILENAME_V5))
	if include_checksum {
		ret = append(ret, path.Join(pth, basictypes.PACKAGE_CHECKSUM_FILENAME))
	}

	{
		tarballs, err := ioutil.ReadDir(self.site.GetDIR_TARBALL())
		if err != nil {
			return ret, err
		}
		for _, i := range tarballs {
			if !i.IsDir() {
				ret = append(ret, path.Join(pth, basictypes.DIR_TARBALL, i.Name()))
			}
		}
	}
	{
		lists, err := ioutil.ReadDir(self.site.GetDIR_LISTS())
		if err != nil {
			return ret, err
		}
		for _, i := range lists {
			if !i.IsDir() && strings.HasSuffix(i.Name(), ".xz") {
				ret = append(ret, path.Join(pth, basictypes.DIR_LISTS, i.Name()))
			}
		}
	}
	return ret, nil
}

func (self Packager) MakeChecksums(log *logger.Logger) error {

	log.Info("Creating checksumms before packaging")

	pth := self.site.path

	result_file := path.Join(self.site.path, "package.sha512")

	list_to_checksum, err := self._ListItemsToPack(false, false)
	if err != nil {
		return err
	}

	result_file_o, err := os.Create(result_file)
	if err != nil {
		return err
	}
	defer result_file_o.Close()

	for _, i := range list_to_checksum {
		// i_stat, err := os.Stat(i)
		// if err != nil {
		// 	return err
		// }
		j_name, err := filepath.Rel(pth, i)
		if err != nil {
			return err
		}

		sum, err := checksums.CalculateFileChecksum(
			i,
			sha512.New(),
			2*(1024^2),
		)
		if err != nil {
			return err
		}

		result_file_o.WriteString(
			fmt.Sprintf(
				"%s *%s\n",
				hex.EncodeToString(sum),
				j_name,
			),
		)

	}

	return nil
}

func (self Packager) Pack(log *logger.Logger) error {

	log.Info("Creating package")

	info, err := self.site.ReadInfo()
	if err != nil {
		return err
	}

	pack_dir, err := self.site.GetOuterAspsDir()
	if err != nil {
		return err
	}

	ts_str, err := basictypes.NewASPTimeStampFromString(info.PackageTimeStamp)
	if err != nil {
		return err
	}

	pack_file_name := fmt.Sprintf(
		"%s.asp",
		(&basictypes.ASPName{
			Name:      info.PackageName,
			Version:   info.PackageVersion,
			Status:    info.PackageStatus,
			TimeStamp: ts_str,
			Host:      info.Host,
			HostArch:  info.HostArch,
			// Target:    info.Target,
		}).String(),
	)

	j_pack_file_name := path.Join(pack_dir, pack_file_name)

	log.Info(fmt.Sprintf("Resulting ASP filename will be: %s", pack_file_name))

	err = os.MkdirAll(pack_dir, 0700)
	if err != nil {
		return err
	}

	list_to_pack, err := self._ListItemsToPack(true, true)
	if err != nil {
		return err
	}

	{
		list_to_pack2 := make([]string, 0)
		var r string
		for _, i := range list_to_pack {
			r, err = filepath.Rel(self.site.path, i)
			if err != nil {
				return err
			}
			list_to_pack2 = append(list_to_pack2, "./"+r)
		}

		list_to_pack = list_to_pack2
	}

	sort.Strings(list_to_pack)

	args := make([]string, 0)
	args = append(args, []string{"-vcf", j_pack_file_name}...)
	args = append(args, list_to_pack...)

	c := exec.Command("tar", args...)
	c.Dir = self.site.path
	c.Stderr = os.Stdout
	return c.Run()

}
