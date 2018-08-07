package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["make"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_make(bs)
	}
}

type Builder_make struct {
	*Builder_std
}

func NewBuilder_make(bs basictypes.BuildingSiteCtlI) (*Builder_make, error) {

	self := new(Builder_make)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_make) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	if !info.ThisIsCrossbuilder() && !info.ThisIsCrossbuilding() {
		ret = append(
			ret,
			[]string{
				"--without-guile",
			}...,
		)
	}

	return ret, nil
}
