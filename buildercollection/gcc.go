package buildercollection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gcc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gcc(bs), nil
	}
}

type Builder_gcc struct {
	*Builder_std
}

func NewBuilder_gcc(bs basictypes.BuildingSiteCtlI) *Builder_gcc {

	self := new(Builder_gcc)

	self.Builder_std = NewBuilder_std(bs)

	self.SeparateBuildDir = true
	self.ForcedTarget = true

	self.AfterExtractCB = self.AfterExtract
	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.EditBuildConcurentJobsCountCB =
		func(log *logger.Logger, ret int) int {
			return 1
		}
	self.EditDistributeEnvCB =
		func(
			log *logger.Logger,
			ret environ.EnvVarEd,
		) (environ.EnvVarEd, error) {

			calc := self.bs.GetBuildingSiteValuesCalculator()

			hostdir, err := calc.CalculateHostDir()
			if err != nil {
				return nil, err
			}

			LD_LIBRARY_PATH := ret.Get("LD_LIBRARY_PATH", "")

			// TODO: yes, this is very bad and should be redone somehow
			for _, i := range basictypes.POSSIBLE_LIBDIR_NAMES {
				np := path.Join(hostdir, "multiarch", "i686-pc-linux-gnu", i)

				if nps, err := os.Stat(np); err == nil && nps.IsDir() {
					LD_LIBRARY_PATH += ":" + np
				}
			}

			ret.Set("LD_LIBRARY_PATH", LD_LIBRARY_PATH)
			ret.Set("LIBRARY_PATH", LD_LIBRARY_PATH)

			return ret, nil
		}

	return self
}

func (self *Builder_gcc) DefineActions() (basictypes.BuilderActions, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret, err := self.Builder_std.DefineActions()
	if err != nil {
		return nil, err
	}

	if info.ThisIsCrossbuilder() {
		ret = ret.Remove("build")
		ret = ret.Remove("distribute")

		ret = ret.AddActionsAfter(
			basictypes.BuilderActions{
				&basictypes.BuilderAction{"build_01", self.BuilderActionBuild_01},
				&basictypes.BuilderAction{"distribute_01", self.BuilderActionDistribute_01},

				&basictypes.BuilderAction{"intermediate_instruction_1", self.BuilderActionIntermediateInstruction_1},

				&basictypes.BuilderAction{"build_02", self.BuilderActionBuild_02},
				&basictypes.BuilderAction{"distribute_02", self.BuilderActionDistribute_02},

				&basictypes.BuilderAction{"intermediate_instruction_2", self.BuilderActionIntermediateInstruction_2},

				&basictypes.BuilderAction{"build_03", self.BuilderActionBuild_03},
				&basictypes.BuilderAction{"distribute_03", self.BuilderActionDistribute_03},
			},
			len(ret)-1,
		)
	} else {
		self.AfterDistributeCB = self.AfterDistribute
	}

	return ret, nil
}

