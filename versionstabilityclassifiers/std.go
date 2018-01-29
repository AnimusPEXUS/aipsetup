package versionstabilityclassifiers

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
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

func (self *ClassifierStd) Check(version []int) (
	tarballstabilityclassification.StabilityClassification,
	error,
) {
	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierStd) IsStable(version []int) (bool, error) {
	cr, err := self.Check(version)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
