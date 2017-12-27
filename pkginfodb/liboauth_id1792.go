package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_liboauth = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://sourceforge.net/projects/liboauth",

	BuilderName: "std",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"sf_hosted:liboauth"},

	TarballVersionTool: "std",

	TarballName:           "liboauth",
	TarballFileNameParser: "std",
	TarballFilters:        []string{},
	TarballProvider:       "sf",
	TarballProviderArguments: []string{
		`liboauth`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
