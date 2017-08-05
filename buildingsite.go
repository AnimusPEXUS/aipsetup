package aipsetup

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	DIR_TARBALL string = "00.TARBALL"
	/*
	   Directory for storing tarballs used in package building. contents is packed
	   into resulting  package as  it is requirements  of most  good licenses
	*/

	DIR_SOURCE string = "01.SOURCE"
	/*
	   Directory for detarred sources, which used for building package. This is not
	   packed  into final  package,  as we  already  have original  tarballs.
	*/

	DIR_PATCHES string = "02.PATCHES"
	/*
	   Patches stored here. packed.
	*/

	DIR_BUILDING string = "03.BUILDING"
	/*
	   Here package are build. not packed.
	*/

	DIR_DESTDIR string = "04.DESTDIR"
	/*
	   Primary root of files for package. those will be installed into target system.
	*/

	DIR_BUILD_LOGS string = "05.BUILD_LOGS"
	/*
	   Various building logs are stored here. Packed.
	*/

	DIR_LISTS string = "06.LISTS"
	/*
	   Various lists stored here. Packed.
	*/

	DIR_TEMP string = "07.TEMP"
	/*
	   Temporary directory used by aipsetup while building package. Throwed away.
	*/

	DIR_ALL []string = []string{
		DIR_TARBALL,
		DIR_SOURCE,
		DIR_PATCHES,
		DIR_BUILDING,
		DIR_DESTDIR,
		DIR_BUILD_LOGS,
		DIR_LISTS,
		DIR_TEMP,
	}
)

type (
	BuildingSitePackageInfo struct {
	}
)

func IsDirRestrictedForWork(path string) bool {
	var err error

	path, err = filepath.Abs(path)

	if err != nil {
		return true
	}

	for _, i := range []string{
		"/bin", "/boot", "/daemons",
		"/dev", "/etc", "/lib", "/proc",
		"/sbin", "/sys",
		"/usr",
	} {
		if path == i || strings.HasPrefix(path, i+"/") {
			return true
		}
	}

	for _, i := range []string{
		"/opt", "/var", "/",
	} {
		if path == i {
			return true
		}
	}

	return false
}

type BuildingSiteCtl struct {
	Path string
}

func NewBuildingSiteCtl(path string) (*BuildingSiteCtl, error) {

	{
		var err error

		path, err = filepath.Abs(path)

		if err != nil {
			return nil, err
		}
	}

	if IsDirRestrictedForWork(path) {
		return nil, errors.New("dir is restricted " + path)
	}

	ret := new(BuildingSiteCtl)

	ret.Path = path

	return ret, nil
}

func (self *BuildingSiteCtl) Init() error {
	fmt.Println("Going to initiate directory", self.Path)
	for _, i := range DIR_ALL {
		j := filepath.Join(self.Path, i)
		f, err := os.Open(j)
		if err != nil {
			err := os.MkdirAll(j, 0700)
			if err != nil {
				return errors.New("Can't create dir " + err.Error())
			}
		} else {
			f.Close()
		}
	}
	return nil
}
