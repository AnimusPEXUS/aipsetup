package buildercollection

import (
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
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

	for i := len(ret) - 1; i != -1; i -= 1 {

		for _, j := range []string{"--enable-shared"} {
			if ret[i] == j {
				ret = append(ret[:i], ret[i+1:]...)
			}
		}

		for _, j := range []string{"CC", "CXX", "GCC"} {
			if strings.HasPrefix(ret[i], j+"=") {
				ret = append(ret[:i], ret[i+1:]...)
			}
		}
	}

	return ret, nil
}

func (self *Builder_qemu) EditConfigureEnv(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error) {
	for _, i := range []string{"CC", "CXX", "GCC"} {
		ret.Del(i)
	}
	return ret, nil
}
