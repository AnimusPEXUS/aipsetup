package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["cups"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_cups(bs)
	}
}

type Builder_cups struct {
	*Builder_std

	server_bin string
}

func NewBuilder_cups(bs basictypes.BuildingSiteCtlI) (*Builder_cups, error) {

	self := new(Builder_cups)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = func(log *logger.Logger, ret []string) ([]string, error) {
		ret = append(
			ret,
			[]string{
				"SERVERBIN=" + self.server_bin,
			}...,
		)
		return ret, nil
	}

	self.EditDistributeArgsCB = func(log *logger.Logger, ret []string) ([]string, error) {
		ret = append(
			ret,
			[]string{
				"SERVERBIN=" + self.server_bin,
				"BUILDROOT=" + self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
			}...,
		)
		return ret, nil
	}

	dst_install_libdir, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallLibDir()
	if err != nil {
		return nil, err
	}

	self.server_bin = path.Join(dst_install_libdir, "cups")

	return self, nil
}
