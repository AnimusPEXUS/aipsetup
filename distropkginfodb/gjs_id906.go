package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_gjs = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://gnome.org/",

	TarballFileNameParser: "std",
	TarballName:           "gjs",
	Filters:               []string{},

	BuilderName: "gjs",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"gnome_project"},

	TarballVersionTool: "gnome",

	TarballProvider: "gnome",
	TarballProviderArguments: []string{
		"gjs"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "gnome",
	TarballProviderVersionSyncDepth: 0,
}
