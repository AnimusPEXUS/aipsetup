package buildercollection

import (
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func Get(name string, bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
	if t, ok := Index[name]; ok {
		return t(bs)
	} else {
		return nil, errors.New("builder not found")
	}
}

var Index = make(
	map[string](func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error)),
	0,
)
