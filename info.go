package aipsetup

type CompletePackageInfo struct {
	OveralPackageInfo
	TarballServerTraversePackageInfo
	BuildingPackageInfo
	CategorizationPackageInfo
}

type OveralPackageInfo struct {
	Description string `json:"description" yaml:"description"`
	HomePage    string `json:"home_page" yaml:"home_page"`
	VersionTool string `json:"version_tool" yaml:"version_tool"`

	Removable          bool `json:"removable" yaml:"removable"`
	Reducible          bool `json:"reducible" yaml:"reducible"`
	NonInstallable     bool `json:"non_installable" yaml:"non_installable"`
	Deprecated         bool `json:"deprecated" yaml:"deprecated"`
	PrimaryInstallOnly bool `json:"only_primary_install" yaml:"only_primary_install"`

	BuildDeps   []string `json:"build_deps" yaml:"build_deps"`
	SODeps      []string `json:"so_deps" yaml:"so_deps"`
	RunTimeDeps []string `json:"runtime_deps" yaml:"runtime_deps"`
}

type TarballServerTraversePackageInfo struct {
	BaseName           string   `json:"basename" yaml:"basename"`
	Filters            string   `json:"filters" yaml:"filters"`
	SourcePathPrefixes []string `json:"source_path_prefixes" yaml:"source_path_prefixes"`
}

type BuildingPackageInfo struct {
	BuilderName string `json:"buildscript" yaml:"buildscript"`
}

type CategorizationPackageInfo struct {
	Tags []string `json:"tags" yaml:"tags"`
}

/*
SAMPLE_PACKAGE_INFO_STRUCTURE = collections.OrderedDict([
    # description
    ('description', ""),

    # not required, but can be useful
    ('home_page', ""),

    # string
    ('buildscript', ''),

    # string
    ('version_tool', ''),

    # file name base
    ('basename', ''),

    # filters. various filters to provide correct list of acceptable tarballs
    # by they filenames
    ('filters', ''),

    # can package be deleted without hazard to aipsetup functionality
    # (including system stability)?
    ('removable', True),

    # can package be updated without hazard to aipsetup functionality
    # (including system stability)?
    ('reducible', True),

    # package can not be installed
    ('non_installable', False),

    # package outdated and need to be removed
    ('deprecated', False),

    # some shitty packages
    # (like python2 and python3: see https://bugs.python.org/issue1294959)
    # can't be forced to be installed in lib64
    ('only_primary_install', False),

    # list of str
    ('tags', []),

    # to make search faster and exclude not related sources
    ('source_path_prefixes', []),

    # following packages required to build this package
    ('build_deps', []),

    # depends on .so files in following packages
    ('so_deps', []),

    # run time dependenties. (man pages reader requiers 'less' command i.e.)
    ('runtime_deps', [])

    ])
*/
