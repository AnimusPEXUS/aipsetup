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
	Index["findutils"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_findutils(bs)
	}
}

type Builder_findutils struct {
	*Builder_std
}

func NewBuilder_findutils(bs basictypes.BuildingSiteCtlI) (*Builder_findutils, error) {

	self := new(Builder_findutils)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_findutils) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"extract",
		"patch", self.BuilderActionPatch,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_findutils) BuilderActionPatch(log *logger.Logger) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageVersion == "4.6.0" {
		log.Info("findutils 4.6.0 requires patching")

		{
			c_files, err := filepath.Glob(path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "gl/lib/*.c"))
			if err != nil {
				return err
			}

			args := []string{"-i", "s/IO_ftrylockfile/IO_EOF_SEEN/"}
			args = append(args, c_files...)

			c := exec.Command("sed", args...)
			c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
			c.Stdout = log.StdoutLbl()
			c.Stderr = log.StderrLbl()

			err = c.Run()
			if err != nil {
				return err
			}
		}

		{
			c := exec.Command(
				"sed",
				"-i",
				"/unistd/a #include <sys/sysmacros.h>",
				"gl/lib/mountlist.c",
			)
			c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
			c.Stdout = log.StdoutLbl()
			c.Stderr = log.StderrLbl()

			err = c.Run()
			if err != nil {
				return err
			}
		}

		{
			stdio_impl_h := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "gl/lib/stdio-impl.h")

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
