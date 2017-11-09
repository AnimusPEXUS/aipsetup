package basictypes

type BuildingSiteInfo struct {
	SystemTitle   string `json:"system_title"`
	SystemVersion string `json:"system_version"`

	Host   string `json:"host"`
	Arch   string `json:"arch"`
	Build  string `json:"build"`
	Target string `json:"target"`

	PackageName string `json:"package_name"`

	Sources []string `json:"sources"`

	MainTarballInfo *PackageInfo `json:"main_tarball_info"`
}

func (self *BuildingSiteInfo) SetInfoLailalo50() {
	self.SystemTitle = "LAILALO"
	self.SystemVersion = "5.0"
}
