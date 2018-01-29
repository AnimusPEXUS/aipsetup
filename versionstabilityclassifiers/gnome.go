package versionstabilityclassifiers

import (
	"errors"

	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["gnome"] = &ClassifierGnome{}
}

type ClassifierGnome struct {
}

func (self *ClassifierGnome) Check(version []int) (
	tarballstabilityclassification.StabilityClassification,
	error,
) {
	if len(version) < 2 {
		return tarballstabilityclassification.Development,
			errors.New("version array to short")
	}

	if version[1]%2 != 0 {
		return tarballstabilityclassification.Development, nil
	}

	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierGnome) IsStable(version []int) (bool, error) {
	cr, err := self.Check(version)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
