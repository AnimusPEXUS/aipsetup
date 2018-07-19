package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gdbm"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gdbm(bs), nil
	}
}

type Builder_gdbm struct {
	*Builder_std
}

func NewBuilder_gdbm(bs basictypes.BuildingSiteCtlI) *Builder_gdbm {

	self := new(Builder_gdbm)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_gdbm) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(ret, "--enable-libgdbm-compat")

	return ret, nil
}
