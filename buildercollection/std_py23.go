package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_py23"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_py23(bs)
	}
}

type Builder_std_py23 struct {
	*Builder_std

	python string
}

func NewBuilder_std_py23(bs basictypes.BuildingSiteCtlI) (*Builder_std_py23, error) {

	self := new(Builder_std_py23)

	self.Builder_std = NewBuilder_std(bs)

	//	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_std_py23) DefineActions() (basictypes.BuilderActions, error) {

	ret := basictypes.BuilderActions{}

	std_actions, err := self.Builder_std.DefineActions()
	if err != nil {
		return nil, err
	}

	append_std_actions := func() error {

		for _, i := range []string{
			"src_cleanup",
			"bld_cleanup",
			"extract",
			"autogen",
			"configure",
			"build",
			"distribute",
		} {
			for _, j := range std_actions {
				if j.Name == i {
					ret, err = ret.AppendSingle(j.Name, j.Callable)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	}

	ret, err = ret.AppendSingle(
		"set_python2",
		func(log *logger.Logger) error {
			self.python = "python2"
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	err = append_std_actions()

	ret, err = ret.AppendSingle(
		"set_python3",
		func(log *logger.Logger) error {
			self.python = "python3"
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	err = append_std_actions()

	for _, i := range []string{"prepack", "pack"} {
		for _, j := range std_actions {
			if j.Name == i {
				ret, err = ret.AppendSingle(j.Name, j.Callable)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return ret, nil
}

func (self *Builder_std_py23) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	exec, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefixExecutable(
		self.python,
	)
	if err != nil {
		return nil, err
	}

	ret = append(ret, "PYTHON="+exec)

	return ret, nil
}
