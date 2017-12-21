package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_ORBit2 = &basictypes.PackageInfo{

	Description: `ORBit2 is a CORBA 2.4-compliant Object Request Broker (ORB) featuring 
mature C, C++ and Python bindings. Bindings (in various degrees of completeness) are also available for Perl, Lisp, Pascal, Ruby, and TCL; others are in-progress. It supports POA, DII, DSI, TypeCode, Any, IR and IIOP. Optional features including INS and threading are available. ORBit2 is engineered for the desktop workstation environment, with a focus on performance, low resource usage, and security. The core ORB is written in C, and runs under Linux, UNIX (BSD, Solaris, HP-UX, ...), and Windows. ORBit2 is developed and released as open source software under GPL/LGPL.

It is supported by Red Hat and Ximian as middleware of the GNOME project.

required by gconf

deprecated! use --disable-orbit for GConf`,
	HomePage: "https://gnome.org/",

	BuilderName: "orbit2",

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
	TarballName:           "ORBit2",
	TarballFileNameParser: "std",
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		"https://ftp.gnome.org/mirror/gnome.org/"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "gnome",
	TarballProviderVersionSyncDepth: 0,
}
