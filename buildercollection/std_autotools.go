package buildercollection

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_autotools"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderStdAutotools(bs), nil
	}
}

type CrossBuildEnum uint

const (
	NoAction CrossBuildEnum = iota
	Force
	Forbid
)

type BuilderStdAutotools struct {

	// NOTE: some comments in this file are left from python time and may be not
	//       correspond to situation. (2018-03-12)

	// # this is for builder_action_autogen() method
	ForcedAutogen                bool
	SeparateBuildDir             bool
	SourceConfigureRelPath       string
	ForcedTarget                 bool
	ApplyHostSpecCompilerOptions bool

	// # None - not used, bool - force value
	ForceCrossbuilder CrossBuildEnum
	ForceCrossbuild   CrossBuildEnum

	EditActionsCB                    func(basictypes.BuilderActions) (basictypes.BuilderActions, error)
	AfterExtractCB                   func(log *logger.Logger, ret error) error
	EditConfigureEnvCB               func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditConfigureArgsCB              func(log *logger.Logger, ret []string) ([]string, error)
	EditConfigureScriptNameCB        func(log *logger.Logger, ret string) (string, error)
	EditConfigureDirCB               func(log *logger.Logger, ret string) (string, error)
	EditConfigureWorkingDirCB        func(log *logger.Logger, ret string) (string, error)
	EditConfigureRelativeExecutionCB func(log *logger.Logger, ret bool) (bool, error)
	EditConfigureIsArgToShellCB      func(log *logger.Logger, ret bool) (bool, error)
	EditBuildConcurentJobsCountCB    func(log *logger.Logger, ret int) int
	EditBuildEnvCB                   func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditBuildArgsCB                  func(log *logger.Logger, ret []string) ([]string, error)
	EditBuildMakefileNameCB          func(log *logger.Logger, ret string) (string, error)
	EditBuildMakefileDirCB           func(log *logger.Logger, ret string) (string, error)
	EditBuildWorkingDirCB            func(log *logger.Logger, ret string) (string, error)
	EditDistributeEnvCB              func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditDistributeDESTDIRCB          func(log *logger.Logger, ret string) (string, error)
	EditDistributeArgsCB             func(log *logger.Logger, ret []string) ([]string, error)
	EditDistributeMakefileNameCB     func(log *logger.Logger, ret string) (string, error)
	EditDistributeMakefileCB         func(log *logger.Logger, ret string) (string, error)
	EditDistributeWorkingDirCB       func(log *logger.Logger, ret string) (string, error)

	bs basictypes.BuildingSiteCtlI
}

// builders are independent of anything so have no moto to return errors
func NewBuilderStdAutotools(buildingsite basictypes.BuildingSiteCtlI) *BuilderStdAutotools {
	ret := new(BuilderStdAutotools)

	ret.bs = buildingsite

	ret.ForcedAutogen = false

	ret.SeparateBuildDir = false
	ret.SourceConfigureRelPath = "."
	ret.ForcedTarget = false
	ret.ApplyHostSpecCompilerOptions = true

	ret.ForceCrossbuilder = NoAction
	ret.ForceCrossbuild = NoAction

	return ret
}

// func (self *BuilderStdAutotools) SetBuildingSite(bs basictypes.BuildingSiteCtlI) {
// 	self.site = bs
// }

