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
	BuilderStdAutotools
}

func NewBuilder_mc(bs basictypes.BuildingSiteCtlI) (*Builder_mc, error) {
	self := new(Builder_mc)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	return self, nil
}