func (self *Builder_gcc) AfterExtract(log *logger.Logger, err error) error {

	if err != nil {
		return err
	}

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	a_tools := new(buildingtools.Autotools)
	tar_dir := self.bs.GetDIR_TARBALL()
	files, err := ioutil.ReadDir(tar_dir)
	if err != nil {
		return err
	}

	NEEDED_PACKAGES := []string{
		// TODO: I don't know what to do
		"gmp",
		"mpc", "mpfr", "isl", "cloog",
		// #"gmp",
		// # NOTE: sometimes gcc could not compile with gmp.
		// #       so use system gmp
		// # requires compiler for bootstrap
		// # "binutils", "gdb", "glibc"
	}

	if info.ThisIsCrossbuilder() {
		NEEDED_PACKAGES = append(
			NEEDED_PACKAGES,
			[]string{"binutils", "gdb", "glibc"}...,
		)
	}

	for _, i := range NEEDED_PACKAGES {
		filename := ""
		for _, j := range files {
			b := path.Base(j.Name())
			if strings.HasPrefix(b, i) {
				filename = b
				break
			}
		}
		if filename == "" {
			log.Warning("not found tarball for " + i)
		}
	}

	for _, i := range NEEDED_PACKAGES {
		filename := ""
		for _, j := range files {
			b := path.Base(j.Name())
			if strings.HasPrefix(b, i) {
				filename = b
				break
			}
		}
		if filename != "" {
			err = a_tools.Extract(
				path.Join(tar_dir, filename),
				path.Join(self.bs.GetDIR_SOURCE(), i),
				self.bs.GetDIR_TEMP(),
				true,
				false,
				true,
				log,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *Builder_gcc) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	host_builders_dir, err := calc.CalculateHostCrossbuildersDir()
	if err != nil {
		return nil, err
	}

	hbt_opts, err := calc.CalculateAutotoolsHBTOptions()
	if err != nil {
		return nil, err
	}

	// 1
	if info.ThisIsCrossbuilder() {

		prefix := path.Join(host_builders_dir, info.CrossbuilderTarget)

		ret = make([]string, 0)
		ret = append(
			ret,
			[]string{
				"--prefix=" + prefix,
				"--mandir=" + path.Join(prefix, basictypes.DIRNAME_SHARE, "man"),
				"--sysconfdir=/etc",
				"--localstatedir=/var",
				"--enable-shared",
				"--disable-gold",
			}...,
		)
	}

	// 2
	if info.ThisIsCrossbuilder() {

		sysroot := path.Join(host_builders_dir, info.CrossbuilderTarget)

		ret = append(
			ret,
			[]string{
				"--disable-gold",

				"--enable-tls",
				"--enable-nls",
				"--enable-__cxa_atexit",
				"--enable-languages=c,c++,objc,obj-c++,fortran,ada",

				"--disable-bootstrap",
				// "--enable-bootstrap",

				"--enable-threads=posix",

				"--disable-multiarch",
				"--disable-multilib",

				"--enable-checking=release",
				"--enable-libada",
				"--enable-shared",

				// # use it when you haven"t built glibc basic parts yet
				// # "--without-headers",

				// # use it when you already have glibc headers and basic parts
				// # installed
				// # using this parameter may reqire creating hacky symlink
				// # pointing to /multiarch dir - you"ll see error what file not
				// # found.
				// # so after gcc and glibc built and installed - rebuild gcc both
				// # without --with-sysroot= and without --without-headers options
				"--with-sysroot=" + sysroot,

				// # TODO: need to try building without --with-sysroot if possible

			}...,
		)
		ret = append(
			ret,
			hbt_opts...,
		)

	}

	// 3
	if info.ThisIsCrossbuilding() {
		ret = append(
			ret,
			[]string{
				"--enable-tls",
				"--enable-nls",
				"--enable-__cxa_atexit",
				"--enable-languages=c,c++,objc,obj-c++,fortran,ada",
				"--disable-bootstrap",
				"--enable-threads=posix",

				"--disable-multiarch",
				"--disable-multilib",

				"--enable-checking=release",
				"--enable-libada",
				"--enable-shared",

				"--disable-gold",
			}...,
		)
	}

	// 4
	if !info.ThisIsCrossbuilding() && !info.ThisIsCrossbuilder() {
		host_dir, err := calc.CalculateHostDir()
		if err != nil {
			return nil, err
		}

		ret = append(
			ret,
			[]string{
				"--with-system-zlib",
				"--disable-gold",

				"--enable-tls",
				"--enable-nls",
				"--enable-__cxa_atexit",

				"--enable-languages=c,c++,objc,obj-c++,fortran,ada,go",

				"--disable-bootstrap",
				// "--enable-bootstrap",

				"--enable-threads=posix",

				// # wine Wow64 support requires this
				// # ld: Relocatable linking with relocations from format
				// #     elf64-x86-64 (aclui.Itv5tk.o) to format elf32-i386
				// #     (aclui.pnv73q.o) is not supported
				"--enable-multiarch",
				"--enable-multilib",

				// #"--disable-multiarch",
				// #"--disable-multilib",

				"--enable-checking=release",
				"--enable-libada",
				"--enable-shared",

				// #"--oldincludedir=" +
				// # wayround_i2p.utils.path.join(
				// #    self.get_host_dir(),
				// #    DIRNAME_INCLUDE
				// #    ),

				// #"--with-gxx-include-dir={}".format(
				// #    wayround_i2p.utils.path.join(
				// #        self.get_host_arch_dir(),
				// #        "include-c++",
				// #        #"c++",
				// #        #"5.2.0"
				// #        )
				// #    ),

				// # experimental option for this place
				// #
				// # NOTE: without it gcc tryes to use incompatible
				// #       /lib/crt*.o files.
				// #
				// # NOTE: this is required. else libs will be searched
				// #       in /lib and /usr/lib, but not in
				// #       /multihost/xxx/lib!:
				"--with-sysroot=" + host_dir,

				// # "--with-build-sysroot={}".format(self.get_host_arch_dir())
				"--with-native-system-header-dir=" + path.Join( //.format(
					// wayround_i2p.utils.path.join(
					// #"/",
					// #"multiarch",
					// # self.get_host_from_pkgi(),
					host_dir,
					basictypes.DIRNAME_INCLUDE,
				),
				// ),
				// #"--with-isl={}".format(self.get_host_dir())

			}...,
		)

		if strings.HasPrefix(info.Host, "x86_64") {
			ret = append(ret, "--enable-targets=all")
		}
	}

	return ret, nil
}

func (self *Builder_gcc) BuilderActionBuild_01(
	log *logger.Logger,
) error {

	self.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"all-gcc"}, nil
	}

	return self.BuilderActionBuild(log)
}

