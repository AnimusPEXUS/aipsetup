package aipsetup

type BuildingSiteConstitution struct {
	Host             string   `json:"host"`
	Arch             string   `json:"arch"`
	Build            string   `json:"build"`
	Target           string   `json:"target"`
	MultilibVariants []string `json:"multilib_variants"`
	SystemTitle      string   `json:"system_title"`
	SystemVersion    string   `json:"system_version"`

	// "build": "x86_64-pc-linux-gnu",
	// "host": "x86_64-pc-linux-gnu",
	// "multilib_variants": [
	//   "m64"
	// ],
	// "system_title": "LAILALO",
	// "system_version": "4.0",
	// "target": "x86_64-pc-linux-gnu"

}

func NewBuildingSiteConstitution() *BuildingSiteConstitution {
	ret := new(BuildingSiteConstitution)
	return ret
}

func (self *BuildingSiteConstitution) SetInfoLailalo50() {
	self.SystemTitle = "LAILALO"
	self.SystemVersion = "5.0"
}

type BuildingSitePackageInfo struct {
	Name           string `json:"name"`
	VersionAdopted []uint `json:"version_adopted"`
	Vendor         string `json:"vendor"`
	TarballName    string `json:"tarball_name"`
	TarballVersion string `json:"tarball_version"`
}
