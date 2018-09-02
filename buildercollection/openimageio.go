package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openimageio"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_openimageio(bs)
	}
}

type Builder_openimageio struct {
	*Builder_std_cmake
}

func NewBuilder_openimageio(bs basictypes.BuildingSiteCtlI) (*Builder_openimageio, error) {

	self := new(Builder_openimageio)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_openimageio) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-DSTOP_ON_WARNING=OFF",
		}...,
	)

	return ret, nil
}
