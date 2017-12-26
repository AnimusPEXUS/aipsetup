package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_dbus_glib = &basictypes.PackageInfo{

	Description: `dbus-glib is a deprecated API for use of D-Bus from GLib applications. Do not use it in new code.
Since version 2.26, GLib's accompanying GIO library provides a high-level API for D-Bus, "GDBus", based on an independent reimplementation of the D-Bus protocol. The maintainers of D-Bus recommend that GLib applications should use GDBus instead of dbus-glib.

but some applications still using it (NetworkManager)`,
	HomePage: "http://www.freedesktop.org",

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
		"group:core0"},

	TarballVersionTool: "std",

	TarballFilters:                         []string{},
	TarballName:                     "dbus-glib",
	TarballFileNameParser:           "std",
	TarballProvider:                 "",
	TarballProviderArguments:        []string{},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
