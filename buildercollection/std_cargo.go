package buildercollection

import (
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_cargo"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_cargo(bs)
	}
}

type Builder_std_cargo struct {
	*Builder_std
}

func NewBuilder_std_cargo(bs basictypes.BuildingSiteCtlI) (*Builder_std_cargo, error) {

	self := new(Builder_std_cargo)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_std_cargo) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("patch")
	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	err := ret.Replace(
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

func (self *Builder_std_cargo) BuilderActionDistribute(
	log *logger.Logger,
) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	cargo, err := calc.CalculateInstallPrefixExecutable("cargo")
	if err != nil {
		return err
	}

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_crates_toml := path.Join(dst_install_prefix, ".crates.toml")

	cmd := exec.Command(cargo, "install", "--root="+dst_install_prefix)
	cmd.Dir = self.bs.GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	err = os.Remove(dst_install_prefix_crates_toml)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}
