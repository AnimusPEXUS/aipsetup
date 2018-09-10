package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["gdk_pixbuf"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_gdk_pixbuf(bs), nil
	}
}

type Builder_gdk_pixbuf struct {
	*Builder_std
}

func NewBuilder_gdk_pixbuf(bs basictypes.BuildingSiteCtlI) *Builder_gdk_pixbuf {

	self := new(Builder_gdk_pixbuf)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_gdk_pixbuf) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(ret, "--with-x11")

	return ret, nil
}
