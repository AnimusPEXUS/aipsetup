package buildercollection

import (
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

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_avahi) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
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
	), nil
}
