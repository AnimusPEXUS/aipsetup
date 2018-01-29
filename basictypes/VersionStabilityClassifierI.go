package basictypes

import "github.com/AnimusPEXUS/utils/tarballstabilityclassification"

type VersionStabilityClassifierI interface {
	Check(version []int) (tarballstabilityclassification.StabilityClassification, error)
	IsStable(version []int) (bool, error)
}
