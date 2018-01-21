package buildercollection

import "github.com/AnimusPEXUS/aipsetup/basictypes"

func init() {
	// Index["linux"] = func(bs basictypes.BuildingSiteCtlI) basictypes.BuilderI {
	// 	return NewBuilderLinux(bs)
	// }
}

type BuilderLinux struct {
}

func NewBuilderLinux(bs basictypes.BuildingSiteCtlI) *BuilderLinux {
	self := new(BuilderLinux)

	return self
}
