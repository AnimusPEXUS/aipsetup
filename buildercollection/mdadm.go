package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["mdadm"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_mdadm(bs)
	}
}

type Builder_mdadm struct {
	*Builder_std
}

func NewBuilder_mdadm(bs basictypes.BuildingSiteCtlI) (*Builder_mdadm, error) {

	self := new(Builder_mdadm)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_mdadm) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	return ret, nil
}
