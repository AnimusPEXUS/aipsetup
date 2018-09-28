package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["lvm2"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_lvm2(bs), nil
	}
}

type Builder_lvm2 struct {
	*Builder_std

	python_name string
}

func NewBuilder_lvm2(bs basictypes.BuildingSiteCtlI) *Builder_lvm2 {

	self := new(Builder_lvm2)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_lvm2) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-pkgconfig",
		}...,
	)

	return ret, nil
}
