package buildercollection

import (
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["infozip"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_infozip(bs)
	}
}

type Builder_infozip struct {
	*Builder_std
}

func NewBuilder_infozip(bs basictypes.BuildingSiteCtlI) (*Builder_infozip, error) {

	self := new(Builder_infozip)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_infozip) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	err := ret.ReplaceShort("distribute", self.BuilderActionDistribute)
	if err != nil {
		return nil, err
	}

	ret, err = ret.AddActionAfterNameShort(
		"distribute",
		"after-distribute", self.BuilderActionAfterDistribute,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_infozip) BuilderActionDistribute(log *logger.Logger) error {

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	c := exec.Command(
		"make",
		"-f", "unix/Makefile",
		"generic",
		"install",
		"prefix="+dst_install_prefix,
	)
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()

	err = c.Run()
	if err != nil {
		return err
	}

	return nil

}

func (self *Builder_infozip) BuilderActionAfterDistribute(log *logger.Logger) error {

	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_man := path.Join(dst_install_prefix, "man")

	dst_share := path.Join(dst_install_prefix, "share")
	dst_share_man := path.Join(dst_share, "man")

	err = os.Mkdir(dst_share, 0700)
	if err != nil {
		return err
	}

	err = os.Rename(dst_man, dst_share_man)
	if err != nil {
		return err
	}

	return nil
}
