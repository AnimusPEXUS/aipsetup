package basictypes

type PackageInfo struct {
	Description string
	HomePage    string

	BuilderName string

	Removable          bool
	Reducible          bool
	AutoReduce         bool
	NonBuildable       bool
	NonInstallable     bool
	Deprecated         bool
	PrimaryInstallOnly bool

	BuildPkgDeps []string

	BuildDeps   []string
	SODeps      []string
	RunTimeDeps []string

	Tags     []string
	Category string
	Groups   []string

	TarballFileNameParser string

	TarballName string

	TarballFilters []string

	TarballProvider                 string
	TarballProviderArguments        []string
	TarballProviderVersionSyncDepth int

	TarballStabilityClassifier string
	TarballVersionComparator   string

	DownloadPatches              bool
	PatchesDownloadingScriptText string
}
