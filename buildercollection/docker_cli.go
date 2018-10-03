package buildercollection

import (
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["docker_cli"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_docker_cli(bs)
	}
}

type Builder_docker_cli struct {
	*Builder_std
}

func NewBuilder_docker_cli(bs basictypes.BuildingSiteCtlI) (*Builder_docker_cli, error) {

	self := new(Builder_docker_cli)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_docker_cli) EditActions(ret basictypes.BuilderActions) (
	basictypes.BuilderActions,
	error,
) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")
	ret = ret.Remove("distribute")

	new_actions := basictypes.BuilderActions{
		&basictypes.BuilderAction{"build", self.BuilderActionBuild},
		&basictypes.BuilderAction{"distribute", self.BuilderActionDistribute},
	}

	ret, err := ret.AddActionsBeforeName(new_actions, "prepack")
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_docker_cli) BuilderActionBuild(log *logger.Logger) error {

	c := exec.Command("make", "-f", "docker.Makefile", "binary")
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_docker_cli) BuilderActionDistribute(log *logger.Logger) error {

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	build_dir := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "build")

	dst_install_prefix_bin := path.Join(dst_install_prefix, "bin")

	err = filetools.CopyTree(
		build_dir,
		dst_install_prefix_bin,
		false,
		true,
		false,
		true,
		log,
		true,
		true,
		filetools.CopyWithInfo,
	)
	if err != nil {
		return err
	}

	return nil
}
