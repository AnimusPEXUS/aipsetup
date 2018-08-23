package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_simple_makefile"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_simple_makefile(bs), nil
	}
}

type Builder_std_simple_makefile struct {
	*Builder_std
}

func NewBuilder_std_simple_makefile(bs basictypes.BuildingSiteCtlI) *Builder_std_simple_makefile {
	self := new(Builder_std_simple_makefile)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

func (self *Builder_std_simple_makefile) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")
	return ret, nil
}

func (self *Builder_std_simple_makefile) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"prefix=" + install_prefix,
		}...,
	)

	return ret, nil
}
