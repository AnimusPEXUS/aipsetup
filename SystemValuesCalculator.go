package aipsetup

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var _ basictypes.SystemValuesCalculatorI = &SystemValuesCalculator{}

type SystemValuesCalculator struct {
	sys *System
	opc OverallPathsCalculator
}

func NewSystemValuesCalculator(sys *System) *SystemValuesCalculator {
	self := new(SystemValuesCalculator)
	self.sys = sys
	return self
}

func (self *SystemValuesCalculator) CalculateMultihostDir() string {
	return self.opc.CalculateMultihostDir(self.sys.root)
}

func (self *SystemValuesCalculator) CalculateHostDir(host string) string {
	return self.opc.CalculateHostDir(self.sys.root, host)
}

func (self *SystemValuesCalculator) CalculateHostMultiarchDir(
	host string,
) string {
	return self.opc.CalculateHostMultiarchDir(self.sys.root, host)
}

func (self *SystemValuesCalculator) CalculateHostArchDir(
	host, hostarch string,
) string {
	return self.opc.CalculateHostArchDir(self.sys.root, host, hostarch)
}

// /{hostpath}/corssbuilders
func (self *SystemValuesCalculator) CalculateHostCrossbuildersDir(
	host string,
) string {
	return self.opc.CalculateHostCrossbuildersDir(self.sys.root, host)
}

// /{hostpath}/corssbuilders/{target}
func (self *SystemValuesCalculator) CalculateHostCrossbuilderDir(
	host, target string,
) string {
	return self.opc.CalculateHostCrossbuilderDir(self.sys.root, host, target)
}
