package basictypes

type PackageInfo struct {
	// Name string `json:"name"`

	Description string //`json:"description"`
	HomePage    string //`json:"homepage"`

	BuilderName string //`json:"builder_name"`

	Removable          bool //`json:"removable"`
	Reducible          bool //`json:"reducible"`
	NonInstallable     bool //`json:"non_installable"`
	Deprecated         bool //`json:"deprecated"`
	PrimaryInstallOnly bool //`json:"primary_install_only"`

	BuildDeps   []string //`json:"build_deps"`
	SODeps      []string //`json:"so_deps"`
	RunTimeDeps []string //`json:"runtime_deps"`

	Tags     []string //`json:"tags"`
	Category string
	Groups   []string

	TarballFileNameParser string //`json:"tarball_filename_parser"`

	TarballName string //`json:"tarball_name"`

	TarballFilters []string

	TarballProvider                 string   //`json:"tarball_provider"`
	TarballProviderArguments        []string // `json:"tarball_provider_arguments"`
	TarballProviderVersionSyncDepth int

	TarballVersionTool         string
	TarballStabilityClassifier string
	TarballVersionComparator   string
}
