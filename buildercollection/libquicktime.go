package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libquicktime"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libquicktime(bs)
	}
}

type Builder_libquicktime struct {
	*Builder_std
}

func NewBuilder_libquicktime(bs basictypes.BuildingSiteCtlI) (*Builder_libquicktime, error) {

	self := new(Builder_libquicktime)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_libquicktime) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-libdv",
			"--with-x264",
		}...,
	)

	return ret, nil
}
