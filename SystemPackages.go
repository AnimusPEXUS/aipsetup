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
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/set"
	"github.com/ulikunitz/xz"
)

type SystemPackages struct {
	sys *System
}

func NewSystemPackages(system *System) *SystemPackages {
	self := new(SystemPackages)
	self.sys = system
	return self
}

func (self *SystemPackages) ListAllInstalledASPs() ([]*basictypes.ASPName, error) {

	files, err := ioutil.ReadDir(self.sys.GetInstalledASPDir())
	if err != nil {
		return nil, err
	}

	ret := make([]*basictypes.ASPName, 0)

	for _, i := range files {
		n := path.Base(i.Name())
		if !i.IsDir() && strings.HasSuffix(n, ").xz") {
			nn, err := basictypes.NewASPNameFromString(n)
			if err != nil {
				return nil, errors.New(
					fmt.Sprintf("Can't parse %s inside installed asps dir\n", i.Name()),
				)
			}
			ret = append(ret, nn)
		}
	}

	return ret, nil
}

func (self *SystemPackages) ListFilteredInstalledASPs(
	host, hostarch string,
) ([]*basictypes.ASPName, error) {

	complete_list, err := self.ListAllInstalledASPs()
	if err != nil {
		return nil, err
	}

	asps := make([]*basictypes.ASPName, 0)

	for _, i := range complete_list {

		if host != "" && i.Host != host {
			continue
		}

		if hostarch != "" && i.HostArch != hostarch {
			continue
		}

		asps = append(asps, i)
	}

	return asps, nil
}

