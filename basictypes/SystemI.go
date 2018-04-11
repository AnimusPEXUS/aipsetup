package basictypes

import (
	"github.com/go-ini/ini"
)

type SystemI interface {
	Cfg() *ini.File
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
