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
	Index["binutils"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_binutils(bs), nil
	}
}

type Builder_binutils struct {
	*Builder_std
}

func NewBuilder_binutils(bs basictypes.BuildingSiteCtlI) *Builder_binutils {

	self := new(Builder_binutils)
	self.Builder_std = NewBuilder_std(bs)

	self.AfterExtractCB = self.AfterExtract
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

//func (self *Builder_binutils) DefineActions() (basictypes.BuilderActions, error) {
//	return self.Builder_std.DefineActions()
//}

func (self *Builder_binutils) AfterExtract(log *logger.Logger, err error) error {

	if err != nil {
		return err
	}

	a_tools := new(buildingtools.Autotools)
	tar_dir := self.GetBuildingSiteCtl().GetDIR_TARBALL()
	files, err := ioutil.ReadDir(tar_dir)
	if err != nil {
		return err
	}

	for _, i := range []string{
		"gmp", "mpc", "mpfr", "isl", "cloog",
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
			path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), i),
			self.GetBuildingSiteCtl().GetDIR_TEMP(),
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

func (self *Builder_binutils) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	// cb, err := calc.CalculateIsCrossbuilder()
	// if err != nil {
	// 	return nil, err
	// }

	if info.ThisIsCrossbuilder() {

		host_builders_dir, err := calc.CalculateHostCrossbuildersDir()
		if err != nil {
			return nil, err
		}

		prefix := path.Join(host_builders_dir, info.CrossbuilderTarget)

		hbt_opts, err := calc.CalculateAutotoolsHBTOptions()
		if err != nil {
			return nil, err
		}

		ret = make([]string, 0)
		ret = append(
			ret,
			[]string{
				"--prefix=" + prefix,
				"--mandir=" + path.Join(prefix, basictypes.DIRNAME_SHARE, "man"),
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

	sysrootdir, err := calc.CalculateHostDir()
	if err != nil {
		return nil, err
	}

	// if info.Host != info.HostArch {
	// 	sysrootdir, err = calc.CalculateHostArchDir()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	ret = append(
		ret,
		[]string{

			"--enable-targets=all",

			"--enable-64-bit-bfd",
			"--disable-werror",
			"--enable-libada",
			"--enable-libssp",
			"--enable-objc-gc",

			"--enable-lto",
			"--enable-ld",

			// # NOTE: no google software in Lilith distro
			"--disable-gold",
			"--without-gold",

			// # this is required. else libs will be searched in /lib and
			// # /usr/lib, but not in /multihost/xxx/lib!:
			"--with-sysroot=" + sysrootdir,

			// # more experiment:
			"--enable-multiarch",
			"--enable-multilib",
		}...,
	)

	if info.ThisIsCrossbuilder() {
		ret = append(ret, "--with-sysroot")
	}

	return ret, nil
}
