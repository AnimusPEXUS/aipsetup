package distropkginfodb

import (
  "github.com/AnimusPEXUS/aipsetup"
  // "github.com/AnimusPEXUS/aipsetup/buildercollection"
  // "github.com/AnimusPEXUS/aipsetup/versiontools"
  )

var DistroPackageInfo_jpeg = &aipsetup.CompletePackageInfo{
  OveralPackageInfo: aipsetup.OveralPackageInfo{
    Description: `use openjpeg1 openjpeg2 and libjpeg-turbo`,
    HomePage: "http://www.ijg.org",

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
    Name : "jpeg",
    VersionTool: "std", //versiontools.Standard,
  },

  BuildingPackageInfo: aipsetup.BuildingPackageInfo{
    BuilderName : "std", //buildercollection.Builder_None,
  },


}


