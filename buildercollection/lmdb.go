package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["lmdb"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_lmdb(bs)
	}
}

type Builder_lmdb struct {
	*Builder_std
}

func NewBuilder_lmdb(bs basictypes.BuildingSiteCtlI) (*Builder_lmdb, error) {

	self := new(Builder_lmdb)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "libraries/liblmdb"), nil
	}

	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self, nil
}

func (self *Builder_lmdb) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	return ret, nil
}
func (self *Builder_lmdb) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{
			"prefix=",
			"DESTDIR=",
		},
	)
	if err != nil {
		return nil, err
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"prefix=" + dst_install_prefix,
			"PREFIX=" + dst_install_prefix,
		}...,
	)

	return ret, nil
}
