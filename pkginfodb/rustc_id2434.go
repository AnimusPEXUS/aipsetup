package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_rustc = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://github.com/rust-lang/rust",

	BuilderName: "rustc",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: true,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"github_hosted"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "rustc",
	TarballFileNameParser: "std",
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/rust-lang/rust.git`, `rustc`, `TagPrefixRegExp:^$`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
