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

	// TODO: I don't like this. Need to add some better comparators
	//	if info.PackageVersion == "1.0" || strings.HasPrefix(info.PackageVersion, "1.0.") {
	//		ret = ret.AddActionsAfterName(
	//			basictypes.BuilderActions{
	//				&basictypes.BuilderAction{
	//					Name : "depend",
	//					Callable: self.BuilderActionMakeDepend
	//				},
	//			},
	//			"configure",
	//			)
	//	}

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

	ret = make([]string, 0)

	ret = append(
		ret,
		[]string{
			"--prefix=" + install_prefix,
			"--openssldir=/etc/ssl",
		}...,
	)

	libdir_slice := []string{install_prefix, "lib"}

	if info.PackageVersion == "0.9" || strings.HasPrefix(info.PackageVersion, "0.9.") {
		libdir_slice = append(libdir_slice, "openssl-0.9")
	}

	if info.PackageVersion == "1.0" || strings.HasPrefix(info.PackageVersion, "1.0.") {
		libdir_slice = append(libdir_slice, "openssl-1.0")
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
		"depend",
		"install",
		// FIXME: fix path join
		"MANDIR=" + path.Join(install_prefix, "share", "man"),
		// "MANSUFFIX=ssl",
		//		"INSTALL_PREFIX=" + dst_install_prefix,
		"DESTDIR=" + self.bs.GetDIR_DESTDIR(),
	}

	return ret, nil
}
