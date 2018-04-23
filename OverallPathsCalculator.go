package aipsetup

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

type OverallPathsCalculator struct {
}

// func NewOverallPathsCalculator() OverallPathsCalculator {
// 	return OverallPathsCalculator{}
// }

func (self OverallPathsCalculator) CalculateMultihostDir(root string) string {
	return path.Join(root, basictypes.LILITH_ROOT_MULTIHOST_DIRNAME)
}

func (self OverallPathsCalculator) CalculateHostDir(root, host string) string {
	return path.Join(self.CalculateMultihostDir(root), host)
}

func (self OverallPathsCalculator) CalculateHostMultiarchDir(
	root, host string,
) string {
	return path.Join(
		self.CalculateHostDir(root, host),
		basictypes.LILITH_MULTIHOST_MULTIARCH_DIRNAME,
	)
}

func (self OverallPathsCalculator) CalculateHostArchDir(
	root, host, hostarch string,
) string {
	return path.Join(self.CalculateHostMultiarchDir(root, host), hostarch)
}

// /{hostpath}/corssbuilders
func (self OverallPathsCalculator) CalculateHostCrossbuildersDir(
	root, host string,
) string {
	return path.Join(
		self.CalculateHostDir(root, host),
		basictypes.LILITH_MULTIHOST_CROSSBULDERS_DIRNAME,
	)
}

// /{hostpath}/corssbuilders/{target}
func (self OverallPathsCalculator) CalculateHostCrossbuilderDir(
	root, host, target string,
) string {
	return path.Join(self.CalculateHostCrossbuildersDir(root, host), target)
}
