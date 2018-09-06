package buildercollection

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["nss"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_nss(bs)
	}
}

type Builder_nss struct {
	*Builder_std
}

func NewBuilder_nss(bs basictypes.BuildingSiteCtlI) (*Builder_nss, error) {

	self := new(Builder_nss)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	//	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditBuildArgsCB = self.EditBuildArgs

	self.EditConfigureDirCB = self.EditConfigureDir
	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir

	self.EditBuildConcurentJobsCountCB = self.EditBuildConcurentJobsCount

	return self, nil
}

func (self *Builder_nss) EditBuildConcurentJobsCount(log *logger.Logger, ret int) int {
	return 1
}

func (self *Builder_nss) EditConfigureDir(log *logger.Logger, ret string) (string, error) {
	return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "nss"), nil
}

func (self *Builder_nss) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
	return self.EditConfigureDir(log, ret)
}

func (self *Builder_nss) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	//	ret = ret.Remove("configure")

	ret = ret.Remove("autogen")

	err := ret.Replace(
		"configure",
		&basictypes.BuilderAction{
			Name:     "configure",
			Callable: self.BuilderActionConfigure,
		},
	)
	if err != nil {
		return nil, err
	}

	err = ret.Replace(
		"distribute",
		&basictypes.BuilderAction{
			Name:     "distribute",
			Callable: self.BuilderActionDistribute,
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_nss) BuilderActionConfigure(log *logger.Logger) error {

	makefile := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "nss", "Makefile")

	makefile_text, err := ioutil.ReadFile(makefile)
	if err != nil {
		return err
	}

	makefile_text = bytes.Replace(
		makefile_text,
		[]byte("nss_build_all: build_nspr all latest"),
		[]byte("nss_build_all: all latest"),
		-1,
	)

	err = ioutil.WriteFile(makefile, makefile_text, 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_nss) EditBuildArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	if variant, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateMultilibVariant(); err != nil {
		return nil, err
	} else {
		switch variant {
		case "32":
		case "64":
			ret = append(ret, "USE_64=1")
		default:
			return nil, errors.New("requested multilib variant is not supported")
		}
	}

	ret = append(
		ret,
		[]string{
			"nss_build_all",
			"BUILD_OPT=1",
			"NSPR_INCLUDE_DIR=" + path.Join(install_prefix, "include", "nspr"),
			"USE_SYSTEM_ZLIB=1",
			"ZLIB_LIBS=-lz",
			"NSS_USE_SYSTEM_SQLITE=1",
		}...,
	)

	return ret, nil
}

