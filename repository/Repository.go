package repository

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"time"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/repository/providers"
	"github.com/AnimusPEXUS/aipsetup/repository/types"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
	"github.com/AnimusPEXUS/utils/version/versioncomparators"
)

var _ types.RepositoryI = &Repository{}

type Repository struct {
	sys basictypes.SystemI
	log *logger.Logger
}

func NewRepository(sys basictypes.SystemI, log *logger.Logger) (*Repository, error) {
	ret := new(Repository)
	ret.sys = sys
	ret.log = log
	return ret, nil
}

func (self *Repository) GetRepositoryPath() string {
	return self.sys.GetTarballRepoRootDir()
}

func (self *Repository) GetCachesDir() string {
	return path.Join(self.GetRepositoryPath(), "cache")
}

func (self *Repository) GetPackagePath(name string) string {
	return path.Join(self.GetRepositoryPath(), "packages", name)
}

func (self *Repository) GetPackageSRSPath(name string) string {
	return path.Join(self.GetPackagePath(name), "srs")
}

func (self *Repository) GetPackageTarballsPath(name string) string {
	return path.Join(self.GetPackagePath(name), "tarballs")
}

func (self *Repository) GetPackagePatchesPath(name string) string {
	return path.Join(self.GetPackagePath(name), "patches")
}

func (self *Repository) GetPackageASPsPath(name string) string {
	return path.Join(self.GetPackagePath(name), "asps")
}

func (self *Repository) GetPackageCachePath(name string) string {
	return path.Join(self.GetCachesDir(), "individual", name)
}

func (self *Repository) GetDedicatedCachePath(name string) string {
	return path.Join(self.GetCachesDir(), "dedicated", name)
}

func (self *Repository) GetTarballDoneFilePath(
	package_name string,
	as_filename string,
) string {
	return self.GetTarballFilePath(package_name, as_filename) + ".done"
}

func (self *Repository) GetTarballFilePath(package_name, as_filename string) string {
	as_filename = path.Base(as_filename)
	tarballs_dir := self.GetPackageTarballsPath(package_name)
	return path.Join(tarballs_dir, as_filename)
}

