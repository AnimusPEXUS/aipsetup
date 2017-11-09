package buildercollection

import "github.com/AnimusPEXUS/aipsetup/basictypes"

var Index = map[string](func() basictypes.BuilderI){
	"std": func() basictypes.BuilderI {
		return new(BuilderAutotoolsStd)
	},
}
