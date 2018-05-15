package buildercollection

import (
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["autogen"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_autogen(bs)
	}
}

type Builder_autogen struct {
	BuilderStdAutotools
}

func NewBuilder_autogen(bs basictypes.BuildingSiteCtlI) (*Builder_autogen, error) {
	self := new(Builder_autogen)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_autogen) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	inst_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	guile_prefix := inst_prefix

	guilde_config := path.Join(guile_prefix, "bin", "guile-config")

	var guile_cflags, guile_libs string

	{
		c := exec.Command(guilde_config, "compile")
		out, err := c.Output()
		if err != nil {
			return nil, err
		}
		guile_cflags = string(out)
	}

	{
		c := exec.Command(guilde_config, "link")
		out, err := c.Output()
		if err != nil {
			return nil, err
		}
		guile_libs = string(out)
	}

	ret = append(
		ret,
		[]string{
			"--with-libguilde=" + guile_prefix,
			"--with-libguile-cflags=" + guile_cflags,
			"--with-libguile-libs={}" + guile_libs,
		}...,
	)

	return ret, nil
}
