package aipsetup

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/pkgconfig"
	"github.com/AnimusPEXUS/utils/systemtriplet"
	"github.com/AnimusPEXUS/utils/textlist"
)

var _ basictypes.BuildingSiteValuesCalculatorI = &BuildingSiteValuesCalculator{}

// TODO: possibly some of BuildingSiteValuesCalculator functions should be moved
//		to SystemValuesCalculator

type BuildingSiteValuesCalculator struct {
	site basictypes.BuildingSiteCtlI
	opc  OverallPathsCalculator
}

func NewBuildingSiteValuesCalculator(
	site basictypes.BuildingSiteCtlI,
) *BuildingSiteValuesCalculator {
	ret := new(BuildingSiteValuesCalculator)
	ret.site = site
	return ret
}

func (self *BuildingSiteValuesCalculator) CalculateMultihostDir() string {
	return self.opc.CalculateMultihostDir("/")
}

func (self *BuildingSiteValuesCalculator) CalculateDstMultihostDir() string {
	return self.opc.CalculateMultihostDir(self.site.GetDIR_DESTDIR())
}

func (self *BuildingSiteValuesCalculator) CalculateHostDir() (string, error) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return self.opc.CalculateHostDir("/", info.Host), nil
}

func (self *BuildingSiteValuesCalculator) CalculateDstHostDir() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return self.opc.CalculateHostDir(self.site.GetDIR_DESTDIR(), info.Host), nil
}

func (self *BuildingSiteValuesCalculator) CalculateHostMultiarchDir() (
	string, error,
) {
	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return self.opc.CalculateHostMultiarchDir("/", info.Host), nil
}

func (self *BuildingSiteValuesCalculator) CalculateDstHostMultiarchDir() (
	string, error,
) {
	host, err := self.CalculateHostMultiarchDir()
	if err != nil {
		return "", err
	}
	return self.opc.CalculateHostMultiarchDir(self.site.GetDIR_DESTDIR(), host),
		nil
}

func (self *BuildingSiteValuesCalculator) CalculateHostArchDir() (string, error) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return self.opc.CalculateHostArchDir("/", info.Host, info.HostArch), nil
}

func (self *BuildingSiteValuesCalculator) CalculateDstHostArchDir() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return self.opc.CalculateHostArchDir(
		self.site.GetDIR_DESTDIR(), info.Host, info.HostArch,
	), nil
}

// /{hostpath}/corssbuilders
func (self *BuildingSiteValuesCalculator) CalculateHostCrossbuildersDir() (
	string,
	error,
) {
	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}
	return self.opc.CalculateHostCrossbuildersDir("/", info.Host), nil
}

func (self *BuildingSiteValuesCalculator) CalculateDstHostCrossbuildersDir() (
	string, error,
) {
	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}
	return self.opc.CalculateHostCrossbuildersDir(
			self.site.GetDIR_DESTDIR(),
			info.Host,
		),
		nil
}

// /{hostpath}/corssbuilders/{target}
func (self *BuildingSiteValuesCalculator) CalculateHostCrossbuilderDir() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return self.opc.CalculateHostCrossbuilderDir(
			"/",
			info.Host,
			info.CrossbuilderTarget,
		),
		nil
}

func (self *BuildingSiteValuesCalculator) CalculateDstHostCrossbuilderDir() (
	string, error,
) {
	hostcrossbuilderdir, err := self.CalculateHostCrossbuilderDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), hostcrossbuilderdir), nil
}

func (self *BuildingSiteValuesCalculator) CalculateHostLibDir() (
	string, error,
) {
	hostdir, err := self.CalculateHostDir()
	if err != nil {
		return "", err
	}
	mmldn, err := self.CalculateMainMultiarchLibDirName()
	if err != nil {
		return "", err
	}
	return path.Join(hostdir, mmldn), nil
}

func (self *BuildingSiteValuesCalculator) CalculateDstHostLibDir() (
	string, error,
) {
	host_lib_dir, err := self.CalculateHostLibDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), host_lib_dir), nil
}

func (self *BuildingSiteValuesCalculator) CalculateHostArchLibDir() (
	string, error,
) {
	lib_dir_name, err := self.CalculateMainMultiarchLibDirName()
	if err != nil {
		return "", err
	}

	host_arch_dir, err := self.CalculateHostArchDir()
	if err != nil {
		return "", err
	}

	return path.Join(host_arch_dir, lib_dir_name), nil
}

func (self *BuildingSiteValuesCalculator) CalculateDstHostArchLibDir() (
	string, error,
) {
	v, err := self.CalculateHostArchLibDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), v), nil
}

