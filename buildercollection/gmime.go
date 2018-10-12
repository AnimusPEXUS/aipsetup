package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gmime"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gmime(bs)
	}
}

type Builder_gmime struct {
	*Builder_std
}

func NewBuilder_gmime(bs basictypes.BuildingSiteCtlI) (*Builder_gmime, error) {
	self := new(Builder_gmime)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_gmime) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{
			"--disable-mono",
		}...,
	), nil
}
