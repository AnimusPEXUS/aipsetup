package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["criu"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_criu(bs), nil
	}
}

type Builder_criu struct {
	*Builder_std

	makefile_flags []string
}

func NewBuilder_criu(bs basictypes.BuildingSiteCtlI) *Builder_criu {

	self := new(Builder_criu)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditBuildArgsCB = self.EditBuildArgs
	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

func (self *Builder_criu) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	return ret, nil
}

func (self *Builder_criu) EditBuildArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	variant, err := calc.CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	opts, err := calc.CalculateAutotoolsAllOptionsMap()
	if err != nil {
		return nil, err
	}

	mff := make([]string, 0)

	mff = append(mff, opts.Strings()...)

	mff = append(
		mff,
		[]string{
			"PREFIX=" + install_prefix,
			//			"SBINDIR=" + path.Join(install_prefix, "sbin"),
			//			"KERNEL_INCLUDE" + path.Join(install_prefix, "include"),
			//			"DBM_INCLUDE" + path.Join(install_prefix, "include"),
			"CC=" + info.Host + "-gcc -m" + variant,
		}...,
	)

	self.makefile_flags = mff

	ret = append(ret, self.makefile_flags...)

	return ret, nil
}

func (self *Builder_criu) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(ret, self.makefile_flags...)

	return ret, nil
}
