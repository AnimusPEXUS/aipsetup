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

	PerformPackageSourcesUpdate(name string) error
	PerformPackageTarballsUpdate(name string) error
	PerformPackagePatchesUpdate(name string) error
	GetPackageASPsPath(name string) string

	PerformDownload(package_name string, as_filename string, uri string) error
	PrepareTarballCleanupListing(package_name string, files_to_keep []string) ([]string, error)

	DeleteTarballFile(package_name string, filename string) error
	DeleteTarballFiles(package_name string, filename []string) error

	ListLocalTarballFiles(package_name string) ([]string, error)
	ListLocalTarballs(package_name string, done_only bool) ([]string, error)

	DeleteASPFile(package_name string, filename string) error
	DeleteASPFiles(package_name string, filename []string) error

	ListLocalASPFiles(package_name string) ([]string, error)
	ListLocalASPs(package_name string) ([]string, error)

	CopyTarballToDir(package_name string, tarball string, outdir string) error
	CopyPatchesToDir(package_name string, outdir string) error
	CopyASPToDir(package_name string, asp string, outdir string) error
}
