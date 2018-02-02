package versionstabilityclassifiers

import (
	"errors"

	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers/types"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
)

func init() {
	Index["gnome"] = &ClassifierGnome{}
}

type ClassifierGnome struct {
}

// func (self *ClassifierGnome) Check(version []int) (
// 	tarballstabilityclassification.StabilityClassification,
// 	error,
// ) {
// 	if len(version) < 2 {
// 		return tarballstabilityclassification.Development,
// 			errors.New("version numbers array too short")
// 	}
//
// 	if version[1]%2 != 0 {
// 		return tarballstabilityclassification.Development, nil
// 	}
//
// 	return tarballstabilityclassification.Release, nil
// }
//
// func (self *ClassifierGnome) IsStable(version []int) (bool, error) {
// 	cr, err := self.Check(version)
// 	if err != nil {
// 		return false, err
// 	}
// 	return tarballstabilityclassification.IsStable(cr), nil
// }

func (self *ClassifierGnome) Check(p types.TarballNameParserI, filename string) (
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

	if len(version) < 2 {
		return tarballstabilityclassification.Development,
			errors.New("version numbers array too short")
	}

	if version[1]%2 != 0 {
		return tarballstabilityclassification.Development, nil
	}

	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierGnome) IsStable(p types.TarballNameParserI, filename string) (bool, error) {
	cr, err := self.Check(p, filename)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
