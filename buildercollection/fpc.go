package buildercollection

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["fpc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_fpc(bs)
	}
}

type Builder_fpc struct {
	*Builder_std
}

func NewBuilder_fpc(bs basictypes.BuildingSiteCtlI) (*Builder_fpc, error) {
	self := new(Builder_fpc)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_fpc) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("build")

	err := ret.ReplaceShort("configure", self.BuilderActionConfigure)
	if err != nil {
		return nil, err
	}

	err = ret.ReplaceShort("distribute", self.BuilderActionDistribute)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_fpc) BuilderActionConfigure(log *logger.Logger) error {
	c := exec.Command("fpcmake", "-Tall")
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *Builder_fpc) BuilderActionDistribute(log *logger.Logger) error {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	settings := map[string]string{
		"CPU_TARGET":     "i386",
		"PREFIX":         install_prefix,
		"INSTALL_PREFIX": dst_install_prefix,
		"AS":             "as --32",
		"LD":             "ld -A elf_i386",
		"ppc":            "ppc386",
	}

	{
		vari, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateMultilibVariant()
		if err != nil {
			return err
		}

		if vari == "64" {
			settings = map[string]string{
				"CPU_TARGET":     "x86_64",
				"PREFIX":         install_prefix,
				"INSTALL_PREFIX": dst_install_prefix,
				"AS":             "as --64",
				"LD":             "ld -A elf_x86_64",
				"ppc":            "ppcx64",
			}
		}
	}

	cmd := []string{
		"clean",
		"fpc_info",
		"all",
		"install",
		"CPU_TARGET=" + settings["CPU_TARGET"],
		"PREFIX=" + settings["PREFIX"],
		"INSTALL_PREFIX=" + settings["INSTALL_PREFIX"],
		"AS=" + settings["AS"],
		"LD=" + settings["LD"],
	}

	tmp_dir := self.GetBuildingSiteCtl().GetDIR_TEMP()

	tmp_bin_dir := path.Join(tmp_dir, "bin")

	err = os.MkdirAll(tmp_bin_dir, 0700)
	if err != nil {
		return err
	}

	ld_script_filename := path.Join(
		tmp_bin_dir,
		settings["CPU_TARGET"]+"-linux-ld",
	)

	script := `#!/bin/bash

` + settings["LD"] + ` $@

`

	err = ioutil.WriteFile(ld_script_filename, []byte(script), 0700)
	if err != nil {
		return err
	}

	env := environ.NewFromStrings(os.Environ())

	env.Set("PATH", env.Get("PATH", "/bin")+":"+tmp_bin_dir)

	c := exec.Command("make", cmd...)
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Env = env.Strings()
	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}
