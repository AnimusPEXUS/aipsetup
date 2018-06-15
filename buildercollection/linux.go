package buildercollection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/systemtriplet"
)

func init() {
	Index["linux"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_linux(bs)
	}
}

type Builder_linux struct {
	*Builder_std

	crossbuild_params []string

	src_arch_dir string
	dst_boot_dir string
	dst_man_dir  string
}

func NewBuilder_linux(bs basictypes.BuildingSiteCtlI) (*Builder_linux, error) {

	self := new(Builder_linux)

	self.Builder_std = NewBuilder_std(bs)

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	self.crossbuild_params = make([]string, 0)

	if info.ThisIsCrossbuilding() { // Crossbuilder?
		host := info.Host
		hostarch := info.HostArch

		tri, err := systemtriplet.NewFromString(host)
		if err != nil {
			return nil, err
		}

		if tri.Company != "pc" || tri.Kernel != "linux" || tri.OS != "gnu" {
			return nil, errors.New("invalid value host selected")
		}

		headers_arch := ""

		switch tri.CPU {
		default:
			return nil, errors.New("unsupported cpu value selected")
		case "i486":
			fallthrough
		case "i586":
			fallthrough
		case "i686":
			headers_arch = "x86"
		case "x86_64":
			headers_arch = "x86_64"
		}

		// linux_headers_arch = None
		// if re.match(r'^i[4-6]86$', cpu) or re.match(r'^x86(_32)?$', cpu):
		// 		linux_headers_arch = 'x86'
		// elif re.match(r'^x86_64$', cpu):
		// 		linux_headers_arch = 'x86_64'
		// else:
		// 		logging.error("Don't know which linux ARCH to apply")
		// 		ret = 3

		self.crossbuild_params = append(
			self.crossbuild_params,
			[]string{
				fmt.Sprintf("ARCH=%s", headers_arch),
				fmt.Sprintf("CROSS_COMPILE=%s-", hostarch), // TODO: not sure arch is the valid value here
			}...,
		)
	}

	self.src_arch_dir = path.Join(bs.GetDIR_SOURCE(), "arch")
	self.dst_boot_dir = path.Join(bs.GetDIR_DESTDIR(), "boot")
	self.dst_man_dir = path.Join(
		bs.GetDIR_DESTDIR(),
		"usr", basictypes.DIRNAME_SHARE, "man", "man9",
	)

	return self, nil
}

func (self *Builder_linux) DefineActions() (basictypes.BuilderActions, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret := basictypes.BuilderActions{
		&basictypes.BuilderAction{"dst_cleanup", self.Builder_std.BuilderActionDstCleanup},
		&basictypes.BuilderAction{"src_cleanup", self.Builder_std.BuilderActionSrcCleanup},
		&basictypes.BuilderAction{"bld_cleanup", self.Builder_std.BuilderActionBldCleanup},
		&basictypes.BuilderAction{"extract", self.Builder_std.BuilderActionExtract},
		//&basictypes.BuilderAction{"patch", self.BuilderActionPatch},
		// &basictypes.BuilderAction{"autogen", self.BuilderActionAutogen},
		&basictypes.BuilderAction{"configure", self.BuilderActionConfigure},
		&basictypes.BuilderAction{"build", self.BuilderActionBuild},

		&basictypes.BuilderAction{"distr_kernel", self.BuilderActionDistrKernel},
		&basictypes.BuilderAction{"distr_modules", self.BuilderActionDistrModules},
		// &basictypes.BuilderAction{"distr_firmware", self.BuilderActionDistrFirmware },// NOTE: removed from linux 4.14

		&basictypes.BuilderAction{"distr_headers_all", self.BuilderActionDistrHeadersAll},

		// &basictypes.BuilderAction{"distr_man", self.BuilderActionDistrMan},
		&basictypes.BuilderAction{"distr_source", self.BuilderActionDistrSource},
		&basictypes.BuilderAction{"distr_symlink", self.BuilderActionDistrSymLink},
	}

	if info.ThisIsCrossbuilder() || info.ThisIsSubarchBuilding() { // TODO: simplify
		if info.ThisIsCrossbuilder() {
			self.bs.GetLog().Info("Crossbuilder building detected")
		}

		if info.ThisIsSubarchBuilding() {
			self.bs.GetLog().Info("Subarch building detected")
		}

		self.bs.GetLog().Info(" - only headers will be prepared")

		ret = basictypes.BuilderActions{
			&basictypes.BuilderAction{"dst_cleanup", self.Builder_std.BuilderActionDstCleanup},
			&basictypes.BuilderAction{"src_cleanup", self.Builder_std.BuilderActionSrcCleanup},
			&basictypes.BuilderAction{"extract", self.Builder_std.BuilderActionExtract},

			&basictypes.BuilderAction{"distr_headers_all", self.BuilderActionDistrHeadersAll},
		}

	}

	ret = append(
		ret,
		[]*basictypes.BuilderAction{
			&basictypes.BuilderAction{"prepack", self.Builder_std.BuilderActionPrePack},
			&basictypes.BuilderAction{"pack", self.Builder_std.BuilderActionPack},
		}...,
	)

	return ret, nil
}

