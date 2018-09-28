package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["sdl2"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_sdl2(bs)
	}
}

type Builder_sdl2 struct {
	*Builder_std_cmake
}

func NewBuilder_sdl2(bs basictypes.BuildingSiteCtlI) (*Builder_sdl2, error) {
	self := new(Builder_sdl2)

	Builder_std_cmake, err := NewBuilder_std_cmake(bs)
	if err != nil {
		return nil, err
	}

	self.Builder_std_cmake = Builder_std_cmake

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_sdl2) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-DPULSEAUDIO=ON",
			"-DPULSEAUDIO_SHARED=ON",
			"-DALSA=OFF",
			"-DALSA_SHARED=OFF",
		}...,
	)

	return ret, nil
}
