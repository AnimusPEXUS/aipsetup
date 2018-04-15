package basictypes

type MassBuilderInfo struct {
	Host      string   `json:"host"`
	HostArchs []string `json:"hostarchs"`

	CrossbuilderTarget string `json:"crossbuilder_target"`

	CrossbuildersHost string `json:"crossbuilders_host"`

	InitiatedByHost string `json:"initiated_by_host"`
}
