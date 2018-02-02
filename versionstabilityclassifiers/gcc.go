package versionstabilityclassifiers

import (
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers/types"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["gcc"] = &ClassifierGCC{}
}

type ClassifierGCC struct {
}

func (self *ClassifierGCC) Check(p types.TarballNameParserI, filename string) (
	tarballstabilityclassification.StabilityClassification,
	error,
) {

	parsed, err := p.Parse(filename)
	if err != nil {
		return tarballstabilityclassification.Development, err
	}

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

func (self *ClassifierGCC) IsStable(p types.TarballNameParserI, filename string) (bool, error) {
	cr, err := self.Check(p, filename)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
