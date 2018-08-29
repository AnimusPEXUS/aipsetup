package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["linux_pam"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_linux_pam(bs)
	}
}

type Builder_linux_pam struct {
	*Builder_std
}

func NewBuilder_linux_pam(bs basictypes.BuildingSiteCtlI) (*Builder_linux_pam, error) {
	self := new(Builder_linux_pam)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_linux_pam) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()

	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--disable-nis",
			"--enable-db=ndbm",
			"--enable-read-both-confs",
			"--enable-selinux",

			// it's not a mistake: 'security' dir is here
			// required at least by 'polkit'
			"--includedir=" + path.Join(install_prefix, "include", "security"),
			"--enable-securedir=" + path.Join(install_prefix, "lib", "security"),
		}...,
	)

	return ret, nil
}
