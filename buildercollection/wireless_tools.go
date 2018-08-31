package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["wireless_tools"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_wireless_tools(bs), nil
	}
}

type Builder_wireless_tools struct {
	*Builder_std
}

func NewBuilder_wireless_tools(bs basictypes.BuildingSiteCtlI) *Builder_wireless_tools {

	self := new(Builder_wireless_tools)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "wireless_tools"), nil
	}

	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	self.EditDistributeArgsCB = self.EditDistributeArgs

	//	self.EditDistributeDESTDIRCB = func(log *logger.Logger, ret string) (string, error) {
	//		return "PREFIX", nil
	//	}

	return self
}

func (self *Builder_wireless_tools) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	return ret, nil
}

func (self *Builder_wireless_tools) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

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
			"CC=" + CC,
			"CXX=" + CXX,
			"PREFIX=" + dst_install_prefix,
			"INSTALL_MAN=" + path.Join(dst_install_prefix, "share", "man"),
		}...,
	)
	return ret, nil
}
