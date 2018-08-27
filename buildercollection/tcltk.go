package buildercollection

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["tcltk"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_tcltk(bs)
	}
}

type Builder_tcltk struct {
	*Builder_std
}

func NewBuilder_tcltk(bs basictypes.BuildingSiteCtlI) (*Builder_tcltk, error) {

	self := new(Builder_tcltk)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditConfigureDirCB = self.EditConfigureDir

	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self, nil
}

func (self *Builder_tcltk) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret, err := ret.AddActionAfterNameShort(
		"distribute",
		"make-links", self.BuilderActionMakeSymlinks,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_tcltk) EditConfigureDir(log *logger.Logger, ret string) (string, error) {
	return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "unix"), nil
}

func (self *Builder_tcltk) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
	return self.EditConfigureDir(log, ret)
}

func (self *Builder_tcltk) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	variant, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{
			"--datarootdir=",
			"--libdir=",
			"--docdir=",
		},
	)
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--enable-threads",
			"--enable-wince",

			// NOTE: tcl should be allways installed in 'lib'
			"--libdir=" + path.Join(install_prefix, "lib"),

			"--mandir=" + path.Join(install_prefix, "share", "man"),
		}...,
	)

	if variant == "64" {
		ret = append(
			ret,
			[]string{
				"--enable-64bit",
				"--enable-64bit-vis",
			}...,
		)
	}

	return ret, nil
}

func (self *Builder_tcltk) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"install-private-headers",
		}...,
	)

	return ret, nil
}

func (self *Builder_tcltk) BuilderActionMakeSymlinks(log *logger.Logger) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	bin_dir := path.Join(dst_install_prefix, "bin")

	files, err := ioutil.ReadDir(bin_dir)

	exe_name := "tclsh"
	if info.PackageName == "tk" {
		exe_name = "wish"
	}

	sym_name := path.Join(bin_dir, exe_name)

	real_name := ""
	for _, i := range files {
		if strings.HasPrefix(i.Name(), exe_name) {
			real_name = i.Name()
			break
		}
	}

	if real_name == "" {
		return errors.New("real exec name of " + exe_name + " not found")
	}

	err = os.Symlink(real_name, sym_name)
	if err != nil {
		return err
	}

	return nil
}
