package distropkginfodb

// WARNING: This file is not generated automatically.
//          Keep it safe when copying files generated with "info-db code"
//          command.

import (
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

func Get(name string) (*basictypes.PackageInfo, error) {
	if t, ok := Index[name]; ok {
		return t, nil
	} else {
		return nil, errors.New("package info not found")
	}
}
