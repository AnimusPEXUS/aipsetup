package buildercollection

import (
	"os/exec"

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

	self.EditActionsCB = self.EditActions
	self.EditDistributeDESTDIRCB = self.EditDistributeDESTDIR

	return self, nil
}

func (self *Builder_nasm) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"extract",
		"patch", self.BuilderActionPatch,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_nasm) BuilderActionPatch(log *logger.Logger) error {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageVersion == "2.13.3" {
		log.Info("nasm 2.13.3 requires patching")

		{

			args := []string{
				"-e", "/seg_init/d",
				"-e", "s/pure_func seg_alloc/seg_alloc/",
				"-i", "include/nasmlib.h",
			}

			c := exec.Command("sed", args...)
			c.Dir = self.bs.GetDIR_SOURCE()
			c.Stdout = log.StdoutLbl()
			c.Stderr = log.StderrLbl()

			err = c.Run()
			if err != nil {
				return err
			}
		}

	}

	return nil
}

//func (self *Builder_nasm) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
//	ret = append(
//		ret,
//		[]string{
//			"--disable-werror",
//		}...,
//	)

//	return ret, nil
//}

func (self *Builder_nasm) EditDistributeDESTDIR(log *logger.Logger, ret string) (string, error) {
	return "INSTALLROOT", nil
}
