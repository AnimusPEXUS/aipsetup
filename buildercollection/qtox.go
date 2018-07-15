package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["qtox"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_qtox(bs)
	}
}

type Builder_qtox struct {
	*Builder_std_cmake
}

func NewBuilder_qtox(bs basictypes.BuildingSiteCtlI) (*Builder_qtox, error) {

	self := new(Builder_qtox)
	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.EditConfigureEnvCB = self.EditConfigureEnv

	return self, nil
}

func (self *Builder_qtox) EditConfigureEnv(
	log *logger.Logger,
	ret environ.EnvVarEd,
) (
	environ.EnvVarEd,
	error,
) {
	calc := self.bs.GetBuildingSiteValuesCalculator()
	prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	// TODO: had to do it fast. no hardcode.
	ret.Set("PATH", ret.Get("PATH", "")+":"+prefix+"/opt/qt/5/bin")

	return ret, nil
}
