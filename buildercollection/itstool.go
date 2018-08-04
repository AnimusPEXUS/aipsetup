package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["itstool"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_itstool(bs)
	}
}

type Builder_itstool struct {
	*Builder_std
}

func NewBuilder_itstool(bs basictypes.BuildingSiteCtlI) (*Builder_itstool, error) {

	self := new(Builder_itstool)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_itstool) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	python, err := self.bs.GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("python2")
	if err != nil {
		return nil, err
	}

	ret = append(ret, "PYTHON="+python)

	return ret, nil
}
