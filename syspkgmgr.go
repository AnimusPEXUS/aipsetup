package aipsetup

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
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

func (self *SystemPackages) ListInstalledASPs(host, arch string) []string {

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
		parsed_name := NewASPNameParsedFromString(aspname)
		sys := self.Sys
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

		parsed_name_aspname := NewASPNameParsedFromString(aspname)

		all_asps := self.ListInstalledPackageNameASPs(
			parsed_name_aspname.Name,
			parsed_name_aspname.Host,
			parsed_name_aspname.Arch,
		)

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
			panic("can't evaluate filepath.Abs(i)")
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

func (self *SystemPackages) InstallASP(filename string) int {
	return 1
}
