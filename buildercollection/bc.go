package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["bc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_bc(bs)
	}
}

type Builder_bc struct {
	BuilderStdAutotools
}

func NewBuilder_bc(bs basictypes.BuildingSiteCtlI) (*Builder_bc, error) {
	self := new(Builder_bc)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	self.ForcedAutogen = true
	return self, nil
}
