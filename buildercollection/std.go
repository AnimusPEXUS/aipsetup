package buildercollection

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
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
		return NewBuilder_std(bs), nil
	}
}

type CrossBuildEnum uint

const (
	NoAction CrossBuildEnum = iota
	Force
	Forbid
)

type Builder_std struct {

	// NOTE: some comments in this file are left from python time and may be not
	//       correspond to situation. (2018-03-12)

	EditActionsCB                       func(ret basictypes.BuilderActions) (basictypes.BuilderActions, error)
	AfterExtractCB                      func(log *logger.Logger, ret error) error
	AfterAutogenCB                      func(log *logger.Logger, ret error) error
	EditExtractMoreThanOneExtractedOkCB func(log *logger.Logger, ret bool) (bool, error)
	EditExtractUnwrapCB                 func(log *logger.Logger, ret bool) (bool, error)
	EditAutogenForceCB                  func(log *logger.Logger, ret bool) (bool, error)
	EditAutogenFailIsOkCB               func(log *logger.Logger, ret bool) (bool, error)
	EditConfigureEnvCB                  func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditConfigureArgsCB                 func(log *logger.Logger, ret []string) ([]string, error)
	EditConfigureScriptNameCB           func(log *logger.Logger, ret string) (string, error)
	EditConfigureDirCB                  func(log *logger.Logger, ret string) (string, error)
	EditConfigureWorkingDirCB           func(log *logger.Logger, ret string) (string, error)
	EditConfigureRelativeExecutionCB    func(log *logger.Logger, ret bool) (bool, error)
	EditConfigureShellCB                func(log *logger.Logger, ret string) (string, error)
	EditConfigureIsArgToShellCB         func(log *logger.Logger, ret bool) (bool, error)
	EditBuildConcurentJobsCountCB       func(log *logger.Logger, ret int) int
	EditBuildEnvCB                      func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditBuildArgsCB                     func(log *logger.Logger, ret []string) ([]string, error)
	EditBuildMakefileNameCB             func(log *logger.Logger, ret string) (string, error)
	EditBuildMakefileDirCB              func(log *logger.Logger, ret string) (string, error)
	EditBuildWorkingDirCB               func(log *logger.Logger, ret string) (string, error)
	EditDistributeEnvCB                 func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditDistributeDESTDIRCB             func(log *logger.Logger, ret string) (string, error)
	EditDistributeArgsCB                func(log *logger.Logger, ret []string) ([]string, error)
	EditDistributeMakefileNameCB        func(log *logger.Logger, ret string) (string, error)
	EditDistributeMakefileCB            func(log *logger.Logger, ret string) (string, error)
	EditDistributeWorkingDirCB          func(log *logger.Logger, ret string) (string, error)

	bs_ren basictypes.BuildingSiteCtlI
}

// builders are independent of anything so have no moto to return errors
func NewBuilder_std(bs basictypes.BuildingSiteCtlI) *Builder_std {
	ret := new(Builder_std)

	ret.bs_ren = bs

	return ret
}

func (self *Builder_std) GetBuildingSiteCtl() basictypes.BuildingSiteCtlI {
	return self.bs_ren
}

