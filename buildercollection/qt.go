package buildercollection

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/systemtriplet"
)

func init() {
	Index["qt"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_qt(bs)
	}
}

type Builder_qt struct {
	*Builder_std

	qt_major_version int
}

func NewBuilder_qt(bs basictypes.BuildingSiteCtlI) (*Builder_qt, error) {
	self := new(Builder_qt)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditDistributeDESTDIRCB = func(log *logger.Logger, ret string) (string, error) {
		return "INSTALL_ROOT", nil
	}

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	switch info.PackageName {
	default:
		return nil, errors.New("unsupported qt version")

	case "qt4":
		self.qt_major_version = 4

	case "qt5":
		self.qt_major_version = 5
	}

	return self, nil
}

func (self *Builder_qt) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autoconf")

	ret.ReplaceShort("configure", self.BuilderActionConfigure)

	ret, err := ret.AddActionAfterNameShort(
		"distribute",
		"sys_env", self.BuilderActionSetupSysEnv,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_qt) BuilderActionConfigure(log *logger.Logger) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	tri, err := systemtriplet.NewFromString(info.HostArch)
	if err != nil {
		return err
	}

	platform := ""

	switch tri.Kernel {
	case "linux":
		switch tri.CPU {
		case "i486":
			fallthrough
		case "i586":
			fallthrough
		case "i686":
			platform = "linux-g++-32"

		case "x86_64":
			platform = "linux-g++-64"
		}
	}

	if platform == "" {
		return errors.New("unsupported platform")
	}

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return err
	}

	opts := []string{
		"-opensource",
		"-confirm-license",
		"-prefix", path.Join(
			install_prefix,
			"opt",
			"qt",
			strconv.Itoa(self.qt_major_version),
		),
		"-system-sqlite",

		// # I"m adding this for qt5, don"t know if qt4 has this
		"-dbus-linked",
		"-openssl-linked",

		"-platform", platform,
	}

	if self.qt_major_version == 5 {
		opts = append(
			opts,
			[]string{
				"-no-compile-examples",
				"-nomake", "examples",
				"-skip", "qtwebengine",
				"-pulseaudio",
				"-no-alsa",
			}...,
		)
	}

	c := exec.Command("./configure", opts...)
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_qt) BuilderActionSetupSysEnv(log *logger.Logger) error {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return err
	}

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	set_dir := path.Join(
		self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
		"etc",
		"profile.d",
		"SET",
	)

	set_dir_file := path.Join(
		set_dir,
		fmt.Sprintf(
			"009.qt%d.%s.%s.sh",
			self.qt_major_version,
			info.Host,
			info.HostArch,
		),
	)

	qt_envs := `#!/bin/bash

INT_QTDIR=` + install_prefix + `/opt/qt/` + strconv.Itoa(self.qt_major_version) + `

export PATH=$PATH:$INT_QTDIR/bin

if [ "${#PKG_CONFIG_PATH}" -ne "0" ]; then
    PKG_CONFIG_PATH+=":"
fi
export PKG_CONFIG_PATH+="$INT_QTDIR/lib/pkgconfig"
export PKG_CONFIG_PATH+=":$INT_QTDIR/lib64/pkgconfig"
export PKG_CONFIG_PATH+=":$INT_QTDIR/share/pkgconfig"

if [ "${#PKG_CONFIG_PATH}" -ne "0" ]; then
    LD_LIBRARY_PATH+=":"
fi
export LD_LIBRARY_PATH+="$INT_QTDIR/lib"
export LD_LIBRARY_PATH+=":$INT_QTDIR/lib64"

unset INT_QTDIR
`

	err = os.MkdirAll(set_dir, 0700)
	if err != nil {
		return err
	}

	f, err := os.Create(set_dir_file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(qt_envs)
	if err != nil {
		return err
	}

	return nil
}
