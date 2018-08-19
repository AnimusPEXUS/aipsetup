package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

// TODO: libcap is deprecated and, probably, this builder will never be used..
//		delete?

func init() {
	Index["libcap"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libcap(bs)
	}
}

type Builder_libcap struct {
	*Builder_std
}

func NewBuilder_libcap(bs basictypes.BuildingSiteCtlI) (*Builder_libcap, error) {

	self := new(Builder_libcap)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditDistributeArgsCB = self.EditDistributeArgs
	self.AfterDistributeCB = self.AfterDistribute

	return self, nil
}

func (self *Builder_libcap) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("configure")
	ret = ret.Remove("autogen")
	ret = ret.Remove("build")
	ret = ret.Remove("patch")

	return ret, nil
}

func (self *Builder_libcap) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {
	//            'all',
	//            'install',
	//            'prefix={}'.format(self.calculate_install_prefix()),
	//            'lib={}'.format(lib),
	//            #'exec_prefix={}'.format(self.get_host_dir()),
	//            #'lib_prefix={}'.format(self.get_host_dir()),
	//            #'inc_prefix={}'.format(self.get_host_dir()),
	//            #'man_prefix={}'.format(self.get_host_dir()),
	//            'DESTDIR={}'.format(self.get_dst_dir()),
	//            #'PKGCONFIGDIR={}'.format(
	//            #    wayround_i2p.utils.path.join(
	//            #        self.get_dst_host_dir(),
	//            #       'lib',
	//            #        'pkgconfig'
	//            #        )
	//            #    ),
	//            'RAISE_SETFCAP=no',
	//            'PAM_CAP=yes',
	//            #'SYSTEM_HEADERS={}'.format(
	//            #    wayround_i2p.utils.path.join(
	//            #        self.get_host_dir(),
	//            #        'include'
	//            #        )
	//            #    ),
	//            #'CFLAGS=-I{}'.format(
	//            #    wayround_i2p.utils.path.join(
	//            #        self.get_host_dir(),
	//            #        'include'
	//            #        )
	//            #    )

	install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	main_multiarch_libdir_name, err := self.bs.GetBuildingSiteValuesCalculator().CalculateMainMultiarchLibDirName()
	if err != nil {
		return nil, err
	}

	ret = []string{
		"all",
		"install",
		"prefix=" + install_prefix,
		"lib=" + main_multiarch_libdir_name,
		"DESTDIR=" + self.bs.GetDIR_DESTDIR(),
		"RAISE_SETFCAP=no",
		"PAM_CAP=yes",
	}

	return ret, nil
}

func (self *Builder_libcap) AfterDistribute(log *logger.Logger, err error) error {
	if err != nil {
		return err
	}
	return nil
}