func (self *Repository) PerformPackageSourcesUpdate(name string) error {
	for _, i := range []func(string) error{
		self.PerformPackageTarballsUpdate,
		self.PerformPackagePatchesUpdate,
	} {
		err := i(name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Repository) PerformPackageTarballsUpdate(name string) error {

	info, err := pkginfodb.Get(name)
	if err != nil {
		return err
	}

	log := logger.New()
	log.AddOutput(os.Stdout)

	prov, err := providers.Get(
		info.TarballProvider,
		self,
		name,
		info,
		self.sys,
		self.GetPackageTarballsPath(name),
		log,
	)
	if err != nil {
		return err
	}

	err = prov.PerformUpdate()
	if err != nil {
		return err
	}

	return nil
}

func (self *Repository) PerformPackagePatchesUpdate(name string) error {
	info, err := pkginfodb.Get(name)
	if err != nil {
		return err
	}

	if !info.DownloadPatches {
		return nil
	}

	patches_path := self.GetPackagePatchesPath(name)

	err = os.MkdirAll(patches_path, 0700)
	if err != nil {
		return err
	}

	u, err := user.Current()
	if err != nil {
		return err
	}

	file_tmp_dir_path := path.Join(u.HomeDir, ".config", "aipsetup5", "tmp")
	// file_tmp_dir_path := "/tmp"

	err = os.MkdirAll(file_tmp_dir_path, 0700)
	if err != nil {
		return err
	}

	h := md5.New()
	h.Write([]byte(time.Now().UTC().String()))

	f, err := ioutil.TempFile(
		file_tmp_dir_path,
		hex.EncodeToString(h.Sum([]byte{})),
	)
	if err != nil {
		return err
	}

	file_tmp_dir_file_path := f.Name()

	defer func(f *os.File, pth string) {
		f.Close()
		os.Remove(pth)
	}(f, file_tmp_dir_file_path)

	err = os.Chmod(file_tmp_dir_file_path, 0700)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(info.PatchesDownloadingScriptText))
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	cmd := exec.Command(file_tmp_dir_file_path)
	cmd.Dir = patches_path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Repository) DetermineNewestStableTarball(package_name string) (string, error) {

	name_info, err := pkginfodb.Get(package_name)
	if err != nil {
		return "", err
	}

	tarballs, err := self.ListLocalTarballs(package_name, true)
	if err != nil {
		return "", err
	}

	if len(tarballs) == 0 {
		return "", errors.New("repository does not have tarballs for this package")
	}

	p, err := tarballnameparsers.Get(name_info.TarballFileNameParser)
	if err != nil {
		return "", err
	}

	c, err := versioncomparators.Get(name_info.TarballVersionComparator)
	if err != nil {
		return "", err
	}

	version_tool, err := tarballstabilityclassification.Get(name_info.TarballStabilityClassifier)
	if err != nil {
		return "", err
	}

	err = c.SortStrings(tarballs, p)
	if err != nil {
		return "", err
	}

	{
		tarballs2 := make([]string, 0)
		for _, i := range tarballs {

			parsed, err := p.Parse(i)
			if err != nil {
				return "", err
			}

			isstable, err := version_tool.IsStable(parsed)
			if err != nil {
				return "", err
			}
			if isstable {
				tarballs2 = append(tarballs2, i)
			}
		}
		tarballs = tarballs2
	}

	ret := tarballs[len(tarballs)-1]
	//fmt.Println(t)

	return ret, nil
}

func (self *Repository) ListLocalTarballs(package_name string, done_only bool) ([]string, error) {
	ret := make([]string, 0)

	res, err := self.ListLocalFiles(package_name)
	if err != nil {
		return ret, err
	}

	info, err := pkginfodb.Get(package_name)
	if err != nil {
		return ret, err
	}

	parser, err := tarballnameparsers.Get(info.TarballFileNameParser)
	if err != nil {
		return ret, err
	}

	for _, i := range res {
		err := tarballname.IsPossibleTarballNameErr(i)
		if err != nil {
			continue
		}

		if parse_res, err := parser.Parse(i); err != nil {
			continue
		} else {
			if parse_res.Name != info.TarballName {
				continue
			}
		}
		full_out_path_done := self.GetTarballDoneFilePath(package_name, i)
		if done_only {
			_, err = os.Stat(full_out_path_done)
			if err != nil {
				continue
			}
		}

		ret = append(ret, i)
	}

	return ret, nil
}

func (self *Repository) ListLocalFiles(package_name string) ([]string, error) {
	ret := make([]string, 0)

	pth := self.GetPackageTarballsPath(package_name)

	files, err := ioutil.ReadDir(pth)
	if err != nil {
		return ret, err
	}

	for _, i := range files {
		if !i.IsDir() {
			ret = append(ret, i.Name())
		}
	}

	return ret, nil
}

func (self *Repository) PerformDownload(
	package_name string,
	as_filename string,
	uri string,
) error {
	as_filename = path.Base(as_filename)

	full_out_path := self.GetTarballFilePath(package_name, as_filename)
	full_out_path_done := self.GetTarballDoneFilePath(package_name, as_filename)

	_, err := os.Stat(full_out_path_done)
	if err == nil {
		_, err := os.Stat(full_out_path)
		if err == nil {
			return nil
		}
	}

	err = os.MkdirAll(path.Dir(full_out_path), 0700)
	if err != nil {
		return err
	}
	c := exec.Command("wget", "--max-redirect=100", "--progress=dot", "-c", "-O", full_out_path, uri)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	ret := c.Run()
	if ret == nil {
		f, err := os.Create(full_out_path_done)
		if err != nil {
			return err
		}
		f.Close()
	}
	return ret
}

func (self *Repository) PrepareTarballCleanupListing(
	package_name string,
	files_to_keep []string,
) ([]string, error) {
	lst, err := self.ListLocalTarballs(package_name, false)
	// lst, err := self.ListLocalFiles(package_name, false)
	if err != nil {
		return []string{}, err
	}

	to_delete := make([]string, 0)

	for _, i := range lst {
		found := false
		for _, j := range files_to_keep {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			to_delete = append(to_delete, i)

			done_file_path := self.GetTarballDoneFilePath(package_name, i)
			to_delete = append(to_delete, path.Base(done_file_path))
		}
	}

	return to_delete, nil
}

func (self *Repository) DeleteFile(
	package_name string,
	filename string,
) error {
	var ret error
	tarballs_dir := self.GetPackageTarballsPath(package_name)
	filename = path.Base(filename)
	full_path := path.Join(tarballs_dir, filename)
	if _, err := os.Stat(full_path); err == nil {
		ret = os.Remove(full_path)
	}
	return ret
}

func (self *Repository) DeleteFiles(package_name string, filename []string) error {
	for _, i := range filename {
		if err := self.DeleteFile(package_name, i); err != nil {
			return err
		}
	}
	return nil
}

func (self *Repository) MoveInTarball(filename string, copy bool) error {

	pkgname, _, err := pkginfodb.DetermineTarballPackageInfoSingle(filename)
	if err != nil {
		return err
	}

	tarballs_dir := self.GetPackageTarballsPath(pkgname)
	full_out_path := path.Join(tarballs_dir, path.Base(filename))

	err = os.MkdirAll(tarballs_dir, 0700)
	if err != nil {
		return err
	}

	if copy {
		err = filetools.CopyWithInfo(filename, full_out_path, nil)
		if err != nil {
			return err
		}
	} else {
		err = os.Rename(filename, full_out_path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Repository) MoveInASP(filename string, copy bool) error {
	aspname, err := basictypes.NewASPNameFromString(filename)
	if err != nil {
		return err
	}

	name := aspname.Name

	pth := self.GetPackageASPsPath(name)

	err = os.MkdirAll(pth, 0700)
	if err != nil {
		return err
	}

	filename_j := path.Join(pth, path.Base(filename))

	if copy {
		err = filetools.CopyWithInfo(filename, filename_j, nil)
		if err != nil {
			return err
		}
	} else {
		err = os.Rename(filename, filename_j)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Repository) CopyTarballToDir(
	package_name string,
	tarball string,
	outdir string,
) error {
	src_file_path := self.GetTarballFilePath(package_name, tarball)
	out_file := path.Join(outdir, tarball)

	if _, err := os.Stat(src_file_path); err != nil {
		return err
	}

	err := os.MkdirAll(outdir, 0700)
	if err != nil {
		return err
	}

	out_file_o, err := os.Create(out_file)
	if err != nil {
		return err
	}

	src_file_o, err := os.Open(src_file_path)
	if err != nil {
		return err
	}

	_, err = io.Copy(out_file_o, src_file_o)
	if err != nil {
		return err
	}

	return nil
}

func (self *Repository) CopyPatchesToDir(
	package_name string,
	outdir string,
) error {

	err := filetools.CopyTree(
		self.GetPackagePatchesPath(package_name),
		outdir,
		false,
		true,
		true,
		true,
		self.log,
		filetools.CopyWithInfo,
	)

	if err != nil {
		return err
	}

	return nil
}
