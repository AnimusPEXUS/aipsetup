package buildercollection

import (
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_waf"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_waf(bs)
	}
}

type Builder_std_waf struct {
	builder_std *Builder_std

	EditActionsCB       func(ret basictypes.BuilderActions) (basictypes.BuilderActions, error)
	EditConfigureEnvCB  func(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error)
	EditConfigureArgsCB func(log *logger.Logger, ret []string) ([]string, error)
}

func NewBuilder_std_waf(bs basictypes.BuildingSiteCtlI) (*Builder_std_waf, error) {
	self := new(Builder_std_waf)

	self.builder_std = NewBuilder_std(bs)

	return self, nil
}

func (self *Builder_std_waf) GetBuildingSiteCtl() basictypes.BuildingSiteCtlI {
	return self.builder_std.GetBuildingSiteCtl()
}

func (self *Builder_std_waf) DefineActions() (basictypes.BuilderActions, error) {

	ret := basictypes.BuilderActions{

		&basictypes.BuilderAction{"dst_cleanup", self.builder_std.BuilderActionDstCleanup},
		&basictypes.BuilderAction{"src_cleanup", self.builder_std.BuilderActionSrcCleanup},
		&basictypes.BuilderAction{"bld_cleanup", self.builder_std.BuilderActionBldCleanup},
		&basictypes.BuilderAction{"extract", self.builder_std.BuilderActionExtract},
		&basictypes.BuilderAction{"configure", self.BuilderActionConfigure},
		&basictypes.BuilderAction{"build", self.builder_std.BuilderActionBuild},
		&basictypes.BuilderAction{"distribute", self.builder_std.BuilderActionDistribute},
		&basictypes.BuilderAction{"prepack", self.builder_std.BuilderActionPrePack},
		&basictypes.BuilderAction{"pack", self.builder_std.BuilderActionPack},
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

//func (self *Builder_std_waf) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

//	ret, err := buildingtools.FilterAutotoolsConfigOptions(
//		ret,
//		[]string{"--enable-shared"},
//		[]string{"CC=", "CXX=", "GCC="},
//	)
//	if err != nil {
//		return nil, err
//	}

//	return ret, nil
//}

func (self *Builder_std_waf) BuilderActionConfigureArgsDef(
	log *logger.Logger,
) ([]string, error) {

	ret, err := self.builder_std.BuilderActionConfigureArgsDef(log)
	if err != nil {
		return nil, err
	}

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{"--enable-shared"},
		[]string{"CC=", "CXX=", "GCC="},
		//"--host", "--build", "--docdir"
	)
	if err != nil {
		return nil, err
	}

	if self.EditConfigureArgsCB != nil {
		ret, err = self.EditConfigureArgsCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_std_waf) BuilderActionConfigureEnvDef(
	log *logger.Logger,
) (environ.EnvVarEd, error) {

	ret, err := self.builder_std.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return nil, err
	}

	if self.EditConfigureEnvCB != nil {
		ret, err = self.EditConfigureEnvCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil

}

func (self *Builder_std_waf) BuilderActionConfigure(log *logger.Logger) error {

	// TODO: move this to separate builder tool, like cmake and autotools

	//	env := environ.NewFromStrings(os.Environ())

	//	nenv, err := self.BuilderActionConfigureEnvDef(log)
	//	if err != nil {
	//		return err
	//	}

	//	env.UpdateWith(nenv)

	//	args, err := self.BuilderActionConfigureArgsDef(log)
	//	if err != nil {
	//		return err
	//	}

	//	//	buildtype, err := self.BuilderActionConfigureDefBuildType(log)
	//	//	if err != nil {
	//	//		return err
	//	//	}

	//	args2 := make([]string, 0)
	//	args2 = append(args2, "--buildtype="+buildtype)
	//	args2 = append(args2, args...)
	//	args2 = append(args2, "../01.SOURCE")

	//	log.Info("meson arguments: " + strings.Join(args2, " "))
	//	for _, i := range args2 {
	//		log.Info("   " + i)

	//	}

	//	c := exec.Command(self.meson, args2...)
	//	c.Stdout = log.StdoutLbl()
	//	c.Stderr = log.StderrLbl()
	//	c.Dir = self.GetBuildingSiteCtl().GetDIR_BUILDING()
	//	c.Env = env.Strings()

	//	err = c.Run()
	//	if err != nil {
	//		return err
	//	}

	//	return nil
	return errors.New("TODO")
}
