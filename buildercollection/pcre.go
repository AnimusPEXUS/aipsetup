package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["pcre"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_pcre(bs)
	}
}

type Builder_pcre struct {
	*Builder_std
}

func NewBuilder_pcre(bs basictypes.BuildingSiteCtlI) (*Builder_pcre, error) {

	self := new(Builder_pcre)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_pcre) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-utf",
			"--enable-unicode-properties",
		}...,
	)

	return ret, nil
}
