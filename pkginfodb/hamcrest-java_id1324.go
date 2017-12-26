package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_hamcrest_java = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://github.com/hamcrest/JavaHamcrest",

	BuilderName: "hamcrest_java",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"github_hosted", "group:core1"},

	TarballVersionTool: "std",

	TarballFilters:               []string{},
	TarballName:           "hamcrest-java",
	TarballFileNameParser: "std",
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/hamcrest/JavaHamcrest.git`, `hamcrest-java`, `TagPrefixRegExp:hamcrest-java`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
