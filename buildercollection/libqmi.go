package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libqmi"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libqmi(bs)
	}
}

type Builder_libqmi struct {
	*Builder_std
}

func NewBuilder_libqmi(bs basictypes.BuildingSiteCtlI) (*Builder_libqmi, error) {

	self := new(Builder_libqmi)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_libqmi) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-more-warnings=no",
		}...,
	)

	return ret, nil
}