func (self *BuilderStdAutotools) DefineActions() (basictypes.BuilderActions, error) {

	ret := basictypes.BuilderActions{

		&basictypes.BuilderAction{"dst_cleanup", self.BuilderActionDstCleanup},
		&basictypes.BuilderAction{"src_cleanup", self.BuilderActionSrcCleanup},
		&basictypes.BuilderAction{"bld_cleanup", self.BuilderActionBldCleanup},
		&basictypes.BuilderAction{"primary_extract", self.BuilderActionPrimaryExtract},
		// &basictypes.BuilderAction{				//ret["patch"] = self.BuilderActionPatch},
		// &basictypes.BuilderAction{				//ret["autogen"] = self.BuilderActionAutogen},
		&basictypes.BuilderAction{"configure", self.BuilderActionConfigure},
		&basictypes.BuilderAction{"build", self.BuilderActionBuild},
		&basictypes.BuilderAction{"distribute", self.BuilderActionDistribute},
		&basictypes.BuilderAction{"prepack", self.BuilderActionPrePack},
		&basictypes.BuilderAction{"pack", self.BuilderActionPack},
	}

	if self.EditActionsCB != nil {
		var err error
		ret, err = self.EditActionsCB(ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDstCleanup(
	log *logger.Logger,
) error {
	dst_dir := self.bs.GetDIR_DESTDIR()
	os.RemoveAll(dst_dir)
	os.MkdirAll(dst_dir, 0700)
	return nil
}

func (self *BuilderStdAutotools) BuilderActionSrcCleanup(
	log *logger.Logger,
) error {
	src_dir := self.bs.GetDIR_SOURCE()
	os.RemoveAll(src_dir)
	os.MkdirAll(src_dir, 0700)
	return nil
}
func (self *BuilderStdAutotools) BuilderActionBldCleanup(
	log *logger.Logger,
) error {
	bld_dir := self.bs.GetDIR_BUILDING()
	os.RemoveAll(bld_dir)
	os.MkdirAll(bld_dir, 0700)
	return nil
}

func (self *BuilderStdAutotools) BuilderActionPrimaryExtract(
	log *logger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if len(info.Sources) == 0 {
		return errors.New("no tarballs supplied. primary tarball is required")
	}
	tarball := info.Sources[0]
	tarball = path.Join(self.bs.GetDIR_TARBALL(), tarball)
	err = a_tools.Extract(
		tarball,
		self.bs.GetDIR_SOURCE(),
		path.Join(self.bs.GetDIR_TEMP(), "primary_tarball"),
		true,
		false,
		"",
		false,
		false,
		log,
	)
	if err != nil {
		return err
	}

	if self.AfterExtractCB != nil {
		err = self.AfterExtractCB(log, err)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BuilderStdAutotools) BuilderActionPatch(
	log *logger.Logger,
) error {
	return errors.New("not impl")
}

func (self *BuilderStdAutotools) BuilderActionAutogen(
	log *logger.Logger,
) error {
	return errors.New("not impl")
}

func (self *BuilderStdAutotools) BuilderActionConfigureEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {
	env := environ.New()

	calc := self.bs.ValuesCalculator()

	pkgcp, err := calc.CalculatePkgConfigSearchPaths("")
	if err != nil {
		return env, err
	}

	ldlp, err := calc.Calculate_LD_LIBRARY_PATH([]string{})
	if err != nil {
		return env, err
	}

	lp, err := calc.Calculate_LIBRARY_PATH([]string{})
	if err != nil {
		return env, err
	}

	ci, err := calc.Calculate_C_INCLUDE_PATH([]string{})
	if err != nil {
		return env, err
	}

	path, err := calc.Calculate_PATH("")
	if err != nil {
		return env, err
	}

	cc, err := calc.CalculateAutotoolsCCParameterValue()
	if err != nil {
		return env, err
	}

	cxx, err := calc.CalculateAutotoolsCXXParameterValue()
	if err != nil {
		return env, err
	}

	env.Set("PKG_CONFIG_PATH", strings.Join(pkgcp, ":"))
	env.Set("LD_LIBRARY_PATH", strings.Join(ldlp, ":"))
	env.Set("LIBRARY_PATH", strings.Join(lp, ":"))
	env.Set("C_INCLUDE_PATH", strings.Join(ci, ":"))
	env.Set("PATH", strings.Join(path, ":"))
	env.Set("CC", cc)
	env.Set("CXX", cxx)

	if self.EditConfigureEnvCB != nil {
		env, err = self.EditConfigureEnvCB(log, env)
		if err != nil {
			return nil, err
		}
	}

	return env, nil

}

func (self *BuilderStdAutotools) BuilderActionConfigureArgsDef(
	log *logger.Logger,
) ([]string, error) {

	ret := make([]string, 0)

	calc := self.bs.ValuesCalculator()

	prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return ret, err
	}

	libdir, err := calc.CalculateInstallLibDir()
	if err != nil {
		return ret, err
	}

	opt_map, err := calc.CalculateAllOptionsMap()
	if err != nil {
		return ret, err
	}

	hbt_options, err := calc.CalculateAutotoolsHBTOptions()
	if err != nil {
		return ret, err
	}

	ret = append(
		ret,
		fmt.Sprintf("--prefix=%s", prefix),
		fmt.Sprintf("--libdir=%s", libdir),
		"--sysconfdir=/etc",
		"--localstatedir=/var",
		"--enable-shared",
	)

	ret = append(
		ret,
		hbt_options...,
	)

	ret = append(
		ret,
		opt_map.Strings()...,
	)

	if self.EditConfigureArgsCB != nil {
		ret, err = self.EditConfigureArgsCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureScriptNameDef(
	log *logger.Logger,
) (string, error) {

	ret := "configure"

	if self.EditConfigureScriptNameCB != nil {
		var err error
		ret, err = self.EditConfigureScriptNameCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.bs.GetDIR_SOURCE()

	if self.EditConfigureDirCB != nil {
		var err error
		ret, err = self.EditConfigureDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureWorkingDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.bs.GetDIR_SOURCE()

	if self.EditConfigureWorkingDirCB != nil {
		var err error
		ret, err = self.EditConfigureWorkingDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureRelativeExecutionDef(
	log *logger.Logger,
) (bool, error) {

	ret := true

	if self.EditConfigureRelativeExecutionCB != nil {
		var err error
		ret, err = self.EditConfigureRelativeExecutionCB(log, ret)
		if err != nil {
			return false, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureIsArgToShellDef(
	log *logger.Logger,
) (bool, error) {

	ret := false

	if self.EditConfigureIsArgToShellCB != nil {
		var err error
		ret, err = self.EditConfigureIsArgToShellCB(log, ret)
		if err != nil {
			return false, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigure(
	log *logger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)

	env := environ.New()

	nenv, err := self.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return err
	}

	env.UpdateWith(nenv)

	args, err := self.BuilderActionConfigureArgsDef(log)
	if err != nil {
		return err
	}

	cfg_script_name, err := self.BuilderActionConfigureScriptNameDef(log)
	if err != nil {
		return err
	}

	cd, err := self.BuilderActionConfigureDirDef(log)
	if err != nil {
		return err
	}

	wd, err := self.BuilderActionConfigureWorkingDirDef(log)
	if err != nil {
		return err
	}

	is_rel, err := self.BuilderActionConfigureRelativeExecutionDef(log)
	if err != nil {
		return err
	}

	is_arg_to_shell, err := self.BuilderActionConfigureIsArgToShellDef(log)
	if err != nil {
		return err
	}

	err = a_tools.Configure(
		args,
		env.Strings(),
		buildingtools.Copy,
		cfg_script_name,
		cd,
		wd,
		is_rel,
		is_arg_to_shell,
		"bash",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *BuilderStdAutotools) BuilderActionBuildConcurentJobsCountDef(
	log *logger.Logger,
) int {

	ret := runtime.NumCPU()

	if self.EditBuildConcurentJobsCountCB != nil {
		ret = self.EditBuildConcurentJobsCountCB(log, ret)
	}

	return ret

}

func (self *BuilderStdAutotools) BuilderActionBuildEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {
	log.Info(
		"this builder uses same environment variables for make as for configure",
	)

	ret, err := self.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return nil, err
	}

	if self.EditBuildEnvCB != nil {

		ret, err = self.EditBuildEnvCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionBuildArgsDef(
	log *logger.Logger,
) ([]string, error) {
	ret := make([]string, 0)

	if self.EditBuildArgsCB != nil {
		var err error
		ret, err = self.EditBuildArgsCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionBuildMakefileNameDef(
	log *logger.Logger,
) (string, error) {

	ret := "Makefile"

	if self.EditBuildMakefileNameCB != nil {
		var err error
		ret, err = self.EditBuildMakefileNameCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionBuildMakefileDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.bs.GetDIR_SOURCE()

	if self.EditBuildMakefileDirCB != nil {
		var err error
		ret, err = self.EditBuildMakefileDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionBuildWorkingDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.bs.GetDIR_SOURCE()

	if self.EditBuildWorkingDirCB != nil {
		var err error
		ret, err = self.EditBuildWorkingDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionBuild(
	log *logger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)

	cpu_count := self.BuilderActionBuildConcurentJobsCountDef(log)

	env := environ.New()

	nenv, err := self.BuilderActionBuildEnvDef(log)
	if err != nil {
		return err
	}

	env.UpdateWith(nenv)

	args, err := self.BuilderActionBuildArgsDef(log)
	if err != nil {
		return err
	}

	{
		args2 := make([]string, 0)
		args2 = append(args2, fmt.Sprintf("-j%d", cpu_count))
		args2 = append(args2, args...)
		args = args2
	}

	makefile_name, err := self.BuilderActionBuildMakefileNameDef(log)
	if err != nil {
		return err
	}

	makefile_dir, err := self.BuilderActionBuildMakefileDirDef(log)
	if err != nil {
		return err
	}

	wd, err := self.BuilderActionBuildWorkingDirDef(log)
	if err != nil {
		return err
	}

	err = a_tools.Make(
		args,
		env.Strings(),
		buildingtools.Copy,
		makefile_name,
		makefile_dir,
		wd,
		"make",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {

	// TODO: all those info logs are, probably, should be corrected.. or
	//       detalized in Edit callbacks
	log.Info(
		"this builder uses same environment variables for make as for configure",
	)

	ret, err := self.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return nil, err
	}

	if self.EditDistributeEnvCB != nil {
		ret, err = self.EditDistributeEnvCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeDESTDIRDef(
	log *logger.Logger,
) (string, error) {

	ret := "DESTDIR"

	if self.EditDistributeDESTDIRCB != nil {
		var err error
		ret, err = self.EditDistributeDESTDIRCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeArgsDef(
	log *logger.Logger,
) ([]string, error) {
	destdir_string, err := self.BuilderActionDistributeDESTDIRDef(log)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0)
	ret = append(ret, "install")
	ret = append(
		ret,
		fmt.Sprintf(
			"%s=%s",
			destdir_string,
			self.bs.GetDIR_DESTDIR(),
		),
	)

	if self.EditDistributeArgsCB != nil {
		var err error
		ret, err = self.EditDistributeArgsCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeMakefileNameDef(
	log *logger.Logger,
) (string, error) {

	ret := "Makefile"

	if self.EditDistributeMakefileNameCB != nil {
		var err error
		ret, err = self.EditDistributeMakefileNameCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeMakefileDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.bs.GetDIR_SOURCE()

	if self.EditDistributeMakefileCB != nil {
		var err error
		ret, err = self.EditDistributeMakefileCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeWorkingDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.bs.GetDIR_SOURCE()

	if self.EditDistributeWorkingDirCB != nil {
		var err error
		ret, err = self.EditDistributeWorkingDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDistribute(
	log *logger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)

	env := environ.New()

	nenv, err := self.BuilderActionDistributeEnvDef(log)
	if err != nil {
		return err
	}

	env.UpdateWith(nenv)

	args, err := self.BuilderActionDistributeArgsDef(log)
	if err != nil {
		return err
	}

	makefile_name, err := self.BuilderActionDistributeMakefileNameDef(log)
	if err != nil {
		return err
	}

	makefile_dir, err := self.BuilderActionDistributeMakefileDirDef(log)
	if err != nil {
		return err
	}

	wd, err := self.BuilderActionDistributeWorkingDirDef(log)
	if err != nil {
		return err
	}

	err = a_tools.Make(
		args,
		env.Strings(),
		buildingtools.Copy,
		makefile_name,
		makefile_dir,
		wd,
		"make",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *BuilderStdAutotools) BuilderActionPrePack(
	log *logger.Logger,
) error {
	err := self.bs.PrePackager().Run(log)
	if err != nil {
		return err
	}
	return nil
}

func (self *BuilderStdAutotools) BuilderActionPack(
	log *logger.Logger,
) error {
	err := self.bs.Packager().Run(log)
	if err != nil {
		return err
	}
	return nil
}
