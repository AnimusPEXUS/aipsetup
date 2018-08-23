package buildercollection

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["less"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_less(bs)
	}
}

type Builder_less struct {
	*Builder_std
}

func NewBuilder_less(bs basictypes.BuildingSiteCtlI) (*Builder_less, error) {
	self := new(Builder_less)

	self.Builder_std = NewBuilder_std(bs)

	return self, nil
}

func (self *Builder_less) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"distribute",
		"after-distribute", self.BuilderActionAfterDistribute,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_less) BuilderActionAfterDistribute(log *logger.Logger) error {

	dir := path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), "etc", "profile.d", "SET")
	file := path.Join(dir, "009.LESS.sh")

	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	_, err = f.WriteString(`
#!/bin/bash
export LESS=' -R '
`)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
