package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["harfbuzz"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_harfbuzz(bs)
	}
}

//type Builder_harfbuzz struct {
//	*Builder_std_cmake
//}

//func NewBuilder_harfbuzz(bs basictypes.BuildingSiteCtlI) (*Builder_harfbuzz, error) {
//	self := new(Builder_harfbuzz)

//	if t, err := NewBuilder_std_cmake(bs); err != nil {
//		return nil, err
//	} else {
//		self.Builder_std_cmake = t
//	}

//	//	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
//	//		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "DevIL"), nil
//	//	}

//	self.EditConfigureArgsCB = self.EditConfigureArgs

//	return self, nil
//}

//func (self *Builder_harfbuzz) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

//	ret = append(
//		ret,
//		[]string{
//			//			"-DHB_HAVE_FREETYPE=on",
//			//			"-DHB_HAVE_GLIB=on",
//			//			"-DHB_HAVE_GOBJECT=on",
//			//			"-DHB_HAVE_GRAPHITE2=on",
//			//			"-DHB_HAVE_ICU=on",
//			//			"-DHB_HAVE_INTROSPECTION=on",
//		}...,
//	)

//	return ret, nil
//}

type Builder_harfbuzz struct {
	*Builder_std
}

func NewBuilder_harfbuzz(bs basictypes.BuildingSiteCtlI) (*Builder_harfbuzz, error) {

	self := new(Builder_harfbuzz)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_harfbuzz) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--without-freetype",
			//			"--with-bzip2=yes",
			//			"--with-harfbuzz=yes",
			//			"--with-png=yes",
			//			"--with-zlib=yes",
		}...,
	)

	return ret, nil
}
