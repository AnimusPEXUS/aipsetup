package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_mess_desktop_entries = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://gnome.org/",

	TarballFileNameParser: "std",
	TarballName:           "mess-desktop-entries",
	Filters:               []string{},

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
		"gnome_project", "group:gnome"},

	TarballVersionTool: "gnome",

	TarballProvider: "gnome",
	TarballProviderArguments: []string{
		"mess-desktop-entries"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "gnome",
	TarballProviderVersionSyncDepth: 0,
}
