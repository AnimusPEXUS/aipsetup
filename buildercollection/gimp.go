package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gimp"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gimp(bs)
	}
}

type Builder_gimp struct {
	*Builder_std
}

func NewBuilder_gimp(bs basictypes.BuildingSiteCtlI) (*Builder_gimp, error) {

	self := new(Builder_gimp)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_gimp) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-python",
		}...,
	)

	return ret, nil
}
