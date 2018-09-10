package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

// WARNING: fontconfig's cmake building variant is not for unixes

func init() {
	Index["fontconfig"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_fontconfig(bs)
	}
}

type Builder_fontconfig struct {
	*Builder_std
}

func NewBuilder_fontconfig(bs basictypes.BuildingSiteCtlI) (*Builder_fontconfig, error) {

	self := new(Builder_fontconfig)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_fontconfig) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{}...,
	)

	return ret, nil
}
