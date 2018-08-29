package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["iw"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_iw(bs)
	}
}

type Builder_iw struct {
	*Builder_std
}

func NewBuilder_iw(bs basictypes.BuildingSiteCtlI) (*Builder_iw, error) {

	self := new(Builder_iw)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_iw) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	return ret, nil
}
