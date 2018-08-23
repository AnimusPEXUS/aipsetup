package buildercollection

// TODO: Scons builder requires many and many thinking

import (
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["scons"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_scons(bs), nil
	}
}

type Builder_scons struct {
	*Builder_std
}

func NewBuilder_scons(bs basictypes.BuildingSiteCtlI) *Builder_scons {
	self := new(Builder_scons)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self
}

func (self *Builder_scons) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

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

func (self *Builder_scons) BuilderActionBuild(
	log *logger.Logger,
) error {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	python, err := calc.CalculateInstallPrefixExecutable("python2")
	if err != nil {
		return err
	}

	log.Info("python in use: " + python)

	cmd := exec.Command(
		python,
		[]string{
			"./bootstrap.py",
			self.GetBuildingSiteCtl().GetDIR_SOURCE(),
			"build",
			//			"scons",
		}...,
	)
	cmd.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_scons) BuilderActionDistribute(
	log *logger.Logger,
) error {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	python, err := calc.CalculateInstallPrefixExecutable("python2")
	if err != nil {
		return err
	}

	dst_instll_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	log.Info("python in use: " + python)

	cmd := exec.Command(
		python,
		[]string{
			"setup.py",
			"install",
			"--prefix=" + dst_instll_prefix,
		}...,
	)
	cmd.Dir = path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "build", "scons")
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	dst_instll_prefix_man := path.Join(dst_instll_prefix, "man")
	dst_instll_prefix_share := path.Join(dst_instll_prefix, "share")
	dst_instll_prefix_share_man := path.Join(dst_instll_prefix_share, "man")

	err = os.MkdirAll(dst_instll_prefix_share, 0700)
	if err != nil {
		return err
	}

	err = os.Rename(dst_instll_prefix_man, dst_instll_prefix_share_man)
	if err != nil {
		return err
	}

	return nil
}
