package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["autoconf"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_autoconf(bs)
	}
}

type Builder_autoconf struct {
	*Builder_std
}

func NewBuilder_autoconf(bs basictypes.BuildingSiteCtlI) (*Builder_autoconf, error) {

	self := new(Builder_autoconf)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_autoconf) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--datarootdir=" + path.Join(install_prefix, "share"),
		}...,
	)

	return ret, nil
}
