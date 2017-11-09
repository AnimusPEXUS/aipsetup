package basictypes

import "github.com/AnimusPEXUS/gologger"

type BuilderActions map[string](func(log *gologger.Logger) error)

type BuilderI interface {
	SetBuildingSite(building_site BuildingSiteCtlI)
	DefineActions() ([]string, BuilderActions)
}
