package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["systemd"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_systemd(bs), nil
	}
}

type Builder_systemd struct {
	*Builder_std
}

func NewBuilder_systemd(bs basictypes.BuildingSiteCtlI) *Builder_systemd {
	self := new(Builder_systemd)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_systemd) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	python, err := self.bs.GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("python")
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			// # '--disable-silent-rules',
			"--enable-gudev=auto",
			"--enable-gtk-doc=auto",
			"--enable-logind=auto",
			"--enable-microhttpd=auto",
			"--enable-qrencode=auto",
			// # '--enable-static',
			// # '--disable-tests',
			// # '--disable-coverage',
			"--enable-shared",
			"--enable-compat-libs",
			// #'--with-libgcrypt-prefix={}'.format(self.get_host_dir()),
			// #'--with-rootprefix={}'.format(self.get_host_dir()),
		}...,
	)

	ret = append(
		ret,
		[]string{
			"--with-pamlibdir=" + path.Join(install_prefix, "lib", "security"),
		}...,
	)

	ret = append(
		ret,
		[]string{
			"PYTHON=" + python,
		}...,
	)

	return ret, nil
}
