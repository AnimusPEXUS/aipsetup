package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["poppler"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_poppler(bs)
	}
}

type Builder_poppler struct {
	*Builder_std_cmake
}

func NewBuilder_poppler(bs basictypes.BuildingSiteCtlI) (*Builder_poppler, error) {

	self := new(Builder_poppler)

	t, err := NewBuilder_std_cmake(bs)
	if err != nil {
		return nil, err
	}
	self.Builder_std_cmake = t

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_poppler) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-DENABLE_XPDF_HEADERS=yes",
			"-DENABLE_ZLIB=yes",
			//			"--with-x",
			//			"--enable-xpdf-headers",
			//			"--enable-zlib",
			//			"--enable-libopenjpeg=openjpeg2",
		}...,
	)

	return ret, nil
}
