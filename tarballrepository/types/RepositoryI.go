package types

import "github.com/AnimusPEXUS/utils/cache01"

type RepositoryI interface {
	PerformPackageTarballsUpdate(name string) error
	CreateCacheObjectForPackage(name string) (
		*cache01.CacheDir,
		error,
	)
}
