package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_json_c = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://github.com/json-c/json-c",

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
		"github_hosted", "group:core1"},

	TarballVersionTool: "std",

	TarballName:           "json-c",
	TarballFileNameParser: "std",
	TarballFilters:        []string{},
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/json-c/json-c.git`, `json-c`, `TagName:json-c`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
