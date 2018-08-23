package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["glib"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_glib(bs), nil
	}
}

type Builder_glib struct {
	*Builder_std
}

func NewBuilder_glib(bs basictypes.BuildingSiteCtlI) *Builder_glib {

	self := new(Builder_glib)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_glib) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	res, err := calc.CalculateInstallPrefixExecutable("python3")
	if err != nil {
		return nil, err
	}

	ret = append(ret, "--with-python="+res)

	return ret, nil
}
