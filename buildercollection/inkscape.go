package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["inkscape"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_inkscape(bs)
	}
}

type Builder_inkscape struct {
	*Builder_std
}

func NewBuilder_inkscape(bs basictypes.BuildingSiteCtlI) (*Builder_inkscape, error) {

	self := new(Builder_inkscape)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_inkscape) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(ret, "--enable-gtk3-experimental")

	return ret, nil
}