func (self *BuildingSiteValuesCalculator) CalculateInstallPrefix() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	if info.Host == info.HostArch {
		return self.CalculateHostDir()
	} else {
		return self.CalculateHostArchDir()
	}
}

func (self *BuildingSiteValuesCalculator) CalculateDstInstallPrefix() (
	string, error,
) {
	v, err := self.CalculateInstallPrefix()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), v), nil
}

func (self *BuildingSiteValuesCalculator) CalculateInstallLibDir() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	if info.Host == info.HostArch {
		return self.CalculateHostLibDir()
	} else {
		return self.CalculateHostArchLibDir()
	}
}

func (self *BuildingSiteValuesCalculator) CalculateDstInstallLibDir() (
	string, error,
) {
	v, err := self.CalculateInstallLibDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), v), nil
}

func (self *BuildingSiteValuesCalculator) CalculateMainMultiarchLibDirName() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	switch info.Host {

	case basictypes.I686_PC_LINUX_GNU:
		switch info.HostArch {

		case basictypes.I686_PC_LINUX_GNU:
			return basictypes.DIRNAME_LIB, nil

		}

	case basictypes.X86_64_PC_LINUX_GNU:
		switch info.HostArch {

		case basictypes.I686_PC_LINUX_GNU:
			return basictypes.DIRNAME_LIB, nil

		case basictypes.X86_64_PC_LINUX_GNU:
			return basictypes.DIRNAME_LIB64, nil
		}
	}

	return "", errors.New("host or [host/hostarch] value not supported")
}

//func (self *BuildingSiteValuesCalculator) CalculatePkgConfigSearchPaths() ([]string, error) {

//	inst_prefix, err := self.CalculateInstallPrefix()
//	if err != nil {
//		return []string{}, err
//	}

//	ret := make([]string, 0)

//	prefix, err := self.CalculateInstallPrefix()
//	if err != nil {
//		return []string{}, nil
//	}

//	for _, i := range []string{
//		path.Join(prefix, basictypes.DIRNAME_SHARE, "pkgconfig"),
//		path.Join(prefix, basictypes.DIRNAME_LIB, "pkgconfig"),
//		path.Join(prefix, basictypes.DIRNAME_LIB64, "pkgconfig"),
//		path.Join(inst_prefix, basictypes.DIRNAME_SHARE, "pkgconfig"),
//		path.Join(inst_prefix, basictypes.DIRNAME_LIB, "pkgconfig"),
//		path.Join(inst_prefix, basictypes.DIRNAME_LIB64, "pkgconfig"),
//	} {
//		if s, err := os.Stat(i); err == nil && s.IsDir() {
//			ret = append(ret, i)
//		}
//	}

//	ret = textlist.RemoveDuplicatedStrings(ret)

//	return ret, nil
//}

func (self *BuildingSiteValuesCalculator) Calculate_LD_LIBRARY_PATH() (
	[]string,
	error,
) {

	// info, err := self.site.ReadInfo()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// host_dir, err := self.CalculateHostDir()
	// if err != nil {
	// 	return []string{}, err
	// }
	//
	// if info.Host != info.HostArch {
	// 	host_dir, err = self.CalculateHostArchDir()
	// 	if err != nil {
	// 		return []string{}, err
	// 	}
	// }
	//
	// ret := make([]string, 0)
	//
	// search_roots := make([]string, 0)
	//
	// search_roots = append(search_roots, host_dir)
	// search_roots = append(search_roots, prefixes...)
	//
	// search_roots = textlist.RemoveDuplicatedStrings(search_roots)
	//
	// for _, i := range search_roots {
	// 	for _, j := range basictypes.POSSIBLE_LIBDIR_NAMES {
	// 		joined := path.Join(i, j)
	// 		if s, err := os.Stat(joined); err == nil && s.IsDir() {
	// 			ret = append(ret, joined)
	// 		}
	// 	}
	// }
	//
	// ret = textlist.RemoveDuplicatedStrings(ret)

	sys := self.site.GetSystem()

	host, err := sys.Host()
	if err != nil {
		return nil, err
	}

	archs, err := sys.Archs()
	if err != nil {
		return nil, err
	}

	info, err := self.site.ReadInfo()
	if err != nil {
		return nil, err
	}

	prefixes := make([]string, 0)

	prefixes = append(
		prefixes,
		self.opc.CalculateHostArchDir("/", info.Host, info.HostArch),
	)

	for _, arch := range archs {
		// TODO: something wrong here
		prefixes = append(
			prefixes,
			self.opc.CalculateHostArchDir("/", host, arch),
		)
	}

	ret := make([]string, 0)

	for _, i := range prefixes {
		for _, j := range basictypes.POSSIBLE_LIBDIR_NAMES {
			joined := path.Join(i, j)
			s, err := os.Stat(joined)
			if err == nil && s.IsDir() {
				ret = append(ret, joined)
			}
		}
	}

	return ret, nil
}

