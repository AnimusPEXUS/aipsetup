package versionstabilityclassifiers

import (
	"errors"

	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["gnome"] = &ClassifierGnome{}
}

type ClassifierGnome struct {
}

func (self *ClassifierGnome) Check(parsed *tarballname.ParsedTarballName) (
	tarballstabilityclassification.StabilityClassification,
	error,
) {

	if parsed.Status.Str != "" {
		return tarballstabilityclassification.Development, nil
	}

	version, err := parsed.Version.ArrInt()
	if err != nil {
		return tarballstabilityclassification.Development, err
	}

	if len(version) < 2 {
		return tarballstabilityclassification.Development,
			errors.New("version numbers array too short")
	}

	if version[1]%2 != 0 {
		return tarballstabilityclassification.Development, nil
	}

	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierGnome) IsStable(parsed *tarballname.ParsedTarballName) (bool, error) {
	cr, err := self.Check(parsed)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
