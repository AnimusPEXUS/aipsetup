package distropkginfodb

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/textlist"
)

var DistroPackageInfo_autoconf2_13 = &basictypes.PackageInfo{
	Description: ``,
	HomePage:    "",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	TarballFileNameParser: "std",
	TarballName:           "autoconf",

	Filters: textlist.ParseFilterTextLinesMust(
		[]string{
			"- version-!= 2.13",
		},
	),

	TarballVersionTool: "std", //versiontools.Standard,

	BuilderName: "std", //buildercollection.Builder_std,

}
