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
	Builder_std
}

func NewBuilder_bc(bs basictypes.BuildingSiteCtlI) (*Builder_bc, error) {
	self := new(Builder_bc)

	self.Builder_std = *NewBuilder_std(bs)

	self.ForcedAutogen = true
	return self, nil
}
