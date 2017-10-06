package buildercollection

import "errors"

type BuilderI interface {
	DefineActions() []string
}

func NewBuilder(name string) (BuilderI, error) {
	switch name {
	default:
		return nil, errors.New("no such builder")
	case "std":
		return new(BuilderAutotoolsStd), nil
	}
}
