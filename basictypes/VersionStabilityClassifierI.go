package basictypes

import (
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers/types"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

type VersionStabilityClassifierI interface {
	// Check(version []int) (tarballstabilityclassification.StabilityClassification, error)
	// IsStable(version []int) (bool, error)
	Check(p types.TarballNameParserI, filename string) (tarballstabilityclassification.StabilityClassification, error)
	IsStable(p types.TarballNameParserI, filename string) (bool, error)
}
