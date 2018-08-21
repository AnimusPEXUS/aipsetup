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
	Index["gzip"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_gzip(bs)
	}
}

type Builder_gzip struct {
	*Builder_std
}

func NewBuilder_gzip(bs basictypes.BuildingSiteCtlI) (*Builder_gzip, error) {

	self := new(Builder_gzip)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_gzip) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"extract",
		"patch", self.BuilderActionPatch,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_gzip) BuilderActionPatch(log *logger.Logger) error {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageVersion == "1.9" {
		log.Info("gzip 1.9 requires patching")

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
