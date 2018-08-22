package buildercollection

import (
	"fmt"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libressl"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libressl(bs)
	}
}

type Builder_libressl struct {
	*Builder_std
}

func NewBuilder_libressl(bs basictypes.BuildingSiteCtlI) (*Builder_libressl, error) {

	self := new(Builder_libressl)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditConfigureScriptNameCB = self.EditConfigureScriptName
	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditConfigureIsArgToShellCB = self.EditConfigureIsArgToShell

	self.EditConfigureShellCB = self.EditConfigureShell
	//	self.AfterDistributeCB = self.AfterDistribute

	return self, nil
}

func (self *Builder_libressl) veredName() (string, error) {
	info, err := self.bs.ReadInfo()
	if err != nil {
		return "", err
	}

	spl := strings.Split(info.PackageVersion, ".")

	ret := fmt.Sprintf(
		"libressl-%s.%s",
		spl[0],
		spl[1],
	)

	return ret, nil
}

func (self *Builder_libressl) EditConfigureShell(log *logger.Logger, ret string) (string, error) {
	return "perl", nil
}

func (self *Builder_libressl) EditConfigureIsArgToShell(log *logger.Logger, ret bool) (bool, error) {
	return true, nil
}

func (self *Builder_libressl) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	platform := "linux-generic32"
	if strings.HasPrefix(info.HostArch, "x86_64") {
		platform = "linux-x86_64"
	}

	install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = make([]string, 0)

	ret = append(
		ret,
		[]string{
			"--prefix=" + install_prefix,
			"--libressldir=/etc/ssl",
		}...,
	)

	libdir_slice := []string{"lib"}

	{

		ver_name, err := self.veredName()
		if err != nil {
			return nil, err
		}

		// TODO: I don't like this. Need to add some better comparators
		spl := strings.Split(info.PackageVersion, ".")
		if !(spl[0] == "1" && spl[1] == "1") {
			libdir_slice = append(libdir_slice, ver_name)
		}
	}

	ret = append(
		ret,
		[]string{
			"--libdir=" + path.Join(libdir_slice...),
		}...,
	)

	ret = append(
		ret,
		[]string{
			"shared",
			"zlib-dynamic",
			platform,
		}...,
	)

	return ret, nil
}

func (self *Builder_libressl) EditConfigureScriptName(log *logger.Logger, ret string) (string, error) {
	return "Configure", nil
}

func (self *Builder_libressl) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	//	dst_install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	//	if err != nil {
	//		return nil, err
	//	}

	ret = []string{
		"depend",
		"install",
		// FIXME: fix path join
		"MANDIR=" + path.Join(install_prefix, "share", "man"),
		// "MANSUFFIX=ssl",
		"INSTALL_PREFIX=" + self.bs.GetDIR_DESTDIR(),
		"DESTDIR=" + self.bs.GetDIR_DESTDIR(),
	}

	return ret, nil
}
