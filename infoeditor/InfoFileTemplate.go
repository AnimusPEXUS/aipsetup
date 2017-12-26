package infoeditor

var InfoFileTemplate = `package pkginfodb

// WARNING: Some parts of this may be generated automatically using infoeditor.
//          Be mindfull and make automatic parts changes to infoeditor,
//          compile and use "info-db code" cmd for regenerating.

import (
  "github.com/AnimusPEXUS/aipsetup/basictypes"
)

var %s = &basictypes.PackageInfo{

  Description: %s,
  HomePage   : %s,

  BuilderName  : %s,

  Removable           : %t,
  Reducible           : %t,
  NonInstallable      : %t,
  Deprecated          : %t,
  PrimaryInstallOnly  : %t,

  BuildDeps    : %s,
  SODeps       : %s,
  RunTimeDeps  : %s,

  Tags  : %s,

  TarballVersionTool : %s,

  TarballName                     : %s,
  TarballFileNameParser           : %s,
  TarballFilters                  : %s,
  TarballProvider                 : %s,
  TarballProviderArguments        : %s,
  TarballProviderUseCache         : %t,
  TarballProviderCachePresetName  : %s,
  TarballProviderVersionSyncDepth : %d,

}

`
