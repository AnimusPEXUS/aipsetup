package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_cloog = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "",

	TarballFileNameParser: "std",
	TarballName:           "cloog",
	Filters:               []string{},

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

	TarballProvider: "https",
	TarballProviderArguments: []string{
		"https://www.bastoul.net/cloog/pages/download/"},
	TarballProviderUseCache:         true,
	TarballProviderCachePresetName:  "bastoul.net",
	TarballProviderVersionSyncDepth: 0,
}
