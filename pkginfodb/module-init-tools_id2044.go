package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_module_init_tools = &basictypes.PackageInfo{

	Description: `use kmod`,
	HomePage:    "http://www.kernel.org",

	BuilderName: "moduleinittools",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         true,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	Tags: []string{},

	TarballVersionTool: "std",

	Filters:                         []string{},
	TarballName:                     "module-init-tools",
	TarballFileNameParser:           "std",
	TarballProvider:                 "",
	TarballProviderArguments:        []string{},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "",
	TarballProviderVersionSyncDepth: 0,
}
