package buildercollection

import (
	"errors"
	"io/ioutil"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gcc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderGCC(bs), nil
	}
}

type BuilderGCC struct {
	bs basictypes.BuildingSiteCtlI

	std_builder *BuilderStdAutotools
}

func NewBuilderGCC(bs basictypes.BuildingSiteCtlI) *BuilderGCC {

	self := new(BuilderGCC)
	self.bs = bs

	self.std_builder = NewBuilderStdAutotools(bs)

	self.std_builder.SeparateBuildDir = true
	self.std_builder.ForcedTarget = true

	self.std_builder.AfterExtractCB = self.AfterExtract
	self.std_builder.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *BuilderGCC) DefineActions() (basictypes.BuilderActions, error) {

	calc := self.bs.ValuesCalculator()

	cb, err := calc.CalculateIsCrossbuilder()
	if err != nil {
		return nil, err
	}

	ret, err := self.std_builder.DefineActions()
	if err != nil {
		return nil, err
	}

	ret = ret.AddBefore(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{"edit_package_info", self.BuilderActionEditInfo},
		},
		0,
	)

	if cb {
		ret = ret.Remove("build")
		ret = ret.Remove("distribute")

		ret = ret.AddAfter(
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
	}

	return ret, nil
}

func (self *BuilderGCC) BuilderActionEditInfo(
	log *logger.Logger,
) error {

	log.Info("Checking info file editing need")

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	// calc := self.bs.ValuesCalculator()

	// cb, err := calc.CalculateIsCrossbuilder()
	// if err != nil {
	// 	return err
	// }

	// if cb {
	// 	info.PackageName = fmt.Sprintf("cb-gcc-%s", info.Target)
	// } else {
	// 	info.PackageName = "gcc"
	// }

	err = self.bs.WriteInfo(info)
	if err != nil {
		return err
	}

	return nil
}

