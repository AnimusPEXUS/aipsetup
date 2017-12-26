package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_xulrunner = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "http://mozilla.org",

	BuilderName: "xulrunner",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{},

	TarballVersionTool: "std",

	TarballFilters:                         []string{},
	TarballName:                     "xulrunner",
	TarballFileNameParser:           "std",
	TarballProvider:                 "",
	TarballProviderArguments:        []string{},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
