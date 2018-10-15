package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libnet"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_libnet(bs), nil
	}
}

type Builder_libnet struct {
	*Builder_std

	makefile_flags []string
}

func NewBuilder_libnet(bs basictypes.BuildingSiteCtlI) *Builder_libnet {

	self := new(Builder_libnet)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "libnet"), nil
	}

	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	return self
}
