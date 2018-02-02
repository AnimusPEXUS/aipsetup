package versionstabilityclassifiers

import (
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["gcc"] = &ClassifierGCC{}
}

type ClassifierGCC struct {
}

func (self *ClassifierGCC) Check(parsed *tarballname.ParsedTarballName) (
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

func (self *ClassifierGCC) IsStable(parsed *tarballname.ParsedTarballName) (bool, error) {
	cr, err := self.Check(parsed)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
