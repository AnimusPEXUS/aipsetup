package basictypes

type SystemValuesCalculatorI interface {
	CalculateMultihostDir() string
	CalculateHostDir(host string) string
	CalculateHostMultiarchDir(host string) string
	CalculateHostArchDir(host, hostarch string) string
	CalculateHostCrossbuildersDir(host string) string
	CalculateHostCrossbuilderDir(host, target string) string
}
