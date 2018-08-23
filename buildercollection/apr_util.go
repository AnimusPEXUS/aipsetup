package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["apr_util"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_apr_util(bs), nil
	}
}

type Builder_apr_util struct {
	*Builder_std
}

func NewBuilder_apr_util(bs basictypes.BuildingSiteCtlI) *Builder_apr_util {
	self := new(Builder_apr_util)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self
}

func (self *Builder_apr_util) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	apr_1_config, err := calc.CalculateInstallPrefixExecutable("apr-1-config")
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--with-apr=" + apr_1_config,
			"--with-berkeley-db=" + install_prefix,
		}...,
	)

	return ret, nil
}
