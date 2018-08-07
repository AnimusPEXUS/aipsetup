package buildercollection

import (
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openssl"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_openssl(bs)
	}
}

type Builder_openssl struct {
	*Builder_std
}

func NewBuilder_openssl(bs basictypes.BuildingSiteCtlI) (*Builder_openssl, error) {

	self := new(Builder_openssl)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditActionsCB = self.EditActions

	self.EditConfigureScriptNameCB = self.EditConfigureScriptName
	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self, nil
}

func (self *Builder_openssl) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("build")
	return ret, nil
}

func (self *Builder_openssl) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

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

	ret = []string{
		"--prefix=" + install_prefix,
		"--openssldir=/etc/ssl",
		"shared",
		"zlib-dynamic",
		platform,
	}

	return ret, nil
}

func (self *Builder_openssl) EditConfigureScriptName(log *logger.Logger, ret string) (string, error) {
	return "Configure", nil
}

func (self *Builder_openssl) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	//	dst_install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	//	if err != nil {
	//		return nil, err
	//	}

	ret = []string{
		"install",
		// FIXME: fix path join
		"MANDIR=" + path.Join(install_prefix, "share", "man"),
		// "MANSUFFIX=ssl",
		//		"INSTALL_PREFIX=" + dst_install_prefix,
		"DESTDIR=" + self.bs.GetDIR_DESTDIR(),
	}

	return ret, nil
}
