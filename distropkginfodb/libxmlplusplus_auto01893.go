package distropkginfodb

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	// "github.com/AnimusPEXUS/aipsetup/buildercollection"
	// "github.com/AnimusPEXUS/aipsetup/versiontools"
)

var DistroPackageInfo_libxmlplusplus = &basictypes.PackageInfo{
	Description: `WARNING

as libxml++-2.36.0.tar.xz is too complicated for parsing, it need to be renamed to libxmlplusplus-2.36.0.tar.xz in order to properly work with aipsetup`,
	HomePage: "http://www.gnome.org",

	Removable:          true,
	Reducible:          true,
	NonInstallable:     false,
	Deprecated:         false,
	PrimaryInstallOnly: false,

	BuildDeps:   []string{},
	SODeps:      []string{},
	RunTimeDeps: []string{},

	TarballFileNameParser: "std",
	TarballName:           "libxmlplusplus",

	TarballVersionTool: "std", //versiontools.Standard,

	BuilderName: "std", //buildercollection.Builder_std,

}
