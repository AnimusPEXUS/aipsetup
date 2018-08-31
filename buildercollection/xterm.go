package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["xterm"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_xterm(bs)
	}
}

type Builder_xterm struct {
	*Builder_std
}

func NewBuilder_xterm(bs basictypes.BuildingSiteCtlI) (*Builder_xterm, error) {
	self := new(Builder_xterm)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	//	self.EditAutogenForceCB = func(log *logger.Logger, ret bool) (bool, error) {
	//		return true, nil
	//	}

	return self, nil
}

func (self *Builder_xterm) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{"--docdir="},
	)
	if err != nil {
		return nil, err
	}

	pkgconfig, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().GetPrefixPkgConfig()
	if err != nil {
		return nil, err
	}

	ncurses_cflags, err := pkgconfig.CommandOutput("--cflags", "ncursesw")
	if err != nil {
		return nil, err
	}

	ncurses_libs, err := pkgconfig.CommandOutput("--libs", "ncursesw")
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			// TODO: this requires better solution
			"CFLAGS=" + ncurses_cflags,
			"LDFLAGS=" + ncurses_libs,
		}...,
	)

	return ret, nil
}
