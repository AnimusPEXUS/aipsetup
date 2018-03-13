package aipsetup

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildercollection"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
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

const PACKAGE_INFO_FILENAME_V5 = "package_info_v5.json"
const PACKAGE_CHECKSUM_FILENAME = "package.sha512"

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

// forced compile-time type checking: BuildingSiteCtl must fullfill
// BuildingSiteCtlI
var _ basictypes.BuildingSiteCtlI = &BuildingSiteCtl{}

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

func (self *BuildingSiteCtl) ReadInfo() (*basictypes.BuildingSiteInfo, error) {
	fullpath := path.Join(self.Path, PACKAGE_INFO_FILENAME_V5)

	res, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return nil, err
	}

	j_res := new(basictypes.BuildingSiteInfo)

	err = json.Unmarshal(res, j_res)
	if err != nil {
		return nil, err
	}

	return j_res, nil
}

func (self *BuildingSiteCtl) WriteInfo(info *basictypes.BuildingSiteInfo) error {
	fullpath := path.Join(self.Path, PACKAGE_INFO_FILENAME_V5)

	res, err := json.Marshal(info)
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)

	err = json.Indent(b, res, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fullpath, b.Bytes(), 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *BuildingSiteCtl) Init() error {
	fmt.Println("Going to initiate directory", self.Path)
	for _, i := range DIR_ALL {
		j := filepath.Join(self.Path, i)
		f, err := os.Open(j)
		if err != nil {
			err := os.MkdirAll(j, 0700)
			if err != nil {
				return errors.New("Can't create dir: " + err.Error())
			}
		} else {
			f.Close()
		}
	}
	return nil
}

func (self *BuildingSiteCtl) ApplyInitialInfo(
	pkgname string,
	info *basictypes.PackageInfo,
) error {

	b_info := new(basictypes.BuildingSiteInfo)
	b_info.SetInfoLailalo50()

	// b_info.MainTarballInfo = info
	b_info.PackageName = pkgname

	err := self.WriteInfo(b_info)

	return err
}

func (self *BuildingSiteCtl) ApplyHostArchBuildTarget(
	host, arch, build, target string,
) error {
	i, err := self.ReadInfo()
	if err != nil {
		return err
	}

	i.Host = host
	i.Arch = arch
	i.Build = build
	i.Target = target

	err = self.WriteInfo(i)
	if err != nil {
		return err
	}

	return nil
}

func (self *BuildingSiteCtl) CopyInTarballs(filelist []string) error {

	read_buffer := make([]byte, 2*1024*1024)

	for _, i := range filelist {

		b := path.Base(i)

		new_full_path := path.Join(self.GetDIR_TARBALL(), b)

		in_f, err := os.Open(i)
		if err != nil {
			return err
		}
		defer in_f.Close()

		of_f, err := os.Create(new_full_path)
		if err != nil {
			return err
		}
		defer of_f.Close()

		for {

			count, err := in_f.Read(read_buffer)
			if err != nil {
				if err != nil {
					if err == io.EOF {
						break
					} else {
						return err
					}
				}
			}

			_, err = of_f.Write(read_buffer[:count])
			if err != nil {
				return err
			}

		}

		stat, err := os.Stat(i)
		if err != nil {
			return err
		}

		os.Chtimes(new_full_path, stat.ModTime(), stat.ModTime())

	}

	return nil

}

