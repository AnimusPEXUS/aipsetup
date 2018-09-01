package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libjpeg_turbo"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libjpeg_turbo(bs)
	}
}

type Builder_libjpeg_turbo struct {
	*Builder_std_cmake
}

func NewBuilder_libjpeg_turbo(bs basictypes.BuildingSiteCtlI) (*Builder_libjpeg_turbo, error) {

	self := new(Builder_libjpeg_turbo)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_libjpeg_turbo) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-DWITH_JPEG7=yes",
			"-DWITH_JPEG8=yes",
		}...,
	)

	return ret, nil
}
