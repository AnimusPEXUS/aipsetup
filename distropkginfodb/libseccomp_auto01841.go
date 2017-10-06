package distropkginfodb

import (
  "github.com/AnimusPEXUS/aipsetup"
  // "github.com/AnimusPEXUS/aipsetup/buildercollection"
  // "github.com/AnimusPEXUS/aipsetup/versiontools"
  )

var DistroPackageInfo_libseccomp = &aipsetup.CompletePackageInfo{
  OveralPackageInfo: aipsetup.OveralPackageInfo{
    Description: `The libseccomp library provides and easy to use, platform independent, interface to the Linux Kernel's syscall filtering mechanism: seccomp
`,
    HomePage: "http://sourceforge.net/projects/libseccomp/",

    Removable: true,
    Reducible: true,
    NonInstallable: false,
    Deprecated: false,
    PrimaryInstallOnly: false,

    BuildDeps   : []string{},
    SODeps      : []string{},
    RunTimeDeps : []string{},
  },

  TarballPackageInfo: aipsetup.TarballPackageInfo{
    Name : "libseccomp",
    VersionTool: "std", //versiontools.Standard,
  },

  BuildingPackageInfo: aipsetup.BuildingPackageInfo{
    BuilderName : "std", //buildercollection.Builder_libseccomp,
  },


}


