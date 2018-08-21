package buildercollection

import (
	"os/exec"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["serf"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_serf(bs)
	}
}

type Builder_serf struct {
	*Builder_std
}

func NewBuilder_serf(bs basictypes.BuildingSiteCtlI) (*Builder_serf, error) {
	self := new(Builder_serf)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_serf) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

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

//func (self *Builder_serf) BuilderActionConfigure(
//	log *logger.Logger,
//) error {
//	return nil
//}

func (self *Builder_serf) BuilderActionBuild(
	log *logger.Logger,
) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	scons, err := calc.CalculateInstallPrefixExecutable("scons")
	if err != nil {
		return err
	}

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return err
	}

	install_libdir, err := calc.CalculateInstallLibDir()
	if err != nil {
		return err
	}

	params := []string{
		"PREFIX=" + install_prefix,
	}

	params = append(params, []string{
		"APR=" + install_prefix,
		"APU=" + install_prefix,
		"OPENSSL=" + install_prefix,
		"LIBDIR=" + install_libdir,
	}...,
	)

	comp_opt_map, err := calc.CalculateAutotoolsCompilerOptionsMap()
	if err != nil {
		return err
	}

	params = append(params, comp_opt_map.Strings()...)

	cmd := exec.Command(scons, params...)
	cmd.Dir = self.bs.GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_serf) BuilderActionDistribute(
	log *logger.Logger,
) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	scons, err := calc.CalculateInstallPrefixExecutable("scons")
	if err != nil {
		return err
	}

	//	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	//	if err != nil {
	//		return err
	//	}

	params := []string{
		"install",
		"--install-sandbox=" + self.bs.GetDIR_DESTDIR(),
	}

	cmd := exec.Command(scons, params...)
	cmd.Dir = self.bs.GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
