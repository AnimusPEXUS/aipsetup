package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["a52dec"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_a52dec(bs)
	}
}

type Builder_a52dec struct {
	*Builder_std
}

func NewBuilder_a52dec(bs basictypes.BuildingSiteCtlI) (*Builder_a52dec, error) {
	self := new(Builder_a52dec)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_a52dec) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{
			"--with-pic",
		}...,
	), nil
}