func (self *BuildingSiteValuesCalculator) Calculate_LIBRARY_PATH() ([]string, error) {
	// # NOTE: potentially this is different from LD_LIBRARY_PATH.
	// #       LIBRARY_PATH is for GCC and it's friends. so it's possible
	// #       for it to differ also in code, in future, not only in name.
	// ret = self.calculate_LD_LIBRARY_PATH(prefix)
	return self.Calculate_LD_LIBRARY_PATH()
}

func (self *BuildingSiteValuesCalculator) Calculate_C_INCLUDE_PATH() ([]string, error) {

	inst_prefix, err := self.CalculateInstallPrefix()
	if err != nil {
		return []string{}, err
	}

	ret := make([]string, 0)

	search_roots := make([]string, 0)

	search_roots = append(search_roots, inst_prefix)

	search_roots = textlist.RemoveDuplicatedStrings(search_roots)

	for _, i := range search_roots {
		joined := path.Join(i, basictypes.DIRNAME_INCLUDE)
		if s, err := os.Stat(joined); err == nil && s.IsDir() {
			ret = append(ret, joined)
		}
	}

	ret = textlist.RemoveDuplicatedStrings(ret)

	return ret, nil
}

func (self *BuildingSiteValuesCalculator) Calculate_PATH() (
	[]string, error,
) {

	inst_prefix, err := self.CalculateInstallPrefix()
	if err != nil {
		return []string{}, err
	}

	host_dir, err := self.CalculateHostDir()
	if err != nil {
		return []string{}, err
	}

	ret := make([]string, 0)

	search_roots := make([]string, 0)

	search_roots = append(search_roots, inst_prefix)
	search_roots = append(search_roots, host_dir)

	search_roots = textlist.RemoveDuplicatedStrings(search_roots)

	for _, i := range search_roots {
		for _, j := range basictypes.PATH_CALCULATOR_BIN_DIR_NAMES {
			joined := path.Join(i, j)
			if s, err := os.Stat(joined); err == nil && s.IsDir() {
				ret = append(ret, joined)
			}
		}
	}

	ret = textlist.RemoveDuplicatedStrings(ret)

	return ret, nil
}

func (self *BuildingSiteValuesCalculator) Calculate_C_Compiler() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", info.Host, "gcc"), nil
}

func (self *BuildingSiteValuesCalculator) Calculate_CXX_Compiler() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", info.Host, "g++"), nil
}

func (self *BuildingSiteValuesCalculator) CalculateMultilibVariant() (
	string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return "", err
	}

	tr, err := systemtriplet.NewFromString(info.HostArch)
	if err != nil {
		return "", err
	}

	switch tr.CPU {
	case "x86_64":
		return "64", nil
	case "i486":
		fallthrough
	case "i586":
		fallthrough
	case "i686":
		return "32", nil
	}

	return "", errors.New("CalculateMultilibVariant(): not supported cpu")
}

func (self *BuildingSiteValuesCalculator) CalculateAutotoolsHBTOptions() (
	[]string, error,
) {

	info, err := self.site.ReadInfo()
	if err != nil {
		return nil, err
	}

	// NOTE: possibly some builders may require forcing crossbuilder creation
	//       but apperently builder tool world going some other way, while
	//       autotools mainly stays same: configure --target != --host indicates
	//       so
	forced_target := false

	build, err := self.site.GetSystem().Host()
	if err != nil {
		return nil, err
	}

	host := info.Host
	hostarch := info.HostArch
	target := info.CrossbuilderTarget

	if hostarch != "" &&
		(((host == build) && (build == target)) ||
			((hostarch == build) && (build == target)) ||
			((host == build) && (target == ""))) &&
		!forced_target {
		target = ""
	}

	if info.HostArch != info.Host {
		host = info.HostArch
		build = info.HostArch
	}

	ret := make([]string, 0)

	if host != "" {
		ret = append(ret, fmt.Sprintf("--host=%s", host))
	}

	if build != "" {
		ret = append(ret, fmt.Sprintf("--build=%s", build))
	}

	if target != "" {
		ret = append(ret, fmt.Sprintf("--target=%s", target))
	}

	return ret, nil
}

