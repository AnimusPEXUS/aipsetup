package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["liboop"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_liboop(bs)
	}
}

type Builder_liboop struct {
	*Builder_std
}

func NewBuilder_liboop(bs basictypes.BuildingSiteCtlI) (*Builder_liboop, error) {

	self := new(Builder_liboop)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_liboop) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--without-glib",
		}...,
	)

	return ret, nil
}
