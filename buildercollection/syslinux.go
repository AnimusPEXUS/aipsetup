package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["syslinux"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_syslinux(bs)
	}
}

type Builder_syslinux struct {
	*Builder_std
}

func NewBuilder_syslinux(bs basictypes.BuildingSiteCtlI) (*Builder_syslinux, error) {

	self := new(Builder_syslinux)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditBuildConcurentJobsCountCB = self.EditBuildConcurentJobsCount

	return self, nil
}

func (self *Builder_syslinux) EditBuildConcurentJobsCount(log *logger.Logger, ret int) int {
	return 1
}

func (self *Builder_syslinux) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("configure")
	ret = ret.Remove("autogen")
	ret = ret.Remove("build")

	return ret, nil
}

func (self *Builder_syslinux) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = []string{
		"bios", "efi32", "efi64",
		"installer",
		"install",
		"INSTALLROOT=" + self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
	}

	return ret, nil
}
