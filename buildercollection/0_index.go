package buildercollection

import "github.com/AnimusPEXUS/aipsetup/basictypes"

var Index = map[string](func(bs basictypes.BuildingSiteCtlI) basictypes.BuilderI){

	"std": func(bs basictypes.BuildingSiteCtlI) basictypes.BuilderI {
		return NewBuilderStdAutotools(bs)
	},
}
