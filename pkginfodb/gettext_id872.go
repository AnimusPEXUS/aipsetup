package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_gettext = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://www.gnu.org/software/gettext",

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
		"gnu_project", "group:core0", "sf_hosted:warzone2100"},

	TarballVersionTool: "std",

	TarballName:           "gettext",
	TarballFileNameParser: "std",
	TarballFilters:        []string{},
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		`https://ftp.gnu.org/gnu/gettext`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "by_https_host",
	TarballProviderVersionSyncDepth: 0,
}
