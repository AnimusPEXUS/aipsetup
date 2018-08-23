package buildercollection

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/systemtriplet"
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

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	//	ret = ret.Remove("build")

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

func (self *Builder_std_cargo) BuilderActionBuild(
	log *logger.Logger,
) error {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	cargo, err := calc.CalculateInstallPrefixExecutable("cargo")
	if err != nil {
		return err
	}

	log.Info("cargo is: " + cargo)

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	//	target := info.HostArch

	tri, err := systemtriplet.NewFromString(info.HostArch)
	if err != nil {
		return err
	}

	tri.Company = "unknown"

	//	target := "x86_64-unknown-linux-gnu"

	args := []string{"build", "--target", tri.String()}

	log.Info("cargo args: " + strings.Join(args, " "))

	cmd := exec.Command(cargo, args...)
	cmd.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_std_cargo) BuilderActionDistribute(
	log *logger.Logger,
) error {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	cargo, err := calc.CalculateInstallPrefixExecutable("cargo")
	if err != nil {
		return err
	}

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_crates_toml := path.Join(dst_install_prefix, ".crates.toml")

	args := []string{"install", "--root", dst_install_prefix}

	log.Info("cargo args: " + strings.Join(args, " "))

	cmd := exec.Command(cargo, args...)
	cmd.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
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
