package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["inetutils"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_inetutils(bs), nil
	}
}

type Builder_inetutils struct {
	*Builder_std
}

func NewBuilder_inetutils(bs basictypes.BuildingSiteCtlI) *Builder_inetutils {

	self := new(Builder_inetutils)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_inetutils) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

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

	//	include_dir, err := pkgconfig.CommandOutput("--variable=includedir", "ncurses")
	//	if err != nil {
	//		return nil, err
	//	}

	//	include_dirs, err := pkgconfig.GetIncludeDirs("ncurses")
	//	if err != nil {
	//		return nil, err
	//	}

	//	if len(include_dirs) != 1 {
	//		return nil, errors.New("got invalid number of ncurses include dirs")
	//	}

	ret = append(
		ret,
		[]string{
			"--disable-logger",
			//			"--with-ncurses-include-dir=" + include_dir,
		}...,
	)

	ret = append(
		ret,
		[]string{
			"CFLAGS=" + ncurses_cflags,
			"LDFLAGS=" + ncurses_libs,
		}...,
	)

	return ret, nil
}
