package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["fribidi"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_fribidi(bs)
	}
}

type Builder_fribidi struct {
	*Builder_std_meson
}

func NewBuilder_fribidi(bs basictypes.BuildingSiteCtlI) (*Builder_fribidi, error) {

	self := new(Builder_fribidi)

	if t, err := NewBuilder_std_meson(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_meson = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_fribidi) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-Ddocs=false",
		}...,
	)

	return ret, nil
}
