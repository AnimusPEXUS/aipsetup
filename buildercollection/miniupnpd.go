package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["miniupnpd"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_miniupnpd(bs)
	}
}

type Builder_miniupnpd struct {
	*Builder_std
}

func NewBuilder_miniupnpd(bs basictypes.BuildingSiteCtlI) (*Builder_miniupnpd, error) {

	self := new(Builder_miniupnpd)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditBuildMakefileNameCB = func(log *logger.Logger, ret string) (string, error) {
		return "Makefile.linux", nil
	}

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "miniupnpd"), nil
	}

	//	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	self.EditConfigureWorkingDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "miniupnpd"), nil
	}

	return self, nil
}

func (self *Builder_miniupnpd) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	return ret, nil
}
