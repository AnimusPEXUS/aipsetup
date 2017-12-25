package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_ibus = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://github.com/ibus/ibus",

	BuilderName: "ibus",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"github_hosted"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "ibus",
	TarballFileNameParser: "std",
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/ibus/ibus.git`, `ibus`, `TagPrefixRegExp:^$`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
