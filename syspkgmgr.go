package aipsetup

import (
	"archive/tar"
	"bytes"
	"debug/elf"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/set"
	"github.com/ulikunitz/xz"
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

// func (self *SystemPackages) _TestHostArchParameters(host, hostarch, target string) error {
//
// 	if host == "" && arch != "" {
// 		return errors.New("if `host' is empty, `arch' must be empty too")
// 	}
//
// 	if host != "" {
// 		_, err := systemtriplet.NewFromString(host)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	if arch != "" {
// 		_, err := systemtriplet.NewFromString(arch)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }

func (self *SystemPackages) UninstallName(filename string) int {
	return 1
}

func (self *SystemPackages) ListAllInstalledASPs() ([]string, error) {
	pth := self.Sys.GetInstalledASPDir()
	files, err := ioutil.ReadDir(pth)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0)

	for _, i := range files {
		n := path.Base(i.Name())
		if !i.IsDir() && strings.HasSuffix(n, ").xz") {
			nn, err := basictypes.NewASPNameFromString(n)
			if err != nil {
				fmt.Errorf("Can't parse %s inside installed asps dir\n", i.Name())
				continue
			}
			ret = append(ret, nn.String())
		}
	}

	sort.Strings(ret)

	return ret, nil
}

func (self *SystemPackages) ListFilteredInstalledASPs(
	host, hostarch, target string,
) ([]string, error) {

	complete_list, err := self.ListAllInstalledASPs()
	if err != nil {
		return nil, err
	}

	asps := make([]string, 0)

	for _, i := range complete_list {

		parsed_asp_name, err := basictypes.NewASPNameFromString(i)
		if err != nil {
			return nil, errors.New("could not parse string as ASP name: " + i)
		}

		if host != "" && parsed_asp_name.Host != host {
			continue
		}

		if hostarch != "" && parsed_asp_name.Arch != hostarch {
			continue
		}

		if target != "" && parsed_asp_name.Target != target {
			continue
		}

		asps = append(asps, i)
	}

	return asps, nil
}

