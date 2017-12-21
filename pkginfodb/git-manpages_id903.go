package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_git_manpages = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "http://git-scm.com/",

	BuilderName: "git_manpages",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"group:core0", "kernelorg_hosted"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "git-manpages",
	TarballFileNameParser: "std",
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		"https://cdn.kernel.org/pub/"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "by_https_host",
	TarballProviderVersionSyncDepth: 0,
}
