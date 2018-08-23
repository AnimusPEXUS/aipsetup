package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["subversion"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_subversion(bs)
	}
}

type Builder_subversion struct {
	*Builder_std
}

func NewBuilder_subversion(bs basictypes.BuildingSiteCtlI) (*Builder_subversion, error) {
	self := new(Builder_subversion)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_subversion) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--with-serf=" + install_prefix,
			//			"--with-ssl",
			//			"--with-openssl",
			"--without-python",
		}...,
	)
	return ret, nil
}
