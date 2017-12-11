package basictypes

import "github.com/AnimusPEXUS/utils/textlist"

type PackageInfo struct {
	// Name string `json:"name"`

	Description string //`json:"description"`
	HomePage    string //`json:"homepage"`

	TarballFileNameParser string //`json:"tarball_filename_parser"`
	TarballName           string //`json:"tarball_name"`
	Filters               textlist.Filters

	BuilderName string //`json:"builder_name"`

	Removable          bool //`json:"removable"`
	Reducible          bool //`json:"reducible"`
	NonInstallable     bool //`json:"non_installable"`
	Deprecated         bool //`json:"deprecated"`
	PrimaryInstallOnly bool //`json:"primary_install_only"`

	BuildDeps   []string //`json:"build_deps"`
	SODeps      []string //`json:"so_deps"`
	RunTimeDeps []string //`json:"runtime_deps"`

	Tags []string //`json:"tags"`

	TarballVersionTool string //`json:"tarball_version_tool"`

	TarballProvider string //`json:"tarball_provider"`

	// NOTE: some providers, like sf.net, requires to know additional data to
	//       work with. for instance sf.net needs a project name, as one may
	//       spawn multiple tarball names. so, for such additional data, this
	//       field is provided. Providers, on theyr part, should describe
	//       arguments which they require, using for this they'r source
	//       files.
	TarballProviderArguments []string // `json:"tarball_provider_arguments"`

	TarballProviderUseCache        bool
	TarballProviderCachePresetName string
}
