package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openvpn"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_openvpn(bs)
	}
}

type Builder_openvpn struct {
	*Builder_std
}

func NewBuilder_openvpn(bs basictypes.BuildingSiteCtlI) (*Builder_openvpn, error) {

	self := new(Builder_openvpn)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_openvpn) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{
			"--enable-iproute2",
			"--enable-systemd",
		}...,
	)
	return ret, nil
}