func (self *BuildingSiteValuesCalculator) CalculateAutotoolsCCParameterValue() (
	string, error,
) {
	c, err := self.Calculate_C_Compiler()
	if err != nil {
		return "", err
	}
	mlv, err := self.CalculateMultilibVariant()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s -m%s", c, mlv), nil
}

func (self *BuildingSiteValuesCalculator) CalculateAutotoolsCXXParameterValue() (
	string, error,
) {
	c, err := self.Calculate_CXX_Compiler()
	if err != nil {
		return "", err
	}
	mlv, err := self.CalculateMultilibVariant()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s -m%s", c, mlv), nil
}

func (self *BuildingSiteValuesCalculator) CalculateAutotoolsCompilerOptionsMap() (
	environ.EnvVarEd, error,
) {
	ret := environ.New()

	cc_string, err := self.CalculateAutotoolsCCParameterValue()
	if err != nil {
		return ret, err
	}

	cxx_string, err := self.CalculateAutotoolsCXXParameterValue()
	if err != nil {
		return ret, err
	}

	ret.Set("CC", cc_string)
	ret.Set("CXX", cxx_string)

	return ret, nil
}

func (self *BuildingSiteValuesCalculator) CalculateAutotoolsAllOptionsMap() (
	environ.EnvVarEd, error,
) {

	ret := environ.New()

	c, err := self.CalculateAutotoolsCompilerOptionsMap()
	if err != nil {
		return ret, err
	}

	ret.UpdateWith(c)

	return ret, nil
}

func (self *BuildingSiteValuesCalculator) CalculateCmakeAllOptionsMap() (
	environ.EnvVarEd, error,
) {
	//        if not 'CMAKE_C_COMPILER' in d:
	//            d['CMAKE_C_COMPILER'] = []
	//        d['CMAKE_C_COMPILER'].append('{}'.format(self.calculate_CC()))

	//        if not 'CMAKE_CXX_COMPILER' in d:
	//            d['CMAKE_CXX_COMPILER'] = []
	//        d['CMAKE_CXX_COMPILER'].append('{}'.format(self.calculate_CXX()))

	//        if not 'CMAKE_CXX_FLAGS' in d:
	//            d['CMAKE_CXX_FLAGS'] = []
	//        d['CMAKE_CXX_FLAGS'].append(
	//            '-m{}'.format(self.get_multilib_variant_int())
	//            )

	//        if not 'CMAKE_C_FLAGS' in d:
	//            d['CMAKE_C_FLAGS'] = []
	//        d['CMAKE_C_FLAGS'].append(
	//            '-m{}'.format(self.get_multilib_variant_int())
	//            )
	//        # d['CMAKE_C_FLAGS'].append(
	//        #    '-I{}'.format(
	//        #        wayround_i2p.utils.path.join(
	//        #            self.calculate_install_prefix(),
	//        #            'include'
	//        #            )
	//        #        )
	//        #    )

	c_compiler, err := self.Calculate_C_Compiler()
	if err != nil {
		return nil, err
	}

	cxx_compiler, err := self.Calculate_CXX_Compiler()
	if err != nil {
		return nil, err
	}

	multilib_variant, err := self.CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	ret := environ.New()

	ret.Set("CMAKE_C_COMPILER", c_compiler)
	ret.Set("CMAKE_CXX_COMPILER", cxx_compiler)
	ret.Set("CMAKE_C_FLAGS", fmt.Sprintf("-m%s", multilib_variant))
	ret.Set("CMAKE_CXX_FLAGS", fmt.Sprintf("-m%s", multilib_variant))

	return ret, nil
}

func (self *BuildingSiteValuesCalculator) CalculateOptAppDir(name string) string {
	// TODO: finish this
	return ""
}

func (self *BuildingSiteValuesCalculator) CalculateInstallPrefixExecutable(name string) (string, error) {

	PATH, err := self.Calculate_PATH()
	if err != nil {
		return "", err
	}

	program, err := filetools.Which(name, PATH)
	if err != nil {
		return "", err
	}

	return program, nil
}

func (self *BuildingSiteValuesCalculator) GetPrefixPkgConfig() (*pkgconfig.PkgConfig, error) {

	path, err := self.Calculate_PATH()
	if err != nil {
		return nil, err
	}

	ret, err := pkgconfig.NewPkgConfig(path, []string{})
	if err != nil {
		return nil, err
	}

	return ret, nil
}
