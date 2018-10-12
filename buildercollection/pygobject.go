package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["pygobject"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_pygobject(bs)
	}
}

type Builder_pygobject struct {
	*Builder_std_meson
}

func NewBuilder_pygobject(bs basictypes.BuildingSiteCtlI) (*Builder_pygobject, error) {

	self := new(Builder_pygobject)

	Builder_std_meson, err := NewBuilder_std_meson(bs)
	if err != nil {
		return nil, err
	}

	self.Builder_std_meson = Builder_std_meson

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_pygobject) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	python, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("python3")
	if err != nil {
		return nil, err
	}

	//	ret = append(ret, "PYTHON="+python)

	ret = append(
		ret,
		[]string{
			"-Dpython=" + python,
		}...,
	)

	return ret, nil
}
