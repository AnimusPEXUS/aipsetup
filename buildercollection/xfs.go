package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["xfs"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_xfs(bs), nil
	}
}

type Builder_xfs struct {
	*Builder_std
}

func NewBuilder_xfs(bs basictypes.BuildingSiteCtlI) *Builder_xfs {

	self := new(Builder_xfs)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

func (self *Builder_xfs) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	pkgconfig, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().GetPrefixPkgConfig()
	if err != nil {
		return nil, err
	}

	//	ncurses_cflags, err := pkgconfig.CommandOutput("--cflags", "ncurses")
	//	if err != nil {
	//		return nil, err
	//	}

	//	ncurses_libs, err := pkgconfig.CommandOutput("--libs", "ncurses")
	//	if err != nil {
	//		return nil, err
	//	}

	ncursesw_cflags, err := pkgconfig.CommandOutput("--cflags", "ncursesw")
	if err != nil {
		return nil, err
	}

	//	ncursesw_libs, err := pkgconfig.CommandOutput("--libs", "ncursesw")
	//	if err != nil {
	//		return nil, err
	//	}

	ret = append(
		ret,
		// NOTE: I don't like this very match!
		[]string{
			"CFLAGS=" + ncursesw_cflags,
		}...,
	)

	return ret, nil
}

func (self *Builder_xfs) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = make([]string, 0)

	ret = append(ret, "install")

	{
		add_install_dev := true
		for _, i := range []string{"acl", "attr"} {
			if info.PackageName == i {
				add_install_dev = false
				break
			}
		}
		if add_install_dev {
			ret = append(ret, "install-dev")
		}
	}

	ret = append(ret, "DESTDIR="+self.GetBuildingSiteCtl().GetDIR_DESTDIR())

	return ret, nil
}