func (self *Builder_gcc) BuilderActionDistribute_01(
	log *logger.Logger,
) error {

	self.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install-gcc", "DESTDIR=" + self.bs.GetDIR_DESTDIR()}, nil
	}

	return self.BuilderActionDistribute(log)
}

func (self *Builder_gcc) BuilderActionIntermediateInstruction_1(
	log *logger.Logger,
) error {
	for _, i := range []string{
		"---------------",
		"Now You have to pack and install this gcc build,",
		"after what install linux-headers and glibc (headers and, maybe, some other parts).",
		"After what - continue building with 'build_02+' action",
		"---------------",
	} {
		log.Info(i)
	}
	return errors.New("user action required")
}

func (self *Builder_gcc) BuilderActionBuild_02(
	log *logger.Logger,
) error {

	self.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"all-target-libgcc"}, nil
	}

	return self.BuilderActionBuild(log)
}

func (self *Builder_gcc) BuilderActionDistribute_02(
	log *logger.Logger,
) error {

	self.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install-target-libgcc", "DESTDIR=" + self.bs.GetDIR_DESTDIR()}, nil
	}

	return self.BuilderActionDistribute(log)
}

func (self *Builder_gcc) BuilderActionIntermediateInstruction_2(
	log *logger.Logger,
) error {
	for _, i := range []string{
		"---------------",
		"Now You have to pack and install this gcc build and then complete",
		"glibc build and install it.",
		"After what - continue building this gcc from 'build_03+' action",
		"---------------",
	} {
		log.Info(i)
	}
	return errors.New("user action required")
}

func (self *Builder_gcc) BuilderActionBuild_03(
	log *logger.Logger,
) error {

	self.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{}, nil
	}

	return self.BuilderActionBuild(log)
}

func (self *Builder_gcc) BuilderActionDistribute_03(
	log *logger.Logger,
) error {

	self.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install", "DESTDIR=" + self.bs.GetDIR_DESTDIR()}, nil
	}

	return self.BuilderActionDistribute(log)
}

func (self *Builder_gcc) AfterDistribute(log *logger.Logger, ret error) error {
	if ret != nil {
		return ret
	}

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	for _, i := range []string{"go", "gofmt"} {
		old_name := path.Join(dst_install_prefix, "bin", i)
		new_name := path.Join(dst_install_prefix, "bin", fmt.Sprintf("gcc_%s", i))

		err = os.Rename(old_name, new_name)
		if err != nil {
			return err
		}
	}

	return nil
}
