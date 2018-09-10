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
//			"-DFT_WITH_BZIP2=y",
//			"-DFT_WITH_HARFBUZZ=y",
//			"-DFT_WITH_PNG=y",
//			"-DFT_WITH_ZLIB=y",
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

	self.EditBuildArgsCB = self.EditBuildArgs

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self, nil
}

func (self *Builder_freetype) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-bzip2",
			"--with-png",
			"--with-zlib",
			"--with-harfbuzz",

			// NOTE: harfbuzz <-> freetype is the circular dep. so it
			//       might be required to build freetype without
			//       harfbuzz once before building harfbuzz on it's
			//       own.
			//
			//			"--without-harfbuzz",

			//"WIN32=",
			//			"RC=",
		}...,
	)

	return ret, nil
}

func (self *Builder_freetype) EditBuildArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{"RC="}...,
	)
	return ret, nil
}

func (self *Builder_freetype) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{"RC="}...,
	)
	return ret, nil
}
