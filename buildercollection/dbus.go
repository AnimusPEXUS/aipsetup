package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["dbus"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_dbus(bs)
	}
}

type Builder_dbus struct {
	*Builder_std
}

func NewBuilder_dbus(bs basictypes.BuildingSiteCtlI) (*Builder_dbus, error) {
	self := new(Builder_dbus)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_dbus) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{
			"--with-x",
			// "--enable-selinux", #lib needed

			//			"--enable-libaudit",
			//			"--enable-dnotify",
			//			"--enable-inotify",

			// '--enable-kqueue', #BSD needed
			// '--enable-launchd', #MacOS needed

			// NOTE: cyrcular dep with systemd.
			//       build without systemd may be required once
			"--enable-user-session",
			"--enable-systemd",
			//#'--disable-systemd',

			// NOTE: cyrcular dep with dbus-glib
			// NOTE: dbus-glib is deprecated
			//# '--without-dbus-glib'

			"--with-system-socket=/run/dbus/system_bus_socket",
			"--with-console-auth-dir=/run/console",
			"--with-system-pid-file=/run/dbus/pid",
		}...,
	), nil
}
