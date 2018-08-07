package basictypes

import (
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/pkgconfig"
)

type BuildingSiteValuesCalculatorI interface {
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

	// /multihost/{host}/{libname}
	CalculateHostLibDir() (string, error)
	CalculateDstHostLibDir() (string, error)

	// /multihost/{host}/multiarch/{arch}/{libname}
	CalculateHostArchLibDir() (string, error)
	CalculateDstHostArchLibDir() (string, error)

	CalculateInstallPrefix() (string, error)
	CalculateDstInstallPrefix() (string, error)

	// /multihost/{host}/{libname}
	// or
	// /multihost/{host}/multiarch/{hostarch}/{libname}
	// depending on host != hostarch
	CalculateInstallLibDir() (string, error)
	CalculateDstInstallLibDir() (string, error)

	CalculateMainMultiarchLibDirName() (string, error)

	//	CalculatePkgConfigSearchPaths() ([]string, error)

	Calculate_LD_LIBRARY_PATH() ([]string, error)
	Calculate_LIBRARY_PATH() ([]string, error)
	Calculate_C_INCLUDE_PATH() ([]string, error)
	Calculate_PATH() ([]string, error)

	Calculate_C_Compiler() (string, error)
	Calculate_CXX_Compiler() (string, error)
	CalculateMultilibVariant() (string, error)

	CalculateAutotoolsCCParameterValue() (string, error)
	CalculateAutotoolsCXXParameterValue() (string, error)

	CalculateAutotoolsCompilerOptionsMap() (environ.EnvVarEd, error)
	CalculateAutotoolsAllOptionsMap() (environ.EnvVarEd, error)

	CalculateAutotoolsHBTOptions() ([]string, error)

	CalculateCmakeAllOptionsMap() (environ.EnvVarEd, error)

	CalculateOptAppDir(name string) string

	CalculateInstallPrefixExecutable(name string) (string, error)

	GetPrefixPkgConfig() (*pkgconfig.PkgConfig, error)
}
