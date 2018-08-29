package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["hdparm"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_hdparm(bs)
	}
}

type Builder_hdparm struct {
	*Builder_std
}

func NewBuilder_hdparm(bs basictypes.BuildingSiteCtlI) (*Builder_hdparm, error) {

	self := new(Builder_hdparm)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_hdparm) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	return ret, nil
}
