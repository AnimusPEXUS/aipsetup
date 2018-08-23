package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["apr"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_apr(bs), nil
	}
}

type Builder_apr struct {
	*Builder_std
}

func NewBuilder_apr(bs basictypes.BuildingSiteCtlI) *Builder_apr {
	self := new(Builder_apr)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self
}

func (self *Builder_apr) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--with-installbuilddir=" +
				path.Join(install_prefix, "share", "apr", "build-1"),
		}...,
	)

	return ret, nil
}
