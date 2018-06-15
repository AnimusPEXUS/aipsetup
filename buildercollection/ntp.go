package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ntp"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_ntp(bs)
	}
}

type Builder_ntp struct {
	*Builder_std
}

func NewBuilder_ntp(bs basictypes.BuildingSiteCtlI) (*Builder_ntp, error) {
	self := new(Builder_ntp)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_ntp) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(ret, "--without-ntpsnmpd"), nil
}
