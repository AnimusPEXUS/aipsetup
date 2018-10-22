package buildercollection

import "github.com/AnimusPEXUS/aipsetup/basictypes"

func init() {
	Index["wine"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_wine(bs), nil
	}
}

func NewBuilder_wine(bs basictypes.BuildingSiteCtlI) *Builder_wine {
	self := new(Builder_wine)

	self.Builder_std = NewBuilder_std(bs)

	//	self.EditActionsCB = self.EditActions

	//	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

type Builder_wine_wow64 struct {
}

type Builder_wine struct {
	*Builder_std
}
