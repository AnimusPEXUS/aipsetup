package versionstabilityclassifiers

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers/types"
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

func (self *ClassifierStd) Check(p types.TarballNameParserI, filename string) (
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

	return tarballstabilityclassification.Release, nil
}

func (self *ClassifierStd) IsStable(
	p types.TarballNameParserI,
	filename string,
) (bool, error) {
	cr, err := self.Check(p, filename)
	if err != nil {
		return false, err
	}
	return tarballstabilityclassification.IsStable(cr), nil
}
