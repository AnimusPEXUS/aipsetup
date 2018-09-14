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
	*Builder_std
}

func NewBuilder_colord(bs basictypes.BuildingSiteCtlI) (*Builder_colord, error) {

	self := new(Builder_colord)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_colord) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-argyllcms-sensor",
			"--enable-vala",
			"--with-daemon-user=colord",
		}...,
	)

	return ret, nil
}
