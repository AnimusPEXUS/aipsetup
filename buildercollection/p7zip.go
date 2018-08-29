package buildercollection

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["p7zip"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_p7zip(bs)
	}
}

type Builder_p7zip struct {
	*Builder_std
}

func NewBuilder_p7zip(bs basictypes.BuildingSiteCtlI) (*Builder_p7zip, error) {

	self := new(Builder_p7zip)

	self.Builder_std = NewBuilder_std(bs)

	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_p7zip) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("build")

	err := ret.ReplaceShort("configure", self.BuilderActionConfigure)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_p7zip) BuilderActionConfigure(log *logger.Logger) error {

	mf1 := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "makefile.linux_any_cpu")
	mf2 := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "makefile.machine")

	os.Remove(mf2)

	err := os.Link(mf1, mf2)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_p7zip) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{"install"},
		[]string{"DESTDIR="},
	)
	if err != nil {
		return nil, err
	}

	env, err := self.BuilderActionDistributeEnvDef(log)
	if err != nil {
		return nil, err
	}

	CC := env.Get("CC", "gcc")
	CXX := env.Get("CXX", "g++")

	args := []string{
		"all3",
		"install",
		"CC=" + CC,
		"CXX=" + CXX,
		"DEST_HOME=" + install_prefix,
		"DEST_DIR=" + path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR()),
		"DEST_MAN=" + path.Join(install_prefix, "share", "man"),
	}

	ret = append(ret, args...)

	return ret, nil
}
