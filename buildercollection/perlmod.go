package buildercollection

import (
	"os/exec"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["perlmod"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_perlmod(bs), nil
	}
}

type Builder_perlmod struct {
	*Builder_std
}

func NewBuilder_perlmod(bs basictypes.BuildingSiteCtlI) *Builder_perlmod {

	self := new(Builder_perlmod)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self
}

func (self *Builder_perlmod) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("autogen")

	err := ret.Replace(
		"configure",
		&basictypes.BuilderAction{
			Name:     "configure",
			Callable: self.BuilderActionConfigure,
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_perlmod) BuilderActionConfigure(log *logger.Logger) error {

	perl, err := self.bs.GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("perl")
	if err != nil {
		return err
	}

	c := exec.Command(perl, "Makefile.PL")
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.bs.GetDIR_SOURCE()

	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}
