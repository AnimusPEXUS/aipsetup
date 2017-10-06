package aipsetup

type CompletePackageInfo struct {
	OveralPackageInfo
	TarballPackageInfo
	BuildingPackageInfo
	CategorizationPackageInfo
	DependenciesPackageInfo
}

type OveralPackageInfo struct {
	Description string
	HomePage    string

	Removable          bool
	Reducible          bool
	NonInstallable     bool
	Deprecated         bool
	PrimaryInstallOnly bool

	BuildDeps   []string
	SODeps      []string
	RunTimeDeps []string
}

type TarballPackageInfo struct {
	Name        string
	VersionTool string
	Filters     string
}

type BuildingPackageInfo struct {
	BuilderName string
}

type CategorizationPackageInfo struct {
	Tags []string
}

type DependenciesPackageInfo struct {
	BuildDeps   []string
	SODeps      []string
	RunTimeDeps []string
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
