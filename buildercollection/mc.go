package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["mc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_mc(bs)
	}
}

type Builder_mc struct {
	Builder_std
}

func NewBuilder_mc(bs basictypes.BuildingSiteCtlI) (*Builder_mc, error) {
	self := new(Builder_mc)
	self.Builder_std = *NewBuilder_std(bs)
	return self, nil
}
