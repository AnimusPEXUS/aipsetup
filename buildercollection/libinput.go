package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libinput"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libinput(bs)
	}
}

type Builder_libinput struct {
	*Builder_std_meson
}

func NewBuilder_libinput(bs basictypes.BuildingSiteCtlI) (*Builder_libinput, error) {

	self := new(Builder_libinput)

	Builder_std_meson, err := NewBuilder_std_meson(bs)
	if err != nil {
		return nil, err
	}

	self.Builder_std_meson = Builder_std_meson

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_libinput) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-Dlibwacom=false",
			"-Ddocumentation=false",
		}...,
	)

	return ret, nil
}
