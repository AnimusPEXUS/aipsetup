package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_e2fsprogs_libs = &basictypes.PackageInfo{

	Description: `subset of e2fsprogs package`,
	HomePage:    "http://sf.net",

	TarballFileNameParser: "std",
	TarballName:           "e2fsprogs-libs",
	Filters:               []string{},

	BuilderName: "std",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     true,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{},

	TarballVersionTool: "std",

	TarballProvider:                 "",
	TarballProviderArguments:        []string{},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
