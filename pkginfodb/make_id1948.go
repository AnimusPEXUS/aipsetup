package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_make = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "http://www.gnu.org",

	BuilderName: "make",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"group:core0"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "make",
	TarballFileNameParser: "std",
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		"https://ftp.gnu.org/gnu/make"},
	TarballProviderUseCache:         true,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
