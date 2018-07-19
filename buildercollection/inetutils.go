package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["inetutils"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_inetutils(bs), nil
	}
}

type Builder_inetutils struct {
	*Builder_std
}

func NewBuilder_inetutils(bs basictypes.BuildingSiteCtlI) *Builder_inetutils {

	self := new(Builder_inetutils)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_inetutils) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(ret, "--disable-logger")

	return ret, nil
}
