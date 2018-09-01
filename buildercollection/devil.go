package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["devil"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_devil(bs)
	}
}

type Builder_devil struct {
	*Builder_std_cmake
}

func NewBuilder_devil(bs basictypes.BuildingSiteCtlI) (*Builder_devil, error) {
	self := new(Builder_devil)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "DevIL"), nil
	}

	//	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

//func (self *Builder_devil) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

//	ret = append(
//		ret,
//		[]string{
//			"-DWITH_STATIC_LIB=ON",
//			"-DBUILD_SHARED_LIBS=ON",
//			"-DENABLE_ZLIB_COMPRESSION=ON",
//		}...,
//	)

//	return ret, nil
//}
