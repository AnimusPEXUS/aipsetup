package basictypes

type SystemValuesCalculatorI interface {
	CalculateMultihostDir() string
	CalculateHostDir(host string) (string, error)
	CalculateHostMultiarchDir(host string) (string, error)
	CalculateHostArchDir(host, hostarch string) (string, error)
	CalculateHostCrossbuildersDir(host string) (string, error)
	CalculateHostCrossbuilderDir(host, target string) (string, error)

	// TODO: Not this time, but maybe this shold be moved here from
	//       BuildingSiteValuesCalculator in the future.
	//
	// CalculateHostLibDir(host) (string, error)
	// CalculateHostArchLibDir() (string, error)
	// CalculateInstallPrefix() (string, error)
	// CalculateInstallLibDir() (string, error)
	// CalculateMainMultiarchLibDirName() (string, error)
	//
	// CalculatePkgConfigSearchPaths(prefix string) ([]string, error)
	//
	// Calculate_LD_LIBRARY_PATH(prefixes []string) ([]string, error)
	// Calculate_LIBRARY_PATH(prefixes []string) ([]string, error)
	// Calculate_C_INCLUDE_PATH(prefixes []string) ([]string, error)
	// Calculate_PATH(prefix string) ([]string, error)
	//
	// Calculate_C_Compiler() (string, error)
	// Calculate_CXX_Compiler() (string, error)
	//
	// CalculateAutotoolsCCParameterValue() (string, error)
	// CalculateAutotoolsCXXParameterValue() (string, error)
	//
	// CalculateAllOptionsMap() (environ.EnvVarEd, error)
	// CalculateCompilerOptionsMap() (environ.EnvVarEd, error)
	//
	// CalculateAutotoolsHBTOptions() ([]string, error)
}
