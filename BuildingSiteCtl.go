package aipsetup

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildercollection"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/repository"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
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
	log  *logger.Logger

	info *basictypes.BuildingSiteInfo

	buildingsitevaluescalculator *BuildingSiteValuesCalculator
}

func NewBuildingSiteCtl(
	path string,
	sys *System,
	log *logger.Logger,
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

	self.log = log

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
		self.info = j_res
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

func (self *BuildingSiteCtl) InitDirs() error {
	for _, i := range basictypes.DIR_ALL {
		j := filepath.Join(self.path, i)
		_, err := os.Stat(j)
		if err != nil {
			if os.IsNotExist(err) {
				self.log.Info("Creating directory: " + i)
				err := os.MkdirAll(j, 0700)
				if err != nil {
					return errors.New("Can't create dir: " + err.Error())
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func (self *BuildingSiteCtl) DetermineMainTarrball() (string, error) {
	info, err := self.ReadInfo()
	if err != nil {
		return "", err
	}

	stats, err := ioutil.ReadDir(self.GetDIR_TARBALL())
	if err != nil {
		return "", err
	}

	var main_tarball string

	for _, i := range stats {
		infoname, _, err := pkginfodb.DetermineTarballPackageInfoSingle(i.Name())
		if err != nil {
			return "", err
		}

		if infoname == info.PackageName {
			main_tarball = path.Base(i.Name())
			break
		}

	}
	if main_tarball == "" {
		return "", errors.New("couldn't find main tarball for package")
	}

	return main_tarball, nil
}

// TODO: somehow I don't like this method
//func (self *BuildingSiteCtl) ApplyMainTarball() error {

//	maintarball, err := self.DetermineMainTarrball()
//	if err != nil {
//		return err
//	}

//	maintarball = path.Base(maintarball)

//	lst, err := ioutil.ReadDir(self.GetDIR_TARBALL())
//	if err != nil {
//		return err
//	}

//	for _, i := range lst {
//		if maintarball == path.Base(i.Name()) {
//			goto maintarball_found
//		}
//	}

//	return errors.New("specified main tarball not found in tarball dir")

//maintarball_found:

//	filelist := make([]string, 0)

//	filelist = append(filelist, maintarball)

//	for _, i := range lst {
//		b := path.Base(i.Name())
//		found := false

//		for _, j := range filelist {
//			if j == b {
//				found = true
//				break
//			}
//		}

//		if !found {
//			filelist = append(filelist, b)
//		}
//	}

//	info, err := self.ReadInfo()
//	if err != nil {
//		return err
//	}

//	tarball_info, err := pkginfodb.Get(info.PackageName)
//	if err != nil {
//		return err
//	}

//	{

//		parser, err := tarballnameparsers.Get(tarball_info.TarballFileNameParser)
//		if err != nil {
//			return err
//		}

//		err = tarballname.IsPossibleTarballNameErr(filelist[0])
//		if err != nil {
//			return err
//		}

//		parsed, err := parser.Parse(filelist[0])
//		if err != nil {
//			return err
//		}

//		if t, err := parsed.Version.IntSliceString("."); err != nil {
//			return err
//		} else {
//			info.PackageVersion = t
//		}

//		if parsed.Status.StrSliceString("") != "" {
//			info.PackageStatus = parsed.Status.StrSliceString("")
//		}

//	}

//	err = self.WriteInfo(info)
//	if err != nil {
//		return err
//	}

//	return nil
//}

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

	info, err := self.ReadInfo()
	if err != nil {
		return err
	}

	err = self.InitDirs()
	if err != nil {
		return err
	}

	l, err := self.CreateLogger("aipsetup bs run log", true)
	if err != nil {
		return err
	}
	defer l.Close()

	tarball_info, err := pkginfodb.Get(info.PackageName)
	if err != nil {
		return err
	}

	l.Info("Determined Package Info")
	{
		json, err := tarball_info.RenderJSON()
		if err != nil {
			return err
		}

		l.Info(json)
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
				self.log.Error("missing requested target actions: " + i)
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
				lo, err := self.CreateLogger(j, true)
				if err != nil {
					return err
				}
				lo.Info(
					fmt.Sprintf(
						"---===="+`//////`+"[ %s : STRT %s ]"+`\\\\\\`+"====---",
						info.PackageName,
						j,
					),
				)
				err = i.Callable(lo)
				if err != nil {
					lo.Error("(aipsetup message) action exited with following aipsetup error:")
					lo.Error(err)
					lo.Error(
						fmt.Sprintf(
							"---===="+`++++++`+"[ %s : FAIL %s ]"+`++++++`+"====---",
							info.PackageName,
							j,
						),
					)
					return err
				}
				lo.Info(
					fmt.Sprintf(
						"---===="+`\\\\\\`+"[ %s : DONE %s ]"+`//////`+"====---",
						info.PackageName,
						j,
					),
				)
				lo.Close()
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

func (self *BuildingSiteCtl) GetLog() *logger.Logger {
	return self.log
}

func (self *BuildingSiteCtl) GetBuildingSiteValuesCalculator() basictypes.BuildingSiteValuesCalculatorI {
	return self.buildingsitevaluescalculator
}

func (self *BuildingSiteCtl) GetPath() string {
	return self.path
}

func (self *BuildingSiteCtl) GetOuterTarballsDir() (string, error) {
	info, err := self.ReadInfo()
	if err != nil {
		return "", err
	}

	if info.TarballsDir == "" {
		return "", nil
	}

	if strings.HasPrefix(info.TarballsDir, "/") {
		return info.TarballsDir, nil
	}

	return path.Join(self.path, info.TarballsDir), nil
}

func (self *BuildingSiteCtl) GetOuterAspsDir() (string, error) {
	info, err := self.ReadInfo()
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(info.AspsDir, "/") {
		return info.AspsDir, nil
	}

	return path.Join(self.path, info.AspsDir), nil
}

func (self *BuildingSiteCtl) GetSources() error {
	err := self.GetTarballs()
	if err != nil {
		return err
	}

	err = self.GetPatches()
	if err != nil {
		return err
	}

	return nil
}

func (self *BuildingSiteCtl) GetTarballs() error {

	bs_info, err := self.ReadInfo()
	if err != nil {
		return err
	}

	from_path := ""

	if outer_tarballs_dir, err := self.GetOuterTarballsDir(); err != nil {
		return err
	} else {
		if outer_tarballs_dir != "" {
			from_path = bs_info.TarballsDir
			if !strings.HasPrefix(from_path, "/") {
				from_path, err = filepath.Abs(path.Join(self.path, from_path))
				if err != nil {
					return err
				}
			}
		}
	}

	if len(from_path) != 0 {
		if s, err := os.Stat(from_path); err != nil {
			return err
		} else {
			if !s.IsDir() {
				return errors.New("error using directory with tarballs: not a directory")
			}
		}
	}

	repo, err := repository.NewRepository(self.GetSystem(), self.log)
	if err != nil {
		return err
	}

	info, err := pkginfodb.Get(bs_info.PackageName)
	if err != nil {
		return err
	}

	pkg_list := make([]string, 0)

	pkg_list = append(pkg_list, bs_info.PackageName)

	pkg_list = append(pkg_list, info.BuildPkgDeps...)

	tar_dir := self.GetDIR_TARBALL()

	if len(from_path) == 0 {
		all_needed_tarballs_gotted := true

		for _, i := range pkg_list {
			t, err := repo.DetermineNewestStableTarball(i)
			if err != nil {
				all_needed_tarballs_gotted = false
				continue
			}

			err = repo.CopyTarballToDir(i, t, tar_dir)
			if err != nil {
				all_needed_tarballs_gotted = false
				continue
			}
		}
		if !all_needed_tarballs_gotted {
			return errors.New("couldn't get all the needed tarballs")
		}
	} else {
		stats, err := ioutil.ReadDir(from_path)
		if err != nil {
			return err
		}

		all_needed_stats_found := true

		for _, i := range pkg_list {
			self.log.Info("Searching tarball of " + i + " package")
			needed_stat_found := false
			for _, j := range stats {

				if j.IsDir() {
					continue
				}

				i_info, err := pkginfodb.Get(i)
				if err != nil {
					return err
				}

				matches, err := pkginfodb.CheckTarballMatchesInfoByInfoName(j.Name(), i)
				if err != nil {
					return err
				}

				if bs_info.PackageName == i {
					// check version too

					// TODO: this is some fast shitty hack and I don't like this.
					//       probably building site info have to contain reqired
					//       versions along with info.BuildPkgDeps too
					if matches {
						parser, err := tarballnameparsers.Get(
							i_info.TarballFileNameParser,
						)
						if err != nil {
							return err
						}

						parsed, err := parser.Parse(j.Name())
						if err != nil {
							return err
						}

						matches = parsed.Version.StrSliceString(".") ==
							bs_info.PackageVersion

					}
				}

				if matches {
					fs := path.Join(from_path, j.Name())
					ts := path.Join(tar_dir, j.Name())
					self.log.Info("  copy")
					self.log.Info("    " + fs)
					self.log.Info("    to " + ts)
					err := filetools.CopyWithInfo(
						fs,
						ts,
						self.log,
					)
					if err != nil {
						break
					}
					needed_stat_found = true
					break
				}

			}
			if needed_stat_found {
				self.log.Info("  found")
			} else {
				self.log.Error("  not found")
			}
			if !needed_stat_found {
				all_needed_stats_found = false
			}
		}
		if !all_needed_stats_found {
			return errors.New("not found all the needed tarballs in given directory")
		}
	}

	return nil
}

func (self *BuildingSiteCtl) GetPatches() error {

	bs_info, err := self.ReadInfo()
	if err != nil {
		return err
	}

	repo, err := repository.NewRepository(self.GetSystem(), self.log)
	if err != nil {
		return err
	}

	info, err := pkginfodb.Get(bs_info.PackageName)
	if err != nil {
		return err
	}

	// TODO: 1. probably, must be created separate PackageInfo setting,
	//          like ApplyPatches.. not using for this DownloadPatches setting;
	//       2. stop using len(info.PatchesDownloadingScriptText) == 0 as
	//          secondary condition.
	if !info.DownloadPatches || len(info.PatchesDownloadingScriptText) == 0 {
		return nil
	}

	pkg_list := make([]string, 0)

	pkg_list = append(pkg_list, bs_info.PackageName)
	pkg_list = append(pkg_list, info.BuildPkgDeps...)

	patches_dir := self.GetDIR_PATCHES()

	all_needed_patches_gotted := true
	for _, i := range pkg_list {

		err = repo.CopyPatchesToDir(i, patches_dir)
		if err != nil {
			all_needed_patches_gotted = false
			continue
		}
	}
	if !all_needed_patches_gotted {
		return errors.New("couldn't get all the needed patches")
	}

	return nil
}

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
	fmt.Println("Build            ", host)
	fmt.Println()
	fmt.Println("This building site initiated by:")
	fmt.Println()
	fmt.Println("InitiatedByHost  ", i.InitiatedByHost)
	fmt.Println()
	fmt.Println("Values gotten from building site config")
	fmt.Println("(resulting package going to be installed as):")
	fmt.Println()
	fmt.Println("Host             ", i.Host)
	fmt.Println("HostArch         ", i.HostArch)
	fmt.Println()
	fmt.Println()
	fmt.Println("Calculated or guessed stuff:")
	fmt.Println()
	fmt.Println("crossbuilding?   ", i.ThisIsCrossbuilding())
	fmt.Println("crossbuilder?    ", i.ThisIsCrossbuilder())
	fmt.Println("subarch building?", i.ThisIsSubarchBuilding())

	fmt.Println()
	fmt.Println(
		"Calculated --host --build --target options for autotools configurations:",
	)

	if ops, err := self.GetBuildingSiteValuesCalculator().
		CalculateAutotoolsHBTOptions(); err != nil {
		fmt.Println(" error:", err)
	} else {
		for _, i := range ops {
			fmt.Println("   ", i)
		}
	}

	return nil
}
