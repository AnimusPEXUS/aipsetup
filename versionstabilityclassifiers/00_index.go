package versionstabilityclassifiers

import (
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func Get(name string) (basictypes.VersionStabilityClassifierI, error) {
	if t, ok := Index[name]; ok {
		return t, nil
	} else {
		return nil, errors.New("classifier not found")
	}
}

var Index = map[string](basictypes.VersionStabilityClassifierI){}
