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
	ret, err := ret.AddActionsBeforeName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "rename_configs",
				Callable: self.RenameConfigs,
			},
		},
		"prepack",
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_openssh) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-tcp-wrappers",
			"--with-pam",
			"--sysconfdir=/etc/ssh",
		}...,
	)

	return ret, nil
}

func (self *Builder_openssh) RenameConfigs(log *logger.Logger) error {

	ssh_dir := path.Join(self.bs.GetDIR_DESTDIR(), "etc", "ssh")

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
