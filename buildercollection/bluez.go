package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["bluez"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_bluez(bs), nil
	}
}

type Builder_bluez struct {
	Builder_std
}

func NewBuilder_bluez(bs basictypes.BuildingSiteCtlI) *Builder_bluez {
	self := new(Builder_bluez)

	self.Builder_std = *NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

// TODO: building may require 'LDFLAGS=-ltinfow' - testing needed

func (self *Builder_bluez) EditConfigureArgs(log *logger.Logger, ret []string) (
	[]string, error,
) {
	ret = append(ret, "--enable-library")

	return ret, nil
}
