package aipsetup

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/systemtriplet"
	"github.com/AnimusPEXUS/utils/textlist"
)

var _ basictypes.ValuesCalculatorI = &ValuesCalculator{}

type ValuesCalculator struct {
	site    *BuildingSiteCtl
	builder basictypes.BuilderI
}

func NewValuesCalculator(site *BuildingSiteCtl) *ValuesCalculator {
	ret := new(ValuesCalculator)
	ret.site = site
	//ret.builder = builder
	return ret
}

/*
	returns true, if building site configured to build for
	[host] which not equal to [host aipsetup configured in /etc/aipsetup5.system.ini]
*/
func (self *ValuesCalculator) CalculateIsCrossbuild() (bool, error) {
	host, _, _, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return false, err
	}

	build := self.site.sys.Host()

	ret := host != build

	return ret, nil
}

/*
	returns true if target != host
*/
func (self *ValuesCalculator) CalculateIsCrossbuilder() (bool, error) {

	host, hostarch, target, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return false, err
	}

	ret := target != host

	if host != hostarch {
		return false, errors.New("invalid configuration")
	}

	return ret, nil
}

func (self *ValuesCalculator) CalculateIsBuildingForSameHostButDifferentArch() (bool, error) {
	host, hostarch, _, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return false, err
	}

	build := self.site.sys.Host()

	ret := (host == build) && hostarch != host

	return ret, err
}

func (self *ValuesCalculator) CalculateMultihostDir() string {
	return path.Join(string(os.PathSeparator), LAILALO_ROOT_MULTIHOST_DIRNAME)
}

func (self *ValuesCalculator) CalculateDstMultihostDir() string {
	return path.Join(self.site.GetDIR_DESTDIR(), self.CalculateMultihostDir())
}

func (self *ValuesCalculator) CalculateHostDir() (string, error) {
	host, err := self.site.GetConfiguredHost()
	if err != nil {
		return "", err
	}
	return path.Join(self.CalculateMultihostDir(), host), nil
}

func (self *ValuesCalculator) CalculateDstHostDir() (string, error) {
	hostdir, err := self.CalculateHostDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), hostdir), nil
}

func (self *ValuesCalculator) CalculateHostMultiarchDir() (string, error) {
	hostdir, err := self.CalculateHostDir()
	if err != nil {
		return "", err
	}
	return path.Join(hostdir, LAILALO_MULTIHOST_MULTIARCH_DIRNAME), nil
}

func (self *ValuesCalculator) CalculateDstHostMultiarchDir() (string, error) {
	hostmultiarchdir, err := self.CalculateHostMultiarchDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), hostmultiarchdir), nil
}

func (self *ValuesCalculator) CalculateHostArchDir() (string, error) {
	hostmultiarchdir, err := self.CalculateHostMultiarchDir()
	if err != nil {
		return "", err
	}

	arch, err := self.site.GetConfiguredHostArch()
	if err != nil {
		return "", err
	}

	return path.Join(hostmultiarchdir, arch), nil
}

func (self *ValuesCalculator) CalculateDstHostArchDir() (string, error) {
	hostarchdir, err := self.CalculateHostArchDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), hostarchdir), nil
}

// /{hostpath}/corssbuilders
func (self *ValuesCalculator) CalculateHostCrossbuildersDir() (string, error) {
	hostdir, err := self.CalculateHostDir()
	if err != nil {
		return "", err
	}
	return path.Join(hostdir, LAILALO_MULTIHOST_CROSSBULDERS_DIRNAME), nil
}

func (self *ValuesCalculator) CalculateDstHostCrossbuildersDir() (string, error) {
	hostcrossbuildersdir, err := self.CalculateHostCrossbuildersDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), hostcrossbuildersdir), nil
}

// /{hostpath}/corssbuilders/{target}
func (self *ValuesCalculator) CalculateHostCrossbuilderDir() (string, error) {
	hostcrossbuildersdir, err := self.CalculateHostCrossbuildersDir()
	if err != nil {
		return "", err
	}
	target, err := self.site.GetConfiguredTarget()
	if err != nil {
		return "", err
	}
	return path.Join(hostcrossbuildersdir, target), nil
}

func (self *ValuesCalculator) CalculateDstHostCrossbuilderDir() (string, error) {
	hostcrossbuilderdir, err := self.CalculateHostCrossbuilderDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), hostcrossbuilderdir), nil
}

