#!/usr/bin/python3


import json
import os.path
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

f1 = open("123.json")
f1_j = json.loads(f1.read())
f1.close()

tgt_dir = os.path.join(".","convert_result_001")

if os.path.isdir(tgt_dir):
    shutil.rmtree(tgt_dir)
os.makedirs(tgt_dir, exist_ok=True)



index_file = open(os.path.join(tgt_dir, "0000000000000000.go"), "w")
index_file.write('''package distropkginfodb

import "github.com/AnimusPEXUS/aipsetup"

var Index = map[string]*aipsetup.PackageInfo{
'''
)

counter = 0

for i in sorted(list(f1_j.keys())):

    counter += 1

    pkgname = i
    pkg_name_norm=normname(pkgname)


    parsed = f1_j[i]

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
    f.write('''package distropkginfodb

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

    TarballFilenameParser: "std",
    TarballFilenameParserParameters: map[string]string{{
        "prefix": "{pkg_basename}",
    }},

    TarballVersionTool: "std", //{pkg_versiontool},

    BuilderName : "std", //buildercollection.Builder_{builder_name},

}}


'''.format(
    pkgname=pkg_name_norm,
    pkg_descript=parsed['description'],
    pkg_homepage=parsed['homepage'],
    pkg_versiontool="versiontools.Standard",
    pkg_removable=boolstr(parsed['removable']),
    pkg_reducible=boolstr(parsed['reducible']),
    pkg_non_installable=boolstr(parsed['non_installable']),
    pkg_deprecated=boolstr(parsed['deprecated']),
    pkg_only_primary=boolstr(parsed['primary_install_only']),
    pkg_build_deps="",
    pkg_so_deps="",
    pkg_runtime_deps="",
    pkg_basename=parsed['tarball_name'],
    builder_name=parsed['builder_name']
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
