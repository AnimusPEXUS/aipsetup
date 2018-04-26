package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ntp"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderNTP(bs)
	}
}

type BuilderNTP struct {
	BuilderStdAutotools
}

func NewBuilderNTP(bs basictypes.BuildingSiteCtlI) (*BuilderNTP, error) {
	self := new(BuilderNTP)
	self.BuilderStdAutotools = *NewBuilderStdAutotools(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *BuilderNTP) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(ret, "--without-ntpsnmpd"), nil
}
