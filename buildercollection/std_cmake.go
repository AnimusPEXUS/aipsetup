package buildercollection

import (
	"fmt"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/cmake"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_cmake"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_cmake(bs)
	}
}

type Builder_std_cmake struct {
	*Builder_std

	EditConfigureArgsCB func(log *logger.Logger, ret []string) ([]string, error)
}

func NewBuilder_std_cmake(bs basictypes.BuildingSiteCtlI) (*Builder_std_cmake, error) {
	self := new(Builder_std_cmake)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir

	return self, nil
}

func (self *Builder_std_cmake) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
	return self.bs.GetDIR_BUILDING(), nil
}

func (self *Builder_std_cmake) EditActions(ret basictypes.BuilderActions) (
	basictypes.BuilderActions,
	error,
) {

	ret.Replace(
		"configure",
		&basictypes.BuilderAction{
			"configure",
			self.BuilderActionConfigure,
		},
	)

	ret = ret.Remove("autogen")

	return ret, nil
}

func (self *Builder_std_cmake) BuilderActionConfigureArgsDef(
	log *logger.Logger,
) ([]string, error) {
	calc := self.bs.GetBuildingSiteValuesCalculator()

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	//	host_dir := self.bs.GetSystem().GetSystemValuesCalculator().CalculateHostDir(info.Host)

	//	cm, err := cmake.NewCMake("", host_dir, "")
	//	if err != nil {
	//		return nil, err

	//	}

	//	cmake_root, err := cm.Get_CMAKE_ROOT()
	//	if err != nil {
	//		return nil, err
	//	}

	opt_map, err := calc.CalculateAutotoolsAllOptionsMap()
	if err != nil {
		return nil, err
	}

	minus_d_list := make([]string, 0)
	for _, i := range opt_map.Strings() {
		minus_d_list = append(minus_d_list, fmt.Sprintf("-D%s", i))
	}

	//        ret = [
	//            #'-DCMAKE_INSTALL_PREFIX={}'.format(
	//            #    self.calculate_install_prefix()
	//            #    ),
	//            #
	//            #'-DCMAKE_SYSROOT={}'.format(self.calculate_install_prefix()),
	//            '-DSYSCONFDIR=/etc',
	//            '-DLOCALSTATEDIR=/var',
	//            #'-DCMAKE_SYSTEM_PREFIX_PATH={}'.format(
	//            #    self.calculate_install_prefix()
	//            #    ),
	//            #'-DCMAKE_SYSTEM_INCLUDE_PATH={}'.format(
	//            #    wayround_i2p.utils.path.join(
	//            #        self.calculate_install_prefix(),
	//            #        'include'
	//            #        )
	//            #    ),
	//            # '-DCMAKE_FIND_ROOT_PATH={}'.format(
	//            #    self.calculate_install_prefix()
	//            #    ),
	//            ]

	ret := make([]string, 0)

	ret = append(
		ret,
		[]string{
			//			"-DCMAKE_ROOT=" + cmake_root,
			"-DSYSCONFDIR=/etc",
			"-DLOCALSTATEDIR=/var",
		}...,
	)

	std_opts, err := self.Builder_std.BuilderActionConfigureArgsDef(log)
	if err != nil {
		return nil, err
	}

	for _, i := range []string{
		"PREFIX",
		"BINDIR",
		"SBINDIR",
		"LIBEXECDIR",
		"SYSCONFDIR",
		"SHAREDSTATEDIR",
		"LOCALSTATEDIR",
		"LIBDIR",
		"INCLUDEDIR",
		"OLDINCLUDEDIR",
		"DATAROOTDIR",
		"DATADIR",
		"MANDIR",
		"DOCDIR",
	} {
		i_l_n := fmt.Sprintf("--%s=", strings.ToLower(i))
		for _, j := range std_opts {
			if strings.HasPrefix(j, i_l_n) {
				splitted := strings.SplitN(j, "=", 2)
				if len(splitted) == 2 {
					ret = append(
						ret,
						fmt.Sprintf(
							"-DCMAKE_INSTALL_%s=%s",
							i,
							splitted[1],
						),
					)
				}
			}
		}
	}

	ret = append(ret, minus_d_list...)

	// TODO: better calculation required
	if strings.HasPrefix(info.HostArch, "x86_64") {
		ret = append(
			ret,
			[]string{
				"-DLIB_SUFFIX=64",
				"-DLIBDIR_SUFFIX=64",
			}...,
		)
	}

	if self.EditConfigureArgsCB != nil {

		ret, err = self.EditConfigureArgsCB(log, ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_std_cmake) BuilderActionConfigure(
	log *logger.Logger,
) error {

	env, err := self.BuilderActionConfigureEnvDef(log)
	if err != nil {
		return err
	}

	args, err := self.BuilderActionConfigureArgsDef(log)
	if err != nil {
		return err
	}

	calc := self.bs.GetBuildingSiteValuesCalculator()

	prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return err
	}

	cm, err := cmake.NewCMake("", prefix, "")
	if err != nil {
		return err
	}

	cd, err := self.BuilderActionConfigureDirDef(log)
	if err != nil {
		return err
	}

	wd, err := self.BuilderActionConfigureWorkingDirDef(log)
	if err != nil {
		return err
	}

	cmake := new(buildingtools.CMake)

	err = cmake.CMake(
		args,
		env.Strings(),
		buildingtools.Copy,
		"",
		cd,
		wd,
		cm.GetExecutable(),
		log,
	)
	if err != nil {
		return err
	}

	return nil
}
