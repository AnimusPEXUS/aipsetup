package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["shadow"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_shadow(bs)
	}
}

type Builder_shadow struct {
	*Builder_std
}

func NewBuilder_shadow(bs basictypes.BuildingSiteCtlI) (*Builder_shadow, error) {
	self := new(Builder_shadow)
	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditAutogenFailIsOkCB = self.EditAutogenFailIsOk

	return self, nil
}

func (self *Builder_shadow) EditAutogenFailIsOk(log *logger.Logger, ret bool) (bool, error) {
	// NOTE: yes, do overriding error state
	return true, nil
}

func (self *Builder_shadow) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			// TODO: I have to make selinux support someday
			"--without-selinux",
			"--enable-man",
		}...,
	)

	return ret, nil
}
