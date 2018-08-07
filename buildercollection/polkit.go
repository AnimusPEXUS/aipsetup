package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["polkit"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_polkit(bs)
	}
}

type Builder_polkit struct {
	*Builder_std
}

func NewBuilder_polkit(bs basictypes.BuildingSiteCtlI) (*Builder_polkit, error) {

	self := new(Builder_polkit)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_polkit) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-libsystemd-login=yes",
			"--enable-introspection=yes",
			"--with-mozjs=mozjs-17.0",
		}...,
	)

	return ret, nil
}
