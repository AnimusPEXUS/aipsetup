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
	Index["std"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
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

	// # this is for builder_action_autogen() method
	ForcedAutogen                bool
	SeparateBuildDir             bool
	SourceConfigureRelPath       string
	ForcedTarget                 bool
	ApplyHostSpecCompilerOptions bool

	// # None - not used, bool - force value
	ForceCrossbuilder CrossBuildEnum
	ForceCrossbuild   CrossBuildEnum

	site basictypes.BuildingSiteCtlI
}

// builders are independent of anything so have no moto to return errors
func NewBuilderStdAutotools(buildingsite basictypes.BuildingSiteCtlI) *BuilderStdAutotools {
	ret := new(BuilderStdAutotools)

	ret.site = buildingsite

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
	}

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDstCleanup(
	log *logger.Logger,
) error {
	dst_dir := self.site.GetDIR_DESTDIR()
	os.RemoveAll(dst_dir)
	os.MkdirAll(dst_dir, 0700)
	return nil
}

func (self *BuilderStdAutotools) BuilderActionSrcCleanup(
	log *logger.Logger,
) error {
	src_dir := self.site.GetDIR_SOURCE()
	os.RemoveAll(src_dir)
	os.MkdirAll(src_dir, 0700)
	return nil
}
func (self *BuilderStdAutotools) BuilderActionBldCleanup(
	log *logger.Logger,
) error {
	bld_dir := self.site.GetDIR_BUILDING()
	os.RemoveAll(bld_dir)
	os.MkdirAll(bld_dir, 0700)
	return nil
}

func (self *BuilderStdAutotools) BuilderActionPrimaryExtract(
	log *logger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)

	info, err := self.site.ReadInfo()
	if err != nil {
		return err
	}

	if len(info.Sources) == 0 {
		return errors.New("no tarballs supplied. primary tarball is required")
	}
	tarball := info.Sources[0]
	tarball = path.Join(self.site.GetDIR_TARBALL(), tarball)
	err = a_tools.Extract(
		tarball,
		self.site.GetDIR_SOURCE(),
		path.Join(self.site.GetDIR_TEMP(), "primary_tarball"),
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

	calc := self.site.ValuesCalculator()

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

	return env, nil

}

func (self *BuilderStdAutotools) BuilderActionConfigureArgsDef(
	log *logger.Logger,
) ([]string, error) {

	ret := make([]string, 0)

	calc := self.site.ValuesCalculator()

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

	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureScriptNameDef(
	log *logger.Logger,
) (string, error) {
	return "configure", nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureDirDef(
	log *logger.Logger,
) (string, error) {
	return self.site.GetDIR_SOURCE(), nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureWorkingDirDef(
	log *logger.Logger,
) (string, error) {
	return self.site.GetDIR_SOURCE(), nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureRelativeExecutionDef(
	log *logger.Logger,
) (bool, error) {

	return true, nil
}

func (self *BuilderStdAutotools) BuilderActionConfigureIsArgToShellDef(
	log *logger.Logger,
) (bool, error) {
	return false, nil
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

	return runtime.NumCPU()

}

func (self *BuilderStdAutotools) BuilderActionBuildEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {
	log.Info(
		"this builder uses same environment variables for make as for configure",
	)
	return self.BuilderActionConfigureEnvDef(log)
}

func (self *BuilderStdAutotools) BuilderActionBuildArgsDef(
	log *logger.Logger,
) ([]string, error) {
	ret := make([]string, 0)
	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionBuildMakefileNameDef(
	log *logger.Logger,
) (string, error) {
	return "Makefile", nil
}

func (self *BuilderStdAutotools) BuilderActionBuildMakefileDirDef(
	log *logger.Logger,
) (string, error) {
	return self.site.GetDIR_SOURCE(), nil
}

func (self *BuilderStdAutotools) BuilderActionBuildWorkingDirDef(
	log *logger.Logger,
) (string, error) {
	return self.site.GetDIR_SOURCE(), nil
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
	log.Info(
		"this builder uses same environment variables for make as for configure",
	)
	return self.BuilderActionConfigureEnvDef(log)
}

func (self *BuilderStdAutotools) BuilderActionDistributeDESTDIRDef(
	log *logger.Logger,
) string {
	return "DESTDIR"
}

func (self *BuilderStdAutotools) BuilderActionDistributeArgsDef(
	log *logger.Logger,
) ([]string, error) {
	ret := make([]string, 0)
	ret = append(ret, "install")
	ret = append(
		ret,
		fmt.Sprintf(
			"%s=%s",
			self.BuilderActionDistributeDESTDIRDef(log),
			self.site.GetDIR_DESTDIR(),
		),
	)
	return ret, nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeMakefileNameDef(
	log *logger.Logger,
) (string, error) {
	return "Makefile", nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeMakefileDirDef(
	log *logger.Logger,
) (string, error) {
	return self.site.GetDIR_SOURCE(), nil
}

func (self *BuilderStdAutotools) BuilderActionDistributeWorkingDirDef(
	log *logger.Logger,
) (string, error) {
	return self.site.GetDIR_SOURCE(), nil
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
