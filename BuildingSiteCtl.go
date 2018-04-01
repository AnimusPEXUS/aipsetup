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

var _ basictypes.BuildingSiteCtlI = &BuildingSiteCtl{}

type BuildingSiteCtl struct {
	sys  *System
	path string
	info *basictypes.BuildingSiteInfo

	buildingsitevaluescalculator *BuildingSiteValuesCalculator
}

func NewBuildingSiteCtl(
	sys *System,
	path string,
) (*BuildingSiteCtl, error) {

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

	self := new(BuildingSiteCtl)

	self.buildingsitevaluescalculator = NewBuildingSiteValuesCalculator(self)

	self.path = path
	self.sys = sys

	return self, nil
}

func (self *BuildingSiteCtl) ReadInfo() (*basictypes.BuildingSiteInfo, error) {

	if self.info == nil {
		fullpath := path.Join(self.path, basictypes.PACKAGE_INFO_FILENAME_V5)

		res, err := ioutil.ReadFile(fullpath)
		if err != nil {
			return nil, err
		}

		j_res := new(basictypes.BuildingSiteInfo)

		err = json.Unmarshal(res, j_res)
		if err != nil {
			return nil, err
		}
	}

	return self.info, nil
}

func (self *BuildingSiteCtl) WriteInfo(info *basictypes.BuildingSiteInfo) error {

	fullpath := path.Join(self.path, basictypes.PACKAGE_INFO_FILENAME_V5)

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

	self.info = info

	return nil
}

func (self *BuildingSiteCtl) Init() error {
	fmt.Println("Going to initiate directory", self.path)
	for _, i := range basictypes.DIR_ALL {
		j := filepath.Join(self.path, i)
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

func (self *BuildingSiteCtl) ApplyHostHostArch(
	host, hostarch string,
) error {
	i, err := self.ReadInfo()
	if err != nil {
		return err
	}

	i.Host = host
	i.HostArch = hostarch

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

func (self *BuildingSiteCtl) getDIR_x(x string) string {
	return path.Join(self.path, x)
}

func (self *BuildingSiteCtl) GetDIR_TARBALL() string {
	return self.getDIR_x(basictypes.DIR_TARBALL)
}

func (self *BuildingSiteCtl) GetDIR_SOURCE() string {
	return self.getDIR_x(basictypes.DIR_SOURCE)
}

func (self *BuildingSiteCtl) GetDIR_PATCHES() string {
	return self.getDIR_x(basictypes.DIR_PATCHES)
}

func (self *BuildingSiteCtl) GetDIR_BUILDING() string {
	return self.getDIR_x(basictypes.DIR_BUILDING)
}

func (self *BuildingSiteCtl) GetDIR_DESTDIR() string {
	return self.getDIR_x(basictypes.DIR_DESTDIR)
}

func (self *BuildingSiteCtl) GetDIR_BUILD_LOGS() string {
	return self.getDIR_x(basictypes.DIR_BUILD_LOGS)
}

func (self *BuildingSiteCtl) GetDIR_LISTS() string {
	return self.getDIR_x(basictypes.DIR_LISTS)
}

func (self *BuildingSiteCtl) GetDIR_TEMP() string {
	return self.getDIR_x(basictypes.DIR_TEMP)
}

func (self *BuildingSiteCtl) IsDirRestrictedForWork() bool {
	return IsDirRestrictedForWork(self.path)
}

func (self *BuildingSiteCtl) IsBuildingSite() bool {
	pkg_file := path.Join(self.path, basictypes.PACKAGE_INFO_FILENAME_V5)
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
				l.Info(fmt.Sprintf("---====//////[START %s]\\\\\\\\\\\\====---", j))
				lo, err := self.CreateLogger(j, true)
				if err != nil {
					return err
				}
				err = i.Callable(lo)
				lo.Close()
				if err != nil {
					return err
				}
				l.Info(fmt.Sprintf("---====\\\\\\\\\\\\[STOP  %s]//////====---", j))
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

func (self *BuildingSiteCtl) GetSystem() basictypes.SystemI {
	return self.sys
}

func (self *BuildingSiteCtl) GetBuildingSiteValuesCalculator() basictypes.BuildingSiteValuesCalculatorI {
	return self.buildingsitevaluescalculator
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

func (self *BuildingSiteCtl) PrintCalculations() error {
	fmt.Println(
		"Calculating values",
	)

	fmt.Println("-----------------------------------")

	i, err := self.ReadInfo()
	if err != nil {
		fmt.Println("Error reading building site configuration")
		return err
	}

	fmt.Println("Building package", i.PackageName)
	fmt.Println()
	fmt.Println("Now running and building on system (calculated at runtime):")
	fmt.Println()
	host, err := self.sys.Host()
	if err != nil {
		fmt.Println("error: can't determine current host")
		return err
	}
	fmt.Println("Build        ", host)
	fmt.Println()
	fmt.Println()
	fmt.Println("Values gotten from building site config")
	fmt.Println("(resulting package going to be installed as):")
	fmt.Println()
	fmt.Println("Host         ", i.Host)
	fmt.Println("HostArch     ", i.HostArch)
	fmt.Println()
	fmt.Println()
	fmt.Println("Calculated or guessed stuff:")
	fmt.Println()
	fmt.Println("crossbuilding?  ", i.ThisIsCrossbuilding)
	fmt.Println("crossbuilder?    ", i.ThisIsCrossbuilder)
	fmt.Println("subarch building?", i.ThisIsSubarchBuilding)

	fmt.Println()
	fmt.Println(
		"Calculated --host --build --target options for autotools configurations:",
	)

	if ops, err := self.GetBuildingSiteValuesCalculator().CalculateAutotoolsHBTOptions(); err != nil {
		fmt.Println(" error:", err)
	} else {
		for _, i := range ops {
			fmt.Println("   ", i)
		}
	}

	return nil
}