func (self *SystemPackages) ListInstalledPackageNames(
	host, hostarch, target string,
) ([]string, error) {

	res, err := self.ListFilteredInstalledASPs(host, hostarch, target)
	if err != nil {
		return make([]string, 0),
			errors.New(
				"Error listing installed package names: " + err.Error(),
			)
	}

	names := []string{}

searching_missing_names:
	for _, i := range res {

		parsed_asp_name, err := basictypes.NewASPNameFromString(i)
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
	host, hostarch, target string,
) ([]string, error) {

	ret := []string{}

	res, err := self.ListFilteredInstalledASPs(host, hostarch, target)
	if err != nil {
		return make([]string, 0), errors.New(
			"Error listing installed package names: " + err.Error(),
		)
	}

search:
	for _, i := range res {

		parsed_asp_name, err := basictypes.NewASPNameFromString(i)
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
	aspname *basictypes.ASPName,
) (string, error) {
	return path.Join(
		self.Sys.GetInstalledASPDir(),
		aspname.String(),
	) + ".xz", nil
}

func (self *SystemPackages) IsASPInstalled(
	aspname *basictypes.ASPName,
) (bool, error) {
	fullname, err := self.GenASPFileListPath(aspname)
	if err != nil {
		return false, err
	}
	f, err := os.Open(fullname)
	if err != nil {
		return false, err
	}
	if err == nil {
		f.Close()
	}
	return err == nil, nil
}

// type ListInstalledASPFilesResult struct {
// 	FileList   []string
// 	ParsedName *ASPName
// 	Sys        *System
// }

func (self *SystemPackages) ListInstalledASPFiles(
	aspname *basictypes.ASPName,
) ([]string, error) {

	ret := make([]string, 0)

	fullname, err := self.GenASPFileListPath(aspname)
	if err != nil {
		return []string{}, err
	}

	file, err := os.Open(fullname)
	if err != nil {
		return []string{}, err
	}

	defer file.Close()

	reader, err := xz.NewReader(file)

	if err != nil {
		return []string{}, err
	}

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(reader)

	if err != nil {
		return []string{}, err
	}

reading_lines:
	for {
		line, err := b.ReadString(0xa)

		for {
			if strings.HasSuffix(line, "\n") {
				line = line[:len(line)-1]
			} else {
				break
			}
		}

		if len(line) != 0 {
			ret = append(ret, line)
		}

		if err != nil {
			if err != io.EOF {
				return []string{}, err
			}
			break reading_lines
		}

	}

	return ret, nil
}

func (self *SystemPackages) RemoveASP_DestDir(
	aspname *basictypes.ASPName,
	exclude_files []string,
) error {

	lib_dirs := make([]string, 0)

	for _, i := range []string{"lib", "lib64", "lib32", "libx32"} {
		res, err := filepath.Glob(
			path.Join(self.Sys.Root(), "/", "multihost", "*", i),
		)
		if err != nil {
			return err
		}
		lib_dirs = append(lib_dirs, res...)
	}

	for _, i := range []string{"lib", "lib64", "lib32", "libx32"} {
		res, err := filepath.Glob(
			path.Join(self.Sys.Root(), "/", "multihost", "*", i),
		)
		if err != nil {
			return err
		}
		lib_dirs = append(lib_dirs, res...)
	}

	// TODO:  is this variant with multiarch ok? we only have multihost dir
	//        directly under root. may be /multiarch variant is explicit
	// FIXME: so may be this is error

	// NOTE: this was commented at 2017-08-20
	// for _, i := range []string{"lib", "lib64", "lib32", "libx32"} {
	// 	lib_dirs = append(
	// 		lib_dirs,
	// 		path.Join(self.Sys.Root(), "multiarch", "*", i),
	// 	)
	// }

	// TODO: and why there is no variant with /multihost/*/multiarch/*/lib* ?
	// FIXME: so may be this is error too

	// NOTE: variants with /multihost/*/multiarch/lib* added at 2017-08-20
	for _, i := range []string{"lib", "lib64", "lib32", "libx32"} {
		res, err := filepath.Glob(
			path.Join(
				self.Sys.Root(),
				"/",
				"multihost",
				"*",
				"multiarch",
				"*",
				i,
			),
		)
		if err != nil {
			return err
		}
		lib_dirs = append(lib_dirs, res...)
	}

	dirs := set.NewSetString()
	{
		res, err := self.ListInstalledASPFiles(aspname)
		if err != nil {
			return err
		}
		{
			lst := set.NewSetString()
			for _, i := range res {
				lst.Add(i)
			}
			res = lst.ListStrings()
		}

		{
			sort.Sort(sort.Reverse(sort.StringSlice(res)))
		}

	deleting_files:
		for _, i := range res {

			i_joined := path.Join(self.Sys.Root(), "/", i)
			i_joined_dir := path.Dir(i_joined)

			// exclude shared objects
			// TODO: todo
			for _, j := range lib_dirs {
				j_joined := path.Join(self.Sys.Root(), "/", j)
				if _t, err := filepath.Abs(j); err != nil {
					return err
				} else {
					j_joined = _t
				}

				if i_joined_dir == j_joined {
					elf_obj, err := elf.Open(i_joined)
					if err == nil && elf_obj.Type == elf.ET_DYN {
						elf_obj.Close()
						fmt.Println(" SO!", i_joined)
						continue deleting_files
					}
				}

			}

			for _, j := range exclude_files {
				if j == i {
					fmt.Println(" EX!", i_joined)
					continue deleting_files
				}
			}

			dirs.Add(path.Dir(i_joined))

			fmt.Println(" -  ", i_joined)
			err := os.Remove(i_joined)
			if err != nil {
				return err
			}
		}

		res = dirs.ListStrings()

		for {
			removed := 0

			for _, i := range res {
				d := i
				for {

					if err := os.Remove(d); err == nil {
						fmt.Println(" -  ", d)
						removed++
					}

					d = path.Dir(d)

					if d == "/" || d == self.Sys.Root() {
						break
					}
				}
			}

			if removed == 0 {
				break
			}
		}
	}

	return nil
}

func (self *SystemPackages) RemoveASP_FileLists(
	aspname *basictypes.ASPName,
) error {
	for _, i := range [][3]string{
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
		{
			"./06.LISTS/DESTDIR.lst.xz",
			self.Sys.GetInstalledASPDir(),
			"package's file list",
		},
	} {
		var dst_dir string

		dst_dir = i[1]

		var dst_filename string

		if i[0] == "./05.BUILD_LOGS.tar.xz" {
			dst_filename = fmt.Sprintf("%s.tar.xz", aspname.String())
		} else {
			dst_filename = fmt.Sprintf("%s.xz", aspname.String())
		}

		dst_filename = path.Join(dst_dir, dst_filename)

		err := os.Remove(dst_filename)
		if err != nil {
			return err
		}

	}
	return nil
}

func (self *SystemPackages) RemoveASP(
	aspname *basictypes.ASPName,
	unregister_only bool,
	exclude_files []string,
) error {

	var err error = nil

	if !unregister_only {
		err = self.RemoveASP_DestDir(aspname, exclude_files)
		if err != nil {
			return err
		}
	}

	err = self.RemoveASP_FileLists(aspname)
	if err != nil {
		return err
	}

	return nil
}

func (self *SystemPackages) ReduceASP(
	reduce_to *basictypes.ASPName,
	reduce_what []*basictypes.ASPName,
	// host, hostarch, target string, // NOTE: abowe parameters already have this info
) error {

	reduce_what_copy := make([]*basictypes.ASPName, 0)
	reduce_what_copy = append(reduce_what_copy, reduce_what...)

	if yes, err := self.IsASPInstalled(reduce_to); err != nil {
		return err
	} else if !yes {
		return errors.New("asp not installed")
	}

	for _, ii := range reduce_what_copy {
		if yes, err := self.IsASPInstalled(ii); err != nil {
			return err
		} else if !yes {
			return errors.New("asp not installed")
		}
	}

	for i := range reduce_what_copy {
		reduce_what_i := reduce_what_copy[i]
		// TODO: something strange. check this
		if reduce_what_i.String() == reduce_to.String() {
			reduce_what_copy =
				append(reduce_what_copy[0:i], reduce_what_copy[:i+1]...)
		}
	}

	fiba, err := self.ListInstalledASPFiles(reduce_to)
	if err != nil {
		return err
	}

	errors_while_reducing_asps := make([]*basictypes.ASPName, 0)
	for _, i := range reduce_what_copy {
		err := self.RemoveASP(i, false, fiba)
		if err != nil {
			// NOTE: error should be reported, but process should continue.
			//       in the end function should exit with error
			errors_while_reducing_asps = append(errors_while_reducing_asps, i)
		}
	}

	if len(errors_while_reducing_asps) > 0 {
		return errors.New("error removing packages")
	}

	return nil
}

func (self *SystemPackages) InstallASP_FileLists(
	filename string,
	parsed *basictypes.ASPName,
) error {

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
				dst_dir = i[1]
				os.MkdirAll(dst_dir, 0755)

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

				if _, err := io.Copy(dst_file_obj, tar_obj); err != nil {
					return err
				}

				dst_file_obj.Close()

				os.Chtimes(dst_filename, head.AccessTime, head.ModTime)
			}

		}
	}
	return nil
}

