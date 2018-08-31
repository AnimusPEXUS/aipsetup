package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["zbar"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_zbar(bs)
	}
}

type Builder_zbar struct {
	*Builder_std
}

func NewBuilder_zbar(bs basictypes.BuildingSiteCtlI) (*Builder_zbar, error) {
	self := new(Builder_zbar)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_zbar) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--disable-video",
			//             '--without-gtk',
			"--without-python",
			"--without-qt",
		}...,
	)

	return ret, nil
}
