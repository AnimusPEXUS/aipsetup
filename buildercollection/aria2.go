package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["aria2"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_aria2(bs)
	}
}

type Builder_aria2 struct {
	BuilderStdAutotools
}

func NewBuilder_aria2(bs basictypes.BuildingSiteCtlI) (*Builder_aria2, error) {
	self := new(Builder_aria2)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_aria2) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{
			"--enable-bittorrent",
			"--enable-metalink",
			"--enable-epoll",
			"--with-gnutls",
			"--with-openssl",
			"--with-sqlite3",
			"--with-libxml2",
			"--with-libexpat",
		}...,
	), nil
}
