package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["pulseaudio"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_pulseaudio(bs)
	}
}

type Builder_pulseaudio struct {
	*Builder_std
}

func NewBuilder_pulseaudio(bs basictypes.BuildingSiteCtlI) (*Builder_pulseaudio, error) {
	self := new(Builder_pulseaudio)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditAutogenForceCB = func(log *logger.Logger, ret bool) (bool, error) {
		return true, nil
	}
	return self, nil
}

func (self *Builder_pulseaudio) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-database=gdbm",
			"--with-speex",
			"--enable-speex",
		}...,
	)
	return ret, nil
}
