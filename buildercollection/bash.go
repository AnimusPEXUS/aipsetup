package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["bash"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_bash(bs)
	}
}

type Builder_bash struct {
	BuilderStdAutotools
}

func NewBuilder_bash(bs basictypes.BuildingSiteCtlI) (*Builder_bash, error) {
	self := new(Builder_bash)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	self.PatchCB = self.BuilderActionPatch
	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_bash) BuilderActionPatch(log *logger.Logger) error {
	//
	// info, err := self.bs.ReadInfo()
	// if err != nil {
	// 	return err
	// }
	//
	// version_parsed := versionorstatus
	//
	// bash_patch_prefix := fmt.Sprintf("bash%d%d-")

	return nil
}

func (self *Builder_bash) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = append(ret, "--enable-multibyte")

	if info.ThisIsCrossbuilder() || info.ThisIsCrossbuilding() {
		ret = append(ret, "--without-curses")
	} else {
		ret = append(ret, "--with-curses")
	}

	return ret, nil
}
