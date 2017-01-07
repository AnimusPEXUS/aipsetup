package aipsetup

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
	
	augfilepath "github.com/AnimusPEXUS/filepath"
)

type SystemPackages struct {
	system *System
}

func NewSystemPackages(system *System) *SystemPackages {

	var (
		ret *SystemPackages
	)

	ret = new(SystemPackages)
	ret.system = system

	return ret
}

func (self *SystemPackages) InstallASP(filename string) int {
	return 1
}

func (self *SystemPackages) UninstallASP(filename string) int {
	return 1
}

func (self *SystemPackages) ListInstalledASPs(host, arch string) []string {

	if arch != "" && host == "" {
		panic("if host is empty, arch must be empty also")
	}

	asps := make([]string, 0)

	dir, err := os.Open(self.system.GetInstalledASPDir())
	if dir != nil && err == nil {
		files_in_dir, err := dir.Readdir(-1)
		if files_in_dir != nil && err == nil {

			for _, i := range files_in_dir {
				if !i.IsDir() && strings.HasSuffix(i.Name(), ").xz") {

					parsed_asp_name := NewASPNameParsedFromString(i.Name())
					if (host == "") ||
						((host != "" && host == parsed_asp_name.Host) &&
							((arch == "") || (arch != "" && arch == parsed_asp_name.Arch))) {
						asps = append(asps, i.Name())
					}

				}
			}
		}
	}

	return asps
}

func (self *SystemPackages) ListInstalledPackageNames(
	host, arch string,
) []string {

	res := self.ListInstalledASPs(host, arch)

	names := []string{}

searching_missing_names:
	for _, i := range res {

		parsed_asp_name := NewASPNameParsedFromString(i)

		for _, j := range names {
			if parsed_asp_name.Name == j {
				continue searching_missing_names
			}
		}
		names = append(names, parsed_asp_name.Name)
	}

	return names
}

func (self *SystemPackages) ListInstalledPackageNameASPs(
	name string,
	host, arch string,
) []string {

	ret := []string{}

	res := self.ListInstalledASPs(host, arch)

search:
	for _, i := range res {

		parsed_asp_name := NewASPNameParsedFromString(i)

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

	return ret
}

func (self *SystemPackages) GenASPFileListPath(
	aspname string,
) string {
	aspname = NormalizeASPName(aspname)
	return path.Join(self.system.GetInstalledASPDir(), aspname) + ".xz"
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
		parsed_name := NewASPNameParsedFromString(aspname)
		sys := self.system
		if parsed_name != nil {
			ret = &ListInstalledASPFilesResult{
				ret_lst,
				parsed_name,
				sys,
			}
		} else {
			ret_err = errors.New("couldn't parse asp name")
			ret = nil
		}
	}
	return ret, ret_err
}

func (self *SystemPackages) RemoveASP(
	aspname string,
	keepnewest bool,
	exclude_so_files_removal_from_lib_dirs bool,
) error {

	aspname = NormalizeASPName(aspname)

	if !self.IsASPInstalled(aspname) {
		return errors.New("such ASP not presen in the system")
	}

	parsed_name_aspname := NewASPNameParsedFromString(aspname)

	all_asps = self.ListInstalledPackageNameASPs(
		parsed_name_aspname.Name,
		parsed_name_aspname.Host,
		parsed_name_aspname.Arch,
	)

	switch len(all_asps) {
	case 0:
		panic("this shouldn't been happen. programming error")
	case 1:
		if keepnewest {
			return errors.New(
				"last ASP of name can't be removed. use name removal cmd",
			)
		}
	}

	sort.Sort(ASPNameSorter(all_asps))

	//exclusions := ([]*ListInstalledASPFilesResult){}

	if keepnewest {
		newest := all_asps[len(all_asps)-1]

		all_asps = all_asps[:len(all_asps)-1]

	}

	if aspname == newest && keepnewest {
		return errors.New(
			"can not remove already latest asp, " +
				"while keeping lates at the same time",
		)
	}

	// newest_files :=

	return nil
}

func (self *SystemPackages) RemoveASPFiles(
	files_which_need_to_be_removed *ListInstalledASPFilesResult,
	files_which_need_to_be_keeped *ListInstalledASPFilesResult,
	exclude_so_files_removal_from_lib_dirs bool,
) error {

	ret := error(nil)

	{
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
	}

	if ret == nil {
	removing_files:

		for _, i := range files_which_need_to_be_removed.FileList {

			if files_which_need_to_be_keeped != nil {

				for _, j := range files_which_need_to_be_keeped.FileList {

					if i == j {
						continue removing_files
					}

				}

			}

			abs_i := filepath.Abs(i)
			dir_abs_i := filepath.Dir(abs_i)
			base_dir_abs_i := filepath.Base(dir_abs_i)

			if exclude_so_files_removal_from_lib_dirs {

				if IsPathAHostRoot(dir_abs_i) || IsPathAnArchDir(dir_abs_i) {

					if strings.HasPrefix(base_dir_abs_i, "lib") {

					}

				}

			}

			fmt.Println("TODO: remove file", filepath.Join(self.Sys.Root(), i))

		}

	}

	return ret
}
