package buildercollection

import (
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["mozjs"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_mozjs(bs)
	}
}

type Builder_mozjs struct {
	*Builder_std
}

func NewBuilder_mozjs(bs basictypes.BuildingSiteCtlI) (*Builder_mozjs, error) {
	self := new(Builder_mozjs)
	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditExtractMoreThanOneExtractedOkCB = self.EditExtractMoreThanOneExtractedOk
	self.EditExtractUnwrapCB = self.EditExtractUnwrap

	self.EditConfigureDirCB = self.EditConfigureDir
	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_mozjs) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("autogen")
	return ret, nil
}

func (self *Builder_mozjs) EditExtractMoreThanOneExtractedOk(log *logger.Logger, value bool) (bool, error) {
	return true, nil
}

func (self *Builder_mozjs) EditExtractUnwrap(log *logger.Logger, value bool) (bool, error) {
	return false, nil
}

func (self *Builder_mozjs) EditConfigureDir(log *logger.Logger, ret string) (string, error) {
	ret = path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "js", "src")
	return ret, nil
}

func (self *Builder_mozjs) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
	return self.EditConfigureDir(log, ret)
}

func (self *Builder_mozjs) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	for i := len(ret) - 1; i != -1; i-- {
		for _, j := range []string{} {
			if strings.HasPrefix(ret[i], j) {
				ret = append(ret[:i], ret[i+1:]...)
			}
		}

	}

	return ret, nil
}