func (self *SystemPackages) ListInstalledPackageNames(
	host, hostarch string,
) ([]string, error) {

	res, err := self.ListFilteredInstalledASPs(host, hostarch)
	if err != nil {
		return nil,
			errors.New(
				"Error listing installed package names: " + err.Error(),
			)
	}

	names := make([]string, 0)

searching_missing_names:
	for _, parsed_asp_name := range res {

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
	host, hostarch string,
) ([]*basictypes.ASPName, error) {

	ret := make([]*basictypes.ASPName, 0)

	res, err := self.ListFilteredInstalledASPs(host, hostarch)
	if err != nil {
		return nil, errors.New(
			"Error listing installed package names: " + err.Error(),
		)
	}

search:
	for _, i := range res {

		if i.Name != name {
			continue
		}

		for _, j := range ret {
			if i.IsEqual(j) {
				continue search
			}
		}
		ret = append(ret, i)
	}

	return ret, nil
}

func (self *SystemPackages) GenASPFileListPath(
	aspname *basictypes.ASPName,
) (string, error) {
	reg_info, err := self.FindSystemPackageRegistrationByName(aspname)
	if err != nil {
		return "", err
	}

	if !reg_info.Found() {
		return "",
			errors.New(
				"path to filelist of " +
					aspname.String() +
					" isn't found",
			)
	}

	// I don't like this.
	// One package must have exectly one file list (and exactly one registration)
	// But this construction is for backword comparability with old aipsetup's
	// asp naming format + it's assumed I watched carefuly for this single-single
	// rule
	ret := path.Join(
		self.sys.GetInstalledASPDir(),
		reg_info.Pkg[0],
	)

	return ret, nil
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
	keep_shared_objects bool,
) error {

	lib_dirs := make([]string, 0)

	for _, i := range []string{
		basictypes.DIRNAME_LIB,
		basictypes.DIRNAME_LIB64,
		basictypes.DIRNAME_LIB32,
		basictypes.DIRNAME_LIBX32,
	} {
		res, err := filepath.Glob(
			path.Join(self.sys.Root(), "/", "multihost", "*", i),
		)
		if err != nil {
			return err
		}
		lib_dirs = append(lib_dirs, res...)
	}

	for _, i := range []string{
		basictypes.DIRNAME_LIB,
		basictypes.DIRNAME_LIB64,
		basictypes.DIRNAME_LIB32,
		basictypes.DIRNAME_LIBX32,
	} {
		res, err := filepath.Glob(
			path.Join(self.sys.Root(), "/", "multihost", "*", i),
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
	// for _, i := range []string{DIRNAME_LIB, DIRNAME_LIB64, DIRNAME_LIB32, DIRNAME_LIBX32} {
	// 	lib_dirs = append(
	// 		lib_dirs,
	// 		path.Join(self.sys.Root(), "multiarch", "*", i),
	// 	)
	// }

	// OLDTODO: and why there is no variant with /multihost/*/multiarch/*/lib* ?
	// FIXME: so may be this is error too

	// NOTE: variants with /multihost/*/multiarch/*/lib* added at 2017-08-20
	for _, i := range []string{
		basictypes.DIRNAME_LIB,
		basictypes.DIRNAME_LIB64,
		basictypes.DIRNAME_LIB32,
		basictypes.DIRNAME_LIBX32,
	} {
		res, err := filepath.Glob(
			path.Join(
				self.sys.Root(),
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

		sort.Sort(sort.Reverse(sort.StringSlice(res)))

	deleting_files:
		for _, i := range res {

			i_joined := path.Join(self.sys.Root(), "/", i)
			i_joined_dir := path.Dir(i_joined)

			if keep_shared_objects {
				// exclude shared objects
				for _, j := range lib_dirs {

					j_joined := path.Join(self.sys.Root(), "/", j)

					if _t, err := filepath.Abs(j); err != nil {
						return err
					} else {
						j_joined = _t
					}

					if i_joined_dir == j_joined {
						elf_obj, err := elf.Open(i_joined)
						if err == nil && elf_obj.Type == elf.ET_DYN {
							elf_obj.Close()
							self.sys.log.Info(" Shared OBJ Exclusion: " + i_joined)
							continue deleting_files
						}
					}

				}
			}

			for _, j := range exclude_files {
				if j == i {
					// self.sys.log.Info(" Reduction Exclusion: " + i_joined)
					continue deleting_files
				}
			}

			dirs.Add(path.Dir(i_joined))

			_, err := os.Lstat(i_joined)
			if err == nil {
				self.sys.log.Info(" Removing: " + i_joined)
				err := os.Remove(i_joined)
				if err != nil {
					if !os.IsNotExist(err) {
						// It's ok if file already missing
						return err
					}
				}
			} else {
				if !os.IsNotExist(err) {
					return err
				}
			}

		}

		res = dirs.ListStrings()

		for {
			removed := 0

			for _, i := range res {
				d := i
				for {

					if err := os.Remove(d); err == nil {
						self.sys.log.Info(" Removed directory: " + d)
						removed++
					}

					d = path.Dir(d)

					if d == "/" || d == self.sys.Root() {
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
	res, err := self.FindSystemPackageRegistrationByName(aspname)
	if err != nil {
		return err
	}

	self.sys.log.Info("Going to unregister package " + aspname.String())
	self.sys.log.Info("  going to remove files:")
	self.sys.log.Info("   build log:")
	self.sys.log.Info(fmt.Sprintf("    %v", res.Buildlogs))
	self.sys.log.Info("   sums:")
	self.sys.log.Info(fmt.Sprintf("    %v", res.Sums))
	self.sys.log.Info("   dep list:")
	self.sys.log.Info(fmt.Sprintf("    %v", res.Deps))
	self.sys.log.Info("   file list:")
	self.sys.log.Info(fmt.Sprintf("    %v", res.Pkg))

	err = res.DeleteAll()
	if err != nil {
		return err
	}

	return nil
}

func (self *SystemPackages) RemoveASP(
	aspname *basictypes.ASPName,
	unregister_only bool,
	called_by_reduce bool,
	reduce_exclude_files []string,
	force_removal_if_reduce_is_possible bool,
) error {

	pkginfo, err := pkginfodb.Get(aspname.Name)
	if err != nil {
		return err
	}

	if !called_by_reduce && !pkginfo.Removable {
		return errors.New("this package is not removable")
	}

	if pkginfo.Reducible && !called_by_reduce && !force_removal_if_reduce_is_possible {
		return errors.New(
			"package is reducible. so can be removed only by reducing doe " +
				"to installing new one",
		)
	}

	if !unregister_only {
		err = self.RemoveASP_DestDir(
			aspname,
			reduce_exclude_files,
			!pkginfo.DontPreserveSharedObjects,
		)
		if err != nil {
			return err
		}
	}

	// unregister must be final step
	err = self.RemoveASP_FileLists(aspname)
	if err != nil {
		return err
	}

	return nil
}

func (self *SystemPackages) ReduceASP(
	reduce_to *basictypes.ASPName,
	reduce_what []*basictypes.ASPName,
) error {

	reduce_what_copy := make([]*basictypes.ASPName, 0)
	reduce_what_copy = append(reduce_what_copy, reduce_what...)

	for _, i := range reduce_what {
		if i.Name != reduce_to.Name {
			// this is programming error, so here is a panic
			panic("reduce_to.Name is different to names in reduce_what")
		}
	}

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

	for i := len(reduce_what_copy) - 1; i != -1; i-- {
		if reduce_what_copy[i].IsEqual(reduce_to) {
			reduce_what_copy =
				append(
					reduce_what_copy[:i],
					reduce_what_copy[:i+1]...,
				)
		}
	}

	fiba, err := self.ListInstalledASPFiles(reduce_to)
	if err != nil {
		return err
	}

	errors_while_reducing_asps := make([]*basictypes.ASPName, 0)
	for _, i := range reduce_what_copy {
		self.sys.log.Info("Reducing asp: " + i.String())
		err := self.RemoveASP(i, false, true, fiba, false)
		if err != nil {
			self.sys.log.Error(err.Error())
			// error should be reported, but process should continue.
			// in the end function should exit with error
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
				self.sys.GetInstalledASPDir(),
				"package's file list",
			},
			{
				"./06.LISTS/DESTDIR.sha512.xz",
				self.sys.GetInstalledASPSumsDir(),
				"package's check sums",
			},
			{
				"./05.BUILD_LOGS.tar.xz",
				self.sys.GetInstalledASPBuildLogsDir(),
				"package's buildlogs",
			},
			{
				"./06.LISTS/DESTDIR.dep_c.xz",
				self.sys.GetInstalledASPDepsDir(),
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
	hardlinks := make([][2]string, 0)

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
						self.sys.log.Error("   not allowed Name")
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
					self.sys.Root(),
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
					self.sys.log.Info(" Writing: " + new_file_path)

					err := os.MkdirAll(new_file_dir, 0755)
					if err != nil {
						return err
					}

					directories.Add(new_file_dir)

					os.Remove(new_file_path)

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

					err = os.Chown(new_file_path, 0, 0)
					if err != nil {
						return err
					}

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
					// fmt.Println("xztar_head.Linkname", xztar_head.Linkname)
					// fmt.Println("new_file_path", new_file_path)
					// if !strings.HasPrefix("/", ln_value) {
					// 	ln_value = path.Join(path.Dir(new_file_path), ln_value)
					// 	abs, err := filepath.Abs(ln_value)
					// 	if err != nil {
					// 		return err
					// 	}
					// 	ln_value = abs
					// }
					// TODO: this was fixed in a horry. need rechecks
					if !strings.HasPrefix("/", ln_value) {
						ln_value = path.Join(self.sys.Root(), ln_value)
					}

					hardlinks = append(hardlinks, [2]string{new_file_path, ln_value})

				case tar.TypeSymlink:

					err := os.MkdirAll(new_file_dir, 0755)
					if err != nil {
						return err
					}

					directories.Add(new_file_dir)

					self.sys.log.Info(" Symlinking")

					self.sys.log.Info(fmt.Sprintf("  %s", new_file_path))
					self.sys.log.Info("  ->")
					self.sys.log.Info(fmt.Sprintf("  %s", xztar_head.Linkname))

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
				err = os.Chmod(i, 0755)
				if err != nil {
					return err
				}
			}

			for _, i := range hardlinks {
				key := i[0]
				val := i[1]
				self.sys.log.Info(" Hardlinking")

				self.sys.log.Info(fmt.Sprintf("  %s", val))
				self.sys.log.Info("  ->")
				self.sys.log.Info(fmt.Sprintf("  %s", key))

				os.Remove(key)

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

	if ok, err := CheckAspPackageByFilename(filename); err != nil {
		return err
	} else {
		if !ok {
			return errors.New("given file didn't passed package check")
		}
	}

	host := parsed.Host
	hostarch := parsed.HostArch
	// target := parsed.Target

	if host == "" || hostarch == "" {
		return errors.New("Invalid value for host or arch")
	}

	pkginfo, err := pkginfodb.Get(parsed.Name)
	if err != nil {
		return err
	}

	if pkginfo.Deprecated {
		return errors.New("package " + parsed.Name + " is deprecated")
	}

	if pkginfo.NonInstallable {
		return errors.New("package " + parsed.Name + " is non installable")
	}

	if pkginfo.PrimaryInstallOnly {
		if host != hostarch {
			return errors.New("package " + parsed.Name + " is only for primary install")
		}
	}

	err = self.InstallASP_FileLists(filename, parsed)
	if err != nil {
		return err
	}

	err = self.InstallASP_DestDir(filename)
	if err != nil {
		return err
	}

	if pkginfo.Reducible {

		lst, err := self.ListInstalledPackageNameASPs(parsed.Name, host, hostarch)
		if err != nil {
			return err
		}

		for i := len(lst) - 1; i != -1; i-- {
			if parsed.IsEqual(lst[i]) {
				lst = append(lst[:i], lst[i+1:]...)
			}
		}

		if len(lst) != 0 {

			self.sys.log.Info("The list of other installations of this package")
			for ii, i := range lst {
				self.sys.log.Info(fmt.Sprintf(" %3d: %s", ii, i.String()))
			}

			if pkginfo.AutoReduce {
				self.sys.log.Info(
					"package is reducable and autoreduce is enabled.. reducing..",
				)

				err = self.InstallASPReduceToSubRoutine(filename, false)
				if err != nil {
					return err
				}

			} else {
				self.sys.log.Warning(
					"This package is reducable, but autoreduce is disabled. " +
						"You'll have to command reduction of older installations separately.",
				)
			}
		}
	}

	self.sys.log.Info("Installation Finished. Looks Ok")

	return nil
}

func (self *SystemPackages) InstallASPReduceToSubRoutine(
	filename string,
	print_list bool,
) error {

	self.sys.log.Info("Reduction to " + filename + " asked")

	parsed, err := basictypes.NewASPNameFromString(filename)
	if err != nil {
		return err
	}

	pkginfo, err := pkginfodb.Get(parsed.Name)
	if err != nil {
		return err
	}

	lst, err := self.ListInstalledPackageNameASPs(
		parsed.Name,
		parsed.Host,
		parsed.HostArch,
	)
	if err != nil {
		return err
	}

	for i := len(lst) - 1; i != -1; i-- {
		if parsed.IsEqual(lst[i]) {
			lst = append(lst[:i], lst[i+1:]...)
		}
	}

	if print_list {
		if len(lst) != 0 {
			self.sys.log.Info("The list of other installations of this package")
			for ii, i := range lst {
				self.sys.log.Info(fmt.Sprintf(" %3d: %s", ii, i.String()))
			}
		} else {
			self.sys.log.Info("Not found other installations of this package")
		}
	}

	if !pkginfo.Reducible {
		return errors.New("this package is not reducible")
	}

	if len(lst) != 0 {
		err = self.ReduceASP(parsed, lst)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *SystemPackages) FindSystemPackageRegistrationByName(
	aspname *basictypes.ASPName,
) (*SystemPackageRegistration, error) {
	return FindSystemPackageRegistrationByName(aspname, self.sys)
}

// TODO: mode should be enum
func (self *SystemPackages) FindPackagesByInstalledFilenames(
	mode string,
	pattern string,
	progress_cb func(perc float64, ind, count int),
) (map[string]([]string), error) {

	asps, err := self.ListAllInstalledASPs()
	if err != nil {
		return nil, err
	}

	ret := make(map[string]([]string))

	m := sync.Mutex{}
	m1 := sync.Mutex{}
	g := sync.WaitGroup{}

	register_found_file := func(pkg string, file string) {
		m.Lock()
		defer m.Unlock()
		_, ok := ret[pkg]
		if !ok {
			ret[pkg] = make([]string, 0)
		}
		ret[pkg] = append(ret[pkg], file)
	}

	asps_l := len(asps)
	thread_errors := make([]error, 0)
	guard := make(chan struct{}, runtime.NumCPU()*2+2)
	for ii, i := range asps {

		g.Add(1)

		guard <- struct{}{}
		go func() {

			defer g.Done()
			defer func() { <-guard }()

			if progress_cb != nil {
				m1.Lock()
				progress_cb(
					100.0/(float64(asps_l)/float64(ii)),
					ii,
					asps_l,
				)
				m1.Unlock()
			}

			i_asp_str := i.String()
			files, err := self.ListInstalledASPFiles(i)
			if err != nil {
				thread_errors = append(thread_errors, err)
				return
			}
			for _, j := range files {
				switch mode {
				default:
					panic("programming error")
				case "sub":
					if strings.Contains(j, pattern) {
						register_found_file(i_asp_str, j)
					}
				case "re":
					ok, err := regexp.MatchString(pattern, j)
					if err != nil {
						thread_errors = append(thread_errors, err)
						return
					}

					if ok {
						register_found_file(i_asp_str, j)
					}

				case "fm":
					ok, err := filepath.Match(pattern, j)
					if err != nil {
						thread_errors = append(thread_errors, err)
						return
					}

					if ok {
						register_found_file(i_asp_str, j)
					}
				}
			}
		}()
	}

	g.Wait()

	if len(thread_errors) != 0 {
		return nil, thread_errors[0]
	}

	return ret, nil
}
