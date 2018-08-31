package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["squashfs"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_squashfs(bs), nil
	}
}

type Builder_squashfs struct {
	*Builder_std
}

func NewBuilder_squashfs(bs basictypes.BuildingSiteCtlI) *Builder_squashfs {

	self := new(Builder_squashfs)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "squashfs-tools"), nil
	}

	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	self.EditDistributeArgsCB = self.EditDistributeArgs

	//	self.EditDistributeDESTDIRCB = func(log *logger.Logger, ret string) (string, error) {
	//		return "PREFIX", nil
	//	}

	return self
}

func (self *Builder_squashfs) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	return ret, nil
}

func (self *Builder_squashfs) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{"install"},
		[]string{"DESTDIR="},
	)
	if err != nil {
		return nil, err
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return nil, err
	}

	env, err := self.BuilderActionDistributeEnvDef(log)
	if err != nil {
		return nil, err
	}

	CC := env.Get("CC", "gcc")
	CXX := env.Get("CXX", "g++")

	ret = append(
		ret,
		[]string{
			"all",
			"install",
			"GZIP_SUPPORT=1",
			"XZ_SUPPORT=1",
			"LZO_SUPPORT=1",
			"CC=" + CC,
			"CXX=" + CXX,
			"INSTALL_DIR=" + path.Join(dst_install_prefix, "bin"),
		}...,
	)
	return ret, nil
}
