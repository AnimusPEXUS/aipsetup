package basictypes

type PackageInfo struct {
	Description string
	HomePage    string

	BuilderName string

	Removable          bool
	Reducible          bool
	NonInstallable     bool
	Deprecated         bool
	PrimaryInstallOnly bool

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
}
