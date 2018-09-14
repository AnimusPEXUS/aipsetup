package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["telepathy_gabble"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_telepathy_gabble(bs)
	}
}

type Builder_telepathy_gabble struct {
	*Builder_std
}

func NewBuilder_telepathy_gabble(bs basictypes.BuildingSiteCtlI) (*Builder_telepathy_gabble, error) {

	self := new(Builder_telepathy_gabble)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_telepathy_gabble) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-ca-certificates=/etc/ssl/cert.pem",
		}...,
	)

	return ret, nil
}
