package buildercollection

import (
	"fmt"
	"os/exec"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["mongodb"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_mongodb(bs)
	}
}

type Builder_mongodb struct {
	*Builder_std
}

func NewBuilder_mongodb(bs basictypes.BuildingSiteCtlI) (*Builder_mongodb, error) {
	self := new(Builder_mongodb)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_mongodb) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	//	err := ret.Replace(
	//		"configure",
	//		&basictypes.BuilderAction{
	//			Name:     "configure",
	//			Callable: self.BuilderActionConfigure,
	//		},
	//	)
	//	if err != nil {
	//		return nil, err
	//	}

	err := ret.Replace(
		"build",
		&basictypes.BuilderAction{
			Name:     "build",
			Callable: self.BuilderActionBuild,
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

//func (self *Builder_mongodb) BuilderActionConfigure(
//	log *logger.Logger,
//) error {
//	return nil
//}

func (self *Builder_mongodb) BuilderActionBuild(
	log *logger.Logger,
) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	scons, err := calc.CalculateInstallPrefixExecutable("scons")
	if err != nil {
		return err
	}

	//	python3, err := self.GetBuildingSiteCtl().
	//		GetBuildingSiteValuesCalculator().CalculateInstallPrefixExecutable("python3")
	//	if err != nil {
	//		return err
	//	}

	//	install_prefix, err := calc.CalculateInstallPrefix()
	//	if err != nil {
	//		return err
	//	}

	//	install_libdir, err := calc.CalculateInstallLibDir()
	//	if err != nil {
	//		return err
	//	}

	params := []string{}

	params = append(params, []string{
		"MONGO_VERSION=" + info.PackageVersion,
		//		"--use-system-boost",
		//		"--use-system-sqlite",
		//		"--use-system-zlib",
		//		"--use-system-icu",
		"mongod",
		"mongo",
	}...,
	)

	comp_opt_map, err := calc.CalculateAutotoolsCompilerOptionsMap()
	if err != nil {
		return err
	}

	params = append(params, comp_opt_map.Strings()...)

	log.Info("SCons params: " + fmt.Sprintf("%v", params))

	cmd := exec.Command(scons, params...)
	cmd.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_mongodb) BuilderActionDistribute(
	log *logger.Logger,
) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	scons, err := calc.CalculateInstallPrefixExecutable("scons")
	if err != nil {
		return err
	}

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	params := []string{
		"install",
		"--prefix=" + dst_install_prefix,
	}

	params = append(params, []string{
		"MONGO_VERSION=" + info.PackageVersion,
	}...,
	)

	log.Info("SCons params: " + fmt.Sprintf("%v", params))

	cmd := exec.Command(scons, params...)
	cmd.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
