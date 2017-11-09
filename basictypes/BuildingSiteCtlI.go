package basictypes

type BuildingSiteCtlI interface {
	ReadInfo() (*BuildingSiteInfo, error)
	GetDIR_TARBALL() string
	GetDIR_SOURCE() string
	GetDIR_PATCHES() string
	GetDIR_BUILDING() string
	GetDIR_DESTDIR() string
	GetDIR_BUILD_LOGS() string
	GetDIR_LISTS() string
	GetDIR_TEMP() string

	GetConfiguredHost() string
	GetConfiguredArch() string
	GetConfiguredBuild() string
	GetConfiguredTarget() string
}

type BuildingProcessValuesCalculator interface {
	CalculateIsCrossbuild() bool
	CalculateIsCrossbuilder() bool
	CalculateIsOnlyArchIsDifferent() bool

	CalculateMultihostDir() string
	CalculateDstMultihostDir() string

	CalculateHostMultiarchDir() string
	CalculateDstHostMultiarchDir() string

	CalculateHostArchDir() string
	CalculateDstHostArchDir() string

	// /{hostpath}/corssbuilders
	CalculateHostCrossbuildersDir() string
	CalculateDstHostCrossbuildersDir() string

	// /{hostpath}/corssbuilders/{target}
	CalculateHostCrossbuilderDir() string
	CalculateDstHostCrossbuilderDir() string
}
