package basictypes

type BuilderI interface {
	//SetBuildingSite(building_site BuildingSiteCtlI)
	DefineActions() (BuilderActions, error)
	GetBuildingSiteCtl() BuildingSiteCtlI
}
