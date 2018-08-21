package buildercollection

import (
	"os/exec"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["make"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_make(bs)
	}
}

type Builder_make struct {
	*Builder_std
}

func NewBuilder_make(bs basictypes.BuildingSiteCtlI) (*Builder_make, error) {

	self := new(Builder_make)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_make) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"extract",
		"patch", self.BuilderActionPatch,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_make) BuilderActionPatch(log *logger.Logger) error {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageVersion == "4.2.1" {
		log.Info("make 4.2.1 requires glob/glob.c patching")
		c := exec.Command("sed", "-i", "211,217 d; 219,229 d; 232 d", "glob/glob.c")
		c.Dir = self.bs.GetDIR_SOURCE()
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()

		err = c.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_make) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	if !info.ThisIsCrossbuilder() && !info.ThisIsCrossbuilding() {
		ret = append(
			ret,
			[]string{
				"--without-guile",
			}...,
		)
	}

	return ret, nil
}
