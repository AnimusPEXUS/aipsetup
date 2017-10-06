package distropkginfodb

import (
  "github.com/AnimusPEXUS/aipsetup"
  // "github.com/AnimusPEXUS/aipsetup/buildercollection"
  // "github.com/AnimusPEXUS/aipsetup/versiontools"
  )

var DistroPackageInfo_postgresql = &aipsetup.CompletePackageInfo{
  OveralPackageInfo: aipsetup.OveralPackageInfo{
    Description: `write something here, please`,
    HomePage: "http://www.postgresql.org",

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
    Name : "postgresql",
    VersionTool: "std", //versiontools.Standard,
  },

  BuildingPackageInfo: aipsetup.BuildingPackageInfo{
    BuilderName : "std", //buildercollection.Builder_postgreSQL,
  },


}


