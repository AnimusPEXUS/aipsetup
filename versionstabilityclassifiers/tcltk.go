package versionstabilityclassifiers

import (
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["tcltk"] = &ClassifierTclTk{}
}

type ClassifierTclTk struct {
}

func (self *ClassifierTclTk) Check(parsed *tarballname.ParsedTarballName) (
	tarballstabilityclassification.StabilityClassification,
	error,
) {

	if parsed.Status.Str != "src" {
		return tarballstabilityclassification.Development, nil
	}

	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierTclTk) IsStable(parsed *tarballname.ParsedTarballName) (bool, error) {
	cr, err := self.Check(parsed)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
