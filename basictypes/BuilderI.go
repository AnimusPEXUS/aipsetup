package basictypes

import "github.com/AnimusPEXUS/utils/logger"

type BuilderAction struct {
	Name     string
	Callable func(log *logger.Logger) error
}

type BuilderActions []*BuilderAction

func (self BuilderActions) Get(name string) (*BuilderAction, bool) {
	for _, i := range self {
		if i.Name == name {
			return i, true
		}
	}
	return nil, false
}

func (self BuilderActions) ActionList() []string {
	ret := make([]string, 0)
	for _, i := range self {
		ret = append(ret, i.Name)
	}
	return ret
}

type BuilderI interface {
	//SetBuildingSite(building_site BuildingSiteCtlI)
	DefineActions() (BuilderActions, error)
}
