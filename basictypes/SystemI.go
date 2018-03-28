package basictypes

type SystemI interface {
	Root() string
	Host() (string, error)
	Archs() ([]string, error)
	GetInstalledASPDir() string
	GetInstalledASPSumsDir() string
	GetInstalledASPBuildLogsDir() string
	GetInstalledASPDepsDir() string
	GetTarballRepoRootDir() string

	GetSystemValuesCalculator() SystemValuesCalculatorI
	// GetTarballsRepository() types.RepositoryI
}
