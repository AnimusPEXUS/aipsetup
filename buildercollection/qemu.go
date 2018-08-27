package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["qemu"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_qemu(bs), nil
	}
}

type Builder_qemu struct {
	*Builder_std
}

func NewBuilder_qemu(bs basictypes.BuildingSiteCtlI) *Builder_qemu {

	self := new(Builder_qemu)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_qemu) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-virglrenderer",

			"--disable-gtk",
			// "--with-gtkabi=3.0",

			// "--cpu=x86_64",
			"--audio-drv-list=pa",

			"--enable-sdl",
			"--with-sdlabi=2.0",

			"--enable-kvm",
			"--enable-system",
			"--enable-user",
			"--enable-linux-user",
			// "--enable-bsd-user",
			// "--enable-guest-base",
		}...,
	)

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{"--enable-shared"},
		[]string{"CC=", "CXX=", "GCC="},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_qemu) EditConfigureEnv(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error) {
	for _, i := range []string{"CC", "CXX", "GCC"} {
		ret.Del(i)
	}
	return ret, nil
}
