package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["icu4c"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_icu4c(bs), nil
	}
}

type Builder_icu4c struct {
	*Builder_std
}

func NewBuilder_icu4c(bs basictypes.BuildingSiteCtlI) *Builder_icu4c {

	self := new(Builder_icu4c)

	self.Builder_std = NewBuilder_std(bs)
	//	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditConfigureDirCB = self.EditConfigureDir
	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir

	return self
}

func (self *Builder_icu4c) EditConfigureDir(log *logger.Logger, ret string) (string, error) {
	return path.Join(self.bs.GetDIR_SOURCE(), "source"), nil
}

func (self *Builder_icu4c) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
	return self.EditConfigureDir(log, ret)
}