func (self *SystemPackages) InstallASP_DestDir(filename string) error {

	tar_file_obj, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer tar_file_obj.Close()

	tar_obj := tar.NewReader(tar_file_obj)

	directories := set.NewSetString()
	hardlinks := make(map[string]string)

	var head *tar.Header

	for {
		var err error

		head, err = tar_obj.Next()
		if err != nil {
			break
		}

		if head.Name == "./04.DESTDIR.tar.xz" {
			xz_reader, err := xz.NewReader(tar_obj)
			if err != nil {
				return err
			}

			xztar_obj := tar.NewReader(xz_reader)

			var xztar_head *tar.Header

		continue_xztar_obj:
			for {
				var err error

				xztar_head, err = xztar_obj.Next()
				if err != nil {
					if err == io.EOF {
						break
					} else {
						return err
					}
				}

				for _, i := range []byte{tar.TypeDir} {
					if xztar_head.Typeflag == i {
						continue continue_xztar_obj
					}
				}

				{
					if !strings.HasPrefix(xztar_head.Name, "./") {
						fmt.Println("   not allowed Name")
						return errors.New(
							"tar file provided forbidden name elements",
						)
					}

					test_abs, err := filepath.Abs(xztar_head.Name[1:])
					if err != nil {
						return err
					}

					if test_abs != xztar_head.Name[1:] {
						return errors.New(
							"tar file provided forbidden name elements",
						)
					}
				}

				new_file_path := path.Join(
					self.Sys.Root(),
					"/",
					xztar_head.Name[1:],
				)
				new_file_path, err = filepath.Abs(new_file_path)
				if err != nil {
					return err
				}

				new_file_dir := path.Dir(new_file_path)

				switch xztar_head.Typeflag {
				default:
					return errors.New(
						fmt.Sprintf("file type not supported: %v",
							xztar_head.Typeflag,
						),
					)
				case tar.TypeReg:
					fallthrough
				case tar.TypeRegA:
					fmt.Println(" +", new_file_path)

					err := os.MkdirAll(new_file_dir, 0755)
					if err != nil {
						return err
					}

					directories.Add(new_file_dir)
					new_file, err := os.Create(new_file_path)
					if err != nil {
						return err
					}

					_, err = io.Copy(new_file, xztar_obj)
					if err != nil {
						new_file.Close()
						return err
					}
					new_file.Close()

					/* NOTE: this should be uncommented when this functionality is
								 ready to work with root
					err = os.Chown(new_file_path, 0, 0)
					if err != nil {
						return err
					}
					*/

					err = os.Chmod(new_file_path, 0755)
					if err != nil {
						return err
					}

					err = os.Chtimes(
						new_file_path,
						xztar_head.AccessTime,
						xztar_head.ModTime,
					)
					if err != nil {
						return err
					}

				case tar.TypeLink:
					err := os.MkdirAll(new_file_dir, 0755)
					if err != nil {
						return err
					}

					directories.Add(new_file_dir)

					ln_value := xztar_head.Linkname
					if !strings.HasPrefix("/", ln_value) {
						ln_value := path.Join(path.Dir(new_file_path), ln_value)
						abs, err := filepath.Abs(ln_value)
						if err != nil {
							return err
						}
						ln_value = abs
					}

					hardlinks[new_file_path] = ln_value
				case tar.TypeSymlink:

					err := os.MkdirAll(new_file_dir, 0755)
					if err != nil {
						return err
					}

					directories.Add(new_file_dir)

					fmt.Printf(
						" + %s\n  s -> %s\n",
						new_file_path,
						xztar_head.Linkname,
					)

					_, err = os.Lstat(new_file_path)

					if err != nil {
						//if !strings.HasSuffix(err.Error(), "no such file or directory") {
						if !os.IsNotExist(err) {
							return err
						}
					}

					if err == nil {
						err = os.Remove(new_file_path)
						if err != nil {
							return err
						}
					}

					err = os.Symlink(xztar_head.Linkname, new_file_path)
					if err != nil {
						return err
					}
				}

			}

			for _, i := range directories.ListStrings() {
				/*
					err := os.Chown(i, 0, 0)
					if err != nil {
						return err
					}
				*/
				err := os.Chmod(i, 0755)
				if err != nil {
					return err
				}
			}

			for key, val := range hardlinks {
				fmt.Printf(
					" + %s\n  h -> %s\n",
					val,
					key,
				)
				err = os.Link(val, key)
				if err != nil {
					return err
				}
			}

			break
		}

	}
	return nil
}

func (self *SystemPackages) InstallASP(filename string) error {

	parsed, err := basictypes.NewASPNameFromString(filename)
	if err != nil {
		return err
	}

	fmt.Println("parse result\n", parsed.StringD())

	host := parsed.Host
	arch := parsed.Arch

	if host == "" || arch == "" {
		return errors.New("Invalid value for host or arch")
	}

	err = self.InstallASP_FileLists(filename, parsed)
	if err != nil {
		return err
	}

	err = self.InstallASP_DestDir(filename)
	if err != nil {
		return err
	}

	return nil
}
