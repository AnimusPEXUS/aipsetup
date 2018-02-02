package basictypes

import (
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

type VersionStabilityClassifierI interface {
	Check(parsed *tarballname.ParsedTarballName) (tarballstabilityclassification.StabilityClassification, error)
	IsStable(parsed *tarballname.ParsedTarballName) (bool, error)
}
