package basictypes

type BuildingSiteCtlI interface {
	GetDIR_TARBALL() string
	GetDIR_SOURCE() string
	GetDIR_PATCHES() string
	GetDIR_BUILDING() string
	GetDIR_DESTDIR() string
	GetDIR_BUILD_LOGS() string
	GetDIR_LISTS() string
	GetDIR_TEMP() string

	ReadInfo() (*BuildingSiteInfo, error)

	GetConfiguredHost() (string, error)
	GetConfiguredArch() (string, error)
	GetConfiguredBuild() (string, error)
	GetConfiguredTarget() (string, error)
	GetConfiguredHABT() (string, string, string, string, error)

	ValuesCalculator() ValuesCalculatorI

	PrePackager() PrePackagerI
	Packager() PackagerI
}
