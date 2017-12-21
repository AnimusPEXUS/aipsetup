#!/usr/bin/python3

import os.path
import json
import shutil

def boolstr(value):
    if value:
        return "true"
    else:
        return "false"

def normname(value):
    ret = value.replace("-","_")
    ret = ret.replace("+","plus")
    ret = ret.replace(".","_")
    ret = ret.replace(".","_")
    ret = ret.replace("@","_")
    return ret

src_dir = os.path.join(".","pkg_info")
tgt_dir = os.path.join(".","go_convert_result")

shutil.rmtree(tgt_dir)
os.makedirs(tgt_dir, exist_ok=True)

lst=sorted(os.listdir(src_dir))


index_file = open(os.path.join(tgt_dir, "0000000000000000.go"), "w")
index_file.write('''package pkginfodb

import "github.com/AnimusPEXUS/aipsetup"

var Index = map[string]*aipsetup.PackageInfo{
'''
)

counter = 0

for i in lst :
    if i.endswith(".json"):

        counter += 1

        pkgname = i[:-5]
        pkg_name_norm=normname(pkgname)

        jf = open(os.path.join(src_dir, i))
        jt = jf.read()
        jf.close()

        parsed = json.loads(jt)

        f=open(
            os.path.join(
                tgt_dir,
                # "{}.go".format(pkg_name_norm)
                "{n}_auto{c:05}.go".format(
                    n=pkg_name_norm,
                    c=counter
                    )
                ),
            "w"
            )
        f.write('''package pkginfodb

import (
  "github.com/AnimusPEXUS/aipsetup"
  // "github.com/AnimusPEXUS/aipsetup/buildercollection"
  // "github.com/AnimusPEXUS/aipsetup/versiontools"
  )

var DistroPackageInfo_{pkgname} = &aipsetup.PackageInfo{{
    Description: `{pkg_descript}`,
    HomePage: "{pkg_homepage}",

    Removable: {pkg_removable},
    Reducible: {pkg_reducible},
    NonInstallable: {pkg_non_installable},
    Deprecated: {pkg_deprecated},
    PrimaryInstallOnly: {pkg_only_primary},

    BuildDeps   : []string{{}},
    SODeps      : []string{{}},
    RunTimeDeps : []string{{}},

    TarballName : "{pkg_basename}",
    TarballVersionTool: "std", //{pkg_versiontool},

    BuilderName : "std", //buildercollection.Builder_{builder_name},

}}


'''.format(
    pkgname=pkg_name_norm,
    pkg_descript=parsed['description'],
    pkg_homepage=parsed['home_page'],
    pkg_versiontool="versiontools.Standard",
    pkg_removable=boolstr(parsed['removable']),
    pkg_reducible=boolstr(parsed['reducible']),
    pkg_non_installable=boolstr(parsed['non_installable']),
    pkg_deprecated=boolstr(parsed['deprecated']),
    pkg_only_primary=boolstr(parsed['only_primary_install']),
    pkg_build_deps="",
    pkg_so_deps="",
    pkg_runtime_deps="",
    pkg_basename=parsed['basename'],
    builder_name=parsed['buildscript']
)
        )
        f.close()

        index_file.write(
        "  \"{pkgname}\": DistroPackageInfo_{pkg_name_norm},\n".format(
            pkgname=pkgname,
            pkg_name_norm=pkg_name_norm
            )
        )
index_file.write("}\n")
index_file.close()
