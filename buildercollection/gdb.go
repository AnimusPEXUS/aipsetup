package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gdb"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gdb(bs), nil
	}
}

type Builder_gdb struct {
	*Builder_std
}

func NewBuilder_gdb(bs basictypes.BuildingSiteCtlI) *Builder_gdb {

	self := new(Builder_gdb)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_gdb) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-system-readline",
			"--with-curses",
			//"--disable-tui" ,// # TODO: can't compile with tui: 2016-May-10
		}...,
	)

	return ret, nil
}
