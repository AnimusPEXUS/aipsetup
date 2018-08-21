package buildercollection

import (
	"io/ioutil"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["m4"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_m4(bs)
	}
}

type Builder_m4 struct {
	*Builder_std
}

func NewBuilder_m4(bs basictypes.BuildingSiteCtlI) (*Builder_m4, error) {

	self := new(Builder_m4)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_m4) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"extract",
		"patch", self.BuilderActionPatch,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_m4) BuilderActionPatch(log *logger.Logger) error {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageVersion == "1.4.18" {
		log.Info("m4 1.4.18 requires patching")

		{
			c_files, err := filepath.Glob(path.Join(self.bs.GetDIR_SOURCE(), "lib/*.c"))
			if err != nil {
				return err
			}

			args := []string{"-i", "s/IO_ftrylockfile/IO_EOF_SEEN/"}
			args = append(args, c_files...)

			c := exec.Command("sed", args...)
			c.Dir = self.bs.GetDIR_SOURCE()
			c.Stdout = log.StdoutLbl()
			c.Stderr = log.StderrLbl()

			err = c.Run()
			if err != nil {
				return err
			}
		}

		{
			stdio_impl_h := path.Join(self.bs.GetDIR_SOURCE(), "lib/stdio-impl.h")

			data, err := ioutil.ReadFile(stdio_impl_h)
			if err != nil {
				return err
			}

			data = append(
				data,
				[]byte("#define _IO_IN_BACKUP 0x100\n")...,
			)

			err = ioutil.WriteFile(stdio_impl_h, data, 0700)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
