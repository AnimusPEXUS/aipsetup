package buildercollection

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["avahi"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_avahi(bs)
	}
}

type Builder_avahi struct {
	*Builder_std
}

func NewBuilder_avahi(bs basictypes.BuildingSiteCtlI) (*Builder_avahi, error) {

	self := new(Builder_avahi)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_avahi) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"distribute",
		"after-distribute", self.BuilderActionAfterDistribute,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_avahi) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{
			"--disable-gtk",
			"--disable-gtk3",
			"--enable-glib",
			"--enable-gobject",
			"--enable-python",
			"--enable-introspection",
			"--disable-mono",
			// '--disable-python-dbus',
			"--disable-pygtk",
			"--disable-qt3",
			"--disable-qt4",
			"--with-distro=lfs",
			//                    '--with-distro=' +
			//                        pkg_info['constitution']['system_title'],
			//                    '--with-dist-version=2.00',
			//                    '--without-systemdsystemunitdir',
		}...,
	)
	return ret, nil
}

func (self *Builder_avahi) BuilderActionAfterDistribute(log *logger.Logger) error {
	dst_run := path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), "run")
	os.Remove(dst_run)
	return nil
}
