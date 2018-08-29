package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openldap"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_openldap(bs)
	}
}

type Builder_openldap struct {
	*Builder_std
}

func NewBuilder_openldap(bs basictypes.BuildingSiteCtlI) (*Builder_openldap, error) {

	self := new(Builder_openldap)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_openldap) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{
			"--with-tls=openssl",
			"--disable-bdb",
			"--disable-hdb",
		}...,
	)
	return ret, nil
}
