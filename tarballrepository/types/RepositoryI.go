package types

import "github.com/AnimusPEXUS/utils/cache01"

type RepositoryI interface {
	// TODO: probably all Get.* Repository struct's function should be displayed
	//			 here

	// GetPackageTarballsPath(name string) string
	GetPackageSRSPath(name string) string
	GetPackageTarballsPath(name string) string

	PerformPackageTarballsUpdate(name string) error
	CreateCacheObjectForPackage(name string) (
		*cache01.CacheDir,
		error,
	)
	PerformDownload(
		package_name string,
		as_filename string,
		uri string,
	) error
	DeleteFile(
		package_name string,
		filename string,
	) error
	ListLocalTarballs(package_name string) ([]string, error)
	ListLocalFiles(package_name string) ([]string, error)
}
