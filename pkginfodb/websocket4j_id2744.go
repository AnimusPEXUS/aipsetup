package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var DistroPackageInfo_websocket4j = &basictypes.PackageInfo{

	Description: ``,
	HomePage:    "https://www.gnu.org/software/websocket4j",

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
		"gnu_project"},

	TarballVersionTool: "std",

	Filters:               []string{},
	TarballName:           "websocket4j",
	TarballFileNameParser: "std",
	TarballProvider:       "https",
	TarballProviderArguments: []string{
		`https://ftp.gnu.org/gnu/websocket4j`},
	TarballProviderUseCache:         false,
	TarballProviderCachePresetName:  "by_https_host",
	TarballProviderVersionSyncDepth: 0,
}