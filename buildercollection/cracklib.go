package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["cracklib"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_cracklib(bs)
	}
}

type Builder_cracklib struct {
	*Builder_std
}

func NewBuilder_cracklib(bs basictypes.BuildingSiteCtlI) (*Builder_cracklib, error) {

	self := new(Builder_cracklib)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "src"), nil
	}

	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	return self, nil
}
