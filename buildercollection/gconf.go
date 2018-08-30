package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gconf"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gconf(bs)
	}
}

type Builder_gconf struct {
	*Builder_std
}

func NewBuilder_gconf(bs basictypes.BuildingSiteCtlI) (*Builder_gconf, error) {
	self := new(Builder_gconf)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_gconf) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-orbit",
		}...,
	)
	return ret, nil
}
