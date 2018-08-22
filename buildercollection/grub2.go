package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["grub2"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_grub2(bs), nil
	}
}

type Builder_grub2 struct {
	*Builder_std
}

func NewBuilder_grub2(bs basictypes.BuildingSiteCtlI) *Builder_grub2 {

	self := new(Builder_grub2)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_grub2) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-werror",
			"--disable-efiemu",
		}...,
	)

	return ret, nil
}
