package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["colord"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_colord(bs)
	}
}

type Builder_colord struct {
	*Builder_std_meson
}

func NewBuilder_colord(bs basictypes.BuildingSiteCtlI) (*Builder_colord, error) {

	self := new(Builder_colord)

	Builder_std_meson, err := NewBuilder_std_meson(bs)
	if err != nil {
		return nil, err
	}

	self.Builder_std_meson = Builder_std_meson

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_colord) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-Dargyllcms_sensor=false",
			"-Dvapi=true",
			"-Ddaemon_user=colord",
			"-Dman=false",
			"-Ddocs=false",
		}...,
	)

	return ret, nil
}
