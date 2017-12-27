package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_fuse = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://github.com/libfuse/libfuse",

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
		"github_hosted", "group:core0"},

	TarballVersionTool: "std",

	TarballName:           "fuse",
	TarballFileNameParser: "std",
	TarballFilters:        []string{},
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/libfuse/libfuse.git`, `fuse`, `TagName:fuse`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
