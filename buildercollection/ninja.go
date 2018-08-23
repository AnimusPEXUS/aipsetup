package buildercollection

import (
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ninja"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_ninja(bs), nil
	}
}

type Builder_ninja struct {
	*Builder_std
}

func NewBuilder_ninja(bs basictypes.BuildingSiteCtlI) *Builder_ninja {
	self := new(Builder_ninja)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self
}

func (self *Builder_ninja) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	err := ret.ReplaceShort("build", self.BuilderActionBuild)
	if err != nil {
		return nil, err
	}

	err = ret.ReplaceShort("distribute", self.BuilderActionDistribute)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_ninja) BuilderActionBuild(log *logger.Logger) error {

	python3, err := self.bs.GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("python3")
	if err != nil {
		return err
	}

	c := exec.Command(python3, "./configure.py", "--bootstrap")
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.bs.GetDIR_SOURCE()

	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_ninja) BuilderActionDistribute(log *logger.Logger) error {

	dst_install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().
		CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_bin := path.Join(dst_install_prefix, "bin")

	err = os.MkdirAll(dst_install_prefix_bin, 0700)
	if err != nil {
		return err
	}

	err = filetools.CopyWithInfo(
		path.Join(self.bs.GetDIR_SOURCE(), "ninja"),
		path.Join(dst_install_prefix_bin, "ninja"),
		log,
	)
	if err != nil {
		return err
	}

	return nil
}
