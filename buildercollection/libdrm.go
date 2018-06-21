package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libdrm"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_libdrm(bs)
	}
}

type Builder_libdrm struct {
	*Builder_std
}

func NewBuilder_libdrm(bs basictypes.BuildingSiteCtlI) (*Builder_libdrm, error) {
	self := new(Builder_libdrm)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_libdrm) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{
			"--enable-udev",
			"--disable-valgrind",
		}...,
	), nil
}
