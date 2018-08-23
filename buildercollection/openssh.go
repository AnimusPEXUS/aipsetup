package buildercollection

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openssh"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_openssh(bs)
	}
}

type Builder_openssh struct {
	*Builder_std
}

func NewBuilder_openssh(bs basictypes.BuildingSiteCtlI) (*Builder_openssh, error) {

	self := new(Builder_openssh)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_openssh) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret, err := ret.AddActionBeforeNameShort(
		"prepack",
		"rename_configs", self.RenameConfigs,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_openssh) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	pkgconfig, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().GetPrefixPkgConfig()
	if err != nil {
		return nil, err
	}

	//	LIBSSH_CFLAGS, err := pkgconfig.CommandOutput("--cflags", "libssh2")
	//	if err != nil {
	//		return nil, err
	//	}

	//	LIBSSH_LIBS, err := pkgconfig.CommandOutput("--libs", "libssh2")
	//	if err != nil {
	//		return nil, err
	//	}

	OPENSSL100_CFLAGS, err := pkgconfig.CommandOutput("--cflags", "openssl-1.0")
	if err != nil {
		return nil, err
	}

	OPENSSL100_LIBS, err := pkgconfig.CommandOutput("--libs", "openssl-1.0")
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--with-tcp-wrappers",
			"--with-pam",
			"--sysconfdir=/etc/ssh",
			"CFLAGS=" + OPENSSL100_CFLAGS,
			"LDFLAGS=" + OPENSSL100_LIBS,
		}...,
	)

	return ret, nil
}

func (self *Builder_openssh) RenameConfigs(log *logger.Logger) error {

	ssh_dir := path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), "etc", "ssh")

	err := os.Rename(
		path.Join(ssh_dir, "sshd_config"),
		path.Join(ssh_dir, "sshd_config.origin"),
	)
	if err != nil {
		return err
	}

	err = os.Rename(
		path.Join(ssh_dir, "ssh_config"),
		path.Join(ssh_dir, "ssh_config.origin"),
	)
	if err != nil {
		return err
	}

	return nil
}
