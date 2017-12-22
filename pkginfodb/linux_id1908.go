package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_linux = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "http://www.kernel.org/",

	BuilderName: "linux",

	Removable:          false,
	Reducible:          false,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"group:cross", "kernelorg_hosted"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "linux",
	TarballFileNameParser: "std",
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		"https://cdn.kernel.org/pub/linux/kernel/"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "by_https_host",
	TarballProviderVersionSyncDepth: 0,
}
