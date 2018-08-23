package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["utillinux"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_utillinux(bs)
	}
}

type Builder_utillinux struct {
	*Builder_std
}

func NewBuilder_utillinux(bs basictypes.BuildingSiteCtlI) (*Builder_utillinux, error) {

	self := new(Builder_utillinux)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_utillinux) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-python=3",
		}...,
	)

	return ret, nil
}
