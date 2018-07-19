package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["iproute2"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_iproute2(bs), nil
	}
}

type Builder_iproute2 struct {
	*Builder_std

	makefile_flags []string
}

func NewBuilder_iproute2(bs basictypes.BuildingSiteCtlI) *Builder_iproute2 {

	self := new(Builder_iproute2)

	self.Builder_std = NewBuilder_std(bs)
	self.EditBuildArgsCB = self.EditBuildArgs
	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

func (self *Builder_iproute2) EditBuildArgs(log *logger.Logger, ret []string) ([]string, error) {
	calc := self.bs.GetBuildingSiteValuesCalculator()

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
			"SBINDIR=" + path.Join(install_prefix, "sbin"),
			"KERNEL_INCLUDE" + path.Join(install_prefix, "include"),
			"DBM_INCLUDE" + path.Join(install_prefix, "include"),
		}...,
	)

	self.makefile_flags = mff

	ret = append(ret, self.makefile_flags...)

	return ret, nil
}

func (self *Builder_iproute2) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(ret, self.makefile_flags...)

	//	calc := self.bs.GetBuildingSiteValuesCalculator()

	//	install_prefix, err := calc.CalculateInstallPrefix()
	//	if err != nil {
	//		return err
	//	}

	//	opts, err := calc.CalculateAutotoolsAllOptionsMap()
	//	if err != nil {
	//		return err
	//	}

	//	ret = append(ret, opts.Strings()...)

	//	ret = append(
	//		ret,
	//		[]string{
	//			"PREFIX=" + install_prefix,
	//			"SBINDIR=" + path.Join(install_prefix, "sbin"),
	//			"KERNEL_INCLUDE" + path.Join(install_prefix, "include"),
	//			"DBM_INCLUDE" + path.Join(install_prefix, "include"),
	//		},
	//	)

	return ret, nil
}
