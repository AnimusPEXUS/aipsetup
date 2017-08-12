package aipsetup

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ulikunitz/xz"

	augfilepath "github.com/AnimusPEXUS/filepath"
)

type SystemPackages struct {
	Sys *System
}

func NewSystemPackages(system *System) *SystemPackages {

	var (
		ret *SystemPackages
	)

	ret = new(SystemPackages)
	ret.Sys = system

	return ret
}

func (self *SystemPackages) UninstallName(filename string) int {
	return 1
}

func (self *SystemPackages) ListInstalledASPs(
	host, arch string,
) ([]string, error) {

	if arch != "" && host == "" {
		panic("if host is empty, arch must be empty also")
	}

	asps := make([]string, 0)

	dir, err := os.Open(self.Sys.GetInstalledASPDir())
	if dir != nil && err == nil {
		files_in_dir, err := dir.Readdir(-1)
		if files_in_dir != nil && err == nil {

			for _, i := range files_in_dir {
				if !i.IsDir() && strings.HasSuffix(i.Name(), ").xz") {

					parsed_asp_name, err := NewASPNameParsedFromString(i.Name())
					if err != nil {
						return make([]string, 0),
							errors.New("could not parse string as ASP name: " + i.Name())
					}

					if (host == "") ||
						((host != "" && host == parsed_asp_name.Host) &&
							((arch == "") || (arch != "" && arch == parsed_asp_name.Arch))) {
						asps = append(asps, i.Name())
					}

				}
			}
		}
	}

	return asps, nil
}

func (self *SystemPackages) ListInstalledPackageNames(
	host, arch string,
) ([]string, error) {

	res, err := self.ListInstalledASPs(host, arch)
	if err != nil {
		return make([]string, 0),
			errors.New(
				"Error listing installed package names: " + err.Error(),
			)
	}

	names := []string{}

searching_missing_names:
	for _, i := range res {

		parsed_asp_name, err := NewASPNameParsedFromString(i)
		if err != nil {
			return make([]string, 0),
				errors.New(
					"Can't parse package name: " + i,
				)
		}

		for _, j := range names {
			if parsed_asp_name.Name == j {
				continue searching_missing_names
			}
		}
		names = append(names, parsed_asp_name.Name)
	}

	return names, nil
}

func (self *SystemPackages) ListInstalledPackageNameASPs(
	name string,
	host, arch string,
) ([]string, error) {

	ret := []string{}

	res, err := self.ListInstalledASPs(host, arch)
	if err != nil {
		return make([]string, 0), errors.New(
			"Error listing installed package names: " + err.Error(),
		)
	}

search:
	for _, i := range res {

		parsed_asp_name, err := NewASPNameParsedFromString(i)
		if err != nil {
			return make([]string, 0),
				errors.New(
					"Can't parse package name: " + i,
				)
		}

		if parsed_asp_name.Name != name {
			continue search
		}

		for _, j := range ret {
			if i == j {
				continue search
			}
		}
		ret = append(ret, parsed_asp_name.String())
	}

	return ret, nil
}

func (self *SystemPackages) GenASPFileListPath(
	aspname string,
) string {
	aspname = NormalizeASPName(aspname)
	return augfilepath.Join(self.Sys.GetInstalledASPDir(), aspname) + ".xz"
}

func (self *SystemPackages) IsASPInstalled(
	aspname string,
) bool {
	fullname := self.GenASPFileListPath(aspname)
	f, err := os.Open(fullname)
	if err == nil {
		f.Close()
	}
	return err == nil
}

type ListInstalledASPFilesResult struct {
	FileList   []string
	ParsedName *ASPNameParsed
	Sys        *System
}

func (self *SystemPackages) ListInstalledASPFiles(
	aspname string,
) (*ListInstalledASPFilesResult, error) {

	var (
		ret     *ListInstalledASPFilesResult
		ret_lst []string
		ret_err error
	)

	fullname := self.GenASPFileListPath(aspname)

	file, err := os.Open(fullname)
	if err != nil {
		ret_err = err
	} else {
		defer file.Close()

		reader, err := xz.NewReader(file)

		if err != nil {
			ret_err = err
		} else {

			b := new(bytes.Buffer)
			_, err := b.ReadFrom(reader)

			if err != nil {
				ret_err = err
			} else {

			reading_lines:
				for true {
					line, err := b.ReadString(0xa)

					for true {
						if strings.HasSuffix(line, "\n") {
							line = line[:len(line)-1]
						} else {
							break
						}
					}

					if len(line) != 0 {
						ret_lst = append(ret_lst, line)
					}

					if err != nil {
						if err != io.EOF {
							ret_err = err
						}
						break reading_lines
					}

				}
			}
		}
	}
	if ret_err == nil {
		parsed_name, err := NewASPNameParsedFromString(aspname)
		if err != nil {
			ret_err = errors.New("couldn't parse asp name")
			ret = nil
		} else {
			sys := self.Sys
			ret = &ListInstalledASPFilesResult{
				ret_lst,
				parsed_name,
				sys,
			}
		}
	}
	return ret, ret_err
}

