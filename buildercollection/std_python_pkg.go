package buildercollection

import (
	"os/exec"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_python2_pkg"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_python_pkg(bs, "2")
	}

	Index["std_python3_pkg"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_python_pkg(bs, "3")
	}
}

type Builder_std_python_pkg struct {
	*Builder_std

	python string
}

func NewBuilder_std_python_pkg(
	bs basictypes.BuildingSiteCtlI,
	python_number string,
) (*Builder_std_python_pkg, error) {

	self := new(Builder_std_python_pkg)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	if python, err := self.bs.GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("python" + python_number); err != nil {
		return nil, err
	} else {
		self.python = python
	}

	return self, nil
}

func (self *Builder_std_python_pkg) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

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

func (self *Builder_std_python_pkg) BuilderActionBuild(log *logger.Logger) error {

	c := exec.Command(self.python, "./setup.py", "--bootstrap")
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.bs.GetDIR_SOURCE()

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_std_python_pkg) BuilderActionDistribute(log *logger.Logger) error {

	c := exec.Command(self.python, "./setup.py", "install")
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.bs.GetDIR_SOURCE()

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
