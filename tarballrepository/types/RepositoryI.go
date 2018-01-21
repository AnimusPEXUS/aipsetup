package types

type RepositoryI interface {
	GetRepositoryPath() string
	GetCachesDir() string
	GetPackagePath(name string) string
	GetPackageSRSPath(name string) string
	GetPackageTarballsPath(name string) string
	GetPackageCachePath(name string) string
	GetDedicatedCachePath(name string) string
	GetTarballDoneFilePath(package_name string, as_filename string) string
	GetTarballFilePath(package_name, as_filename string) string

	PerformPackageTarballsUpdate(name string) error
	// CreateCacheObjectForPackage(name string) (*cache01.CacheDir, error)
	PerformDownload(package_name string, as_filename string, uri string) error
	PrepareTarballCleanupListing(package_name string, files_to_keep []string) ([]string, error)
	DeleteFile(package_name string, filename string) error
	DeleteFiles(package_name string, filename []string) error
	ListLocalTarballs(package_name string, done_only bool) ([]string, error)
	ListLocalFiles(package_name string) ([]string, error)
	CopyTarballToDir(package_name string, tarball string, outdir string) error
}
