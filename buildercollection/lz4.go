package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["lz4"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_lz4(bs)
	}
}

type Builder_lz4 struct {
	*Builder_std_cmake
}

func NewBuilder_lz4(bs basictypes.BuildingSiteCtlI) (*Builder_lz4, error) {

	self := new(Builder_lz4)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	//	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir
	self.EditConfigureDirCB = self.EditConfigureDir

	return self, nil
}

//func (self *Builder_lz4) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
//	return self.GetBuildingSiteCtl().GetDIR_BUILDING(), nil
//}

func (self *Builder_lz4) EditConfigureDir(log *logger.Logger, ret string) (string, error) {
	return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "contrib", "cmake_unofficial"), nil
}
