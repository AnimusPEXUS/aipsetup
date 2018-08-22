package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libssh"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_libssh(bs)
	}
}

type Builder_libssh struct {
	*Builder_std_cmake
}

func NewBuilder_libssh(bs basictypes.BuildingSiteCtlI) (*Builder_libssh, error) {
	self := new(Builder_libssh)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_libssh) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-DWITH_STATIC_LIB=ON",
			"-DBUILD_SHARED_LIBS=ON",
			"-DENABLE_ZLIB_COMPRESSION=ON",
		}...,
	)

	return ret, nil
}
