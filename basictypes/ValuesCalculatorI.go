package basictypes

import "github.com/AnimusPEXUS/utils/environ"

type ValuesCalculatorI interface {
	CalculateIsCrossbuild() (bool, error)
	CalculateIsCrossbuilder() (bool, error)
	CalculateIsBuildingForSameHostButDifferentArch() (bool, error)

	CalculateMultihostDir() string
	CalculateDstMultihostDir() string

	CalculateHostDir() (string, error)
	CalculateDstHostDir() (string, error)

	CalculateHostMultiarchDir() (string, error)
	CalculateDstHostMultiarchDir() (string, error)

	CalculateHostArchDir() (string, error)
	CalculateDstHostArchDir() (string, error)

	// /{hostpath}/corssbuilders
	CalculateHostCrossbuildersDir() (string, error)
	CalculateDstHostCrossbuildersDir() (string, error)

	// /{hostpath}/corssbuilders/{target}
	CalculateHostCrossbuilderDir() (string, error)
	CalculateDstHostCrossbuilderDir() (string, error)

	CalculateHostLibDir() (string, error)
	CalculateDstHostLibDir() (string, error)

	CalculateHostArchLibDir() (string, error)
	CalculateDstHostArchLibDir() (string, error)

	CalculateInstallPrefix() (string, error)
	CalculateDstInstallPrefix() (string, error)

	CalculateInstallLibDir() (string, error)
	CalculateDstInstallLibDir() (string, error)

	CalculateMainMultiarchLibDirName() (string, error)

	CalculatePkgConfigSearchPaths(prefix string) ([]string, error)

	Calculate_LD_LIBRARY_PATH(prefixes []string) ([]string, error)
	Calculate_LIBRARY_PATH(prefixes []string) ([]string, error)
	Calculate_C_INCLUDE_PATH(prefixes []string) ([]string, error)
	Calculate_PATH(prefix string) ([]string, error)

	Calculate_C_Compiler() (string, error)
	Calculate_CXX_Compiler() (string, error)

	CalculateAutotoolsCCParameterValue() (string, error)
	CalculateAutotoolsCXXParameterValue() (string, error)

	CalculateAllOptionsMap() (environ.EnvVarEd, error)
	CalculateCompilerOptionsMap() (environ.EnvVarEd, error)

	CalculateAutotoolsHBTOptions() ([]string, error)
}
