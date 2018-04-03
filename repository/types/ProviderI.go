package types

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

type ProviderI interface {
	ProviderDescription() string

	ArgCount() int

	CanListArg(i int) bool
	ListArg(i int) ([]string, error)

	Tarballs() ([]string, error)
	TarballNames() ([]string, error)

	PerformUpdate() error
}

type NewProviderFunc func(
	repo RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (ProviderI, error)
