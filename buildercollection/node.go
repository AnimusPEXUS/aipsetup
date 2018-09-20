package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["node"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_node(bs)
	}
}

type Builder_node struct {
	*Builder_std
}

func NewBuilder_node(bs basictypes.BuildingSiteCtlI) (*Builder_node, error) {
	self := new(Builder_node)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_node) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{
			"--enable-shared",
		},
		[]string{
			"--libdir=",
			"--docdir=",
			"--sysconfdir",
			"--localstatedir",
			"--host",
			"--build",
			"CC=",
			"CXX=",
		},
	)
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{}...,
	)

	return ret, nil
}
