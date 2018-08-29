package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["mandb"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_mandb(bs)
	}
}

type Builder_mandb struct {
	*Builder_std
}

func NewBuilder_mandb(bs basictypes.BuildingSiteCtlI) (*Builder_mandb, error) {

	self := new(Builder_mandb)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_mandb) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-cache-owner",
			"--disable-setuid",
		}...,
	)

	return ret, nil
}
