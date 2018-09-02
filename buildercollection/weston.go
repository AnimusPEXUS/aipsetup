package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["weston"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_weston(bs)
	}
}

type Builder_weston struct {
	*Builder_std
}

func NewBuilder_weston(bs basictypes.BuildingSiteCtlI) (*Builder_weston, error) {

	self := new(Builder_weston)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_weston) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-setuid-install",
			"--enable-xwayland",
			"--enable-x11-compositor",
			"--enable-drm-compositor",
			"--enable-wayland-compositor",
			"--enable-headless-compositor",
		}...,
	)

	return ret, nil
}