func (self *Builder_std) DefineActions() (basictypes.BuilderActions, error) {

	ret := basictypes.BuilderActions{

		&basictypes.BuilderAction{"dst_cleanup", self.BuilderActionDstCleanup},
		&basictypes.BuilderAction{"src_cleanup", self.BuilderActionSrcCleanup},
		&basictypes.BuilderAction{"bld_cleanup", self.BuilderActionBldCleanup},
		&basictypes.BuilderAction{"extract", self.BuilderActionExtract},
		&basictypes.BuilderAction{"autogen", self.BuilderActionAutogen},
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

func (self *Builder_std) BuilderActionDstCleanup(
	log *logger.Logger,
) error {
	dst_dir := self.GetBuildingSiteCtl().GetDIR_DESTDIR()
	os.RemoveAll(dst_dir)
	os.MkdirAll(dst_dir, 0700)
	return nil
}

func (self *Builder_std) BuilderActionSrcCleanup(
	log *logger.Logger,
) error {
	src_dir := self.GetBuildingSiteCtl().GetDIR_SOURCE()
	os.RemoveAll(src_dir)
	os.MkdirAll(src_dir, 0700)
	return nil
}
func (self *Builder_std) BuilderActionBldCleanup(
	log *logger.Logger,
) error {
	bld_dir := self.GetBuildingSiteCtl().GetDIR_BUILDING()
	os.RemoveAll(bld_dir)
	os.MkdirAll(bld_dir, 0700)
	return nil
}

func (self *Builder_std) BuilderActionExtract(
	log *logger.Logger,
) error {
	a_tools := new(buildingtools.Autotools)

	more_than_one_extracted_ok := false
	unwrap := true

	if self.EditExtractMoreThanOneExtractedOkCB != nil {
		var err error
		more_than_one_extracted_ok, err = self.EditExtractMoreThanOneExtractedOkCB(log, more_than_one_extracted_ok)
		if err != nil {
			return err
		}
	}

	if self.EditExtractUnwrapCB != nil {
		var err error
		unwrap, err = self.EditExtractUnwrapCB(log, unwrap)
		if err != nil {
			return err
		}
	}

	main_tarball, err := self.GetBuildingSiteCtl().DetermineMainTarrball()
	if err != nil {
		return err
	}

	err = a_tools.Extract(
		path.Join(self.GetBuildingSiteCtl().GetDIR_TARBALL(), main_tarball),
		self.GetBuildingSiteCtl().GetDIR_SOURCE(),
		path.Join(self.GetBuildingSiteCtl().GetDIR_TEMP(), "primary_tarball"),
		unwrap,
		more_than_one_extracted_ok,
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

func (self *Builder_std) BuilderActionAutogenForce(log *logger.Logger) (bool, error) {

	ret := false

	if self.EditAutogenForceCB != nil {
		var err error
		ret, err = self.EditAutogenForceCB(log, ret)
		if err != nil {
			return false, err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionAutogen(log *logger.Logger) error {
	needs_autogen := false

	var err error

	fail_is_ok := false
	if self.EditAutogenFailIsOkCB != nil {
		fail_is_ok, err = self.EditAutogenFailIsOkCB(log, fail_is_ok)
		if err != nil {
			return err
		}
	}

	config_script_name, err := self.BuilderActionConfigureScriptNameDef(log)
	if err != nil {
		return err
	}

	configure_dir, err := self.BuilderActionConfigureDirDef(log)
	if err != nil {
		return err
	}

	configure_path := path.Join(
		configure_dir,
		config_script_name,
	)

	_, err = os.Stat(configure_path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			needs_autogen = true
		}
	}

	autogen_force, err := self.BuilderActionAutogenForce(log)
	if err != nil {
		return err
	}

	if !needs_autogen && !autogen_force {
		log.Info("autogen usage not needed and not forced. continuing without it")
		return nil
	}

	if autogen_force {
		log.Info("autogen usage forced")
	}

	if needs_autogen {
		log.Info("detected need to use autogen")
	}

	generated := false

	log.Info("searching for suitable generator")
	for _, i := range [][]string{
		[]string{"makeconf.sh", "./makeconf.sh"},
		[]string{"autogen.sh", "./autogen.sh"},
		[]string{"bootstrap.sh", "./bootstrap.sh"},
		[]string{"bootstrap", "./bootstrap"},
		[]string{"genconfig.sh", "./genconfig.sh"},
		[]string{"configure.ac", "autoreconf", "-i"},
		[]string{"configure.in", "autoreconf", "-i"},
	} {
		log.Info("  " + i[0])
		generator_name := path.Join(configure_dir, i[0])

		_, err := os.Stat(generator_name)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			} else {
				log.Info("    not found")
				continue
			}
		}
		log.Info("    found")

		log.Info(fmt.Sprintf("executing %s %v", i[1], i[2:]))

		c := exec.Command(i[1], i[2:]...)
		c.Dir = configure_dir
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()

		err = c.Run()
		if err != nil {
			log.Error("  error: " + err.Error())
			if fail_is_ok {
				return nil
			} else {
				return err
			}
		}

		log.Info("autogen exited success code")

		generated = true
		break
	}

	if !generated {
		err = errors.New("couldn't find suitable generator")
	}

	if self.AfterAutogenCB != nil {
		err = self.AfterAutogenCB(log, err)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_std) BuilderActionConfigureEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {
	env := environ.New()

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	pkgconfig, err := calc.GetPrefixPkgConfig()
	if err != nil {
		return env, err
	}

	ldlp, err := calc.Calculate_LD_LIBRARY_PATH()
	if err != nil {
		return env, err
	}

	lp, err := calc.Calculate_LIBRARY_PATH()
	if err != nil {
		return env, err
	}

	ci, err := calc.Calculate_C_INCLUDE_PATH()
	if err != nil {
		return env, err
	}

	path, err := calc.Calculate_PATH()
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

	env.Set("PKG_CONFIG_PATH", pkgconfig.GetPKG_CONFIG_PATH())
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

func (self *Builder_std) BuilderActionConfigureArgsDef(
	log *logger.Logger,
) ([]string, error) {

	ret := make([]string, 0)

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	libdir, err := calc.CalculateInstallLibDir()
	if err != nil {
		return nil, err
	}

	opt_map, err := calc.CalculateAutotoolsAllOptionsMap()
	if err != nil {
		return nil, err
	}

	hbt_options, err := calc.CalculateAutotoolsHBTOptions()
	if err != nil {
		return nil, err
	}

	docdir := path.Join(prefix, "share", "doc")

	ret = append(
		ret,
		fmt.Sprintf("--prefix=%s", prefix),
		fmt.Sprintf("--libdir=%s", libdir),
		fmt.Sprintf("--docdir=%s", docdir),
		"--sysconfdir=/etc",
		"--localstatedir=/var",
		"--enable-shared",
	)

	ret = append(
		ret,
		hbt_options...,
	)

	// replacement for std.py's self.all_automatic_flags_as_list()
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

func (self *Builder_std) BuilderActionConfigureScriptNameDef(
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

func (self *Builder_std) BuilderActionConfigureDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.GetBuildingSiteCtl().GetDIR_SOURCE()

	if self.EditConfigureDirCB != nil {
		var err error
		ret, err = self.EditConfigureDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionConfigureWorkingDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.GetBuildingSiteCtl().GetDIR_SOURCE()

	if self.EditConfigureWorkingDirCB != nil {
		var err error
		ret, err = self.EditConfigureWorkingDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionConfigureRelativeExecutionDef(
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

func (self *Builder_std) BuilderActionConfigureShellDef(
	log *logger.Logger,
) (string, error) {

	ret := "bash"

	if self.EditConfigureShellCB != nil {
		var err error
		ret, err = self.EditConfigureShellCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionConfigureIsArgToShellDef(
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

func (self *Builder_std) BuilderActionConfigure(
	log *logger.Logger,
) error {

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

	shell, err := self.BuilderActionConfigureShellDef(log)
	if err != nil {
		return err
	}

	is_arg_to_shell, err := self.BuilderActionConfigureIsArgToShellDef(log)
	if err != nil {
		return err
	}

	a_tools := new(buildingtools.Autotools)

	err = a_tools.Configure(
		args,
		env.Strings(),
		buildingtools.Copy,
		cfg_script_name,
		cd,
		wd,
		is_rel,
		is_arg_to_shell,
		shell,
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_std) BuilderActionBuildConcurentJobsCountDef(
	log *logger.Logger,
) int {

	ret := runtime.NumCPU()

	if self.EditBuildConcurentJobsCountCB != nil {
		ret = self.EditBuildConcurentJobsCountCB(log, ret)
	}

	return ret

}

func (self *Builder_std) BuilderActionBuildEnvDef(
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

func (self *Builder_std) BuilderActionBuildArgsDef(
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

func (self *Builder_std) BuilderActionBuildMakefileNameDef(
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

func (self *Builder_std) BuilderActionBuildMakefileDirDef(
	log *logger.Logger,
) (string, error) {

	ret, err := self.BuilderActionConfigureDirDef(log)
	if err != nil {
		return "", err
	}

	if self.EditBuildMakefileDirCB != nil {
		var err error
		ret, err = self.EditBuildMakefileDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionBuildWorkingDirDef(
	log *logger.Logger,
) (string, error) {

	ret, err := self.BuilderActionConfigureWorkingDirDef(log)
	if err != nil {
		return "", err
	}

	if self.EditBuildWorkingDirCB != nil {
		var err error
		ret, err = self.EditBuildWorkingDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionBuild(
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
		if cpu_count > 0 {
			args2 = append(args2, fmt.Sprintf("-j%d", cpu_count))
		}
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

func (self *Builder_std) BuilderActionDistributeEnvDef(
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

func (self *Builder_std) BuilderActionDistributeDESTDIRDef(
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

func (self *Builder_std) BuilderActionDistributeArgsDef(
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
			self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
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

func (self *Builder_std) BuilderActionDistributeMakefileNameDef(
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

func (self *Builder_std) BuilderActionDistributeMakefileDirDef(
	log *logger.Logger,
) (string, error) {

	ret := self.GetBuildingSiteCtl().GetDIR_SOURCE()

	if self.EditDistributeMakefileCB != nil {
		var err error
		ret, err = self.EditDistributeMakefileCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionDistributeWorkingDirDef(
	log *logger.Logger,
) (string, error) {

	ret, err := self.BuilderActionBuildWorkingDirDef(log)
	if err != nil {
		return "", err
	}

	if self.EditDistributeWorkingDirCB != nil {
		var err error
		ret, err = self.EditDistributeWorkingDirCB(log, ret)
		if err != nil {
			return "", err
		}
	}

	return ret, nil
}

func (self *Builder_std) BuilderActionDistribute(
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

func (self *Builder_std) BuilderActionPrePack(
	log *logger.Logger,
) error {
	err := self.GetBuildingSiteCtl().PrePackager().Run(log)
	if err != nil {
		return err
	}
	return nil
}

func (self *Builder_std) BuilderActionPack(
	log *logger.Logger,
) error {
	err := self.GetBuildingSiteCtl().Packager().Run(log)
	if err != nil {
		return err
	}
	return nil
}
