package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["git"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_git(bs)
	}
}

type Builder_git struct {
	Builder_std
}

func NewBuilder_git(bs basictypes.BuildingSiteCtlI) (*Builder_git, error) {
	self := new(Builder_git)

	self.Builder_std = *NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_git) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{"--with-openssl",
			"--with-curl",
			"--with-expat",
		}...,
	), nil
}