func (self *Builder_nss) BuilderActionDistribute(log *logger.Logger) error {

	build_dist_dir := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "dist")

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	ver := strings.Split(info.PackageVersion, ".")

	for len(ver) < 3 {
		ver = append(ver, "0")
	}

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	{

		build_dist_dir_files, err := ioutil.ReadDir(build_dist_dir)
		if err != nil {
			return err
		}

		OBJ_dir := ""
		for _, i := range build_dist_dir_files {
			if b, err := filepath.Match("Linux*.OBJ", i.Name()); err != nil {
				return err
			} else {
				if b {
					OBJ_dir = path.Join(build_dist_dir, i.Name())
					break
				}
			}
		}

		if OBJ_dir == "" {
			return errors.New("couldn't find files for distribution")
		}

		for _, i := range []string{"bin", "lib"} {
			err = filetools.CopyTree(
				path.Join(OBJ_dir, i),
				path.Join(dst_install_prefix, i),
				false,
				false,
				false,
				true,
				log,
				false,
				true,
				func(src, dst string, log logger.LoggerI) error {
					err := filetools.CopyWithOptions(
						src,
						dst,
						log,
						true,
						true,
					)
					if err != nil {
						return err
					}
					return nil
				},
			)
			if err != nil {
				return err
			}
		}
	}

	{

		for _, i := range []string{"public", "private"} {
			for _, j := range []string{"dbm", "nss"} {

				src_dir := path.Join(build_dist_dir, i, j)

				err = filetools.CopyTree(
					src_dir,
					path.Join(dst_install_prefix, "include", "nss"),
					false,
					false,
					true,
					true,
					log,
					false,
					true,
					func(src, dst string, log logger.LoggerI) error {
						err := filetools.CopyWithOptions(
							src,
							dst,
							log,
							true,
							true,
						)
						if err != nil {
							return err
						}
						return nil
					},
				)
				if err != nil {
					return err
				}

			}
		}

	}

	{

		dst_lib := path.Join(dst_install_prefix, "lib")

		dst_lib_files, err := ioutil.ReadDir(dst_lib)
		if err != nil {
			return err
		}

		lib_names := make([]string, 0)

		for _, i := range dst_lib_files {
			i_name := i.Name()
			if strings.HasPrefix(i_name, "lib") && strings.HasSuffix(i_name, ".so") {
				lib_names = append(lib_names, i_name[3:len(i_name)-3])
			}
		}

		dst_pkgconfig_dir := path.Join(dst_lib, "pkgconfig")
		dst_pkgconfig_dir_file := path.Join(dst_pkgconfig_dir, "nss.pc")

		pkg_config_tpl, err := template.New("pkg_config").Parse(`
prefix={{.Prefix}}
exec_prefix={{.Prefix}}
libdir={{.Libdir}}
includedir=${exec_prefix}/include/nss

Name: NSS
Description: Network Security Services
Version: {{.Nss_major_version}}.{{.Nss_minor_version}}.{{.Nss_patch_version}}
Libs: -L${libdir} {{.Libs}}
Cflags: -I${includedir}
`)
		if err != nil {
			return err
		}

		b := &bytes.Buffer{}

		err = pkg_config_tpl.Execute(
			b,
			struct {
				Prefix            string
				Libdir            string
				Nss_major_version string
				Nss_minor_version string
				Nss_patch_version string
				Libs              string
			}{
				Prefix:            install_prefix,
				Libdir:            path.Join(install_prefix, "lib"),
				Nss_major_version: ver[0],
				Nss_minor_version: ver[1],
				Nss_patch_version: ver[2],
				Libs:              "-l" + strings.Join(lib_names, " -l"),
			},
		)
		if err != nil {
			return err
		}

		err = os.MkdirAll(dst_pkgconfig_dir, 0700)
		if err != nil {
			return err
		}

		f, err := os.Create(dst_pkgconfig_dir_file)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = b.WriteTo(f)
		if err != nil {
			return err
		}

	}

	{

		dst_nss_config_dir := path.Join(dst_install_prefix, "bin")
		dst_nss_config_dir_file := path.Join(dst_nss_config_dir, "nss-config")

		nss_config_tpl, err := template.New("pkg_config").Parse(`#!/bin/sh

prefix={{.Prefix}}
exec_prefix={{.Prefix}}

major_version={{.Nss_major_version}}
minor_version={{.Nss_minor_version}}
patch_version={{.Nss_patch_version}}

usage()
{
    cat <<EOF
Usage: nss-config [OPTIONS] [LIBRARIES]
Options:
    [--prefix[=DIR]]
    [--exec-prefix[=DIR]]
    [--includedir[=DIR]]
    [--libdir[=DIR]]
    [--version]
    [--libs]
    [--cflags]
Dynamic Libraries:
    nss
    nssutil
    smime
    ssl
    softokn
EOF
    exit $1
}

if test $# -eq 0; then
    usage 1 1>&2
fi

lib_nss=yes
lib_nssutil=yes
lib_smime=yes
lib_ssl=yes
lib_softokn=yes

while test $# -gt 0; do
  case "$1" in
  -*=*) optarg=` + "`" + `echo "$1" | sed 's/[-_a-zA-Z0-9]*=//'` + "`" + ` ;;
  *) optarg= ;;
  esac

  case $1 in
    --prefix=*)
      prefix=$optarg
      ;;
    --prefix)
      echo_prefix=yes
      ;;
    --exec-prefix=*)
      exec_prefix=$optarg
      ;;
    --exec-prefix)
      echo_exec_prefix=yes
      ;;
    --includedir=*)
      includedir=$optarg
      ;;
    --includedir)
      echo_includedir=yes
      ;;
    --libdir=*)
      libdir=$optarg
      ;;
    --libdir)
      echo_libdir=yes
      ;;
    --version)
      echo ${major_version}.${minor_version}.${patch_version}
      ;;
    --cflags)
      echo_cflags=yes
      ;;
    --libs)
      echo_libs=yes
      ;;
    nss)
      lib_nss=yes
      ;;
    nssutil)
      lib_nssutil=yes
      ;;
    smime)
      lib_smime=yes
      ;;
    ssl)
      lib_ssl=yes
      ;;
    softokn)
      lib_softokn=yes
      ;;
    *)
      usage 1 1>&2
      ;;
  esac
  shift
done

# Set variables that may be dependent upon other variables
if test -z "$exec_prefix"; then
    exec_prefix=` + "`" + `pkg-config --variable=exec_prefix nss` + "`" + `
fi
if test -z "$includedir"; then
    includedir=` + "`" + `pkg-config --variable=includedir nss` + "`" + `
fi
if test -z "$libdir"; then
    libdir=` + "`" + `pkg-config --variable=libdir nss` + "`" + `
fi

if test "$echo_prefix" = "yes"; then
    echo $prefix
fi

if test "$echo_exec_prefix" = "yes"; then
    echo $exec_prefix
fi

if test "$echo_includedir" = "yes"; then
    echo $includedir
fi

if test "$echo_libdir" = "yes"; then
    echo $libdir
fi

if test "$echo_cflags" = "yes"; then
    echo -I$includedir
fi

if test "$echo_libs" = "yes"; then
    libdirs="-L$libdir"

    if test -n "$lib_nss"; then
        libdirs="$libdirs -lnss${major_version}"
    fi

    if test -n "$lib_nssutil"; then
        libdirs="$libdirs -lnssutil${major_version}"
    fi

    if test -n "$lib_smime"; then
        libdirs="$libdirs -lsmime${major_version}"
    fi

    if test -n "$lib_ssl"; then
        libdirs="$libdirs -lssl${major_version}"
    fi

    if test -n "$lib_softokn"; then
        libdirs="$libdirs -lsoftokn${major_version}"
    fi

    echo $libdirs
fi
`)
		if err != nil {
			return err
		}

		b := &bytes.Buffer{}

		err = nss_config_tpl.Execute(
			b,
			struct {
				Prefix            string
				Libdir            string
				Nss_major_version string
				Nss_minor_version string
				Nss_patch_version string
				Libs              string
			}{
				Prefix: install_prefix,
				//				Libdir:            path.Join(install_prefix, "lib"),
				Nss_major_version: ver[0],
				Nss_minor_version: ver[1],
				Nss_patch_version: ver[2],
				//				Libs:              "-l" + strings.Join(lib_names, " -l"),
			},
		)
		if err != nil {
			return err
		}

		err = os.MkdirAll(dst_nss_config_dir, 0700)
		if err != nil {
			return err
		}

		f, err := os.Create(dst_nss_config_dir_file)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = b.WriteTo(f)
		if err != nil {
			return err
		}
	}
	return nil
}
