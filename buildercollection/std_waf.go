package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_waf"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_waf(bs)
	}
}

type Builder_std_waf struct {
	*Builder_std
}

func NewBuilder_std_waf(bs basictypes.BuildingSiteCtlI) (*Builder_std_waf, error) {
	self := new(Builder_std_waf)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_std_waf) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{"--enable-shared"},
		[]string{"CC=", "CXX=", "GCC="},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
