package distropkginfodb

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/textlist"
	// "github.com/AnimusPEXUS/aipsetup/buildercollection"
	// "github.com/AnimusPEXUS/aipsetup/versiontools"
)

var DistroPackageInfo_autoconf = &basictypes.PackageInfo{
	Description: `write something here, please`,
	HomePage:    "http://www.gnu.org",

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
			"- version-== 2.13",
		},
	),

	TarballVersionTool: "std", //versiontools.Standard,

	BuilderName: "std", //buildercollection.Builder_std,

}
