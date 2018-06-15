package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["cairo"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_cairo(bs), nil
	}
}

type Builder_cairo struct {
	*Builder_std
}

func NewBuilder_cairo(bs basictypes.BuildingSiteCtlI) *Builder_cairo {
	self := new(Builder_cairo)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

// TODO: building may require 'LDFLAGS=-ltinfow' - testing needed

func (self *Builder_cairo) EditConfigureArgs(log *logger.Logger, ret []string) (
	[]string, error,
) {
	ret = append(
		ret,
		[]string{
			//#'--enable-cogl=auto',
			"--enable-directfb=auto",
			//#'--enable-drm=auto',
			"--enable-fc",
			"--enable-ft",
			"--enable-gl",
			//#'--enable-gallium',
			//#'--enable-glesv2',
			"--enable-pdf=yes",
			"--enable-png=yes",
			"--enable-ps=yes",
			"--enable-svg=yes",
			//#                    '--enable-qt',

			"--enable-quartz-font=auto",
			"--enable-quartz-image=auto",
			"--enable-quartz=auto",

			"--enable-script=yes",

			"--enable-tee=yes",
			"--enable-vg=auto",
			"--enable-wg=auto",
			"--enable-xcb",
			"--enable-xcb-shm",
			"--enable-xlib-xcb",
			"--enable-gobject=auto",

			"--enable-egl=auto",
			"--enable-glx=auto",
			// #'--enable-wgl',

			// # xlib is deprecated
			// #                    '--enable-xlib',
			// #                    '--enable-xlib-xcb',
			// #                    '--enable-xlib-xrender',

			"--disable-static",
			"--enable-xml=yes",

			"--with-x",
			//            #'WERROR='
		}...,
	)

	return ret, nil
}
