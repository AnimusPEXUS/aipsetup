package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ghostscript"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_ghostscript(bs)
	}
}

type Builder_ghostscript struct {
	*Builder_std
}

func NewBuilder_ghostscript(bs basictypes.BuildingSiteCtlI) (*Builder_ghostscript, error) {

	self := new(Builder_ghostscript)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self, nil
}

func (self *Builder_ghostscript) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("build")

	return ret, nil
}

func (self *Builder_ghostscript) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-x",
			"--with-install-cups",
			"--with-ijs",
			"--with-drivers=ALL",
		}...,
	)

	return ret, nil
}

func (self *Builder_ghostscript) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"all",
			"install",
			"so",
			"soinstall",
		}...,
	)

	return ret, nil
}
