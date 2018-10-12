package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["cogl"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_cogl(bs)
	}
}

type Builder_cogl struct {
	*Builder_std
}

func NewBuilder_cogl(bs basictypes.BuildingSiteCtlI) (*Builder_cogl, error) {
	self := new(Builder_cogl)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_cogl) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{
			"--enable-waylang-egl-platform",
			"--enable-wayland-egl-server",
			"--enable-kms-egl-platform",
			"--enable-gtk-doc",
		}...,
	)
	return ret, nil
}
