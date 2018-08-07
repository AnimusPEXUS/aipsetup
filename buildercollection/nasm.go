package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["nasm"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_nasm(bs)
	}
}

type Builder_nasm struct {
	*Builder_std
}

func NewBuilder_nasm(bs basictypes.BuildingSiteCtlI) (*Builder_nasm, error) {

	self := new(Builder_nasm)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.EditDistributeDESTDIRCB = self.EditDistributeDESTDIR

	return self, nil
}

func (self *Builder_nasm) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{
			"--disable-werror",
		}...,
	)

	return ret, nil
}

func (self *Builder_nasm) EditDistributeDESTDIR(log *logger.Logger, ret string) (string, error) {
	return "INSTALLROOT", nil
}
