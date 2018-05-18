package basictypes

type BuildingSiteInfo struct {
	SystemTitle   string `json:"system_title"`
	SystemVersion string `json:"system_version"`

	// System which going to run this package
	Host     string `json:"host"`
	HostArch string `json:"hostarch"`

	// This package is crosscompiler to build other packages.
	// This package is going to run under Host
	// CrossbuilderTarget is System which going to run packages built by this crosscompiler
	// ThisIsCrossbuilder bool   `json:"this_is_crossbuilder"`
	CrossbuilderTarget string `json:"crossbuilder_target"`

	// This package is crossbuilding
	// This package is being built to run under Host
	// This package uses CrossbuildersHost's crossbuilder to be built
	// ThisIsCrossbuilding bool   `json:"this_is_crossbuilding"`
	CrossbuildersHost string `json:"crossbuilder_s_host"`

	// ThisIsSubarchBuilding bool `json:"this_is_subarchbuilding"`

	// RunningByHost is host which started this building site
	InitiatedByHost string `json:"initiated_by_host"`

	PackageName      string `json:"package_name"`
	PackageVersion   string `json:"package_version"`
	PackageStatus    string `json:"package_status"`
	PackageTimeStamp string `json:"package_timestamp"`

	TarballsDir string `json:"tarballs_dir"`
	AspsDir     string `json:"asps_dir"`

	//	Sources []string `json:"sources"`

	// MainTarballInfo *PackageInfo `json:"main_tarball_info"`

	ModifyVersionBeforePack bool   `json:"modify_version_before_pack"`
	NewVersion              string `json:"new_version"`
}

func (self *BuildingSiteInfo) SetInfoLilith50() {
	self.SystemTitle = "Lilith"
	self.SystemVersion = "5.0"
}

func (self *BuildingSiteInfo) ThisIsCrossbuilder() bool {
	return self.CrossbuilderTarget != ""
}

func (self *BuildingSiteInfo) ThisIsCrossbuilding() bool {
	return self.CrossbuildersHost != ""
}

func (self *BuildingSiteInfo) ThisIsSubarchBuilding() bool {
	return self.Host != self.HostArch
}
