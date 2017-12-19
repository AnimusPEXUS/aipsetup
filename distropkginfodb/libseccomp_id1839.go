package distropkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_libseccomp = &basictypes.PackageInfo{

	Description: `The libseccomp library provides and easy to use, platform independent, interface to the Linux Kernel's syscall filtering mechanism: seccomp
`,
	HomePage: "https://sourceforge.net/projects/libseccomp",

	BuilderName: "libseccomp",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{
		"group:core1", "sf_project:libseccomp"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "libseccomp",
	TarballFileNameParser: "std",
	TarballProvider:       "sf",
	TarballProviderArguments: []string{
		"libseccomp"},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
