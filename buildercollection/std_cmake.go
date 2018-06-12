package buildercollection

import (
	"fmt"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["std_cmake"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_std_cmake(bs)
	}
}

type Builder_std_cmake struct {
	Builder_std
}

func NewBuilder_std_cmake(bs basictypes.BuildingSiteCtlI) (*Builder_std_cmake, error) {
	self := new(Builder_std_cmake)

	self.Builder_std = *NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions
	return self, nil
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

	return ret, nil
}

func (self *Builder_std_cmake) BuilderActionConfigureArgsDef(
	log *logger.Logger,
) ([]string, error) {
	calc := self.bs.GetBuildingSiteValuesCalculator()

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

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

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

	cmake := new(buildingtools.CMake)

	err = cmake.CMake(
		args,
		env.Strings(),
		buildingtools.Copy,
		"",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_BUILDING(),
		"",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}