func (self *Builder_linux) BuilderActionConfigure(
	log *logger.Logger,
) error {
	log.Info("\n" +
		"Now You have to configure kernel by your needs and\n" +
		"continue building procedure with command:\n" +
		"aipsetup5 build run build+\n")
	return errors.New("user action required")
}

func (self *Builder_linux) BuilderActionBuild(
	log *logger.Logger,
) error {
	at := &buildingtools.Autotools{}

	err := at.Make(
		self.crossbuild_params,
		[]string{},
		buildingtools.Copy,
		"Makefile",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"make",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_linux) BuilderActionDistrKernel(
	log *logger.Logger,
) error {

	err := os.MkdirAll(self.dst_boot_dir, 0700)
	if err != nil {
		return err
	}

	at := &buildingtools.Autotools{}

	args := make([]string, 0)
	args = append(
		args,
		[]string{
			"install",
			fmt.Sprintf("INSTALL_PATH=%s", self.dst_boot_dir),
		}...,
	)
	args = append(
		args,
		self.crossbuild_params...,
	)

	err = at.Make(
		args,
		[]string{},
		buildingtools.Copy,
		"Makefile",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"make",
		log,
	)
	if err != nil {
		return err
	}

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	p1 := path.Join(self.dst_boot_dir, "vmlinuz")
	p2 := path.Join(
		self.dst_boot_dir,
		fmt.Sprintf(
			"vmlinuz-%s-%s",
			info.PackageVersion,
			info.Host,
		),
	)

	log.Info(fmt.Sprintf("Renaming: `%s' to `%s'", p1, p2))

	err = os.Rename(p1, p2)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_linux) BuilderActionDistrModules(
	log *logger.Logger,
) error {

	at := &buildingtools.Autotools{}

	args := make([]string, 0)
	args = append(
		args,
		[]string{
			"modules_install",
			fmt.Sprintf("INSTALL_MOD_PATH=%s", self.bs.GetDIR_DESTDIR()),
		}...,
	)
	args = append(
		args,
		self.crossbuild_params...,
	)

	err := at.Make(
		args,
		[]string{},
		buildingtools.Copy,
		"Makefile",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"make",
		log,
	)
	if err != nil {
		return err
	}

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	modules_dir := path.Join(self.bs.GetDIR_DESTDIR(), basictypes.DIRNAME_LIB, "modules")

	files, err := ioutil.ReadDir(modules_dir)
	if err != nil {
		return err
	}

	if len(files) != 1 {
		log.Error(
			fmt.Sprintf("Can't find  single directory in %s", modules_dir),
		)
		return errors.New("error finding modules directory")
	}

	modules_dir = path.Join(modules_dir, files[0].Name())

	linux_dirname := fmt.Sprintf("linux-%s", info.PackageVersion)

	for _, i := range []string{"build", "source"} {
		new_link := path.Join(modules_dir, i)

		os.Remove(new_link)

		err = os.Symlink(
			path.Join("/", "usr", "src", linux_dirname),
			new_link,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_linux) BuilderActionDistrHeadersAll(
	log *logger.Logger,
) error {

	var install_hdr_path string

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if info.ThisIsCrossbuilder() {
		install_hdr_path = path.Join(
			self.bs.GetDIR_DESTDIR(), "usr", "crossbuilders",
			info.CrossbuilderTarget,
		)

	} else {
		install_hdr_path = path.Join(self.bs.GetDIR_DESTDIR(), "usr")
	}

	at := &buildingtools.Autotools{}

	args := make([]string, 0)
	args = append(
		args,
		[]string{
			"headers_install_all",
			fmt.Sprintf("INSTALL_HDR_PATH=%s", install_hdr_path),
		}...,
	)
	args = append(
		args,
		self.crossbuild_params...,
	)

	err = at.Make(
		args,
		[]string{},
		buildingtools.Copy,
		"Makefile",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"make",
		log,
	)
	if err != nil {
		return err
	}

	user_action_required := false

	var sublog string
	if info.ThisIsCrossbuilder() || info.ThisIsSubarchBuilding() {
		sublog = "and pack this building site - package building completed'"
		user_action_required = true
	} else {
		sublog = "and continue with 'distr_source+' action"
		user_action_required = true
	}

	log.Info(
		"\n" +
			"-----------------\n" +
			"Now You have to create asm symlink inside include dir\n" +
			sublog + "\n" +
			"-----------------",
	)

	if user_action_required {
		return errors.New("user action required")
	}

	return nil
}

func (self *Builder_linux) BuilderActionDistrMan(
	log *logger.Logger,
) error {

	at := &buildingtools.Autotools{}

	args := make([]string, 0)
	args = append(
		args,
		[]string{
			"mandocs",
		}...,
	)
	args = append(
		args,
		self.crossbuild_params...,
	)

	err := at.Make(
		args,
		[]string{},
		buildingtools.Copy,
		"Makefile",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"make",
		log,
	)
	if err != nil {
		return err
	}

	err = os.MkdirAll(self.dst_man_dir, 0700)
	if err != nil {
		return err
	}

	man_files, err := filepath.Glob(
		path.Join(
			self.bs.GetDIR_SOURCE(),
			"Documentation",
			"DocBook",
			"man",
			"*.9.gz",
		),
	)
	if err != nil {
		return err
	}
	sort.Strings(man_files)

	log.Info(fmt.Sprintf("Copying %d man file(s)", len(man_files)))

	for _, i := range man_files {
		base := path.Base(i)
		log.Info(fmt.Sprintf("copying %s", base))

		err = filetools.CopyWithInfo(i, path.Join(self.dst_man_dir, base), nil)
		if err != nil {
			return err
		}

	}

	return nil
}

func (self *Builder_linux) BuilderActionDistrSource(
	log *logger.Logger,
) error {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	source_linux_dir_basename := fmt.Sprintf("linux-%s", info.PackageVersion)

	dst_dir := path.Join(
		self.bs.GetDIR_DESTDIR(),
		"usr",
		"src",
		source_linux_dir_basename,
	)

	err = os.MkdirAll(dst_dir, 0700)
	if err != nil {
		return err
	}

	err = filetools.CopyTree(
		self.bs.GetDIR_SOURCE(),
		dst_dir,
		false,
		true,
		false,
		true,
		log,
		filetools.CopyWithInfo,
	)
	if err != nil {
		return err
	}

	// src_file_list, err := ioutil.ReadDir(self.bs.GetDIR_SOURCE())
	// if err != nil {
	// 	return err
	// }
	//
	// for _, i := range src_file_list {
	//
	// 	// TODO: need to create own copy functions
	// 	cmd := exec.Command(
	// 		"cp",
	// 		"-afv",
	// 		path.Join(self.bs.GetDIR_SOURCE(), i.Name()),
	// 		dst_dir,
	// 	)
	// 	cmd.Stdout = log.StdoutLbl()
	// 	cmd.Stderr = log.StderrLbl()
	// 	err = cmd.Run()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (self *Builder_linux) BuilderActionDistrSymLink(
	log *logger.Logger,
) error {

	dst_usr_src_dir := path.Join(self.bs.GetDIR_DESTDIR(), "usr", "src")
	dst_usr_src_dir_linux := ""

	dst_usr_src_dir_files, err := ioutil.ReadDir(dst_usr_src_dir)
	if err != nil {
		return err
	}

	for _, i := range dst_usr_src_dir_files {
		if i.IsDir() {
			inb := path.Base(i.Name())
			if strings.HasPrefix(inb, "linux-") {
				dst_usr_src_dir_linux = inb
			}
		}
	}

	if dst_usr_src_dir_linux == "" {
		return errors.New("not found linux-* directory with sources in usr/src")
	}

	err = os.Symlink(dst_usr_src_dir_linux, path.Join(dst_usr_src_dir, "linux"))
	if err != nil {
		return err
	}

	return nil
}