func (self *ValuesCalculator) CalculateHostLibDir() (string, error) {
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

func (self *ValuesCalculator) CalculateDstHostLibDir() (string, error) {
	host_lib_dir, err := self.CalculateHostLibDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), host_lib_dir), nil
}

func (self *ValuesCalculator) CalculateHostArchLibDir() (string, error) {
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

func (self *ValuesCalculator) CalculateDstHostArchLibDir() (string, error) {
	v, err := self.CalculateHostArchLibDir()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), v), nil
}

func (self *ValuesCalculator) CalculateInstallPrefix() (string, error) {
	host, hostarch, _, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return "", err
	}

	if host == hostarch {
		return self.CalculateHostDir()
	} else {
		return self.CalculateHostArchDir()
	}
}

func (self *ValuesCalculator) CalculateDstInstallPrefix() (string, error) {
	v, err := self.CalculateInstallPrefix()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), v), nil
}

func (self *ValuesCalculator) CalculateInstallLibDir() (string, error) {
	host, hostarch, _, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return "", err
	}

	if host == hostarch {
		return self.CalculateHostLibDir()
	} else {
		return self.CalculateHostArchLibDir()
	}
}

func (self *ValuesCalculator) CalculateDstInstallLibDir() (string, error) {
	v, err := self.CalculateInstallPrefix()
	if err != nil {
		return "", err
	}
	return path.Join(self.site.GetDIR_DESTDIR(), v), nil
}

// def get_host_arch_list(self):
// 		ret = []
//
// 		lst = os.listdir(self.get_host_multiarch_dir())
//
// 		for i in lst:
// 				jo = wayround_i2p.utils.path.join(
// 						self.get_host_multiarch_dir(),
// 						i
// 						)
// 				if os.path.isdir(jo) and not os.path.islink(jo):
// 						ret.append(i)
//
// 		return sorted(ret)

// # def calculate_default_linker_program(self):
// #    return wayround_i2p.aipsetup.build.find_dl(self.get_host_arch_dir())
//
// # def calculate_default_linker_program_ld_parameter(self):
// #    return '--dynamic-linker={}'.format(
// #        self.calculate_default_linker_program()
// #        )
//
// # def calculate_default_linker_program_gcc_parameter(self):
// #    return '-Wl,{}'.format(
// #        self.calculate_default_linker_program_ld_parameter()
// #        )

func (self *ValuesCalculator) CalculateMainMultiarchLibDirName() (string, error) {
	host, hostarch, _, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return "", err
	}

	switch host {

	case "i686-pc-linux-gnu":
		switch hostarch {
		case "i686-pc-linux-gnu":
			return "lib", nil
		}

	case "x86_64-pc-linux-gnu":
		switch hostarch {
		case "i686-pc-linux-gnu":
			return "lib", nil
		case "x86_64-pc-linux-gnu":
			return "lib64", nil
		}
	}

	return "", errors.New("host or [host/arch] value not supported")
}

func (self *ValuesCalculator) CalculatePkgConfigSearchPaths(prefix string) ([]string, error) {

	inst_prefix, err := self.CalculateInstallPrefix()
	if err != nil {
		return []string{}, err
	}

	ret := make([]string, 0)

	if prefix == "" {
		var err error
		prefix, err = self.CalculateInstallPrefix()
		if err != nil {
			return []string{}, nil
		}
	}

	for _, i := range []string{
		path.Join(prefix, "share", "pkgconfig"),
		path.Join(prefix, "lib", "pkgconfig"),
		path.Join(prefix, "lib64", "pkgconfig"),
		path.Join(inst_prefix, "share", "pkgconfig"),
		path.Join(inst_prefix, "lib", "pkgconfig"),
		path.Join(inst_prefix, "lib64", "pkgconfig"),
	} {
		if s, err := os.Stat(i); err == nil && s.IsDir() {
			ret = append(ret, i)
		}
	}

	ret = textlist.RemoveDuplicatedStrings(ret)

	return ret, nil
}

func (self *ValuesCalculator) Calculate_LD_LIBRARY_PATH(prefixes []string) ([]string, error) {

	// inst_prefix, err := self.CalculateInstallPrefix()
	// if err != nil {
	// 	return []string{}, err
	// }

	host_dir, err := self.CalculateHostDir()
	if err != nil {
		return []string{}, err
	}

	ret := make([]string, 0)

	search_roots := make([]string, 0)

	search_roots = append(search_roots, host_dir)
	search_roots = append(search_roots, prefixes...)

	search_roots = textlist.RemoveDuplicatedStrings(search_roots)

	for _, i := range search_roots {
		for _, j := range POSSIBLE_LIBDIR_NAMES {
			joined := path.Join(i, j)
			if s, err := os.Stat(joined); err == nil && s.IsDir() {
				ret = append(ret, joined)
			}
		}
	}

	ret = textlist.RemoveDuplicatedStrings(ret)

	return ret, nil
}