func (self *SystemPackages) RemoveASP(
	aspname string,
	keepnewest bool,
	exclude_shared_object_files bool,
) error {

	var (
		files_which_need_to_be_removed *ListInstalledASPFilesResult
		files_which_need_to_be_keeped  *ListInstalledASPFilesResult
	)

	{

		var err error

		aspname = NormalizeASPName(aspname)

		if !self.IsASPInstalled(aspname) {
			return errors.New("such ASP not presen in the system")
		}

		parsed_name_aspname, err := NewASPNameParsedFromString(aspname)
		if err != nil {
			return errors.New("Can't parse ASP name: " + aspname)
		}

		all_asps, err := self.ListInstalledPackageNameASPs(
			parsed_name_aspname.Name,
			parsed_name_aspname.Host,
			parsed_name_aspname.Arch,
		)
		if err != nil {
			return err
		}

		sort.Sort(ASPNameSorter(all_asps))

		if len(all_asps) == 0 {
			panic("this shouldn't been happen. programming error")
		}

		files_which_need_to_be_removed, err = self.ListInstalledASPFiles(aspname)
		if files_which_need_to_be_removed == nil || err != nil {
			return errors.New(
				"self.ListInstalledASPFiles(apsname) returned nil. " +
					err.Error(),
			)
		}

		if keepnewest {
			newest := all_asps[len(all_asps)-1]

			if keepnewest && newest == aspname {
				return errors.New(
					"last ASP of name can't be removed. use name removal cmd",
				)
			}

			files_which_need_to_be_keeped, err = self.ListInstalledASPFiles(newest)

			if files_which_need_to_be_keeped == nil || err != nil {
				return errors.New(
					"self.ListInstalledASPFiles(newest) returned nil. " +
						err.Error(),
				)
			}
		}

	}

	return self.RemoveASPFiles(
		files_which_need_to_be_removed,
		files_which_need_to_be_keeped,
		exclude_shared_object_files,
	)
}

func (self *SystemPackages) RemoveASPFiles(
	files_which_need_to_be_removed *ListInstalledASPFilesResult,
	files_which_need_to_be_keeped *ListInstalledASPFilesResult,
	exclude_shared_object_files bool,
) error {

	var (
		ret error
	)

	if files_which_need_to_be_removed == nil {
		panic("files_which_need_to_be_removed == nil")
	}

	if files_which_need_to_be_removed.Sys == nil {
		panic("files_which_need_to_be_removed.Sys == nil")
	}

	if files_which_need_to_be_keeped != nil &&
		files_which_need_to_be_removed.Sys != files_which_need_to_be_keeped.Sys {
		panic(
			"files_which_need_to_be_removed.Sys !=" +
				" files_which_need_to_be_keeped.Sys",
		)
	}

removing_files:
	for _, i := range files_which_need_to_be_removed.FileList {

		rooted_i := augfilepath.Join(self.Sys.Root(), i)

		if files_which_need_to_be_keeped != nil {

			for _, j := range files_which_need_to_be_keeped.FileList {

				if i == j {
					// debug
					//fmt.Println("skipping newer file", j)
					continue removing_files
				}

			}

		}

		abs_i, err := filepath.Abs(i)
		if err != nil {
			panic("can't evaluate absolut path for filename")
		}
		dir_abs_i := filepath.Dir(abs_i)

		if exclude_shared_object_files && IsALibDirPath(dir_abs_i) {

			mutched, err := filepath.Match("*.so*", filepath.Base(i))
			if err != nil {
				panic("looks like unpredicted error")
			}

			if mutched {
				// debug
				//fmt.Println("skipping .so file", rooted_i)

				continue removing_files
			}
		}

		fmt.Println("TODO: remove file", rooted_i)

	}

	return ret
}

func (self *SystemPackages) InstallASP(filename string) error {

	parsed, err := NewASPNameParsedFromString(filename)
	if err != nil {
		return err
	}

	host := parsed.Host
	arch := parsed.Arch

	if host == "" || arch == "" {
		return errors.New("Invalid value for host or arch")
	}

	tar_file_obj, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer tar_file_obj.Close()

	tar_obj := tar.NewReader(tar_file_obj)

	package_name := parsed.String()

	var head *tar.Header

	for {
		var err error

		head, err = tar_obj.Next()
		if err != nil {
			break
		}

		for _, i := range [][3]string{
			{
				"./06.LISTS/DESTDIR.lst.xz",
				self.Sys.GetInstalledASPDir(),
				"package's file list",
			},
			{
				"./06.LISTS/DESTDIR.sha512.xz",
				self.Sys.GetInstalledASPSumsDir(),
				"package's check sums",
			},
			{
				"./05.BUILD_LOGS.tar.xz",
				self.Sys.GetInstalledASPBuildLogsDir(),
				"package's buildlogs",
			},
			{
				"./06.LISTS/DESTDIR.dep_c.xz",
				self.Sys.GetInstalledASPDepsDir(),
				"package's dependencies listing",
			},
		} {
			var dst_dir string

			if head.Name == i[0] {
				dst_dir = path.Dir(i[1])
				os.MkdirAll(dst_dir, 0755)
			}

			var dst_filename string

			if i[0] == "./05.BUILD_LOGS.tar.xz" {
				dst_filename = fmt.Sprintf("%s.tar.xz", package_name)
			} else {
				dst_filename = fmt.Sprintf("%s.xz", package_name)
			}

			dst_filename = path.Join(dst_dir, dst_filename)

			dst_file_obj, err := os.Create(dst_filename)
			if err != nil {
				return err
			}
			io.Copy(dst_file_obj, tar_obj)
		}
	}

	return nil
}
