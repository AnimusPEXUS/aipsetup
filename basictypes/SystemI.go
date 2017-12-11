package basictypes

type SystemI interface {
	Root() string
	Host() string
	Archs() []string
	GetInstalledASPDir() string
	GetInstalledASPSumsDir() string
	GetInstalledASPBuildLogsDir() string
	GetInstalledASPDepsDir() string
	GetTarballRepoRootDir() string
	// GetTarballsRepository() types.RepositoryI
}
