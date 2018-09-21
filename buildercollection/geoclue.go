package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["geoclue"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_geoclue(bs)
	}
}

type Builder_geoclue struct {
	*Builder_std
}

func NewBuilder_geoclue(bs basictypes.BuildingSiteCtlI) (*Builder_geoclue, error) {

	self := new(Builder_geoclue)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_geoclue) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-python",
		}...,
	)

	return ret, nil
}
