package distropkginfodb

import (
  "github.com/AnimusPEXUS/aipsetup"
  // "github.com/AnimusPEXUS/aipsetup/buildercollection"
  // "github.com/AnimusPEXUS/aipsetup/versiontools"
  )

var DistroPackageInfo_jasper = &aipsetup.CompletePackageInfo{
  OveralPackageInfo: aipsetup.OveralPackageInfo{
    Description: ``,
    HomePage: "http://www.ece.uvic.ca",

    Removable: true,
    Reducible: true,
    NonInstallable: false,
    Deprecated: true,
    PrimaryInstallOnly: false,

    BuildDeps   : []string{},
    SODeps      : []string{},
    RunTimeDeps : []string{},
  },

  TarballPackageInfo: aipsetup.TarballPackageInfo{
    Name : "jasper",
    VersionTool: "std", //versiontools.Standard,
  },

  BuildingPackageInfo: aipsetup.BuildingPackageInfo{
    BuilderName : "std", //buildercollection.Builder_std,
  },


}


