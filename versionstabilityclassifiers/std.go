package versionstabilityclassifiers

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["std"] = &ClassifierStd{}
}

type ClassifierStd struct {
}

func NewClassifierStd() basictypes.VersionStabilityClassifierI {
	ret := new(ClassifierStd)
	return ret
}

func (self *ClassifierStd) Check(parsed *tarballname.ParsedTarballName) (
	tarballstabilityclassification.StabilityClassification,
	error,
) {

	if parsed.Status.Str != "" {
		return tarballstabilityclassification.Development, nil
	}

	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierStd) IsStable(parsed *tarballname.ParsedTarballName) (bool, error) {
	cr, err := self.Check(parsed)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
