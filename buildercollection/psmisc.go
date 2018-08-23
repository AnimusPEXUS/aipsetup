package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["psmisc"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_psmisc(bs)
	}
}

type Builder_psmisc struct {
	*Builder_std
}

func NewBuilder_psmisc(bs basictypes.BuildingSiteCtlI) (*Builder_psmisc, error) {

	self := new(Builder_psmisc)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_psmisc) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	pkgconfig, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().GetPrefixPkgConfig()
	if err != nil {
		return nil, err
	}

	ncurses_cflags, err := pkgconfig.CommandOutput("--cflags", "ncurses")
	if err != nil {
		return nil, err
	}

	ncurses_libs, err := pkgconfig.CommandOutput("--libs", "ncurses")
	if err != nil {
		return nil, err
	}

	ncursesw_cflags, err := pkgconfig.CommandOutput("--cflags", "ncursesw")
	if err != nil {
		return nil, err
	}

	ncursesw_libs, err := pkgconfig.CommandOutput("--libs", "ncursesw")
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"NCURSES_CFLAGS=" + ncurses_cflags,
			"NCURSES_LIBS=" + ncurses_libs,
			"NCURSESW_CFLAGS=" + ncursesw_cflags,
			"NCURSESW_LIBS=" + ncursesw_libs,
		}...,
	)

	ret = append(
		ret,
		// NOTE: I don't like this very match!
		[]string{
			"CFLAGS=" + ncursesw_cflags,
		}...,
	)

	return ret, nil
}
