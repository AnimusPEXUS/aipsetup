package distropkginfodb

import (
  "github.com/AnimusPEXUS/aipsetup"
  // "github.com/AnimusPEXUS/aipsetup/buildercollection"
  // "github.com/AnimusPEXUS/aipsetup/versiontools"
  )

var DistroPackageInfo_gr_usrp = &aipsetup.CompletePackageInfo{
  OveralPackageInfo: aipsetup.OveralPackageInfo{
    Description: ``,
    HomePage: "",

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
    Name : "gr-usrp",
    VersionTool: "std", //versiontools.Standard,
  },

  BuildingPackageInfo: aipsetup.BuildingPackageInfo{
    BuilderName : "std", //buildercollection.Builder_std,
  },


}


