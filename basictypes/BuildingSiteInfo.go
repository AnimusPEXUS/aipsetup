package basictypes

type BuildingSiteInfo struct {
	SystemTitle   string `json:"system_title"`
	SystemVersion string `json:"system_version"`

	Host     string `json:"host"`
	HostArch string `json:"hostarch"`
	// Target   string `json:"target"`

	PackageName      string `json:"package_name"`
	PackageVersion   string `json:"package_version"`
	PackageStatus    string `json:"package_status"`
	PackageTimestamp string `json:"package_timestamp"`

	Sources []string `json:"sources"`

	// MainTarballInfo *PackageInfo `json:"main_tarball_info"`
}

func (self *BuildingSiteInfo) SetInfoLailalo50() {
	self.SystemTitle = "LAILALO"
	self.SystemVersion = "5.0"
}