func (self *ValuesCalculator) Calculate_LIBRARY_PATH(prefixes []string) ([]string, error) {
	// # NOTE: potentially this is different from LD_LIBRARY_PATH.
	// #       LIBRARY_PATH is for GCC and it's friends. so it's possible
	// #       for it to differ also in code, in future, not only in name.
	// ret = self.calculate_LD_LIBRARY_PATH(prefix)
	return self.Calculate_LD_LIBRARY_PATH(prefixes)
}

func (self *ValuesCalculator) Calculate_C_INCLUDE_PATH(prefixes []string) ([]string, error) {

	inst_prefix, err := self.CalculateInstallPrefix()
	if err != nil {
		return []string{}, err
	}

	ret := make([]string, 0)

	search_roots := make([]string, 0)

	if len(prefixes) != 0 {
		search_roots = append(search_roots, prefixes...)
	} else {
		search_roots = append(search_roots, inst_prefix)
	}

	search_roots = textlist.RemoveDuplicatedStrings(search_roots)

	for _, i := range search_roots {
		joined := path.Join(i, "include")
		if s, err := os.Stat(joined); err == nil && s.IsDir() {
			ret = append(ret, joined)
		}
	}

	ret = textlist.RemoveDuplicatedStrings(ret)

	return ret, nil
}

func (self *ValuesCalculator) Calculate_PATH(prefix string) ([]string, error) {

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

	if prefix != "" {
		if s, err := os.Stat(prefix); err == nil && s.IsDir() {
			search_roots = append(search_roots, prefix)
		} else {
			return []string{}, err
		}
	}

	search_roots = append(search_roots, inst_prefix)
	search_roots = append(search_roots, host_dir)

	search_roots = textlist.RemoveDuplicatedStrings(search_roots)

	for _, i := range search_roots {
		for _, j := range PATH_CALCULATOR_BIN_DIR_NAMES {
			joined := path.Join(i, j)
			if s, err := os.Stat(joined); err == nil && s.IsDir() {
				ret = append(ret, joined)
			}
		}
	}

	ret = textlist.RemoveDuplicatedStrings(ret)

	return ret, nil
}

func (self *ValuesCalculator) Calculate_C_Compiler() (string, error) {
	host, err := self.site.GetConfiguredHost()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s", host, "gcc"), nil
}

func (self *ValuesCalculator) Calculate_CXX_Compiler() (string, error) {
	host, err := self.site.GetConfiguredHost()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s", host, "g++"), nil
}

func (self *ValuesCalculator) CalculateMultilibVariant() (string, error) {
	_, hostarch, _, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return "", err
	}
	tr, err := systemtriplet.NewFromString(hostarch)
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

func (self *ValuesCalculator) CalculateAutotoolsCCParameterValue() (string, error) {
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

func (self *ValuesCalculator) CalculateAutotoolsCXXParameterValue() (string, error) {
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

func (self *ValuesCalculator) CalculateAllOptionsMap() (environ.EnvVarEd, error) {

	ret := environ.New()

	c, err := self.CalculateCompilerOptionsMap()
	if err != nil {
		return ret, err
	}

	ret.UpdateWith(c)

	return ret, nil
}

func (self *ValuesCalculator) CalculateCompilerOptionsMap() (environ.EnvVarEd, error) {
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

func (self *ValuesCalculator) CalculateAutotoolsHBTOptions() ([]string, error) {
	ret := make([]string, 0)

	host, hostarch, target, err := self.site.GetConfiguredHHAT()
	if err != nil {
		return ret, err
	}

	// NOTE: possibly some builders may require forcing crossbuilder creation
	//       but apperently builder tool world going some other way, while
	//       autotools mainly stays same: configure --target != --host indicates
	//       so
	forced_target := false

	build := self.site.sys.Host()

	if hostarch != "" &&
		(((host == build) && (build == target)) ||
			((hostarch == build) && (build == target)) ||
			((host == build) && (target == ""))) &&
		!forced_target {
		target = ""
	}

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
