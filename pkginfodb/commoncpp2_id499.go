package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_commoncpp2 = &basictypes.PackageInfo{

	Description: `use 'ucommon'`,
	HomePage:    "http://www.gnu.org",

	BuilderName: "std",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     true,
	Deprecated:         true,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{},

	TarballVersionTool: "std",

	TarballFilters:                         []string{},
	TarballName:                     "commoncpp2",
	TarballFileNameParser:           "std",
	TarballProvider:                 "",
	TarballProviderArguments:        []string{},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
