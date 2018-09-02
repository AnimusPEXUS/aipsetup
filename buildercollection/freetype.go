package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["freetype"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_freetype(bs)
	}
}

type Builder_freetype struct {
	*Builder_std_cmake
}

func NewBuilder_freetype(bs basictypes.BuildingSiteCtlI) (*Builder_freetype, error) {

	self := new(Builder_freetype)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_freetype) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-DFT_WITH_BZIP2=ON",
			"-DFT_WITH_HARFBUZZ=ON",
			"-DFT_WITH_PNG=ON",
			"-DFT_WITH_ZLIB=ON",
		}...,
	)

	return ret, nil
}
