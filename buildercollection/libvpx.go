package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libvpx"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libvpx(bs)
	}
}

type Builder_libvpx struct {
	*Builder_std
}

func NewBuilder_libvpx(bs basictypes.BuildingSiteCtlI) (*Builder_libvpx, error) {

	self := new(Builder_libvpx)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_libvpx) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	variant, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	target := "x86-linux-gcc"
	if variant == "64" {
		target = "x86_64-linux-gcc"
	}

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{
			"--includedir=",
			"--mandir=",
			"--docdir=",
			"--sysconfdir=",
			"--localstatedir=",
			"LDFLAGS=",
			"--host=",
			"--build=",
			"--target=",
			"CC=",
			"CXX=",
			"GCC=",
		},
	)
	if err != nil {
		return nil, err
	}

	ret = append(ret, "--target="+target)

	return ret, nil
}
