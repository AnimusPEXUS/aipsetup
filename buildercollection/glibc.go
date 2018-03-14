package buildercollection

import (
	"errors"
	"fmt"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["glibc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderGlibc(bs), nil
	}
}

type BuilderGlibc struct {
	bs basictypes.BuildingSiteCtlI

	std_builder *BuilderStdAutotools
}

func NewBuilderGlibc(bs basictypes.BuildingSiteCtlI) *BuilderGlibc {

	self := new(BuilderGlibc)
	self.bs = bs

	self.std_builder = NewBuilderStdAutotools(bs)

	self.std_builder.SeparateBuildDir = true
	self.std_builder.ForcedTarget = true
	self.std_builder.ApplyHostSpecCompilerOptions = true

	self.std_builder.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *BuilderGlibc) DefineActions() (basictypes.BuilderActions, error) {

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
				&basictypes.BuilderAction{"distribute_01", self.BuilderActionDistribute_01},
				&basictypes.BuilderAction{"distribute_01_2", self.BuilderActionDistribute_01_2},
				&basictypes.BuilderAction{"distribute_01_3", self.BuilderActionDistribute_01_3},
				&basictypes.BuilderAction{"distribute_01_4", self.BuilderActionDistribute_01_4},
				&basictypes.BuilderAction{"distribute_01_5", self.BuilderActionDistribute_01_5},

				&basictypes.BuilderAction{"intermediate_instruction", self.BuilderActionIntermediateInstruction},

				&basictypes.BuilderAction{"build_02", self.BuilderActionBuild_02},
				&basictypes.BuilderAction{"distribute_02", self.BuilderActionDistribute_02},
			},
			len(ret)-1,
		)
	}

	return ret, nil
}

func (self *BuilderGlibc) BuilderActionEditInfo(
	log *logger.Logger,
) error {

	log.Info("Checking info file editing need")

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	calc := self.bs.ValuesCalculator()

	cb, err := calc.CalculateIsCrossbuilder()
	if err != nil {
		return err
	}

	if cb {
		info.PackageName = fmt.Sprintf("cb-glibc-%s", info.Target)
	} else {
		info.PackageName = "glibc"
	}

	err = self.bs.WriteInfo(info)
	if err != nil {
		return err
	}

	return nil
}

func (self *BuilderGlibc) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.ValuesCalculator()

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	cb, err := calc.CalculateIsCrossbuilder()
	if err != nil {
		return nil, err
	}

	var with_headers string

	if t, err := calc.CalculateHostDir(); err != nil {
		return nil, err
	} else {
		with_headers = path.Join(t, "include")
	}

	if cb {

		host_builders_dir, err := calc.CalculateHostCrossbuildersDir()
		if err != nil {
			return nil, err
		}

		prefix := path.Join(
			host_builders_dir,
			info.Target,
		)

		with_headers = path.Join(prefix, "include")

		hbt_opts, err := calc.CalculateAutotoolsHBTOptions()
		if err != nil {
			return nil, err
		}

		ret = make([]string, 0)
		ret = append(
			ret,
			[]string{
				"--prefix=" + prefix,
				"--mandir=" + path.Join(prefix, "share", "man"),
				"--sysconfdir=/etc",
				"--localstatedir=/var",
				"--enable-shared",
			}...,
		)
		ret = append(
			ret,
			hbt_opts...,
		)

	}

	// host_dir, err := calc.CalculateHostDir()
	// if err != nil {
	// 	return nil, err
	// }

	ret = append(
		ret,
		[]string{

			"--enable-obsolete-rpc",
			"--enable-kernel=4.9",
			"--enable-tls",
			"--with-elf",
			// # disabled those 3 items on 2 jul 2015
			// # reenabled those 3 items on 11 aug 2015: sims I need it
			"--enable-multi-arch",
			"--enable-multiarch",
			"--enable-multilib",

			// # this is from configure --help. configure looking for
			// # linux/version.h file
			// #"--with-headers=/usr/src/linux/include",
			"--with-headers=" + with_headers,
			"--enable-shared",

			// # temp
			// #"libc_cv_forced_unwind=yes"
		}...,
	)

	// '''
	// # NOTE: it's not working
	// # NOTE: don't remove this block. it's for informational reason
	// if self.get_arch_from_pkgi().startswith('x86_64'):
	//     ret += ['slibdir=lib64']
	// else:
	//     ret += ['slibdir=lib']
	// '''

	// NOTE: this block is not found in pythonish aipsetup clibc builder, and
	//       is only binutils.go builder copy result. probably can be removed
	//       without any consiquences.
	// if cb {
	// 	ret = append(ret, "--with-sysroot")
	// }

	if cb {
		ret = append(
			ret,
			[]string{
			// TODO: 2018 mar 14. it was easy to comment/uncomment when it was
			//       python script, but from now on this should be done somehow
			//       wiser. to be done..

			// # this can be commented whan gcc fully built and installed
			// #'libc_cv_forced_unwind=yes',
			//
			// # this parameter is required to build `build_02+'
			// # stage.  comment and completle rebuild this glibc
			// # after rebuilding gcc without --without-headers and
			// # with --with-sysroot parameter.
			// #
			// # 'libc_cv_ssp=no'
			// #
			// # else You will see errors like this:
			// #     gethnamaddr.c:185: undefined reference to
			// #     `__stack_chk_guard'

			}...,
		)
	}

	return ret, nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01(
	log *logger.Logger,
) error {
	return errors.New("todo")
	// TODO: high priority TODO
	return nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01_2(
	log *logger.Logger,
) error {
	return nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01_3(
	log *logger.Logger,
) error {
	return nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01_4(
	log *logger.Logger,
) error {
	return nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01_5(
	log *logger.Logger,
) error {
	return nil
}

func (self *BuilderGlibc) BuilderActionIntermediateInstruction(
	log *logger.Logger,
) error {
	return nil
}

func (self *BuilderGlibc) BuilderActionBuild_02(
	log *logger.Logger,
) error {
	return nil
}

func (self *BuilderGlibc) BuilderActionDistribute_02(
	log *logger.Logger,
) error {
	return nil
}
