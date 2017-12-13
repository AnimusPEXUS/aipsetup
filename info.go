package aipsetup

import (
	"errors"
	"fmt"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/textlist"
	"github.com/AnimusPEXUS/utils/version/versionfilterfunctions"
)

func DetermineTarballsBuildinfo(filename string) (
	map[string]*basictypes.PackageInfo,
	error,
) {

	ret := make(map[string]*basictypes.PackageInfo)

	filename_s_base := path.Base(filename)
	filename_s_base_list := []string{filename_s_base}

searching:
	for key, value := range distropkginfodb.Index {

		name_parser_name := value.TarballFileNameParser

		parser_const, ok := tarballnameparsers.Index[name_parser_name]
		if !ok {
			return nil, errors.New("parser " + name_parser_name + "not found")
		}

		parser := parser_const()

		parse_result, err := parser.ParseName(filename_s_base)
		if err != nil {
			fmt.Sprintln(
				"can't parse %s with %s parser",
				filename_s_base,
				name_parser_name,
			)
			continue searching
		}

		if parse_result.Name == value.TarballName {

			fres, err := textlist.FilterList(
				filename_s_base_list,
				value.Filters,
				versionfilterfunctions.StdVersionFilterFunctions,
			)
			if err != nil {
				return nil, err
			}

			switch len(fres) {
			case 0:
			case 1:
				ret[key] = value
			default:
				panic("this shoud be unreachable")
			}

		}

	}

	switch len(ret) {
	case 0:
		return ret, errors.New("not found")
	case 1:
		return ret, nil
	default:
		return ret, errors.New("too many matches")
	}
}
