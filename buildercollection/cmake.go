package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["cmake"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_cmake(bs)
	}
}

// TODO: finish this

type Builder_cmake struct {
	Builder_std_cmake
}

func NewBuilder_cmake(bs basictypes.BuildingSiteCtlI) (*Builder_cmake, error) {
	self := new(Builder_cmake)
	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = *t
	}
	return self, nil
}
