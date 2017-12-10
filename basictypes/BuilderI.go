package basictypes

import "github.com/AnimusPEXUS/utils/logger"

type BuilderActions map[string](func(log *logger.Logger) error)

type BuilderI interface {
	//SetBuildingSite(building_site BuildingSiteCtlI)
	DefineActions() ([]string, BuilderActions)
}
