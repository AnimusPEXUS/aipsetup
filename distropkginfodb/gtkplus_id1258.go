package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_gtkplus = &basictypes.PackageInfo{

	Description: `write something here, please`,
	HomePage:    "https://sourceforge.net/projects/xfce",

	TarballFileNameParser: "std",
	TarballName:           "gtk+",
	Filters:               []string{},

	BuilderName: "gtk3",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"gnome_project", "sf_project:xfce"},

	TarballVersionTool: "gnome",

	TarballProvider: "sf",
	TarballProviderArguments: []string{
		"xfce"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "gnome",
	TarballProviderVersionSyncDepth: 0,
}
