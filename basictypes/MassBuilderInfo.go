package basictypes

type MassBuilderInfo struct {
	Host      string   `json:"host"`
	HostArchs []string `json:"hostarchs"`

	BuildCrossbuilders bool   `json:"build_crossbuilders"`
	CrossbuilderTarget string `json:"crossbuilder_target"`

	BuildCrossbuildings bool   `json:"build_crossbuildings"`
	CrossbuildersHost   string `json:"crossbuilders_host"`

	InitiatedByHost string `json:"initiated_by_host"`
}
