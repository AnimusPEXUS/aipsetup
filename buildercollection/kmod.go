package buildercollection

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["kmod"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_kmod(bs)
	}
}

type Builder_kmod struct {
	*Builder_std
}

func NewBuilder_kmod(bs basictypes.BuildingSiteCtlI) (*Builder_kmod, error) {

	self := new(Builder_kmod)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_kmod) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret, err := ret.AddActionsBeforeName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "make_links",
				Callable: self.MakeLinks,
			},
		},
		"prepack",
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_kmod) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-xz",
			"--with-zlib",
		}...,
	)

	return ret, nil
}

func (self *Builder_kmod) MakeLinks(log *logger.Logger) error {

	dst_install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	bin := path.Join(dst_install_prefix, "bin")
	sbin := path.Join(dst_install_prefix, "sbin")

	os.MkdirAll(bin, 0700)
	os.MkdirAll(sbin, 0700)

	for _, i := range []string{"depmod", "insmod", "modinfo", "modprobe", "rmmod"} {
		err = os.Symlink("../bin/kmod", path.Join(sbin, i))
		if err != nil {
			return err
		}
	}

	err = os.Symlink("./kmod", path.Join(bin, "lsmod"))
	if err != nil {
		return err
	}

	return nil
}
