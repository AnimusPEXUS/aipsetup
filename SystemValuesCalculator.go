package aipsetup

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

var _ basictypes.SystemValuesCalculatorI = &SystemValuesCalculator{}

type SystemValuesCalculator struct {
	sys *System
}

func NewSystemValuesCalculator(sys *System) *SystemValuesCalculator {
	self := new(SystemValuesCalculator)
	self.sys = sys
	return self
}

func (self *SystemValuesCalculator) CalculateMultihostDir() string {
	return path.Join(string(os.PathSeparator), LAILALO_ROOT_MULTIHOST_DIRNAME)
}

func (self *SystemValuesCalculator) CalculateHostDir(host string) (string, error) {
	return path.Join(self.CalculateMultihostDir(), host), nil
}

func (self *SystemValuesCalculator) CalculateHostMultiarchDir(host string) (string, error) {
	d, err := self.CalculateHostDir(host)
	if err != nil {
		return "", err
	}
	return path.Join(d, LAILALO_MULTIHOST_MULTIARCH_DIRNAME), nil
}

func (self *SystemValuesCalculator) CalculateHostArchDir(host, hostarch string) (string, error) {
	d, err := self.CalculateHostMultiarchDir(host)
	if err != nil {
		return "", err
	}

	return path.Join(d, hostarch), nil
}

// /{hostpath}/corssbuilders
func (self *SystemValuesCalculator) CalculateHostCrossbuildersDir(host string) (string, error) {
	hostdir, err := self.CalculateHostDir(host)
	if err != nil {
		return "", err
	}
	return path.Join(hostdir, LAILALO_MULTIHOST_CROSSBULDERS_DIRNAME), nil
}

// /{hostpath}/corssbuilders/{target}
func (self *SystemValuesCalculator) CalculateHostCrossbuilderDir(host, target string) (string, error) {
	hostcrossbuildersdir, err := self.CalculateHostCrossbuildersDir(host)
	if err != nil {
		return "", err
	}
	return path.Join(hostcrossbuildersdir, target), nil
}
