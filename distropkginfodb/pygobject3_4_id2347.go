package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_pygobject3_4 = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://gnome.org/",

	BuilderName: "pygobject",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         true,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"gnome_project"},

	TarballVersionTool: "gnome",

	Filters:               []string{},
	TarballName:           "pygobject3.4",
	TarballFileNameParser: "std",
	TarballProvider:       "gnome",
	TarballProviderArguments: []string{
		"pygobject3.4"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "gnome",
	TarballProviderVersionSyncDepth: 0,
}