func (self *BuilderGCC) AfterExtract(log *logger.Logger, err error) error {

	if err != nil {
		return err
	}

	a_tools := new(buildingtools.Autotools)
	tar_dir := self.bs.GetDIR_TARBALL()
	files, err := ioutil.ReadDir(tar_dir)
	if err != nil {
		return err
	}

	for _, i := range []string{
		"gmp",
		"mpc", "mpfr", "isl", "cloog",
		// #"gmp",
		// # NOTE: sometimes gcc could not compile with gmp.
		// #       so use system gmp
		// # requires compiler for bootstrap
		// # "binutils", "gdb", "glibc"
		// TODO: make "binutils", "gdb", "glibc" autoadd if crosscompiler
		//       building

	} {
		filename := ""
		for _, j := range files {
			b := path.Base(j.Name())
			if strings.HasPrefix(b, i) {
				filename = b
				break
			}
		}
		if filename == "" {
			return errors.New("not found tarball for " + i)
		}

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

	return nil
}

func (self *BuilderGCC) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.ValuesCalculator()

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	cbuilder, err := calc.CalculateIsCrossbuilder()
	if err != nil {
		return nil, err
	}

	cbuild, err := calc.CalculateIsCrossbuild()
	if err != nil {
		return nil, err
	}

	// host_builders_dir, err := calc.CalculateHostCrossbuildersDir()
	// if err != nil {
	// 	return nil, err
	// }

	// hbt_opts, err := calc.CalculateAutotoolsHBTOptions()
	// if err != nil {
	// 	return nil, err
	// }

	// 1
	// if cbuilder {
	//
	// 	prefix := path.Join(
	// 		host_builders_dir,
	// 		info.Target,
	// 	)
	//
	// 	ret = make([]string, 0)
	// 	ret = append(
	// 		ret,
	// 		[]string{
	// 			"--prefix=" + prefix,
	// 			"--mandir=" + path.Join(prefix, "share", "man"),
	// 			"--sysconfdir=/etc",
	// 			"--localstatedir=/var",
	// 			"--enable-shared",
	// 			"--disable-gold",
	// 		}...,
	// 	)
	// }

	// 2
	// if cbuilder {
	//
	// 	sysroot := path.Join(host_builders_dir, info.Target)
	//
	// 	ret = append(
	// 		ret,
	// 		[]string{
	// 			"--disable-gold",
	//
	// 			"--enable-tls",
	// 			"--enable-nls",
	// 			"--enable-__cxa_atexit",
	// 			"--enable-languages=c,c++,objc,obj-c++,fortran,ada",
	// 			"--disable-bootstrap",
	// 			"--enable-threads=posix",
	//
	// 			"--disable-multiarch",
	// 			"--disable-multilib",
	//
	// 			"--enable-checking=release",
	// 			"--enable-libada",
	// 			"--enable-shared",
	//
	// 			// # use it when you haven"t built glibc basic parts yet
	// 			// # "--without-headers",
	//
	// 			// # use it when you already have glibc headers and basic parts
	// 			// # installed
	// 			// # using this parameter may reqire creating hacky symlink
	// 			// # pointing to /multiarch dir - you"ll see error what file not
	// 			// # found.
	// 			// # so after gcc and glibc built and installed - rebuild gcc both
	// 			// # without --with-sysroot= and without --without-headers options
	// 			"--with-sysroot=" + sysroot,
	//
	// 			// # TODO: need to try building without --with-sysroot if possible
	//
	// 		}...,
	// 	)
	// 	ret = append(
	// 		ret,
	// 		hbt_opts...,
	// 	)
	//
	// }

	// 3
	if cbuild {
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

	if !cbuild && !cbuilder {
		host_dir, err := calc.CalculateHostDir()
		if err != nil {
			return nil, err
		}

		ret = append(
			ret,
			[]string{
				"--disable-gold",

				"--enable-tls",
				"--enable-nls",
				"--enable-__cxa_atexit",

				// # NOTE: gcc somtimes fails to crossbuild self with go enabled
				"--enable-languages=c,c++,objc,obj-c++,fortran,ada,go",

				"--disable-bootstrap",

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
				// #    "include"
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
					"include",
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

func (self *BuilderGCC) BuilderActionBuild_01(
	log *logger.Logger,
) error {

	self.std_builder.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"all-gcc"}, nil
	}

	return self.std_builder.BuilderActionBuild(log)
}

func (self *BuilderGCC) BuilderActionDistribute_01(
	log *logger.Logger,
) error {

	self.std_builder.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install-gcc", "DESTDIR=" + self.bs.GetDIR_DESTDIR()}, nil
	}

	return self.std_builder.BuilderActionDistribute(log)
}

func (self *BuilderGCC) BuilderActionIntermediateInstruction_1(
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

func (self *BuilderGCC) BuilderActionBuild_02(
	log *logger.Logger,
) error {

	self.std_builder.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"all-target-libgcc"}, nil
	}

	return self.std_builder.BuilderActionBuild(log)
}

func (self *BuilderGCC) BuilderActionDistribute_02(
	log *logger.Logger,
) error {

	self.std_builder.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install-target-libgcc", "DESTDIR=" + self.bs.GetDIR_DESTDIR()}, nil
	}

	return self.std_builder.BuilderActionDistribute(log)
}

func (self *BuilderGCC) BuilderActionIntermediateInstruction_2(
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

func (self *BuilderGCC) BuilderActionBuild_03(
	log *logger.Logger,
) error {

	self.std_builder.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{}, nil
	}

	return self.std_builder.BuilderActionBuild(log)
}

func (self *BuilderGCC) BuilderActionDistribute_03(
	log *logger.Logger,
) error {

	self.std_builder.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install", "DESTDIR=" + self.bs.GetDIR_DESTDIR()}, nil
	}

	return self.std_builder.BuilderActionDistribute(log)
}
