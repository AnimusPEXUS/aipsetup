package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["llvm_components"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_llvm_components(bs)
	}
}

type Builder_llvm_components struct {
	*Builder_std_cmake
}

func NewBuilder_llvm_components(bs basictypes.BuildingSiteCtlI) (*Builder_llvm_components, error) {

	self := new(Builder_llvm_components)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_llvm_components) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-DCMAKE_BUILD_TYPE=Release",
		}...,
	)

	return ret, nil
}
