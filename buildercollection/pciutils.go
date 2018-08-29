package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["pciutils"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_pciutils(bs)
	}
}

type Builder_pciutils struct {
	*Builder_std
}

func NewBuilder_pciutils(bs basictypes.BuildingSiteCtlI) (*Builder_pciutils, error) {

	self := new(Builder_pciutils)

	self.Builder_std = NewBuilder_std(bs)

	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_pciutils) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	return ret, nil
}

func (self *Builder_pciutils) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	env, err := self.BuilderActionDistributeEnvDef(log)
	if err != nil {
		return nil, err
	}

	CC := env.Get("CC", "gcc")
	CXX := env.Get("CXX", "g++")

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{"install"},
		[]string{"DESTDIR="},
	)
	if err != nil {
		return nil, err
	}

	args := []string{
		"all",
		"install",
		"install-lib",
		"PREFIX=" + install_prefix,
		"DESTDIR=" + self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
		"SHARED=yes",
		"CC=" + CC,
		"CXX=" + CXX,
	}

	ret = append(ret, args...)

	return ret, nil
}
