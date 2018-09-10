package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

// WARNING: freetype's cmake building variant is not for unixes

func init() {
	Index["freetype"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_freetype(bs)
	}
}

//type Builder_freetype struct {
//	*Builder_std_cmake
//}

//func NewBuilder_freetype(bs basictypes.BuildingSiteCtlI) (*Builder_freetype, error) {

//	self := new(Builder_freetype)

//	if t, err := NewBuilder_std_cmake(bs); err != nil {
//		return nil, err
//	} else {
//		self.Builder_std_cmake = t
//	}

//	self.EditConfigureArgsCB = self.EditConfigureArgs

//	return self, nil
//}

//func (self *Builder_freetype) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

//	ret = append(
//		ret,
//		[]string{
//			"-DFT_WITH_BZIP2=ON",
//			"-DFT_WITH_HARFBUZZ=ON",
//			"-DFT_WITH_PNG=ON",
//			"-DFT_WITH_ZLIB=ON",
//		}...,
//	)

//	return ret, nil
//}

type Builder_freetype struct {
	*Builder_std
}

func NewBuilder_freetype(bs basictypes.BuildingSiteCtlI) (*Builder_freetype, error) {

	self := new(Builder_freetype)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_freetype) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			//			"--with-bzip2=yes",
			//			"--with-png=yes",
			//			"--with-zlib=yes",

			// NOTE: harfbuzz <-> freetype is the circular dep. so it
			//       might be required to build freetype without
			//       harfbuzz once before building harfbuzz on it's
			//       own.
			//
			"--with-harfbuzz",
		}...,
	)

	return ret, nil
}
