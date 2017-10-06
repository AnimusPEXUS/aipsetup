package aipsetup

import (
	"errors"
	"fmt"
	"os"
	"path"
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
	BuildingSiteInfo struct {
		Constitution *BuildingSiteConstitution `json:"constitution"`
		PkgInfo      *BuildingSitePackageInfo  `json:"pkg_info"`
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

func (self *BuildingSiteCtl) GetDIR_TARBALL() string {
	return GetDIR_TARBALL(self.Path)
}

func (self *BuildingSiteCtl) GetDIR_SOURCE() string {
	return GetDIR_SOURCE(self.Path)
}

func (self *BuildingSiteCtl) GetDIR_PATCHES() string {
	return GetDIR_PATCHES(self.Path)
}

func (self *BuildingSiteCtl) GetDIR_BUILDING() string {
	return GetDIR_BUILDING(self.Path)
}

func (self *BuildingSiteCtl) GetDIR_DESTDIR() string {
	return GetDIR_DESTDIR(self.Path)
}

func (self *BuildingSiteCtl) GetDIR_BUILD_LOGS() string {
	return GetDIR_BUILD_LOGS(self.Path)
}

func (self *BuildingSiteCtl) GetDIR_LISTS() string {
	return GetDIR_LISTS(self.Path)
}

func (self *BuildingSiteCtl) GetDIR_TEMP() string {
	return GetDIR_TEMP(self.Path)
}

func (self *BuildingSiteCtl) IsWdDirRestricted() bool {
	return IsWdDirRestricted(self.Path)
}

func (self *BuildingSiteCtl) IsDirRestrictedForWork() bool {
	return IsDirRestrictedForWork(self.Path)
}

func getDIR_x(pth string, x string) string {
	res, err := filepath.Abs(path.Join(pth, x))
	if err != nil {
		panic(
			"TODO " +
				"(if you reading this," +
				" write a bugreport) with name 'getDIR_x Abs error'",
		)
	}
	return res
}

func GetDIR_TARBALL(pth string) string {
	return getDIR_x(pth, DIR_TARBALL)
}

func GetDIR_SOURCE(pth string) string {
	return getDIR_x(pth, DIR_SOURCE)
}

func GetDIR_PATCHES(pth string) string {
	return getDIR_x(pth, DIR_PATCHES)
}

func GetDIR_BUILDING(pth string) string {
	return getDIR_x(pth, DIR_BUILDING)
}

func GetDIR_DESTDIR(pth string) string {
	return getDIR_x(pth, DIR_DESTDIR)
}

func GetDIR_BUILD_LOGS(pth string) string {
	return getDIR_x(pth, DIR_BUILD_LOGS)
}

func GetDIR_LISTS(pth string) string {
	return getDIR_x(pth, DIR_LISTS)
}

func GetDIR_TEMP(pth string) string {
	return getDIR_x(pth, DIR_TEMP)
}

func IsWdDirRestricted(pth string) bool {
	panic("use IsDirRestrictedForWork() instead")
	return true
}

func ReadConfig() (*BuildingSiteInfo, error) {
	return nil, nil
}
