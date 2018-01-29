package versionstabilityclassifiers

import (
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["gcc"] = &ClassifierGCC{}
}

type ClassifierGCC struct {
}

func (self *ClassifierGCC) Check(version []int) (
	tarballstabilityclassification.StabilityClassification,
	error,
) {

	if version[0] < 5 {
		return tarballstabilityclassification.Release, nil
	}

	if version[1] == 0 {
		if version[2] == 0 {
			return tarballstabilityclassification.Alpha, nil
		} else {
			return tarballstabilityclassification.Beta, nil
		}
	}

	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierGCC) IsStable(version []int) (bool, error) {
	cr, err := self.Check(version)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
