package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["man_pages"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_man_pages(bs)
	}
}

type Builder_man_pages struct {
	*Builder_std
}

func NewBuilder_man_pages(bs basictypes.BuildingSiteCtlI) (*Builder_man_pages, error) {

	self := new(Builder_man_pages)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_man_pages) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	return ret, nil
}
