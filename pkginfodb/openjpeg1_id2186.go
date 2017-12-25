package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_openjpeg1 = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://github.com/uclouvain/openjpeg",

	BuilderName: "openjpeg",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"github_hosted", "sf_hosted:openjpeg.mirror"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "openjpeg",
	TarballFileNameParser: "std",
	TarballProvider:       "srs",
	TarballProviderArguments: []string{
		`git`, `https://github.com/uclouvain/openjpeg.git`, `openjpeg2`, `TagPrefixRegExp:version.`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "personal",
	TarballProviderVersionSyncDepth: 3,
}
