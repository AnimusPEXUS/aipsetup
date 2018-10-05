package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gst_plugins"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gst_plugins(bs)
	}
}

type Builder_gst_plugins struct {
	*Builder_std
}

func NewBuilder_gst_plugins(bs basictypes.BuildingSiteCtlI) (*Builder_gst_plugins, error) {
	self := new(Builder_gst_plugins)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_gst_plugins) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("autogen")
	return ret, nil
}

func (self *Builder_gst_plugins) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{
			"--docdir=",
		},
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
