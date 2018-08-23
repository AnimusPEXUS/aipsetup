package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["kbd"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_kbd(bs)
	}
}

type Builder_kbd struct {
	*Builder_std
}

func NewBuilder_kbd(bs basictypes.BuildingSiteCtlI) (*Builder_kbd, error) {

	self := new(Builder_kbd)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_kbd) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefix()

	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--enable-nls",
			"--datarootdir=" + path.Join(install_prefix, "share", "kbd"),
		}...,
	)

	return ret, nil
}
