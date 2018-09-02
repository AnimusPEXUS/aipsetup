package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["wayland"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_wayland(bs)
	}
}

type Builder_wayland struct {
	*Builder_std
}

func NewBuilder_wayland(bs basictypes.BuildingSiteCtlI) (*Builder_wayland, error) {

	self := new(Builder_wayland)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_wayland) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-documentation",
		}...,
	)

	return ret, nil
}
