package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["git"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderGit(bs)
	}
}

type BuilderGit struct {
	BuilderStdAutotools
}

func NewBuilderGit(bs basictypes.BuildingSiteCtlI) (*BuilderGit, error) {
	self := new(BuilderGit)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *BuilderGit) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{"--with-openssl",
			"--with-curl",
			"--with-expat",
		}...,
	), nil
}
