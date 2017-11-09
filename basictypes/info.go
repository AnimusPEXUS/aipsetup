package basictypes

import "github.com/AnimusPEXUS/golistfilter"

type PackageInfo struct {
	Description string `json:"description"`
	HomePage    string `json:"homepage"`

	TarballFileNameParser string `json:"tarball_filename_parser"`
	TarballName           string `json:"tarball_name"`
	Filters               golistfilter.Filters

	TarballVersionTool string `json:"tarball_version_tool"`

	BuilderName string `json:"builder_name"`

	Removable          bool `json:"removable"`
	Reducible          bool `json:"reducible"`
	NonInstallable     bool `json:"non_installable"`
	Deprecated         bool `json:"deprecated"`
	PrimaryInstallOnly bool `json:"primary_install_only"`

	BuildDeps   []string `json:"build_deps"`
	SODeps      []string `json:"so_deps"`
	RunTimeDeps []string `json:"runtime_deps"`

	Tags []string `json:"tags"`
}
