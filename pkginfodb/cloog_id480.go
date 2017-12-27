package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_cloog = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "",

	BuilderName: "cloog",

	Removable:          false,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"group:cross"},

	TarballVersionTool: "std",

	TarballName:           "cloog",
	TarballFileNameParser: "std",
	TarballFilters:        []string{},
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		`https://www.bastoul.net/cloog/pages/download/`},
	TarballProviderUseCache:         true,
	TarballProviderCachePresetName:  "by_https_host",
	TarballProviderVersionSyncDepth: 0,
}
