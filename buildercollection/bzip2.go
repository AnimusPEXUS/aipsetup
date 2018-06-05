package buildercollection

import (
	"fmt"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func init() {
	Index["bzip2"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_bzip2(bs)
	}
}

type Builder_bzip2 struct {
	Builder_std

	fixed_CC     string
	fixed_AR     string
	fixed_RANLIB string
}

func NewBuilder_bzip2(bs basictypes.BuildingSiteCtlI) (*Builder_bzip2, error) {
	//        thr['CC'] = 'CC={}-gcc -m{}'.format(
	//            self.get_host_from_pkgi(),
	//            self.get_multilib_variant_int()
	//            )
	//        thr['AR'] = 'AR={}-gcc-ar'.format(self.get_host_from_pkgi())
	//        thr['RANLIB'] = 'RANLIB={}-gcc-ranlib'.format(
	//            self.get_host_from_pkgi()
	//            )

	self := new(Builder_bzip2)

	self.Builder_std = *NewBuilder_std(bs)

	calc := bs.GetBuildingSiteValuesCalculator()

	info, err := bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	mlv, err := calc.CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	self.fixed_CC = fmt.Sprintf("CC=%s-gcc -m%s", info.Host, mlv)
	self.fixed_AR = fmt.Sprintf("AR=%s-gcc-ar -m%s", info.Host, mlv)
	self.fixed_RANLIB = fmt.Sprintf("RANLIB=%s-gcc-ranlib", info.Host)

	return self, nil
}