func (self *BuildingSiteCtl) ApplyTarballs(maintarball string) error {

	maintarball = path.Base(maintarball)

	lst, err := ioutil.ReadDir(self.GetDIR_TARBALL())
	if err != nil {
		return err
	}

	for _, i := range lst {
		if maintarball == path.Base(i.Name()) {
			goto maintarball_found
		}
	}

	return errors.New("specified main tarball not found in tarball dir")

maintarball_found:

	filelist := make([]string, 0)

	filelist = append(filelist, maintarball)

	for _, i := range lst {
		b := path.Base(i.Name())
		found := false

		for _, j := range filelist {
			if j == b {
				found = true
				break
			}
		}

		if !found {
			filelist = append(filelist, b)
		}
	}

	info, err := self.ReadInfo()
	if err != nil {
		return err
	}

	tarball_info, err := pkginfodb.Get(info.PackageName)
	if err != nil {
		return err
	}

	info.Sources = filelist

	{

		parser, err := tarballnameparsers.Get(tarball_info.TarballFileNameParser)
		if err != nil {
			return err
		}

		err = tarballname.IsPossibleTarballNameErr(filelist[0])
		if err != nil {
			return err
		}

		parsed, err := parser.Parse(filelist[0])
		if err != nil {
			return err
		}

		info.PackageVersion = parsed.Version.Str
		if parsed.Status.Str != "" {
			info.PackageStatus = parsed.Status.Str
		}

	}

	err = self.WriteInfo(info)
	if err != nil {
		return err
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

func (self *BuildingSiteCtl) IsBuildingSite() bool {
	pkg_file := path.Join(self.Path, PACKAGE_INFO_FILENAME_V5)
	_, err := os.Stat(pkg_file)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (self *BuildingSiteCtl) Run(targets []string) error {

	l, err := self.CreateLogger("aipsetup bs run log", true)
	if err != nil {
		return err
	}
	defer l.Close()

	info, err := self.ReadInfo()
	if err != nil {
		return err
	}

	tarball_info, err := pkginfodb.Get(info.PackageName)
	if err != nil {
		return err
	}

	builder_name := tarball_info.BuilderName

	builder, err := buildercollection.Get(builder_name, self)
	if err != nil {
		return err
	}

	actions, err := builder.DefineActions()
	if err != nil {
		return err
	}

	{
		actions_missing := false
		for _, i := range targets {
			if _, ok := actions.Get(i); !ok {
				actions_missing = true
				fmt.Println("missing requested target actions:", i)
			}
		}
		if actions_missing {
			return errors.New(
				"some requested targets not found in builder's action list",
			)
		}
	}

main_loop:
	for _, i := range actions {
		for _, j := range targets {
			if i.Name == j {
				l.Info(fmt.Sprintf("---====//////[start %s]\\\\\\\\\\\\====---", j))
				lo, err := self.CreateLogger(j, true)
				if err != nil {
					return err
				}
				// lo.Info(j + " log started")
				err = i.Callable(lo)
				// lo.Info(j + " log ended")
				lo.Close()
				if err != nil {
					return err
				}
				l.Info(fmt.Sprintf("---====\\\\\\\\\\\\[ending %s]//////====---", j))
				continue main_loop
			}
		}
	}

	return nil
}

func (self *BuildingSiteCtl) CreateLogger(name string, console_output bool) (
	*logger.Logger,
	error,
) {
	ret := logger.New()

	t := time.Now().UTC()

	pathname := path.Join(
		self.GetDIR_BUILD_LOGS(),
		fmt.Sprintf(
			"%s %s.txt",
			t.Format(time.RFC3339Nano),
			name,
		),
	)

	f, err := os.Create(pathname)
	if err != nil {
		return nil, err
	}
	ret.AddOutputOpt(
		f,
		&logger.OutputOptions{
			TextIcon:       "",
			InfoIcon:       "[i]",
			WarningIcon:    "[w]",
			ErrorIcon:      "[e]",
			InsertTime:     true,
			TimeLayout:     time.RFC3339Nano,
			ClosedByLogger: true,
		},
	)

	if console_output {
		ret.AddOutputOpt(
			os.Stdout,
			&logger.OutputOptions{
				TextIcon:       "",
				InfoIcon:       "[i]",
				WarningIcon:    "[w]",
				ErrorIcon:      "[e]",
				InsertTime:     true,
				TimeLayout:     time.RFC3339Nano,
				ClosedByLogger: false,
			},
		)
	}

	return ret, nil
}

func (self *BuildingSiteCtl) ListActions() ([]string, error) {

	info, err := self.ReadInfo()
	if err != nil {
		return []string{}, err
	}

	tarball_info, err := pkginfodb.Get(info.PackageName)
	if err != nil {
		return []string{}, err
	}

	b, ok := buildercollection.Index[tarball_info.BuilderName]
	if !ok {
		return []string{}, errors.New("requested builder not found")
	}

	builder, err := b(self)
	if err != nil {
		return []string{}, err
	}

	res, err := builder.DefineActions()
	if err != nil {
		return []string{}, err
	}

	ret := res.ActionList()

	return ret, nil
}

func (self *BuildingSiteCtl) GetConfiguredHost() (string, error) {
	i, err := self.ReadInfo()
	if err != nil {
		return "", err
	}

	return i.Host, nil
}

func (self *BuildingSiteCtl) GetConfiguredArch() (string, error) {
	i, err := self.ReadInfo()
	if err != nil {
		return "", err
	}

	return i.Arch, nil
}

func (self *BuildingSiteCtl) GetConfiguredBuild() (string, error) {
	i, err := self.ReadInfo()
	if err != nil {
		return "", err
	}

	return i.Build, nil
}

func (self *BuildingSiteCtl) GetConfiguredTarget() (string, error) {
	i, err := self.ReadInfo()
	if err != nil {
		return "", err
	}

	return i.Target, nil
}

func (self *BuildingSiteCtl) GetConfiguredHABT() (string, string, string, string, error) {
	i, err := self.ReadInfo()
	if err != nil {
		return "", "", "", "", err
	}

	return i.Host, i.Arch, i.Build, i.Target, nil
}

func (self *BuildingSiteCtl) ValuesCalculator() basictypes.ValuesCalculatorI {
	return NewValuesCalculator(self)
}

// func (self *BuildingSiteCtl) Packager() *Packager {
// 	return NewPackager(self)
// }

func (self *BuildingSiteCtl) Packager() basictypes.PackagerI {
	return NewPackager(self)
}

func (self *BuildingSiteCtl) PrePackager() basictypes.PrePackagerI {
	return NewPrePackager(self)
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
