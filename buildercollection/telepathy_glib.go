package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["telepathy_glib"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_telepathy_glib(bs)
	}
}

type Builder_telepathy_glib struct {
	*Builder_std
}

func NewBuilder_telepathy_glib(bs basictypes.BuildingSiteCtlI) (*Builder_telepathy_glib, error) {

	self := new(Builder_telepathy_glib)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_telepathy_glib) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-introspection",
			"--enable-vala-bindings",
		}...,
	)

	return ret, nil
}
