package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openjpeg"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_openjpeg(bs)
	}
}

type Builder_openjpeg struct {
	*Builder_std_cmake
}

func NewBuilder_openjpeg(bs basictypes.BuildingSiteCtlI) (*Builder_openjpeg, error) {
	self := new(Builder_openjpeg)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_openjpeg) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	libdir_name, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateMainMultiarchLibDirName()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"-DOPENJPEG_INSTALL_LIB_DIR=" + libdir_name,
		}...,
	)

	return ret, nil
}
