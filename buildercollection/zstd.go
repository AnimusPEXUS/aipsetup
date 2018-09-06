package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["zstd"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_zstd(bs), nil
	}
}

type Builder_zstd struct {
	*Builder_std
}

func NewBuilder_zstd(bs basictypes.BuildingSiteCtlI) *Builder_zstd {
	self := new(Builder_zstd)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

func (self *Builder_zstd) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autoconf")
	ret = ret.Remove("configure")

	return ret, nil
}

func (self *Builder_zstd) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{"PREFIX=" + install_prefix}...,
	)

	return ret, nil
}
