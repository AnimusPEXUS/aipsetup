package buildercollection

import (
	"errors"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["wine"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_wine(bs, false)
	}
}

type Builder_wine struct {
	*Builder_std

	info *basictypes.BuildingSiteInfo

	build_with_wow64               bool
	build_with_wow64_part2         bool
	build_with_wow64_part2_builder *Builder_wine
}

func NewBuilder_wine(
	bs basictypes.BuildingSiteCtlI,
	build_with_wow64_part2 bool,
) (*Builder_wine, error) {
	self := new(Builder_wine)

	self.build_with_wow64_part2 = build_with_wow64_part2

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.EditConfigureEnvCB = self.EditConfigureEnv

	//	self.EditDistributeArgsCB = self.EditDistributeArgs

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	self.info = info

	self.build_with_wow64 = info.Host == "x86_64-pc-linux-gnu"

	if self.build_with_wow64 && !build_with_wow64_part2 {
		w, err := NewBuilder_wine(bs, true)
		if err != nil {
			return nil, err
		}
		self.build_with_wow64_part2_builder = w
	}

	subdir := "wine32"
	if self.isWine64() {
		subdir = "wine64"
	}

	self.EditConfigureWorkingDirCB = func(
		log *logger.Logger,
		ret string,
	) (string, error) {
		ret = path.Join(
			self.GetBuildingSiteCtl().GetDIR_SOURCE(),
			subdir,
		)
		err := os.MkdirAll(ret, 0700)
		if err != nil {
			return "", err
		}
		return ret, nil
	}

	return self, nil
}

func (self *Builder_wine) isWine64() bool {
	ret := self.info.Host == "x86_64-pc-linux-gnu" &&
		self.build_with_wow64 &&
		!self.build_with_wow64_part2
	return ret
}

func (self *Builder_wine) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	if self.build_with_wow64_part2 {

		if !self.build_with_wow64 {
			return nil, errors.New("invalid programming")
		}

		ret = basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "configure",
				Callable: self.BuilderActionConfigure,
			},
			&basictypes.BuilderAction{
				Name:     "build",
				Callable: self.BuilderActionBuild,
			},
			&basictypes.BuilderAction{
				Name:     "distribute",
				Callable: self.BuilderActionDistribute,
			},
		}
		return ret, nil
	}

	var wow64_actions basictypes.BuilderActions

	if self.build_with_wow64 {
		t, err := self.build_with_wow64_part2_builder.DefineActions()
		if err != nil {
			return nil, err
		}

		wow64_actions = t

		configure, ok := wow64_actions.Get("configure")
		if !ok {
			return nil, errors.New("error")
		}

		build, ok := wow64_actions.Get("build")
		if !ok {
			return nil, errors.New("error")
		}

		distribute, ok := wow64_actions.Get("distribute")
		if !ok {
			return nil, errors.New("error")
		}

		t, err = ret.AddActionsBeforeName(
			basictypes.BuilderActions{
				&basictypes.BuilderAction{
					Name:     "configure_wow64",
					Callable: configure.Callable,
				},
				&basictypes.BuilderAction{
					Name:     "build_wow64",
					Callable: build.Callable,
				},
			},
			"configure",
		)
		if err != nil {
			return nil, err
		}

		ret = t

		t, err = ret.AddActionAfterNameShort(
			"distribute",
			"distribute_wow64", distribute.Callable,
		)
		if err != nil {
			return nil, err
		}

		ret = t
	}

	return ret, nil
}

func (self *Builder_wine) EditConfigureArgs(
	log *logger.Logger,
	ret []string,
) ([]string, error) {

	if self.isWine64() {
		ret = append(ret, "--enable-win64")
	}

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{
			"--host=",
			"--build=",
			"--prefix=",
			"--libdir=",
			"CC=",
			"CXX=",
			"GCC=",
		},
	)
	if err != nil {
		return nil, err
	}

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	install_prefix := path.Join(
		calc.CalculateMultihostDir(),
		self.info.Host,
	)

	ret = append(ret, "--prefix="+install_prefix)

	libdir := "lib"
	if self.isWine64() {
		libdir = "lib64"
	}

	ret = append(ret, "--libdir="+path.Join(install_prefix, libdir))

	return ret, nil
}

func (self *Builder_wine) EditConfigureEnv(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error) {
	for _, i := range []string{"CC", "CXX", "GCC"} {
		ret.Del(i)
	}
	return ret, nil
}
