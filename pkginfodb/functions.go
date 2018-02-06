package pkginfodb

import (
	"errors"
	"fmt"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/versionfilterfunctions"
	"github.com/AnimusPEXUS/utils/set"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnamefilterfunctions"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/textlist"
)

func Get(name string) (*basictypes.PackageInfo, error) {
	if t, ok := Index[name]; ok {
		return t, nil
	} else {
		return nil, errors.New("package info not found")
	}
}

func DetermineTarballsBuildInfo(filename string) (
	map[string]*basictypes.PackageInfo,
	error,
) {

	ret := make(map[string]*basictypes.PackageInfo)

	filename_s_base := path.Base(filename)
	filename_s_base_list := []string{filename_s_base}

	//parsers_map := make(map[string]types.TarballNameParserI)

searching:
	for key, value := range Index {

		name_parser_name := value.TarballFileNameParser

		parser, err := tarballnameparsers.Get(name_parser_name)
		if err != nil {
			return nil, err
		}

		parse_result, err := parser.Parse(filename_s_base)
		if err != nil {
			fmt.Sprintln(
				"can't parse %s with %s parser",
				filename_s_base,
				name_parser_name,
			)
			continue searching
		}

		if parse_result.Name == value.TarballName {

			fres, err := ApplyInfoFilter(value, filename_s_base_list)
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

func ApplyInfoFilter(
	info *basictypes.PackageInfo,
	data []string,
) ([]string, error) {

	filter_functions := make(textlist.FilterFunctions)

	for _, v1 := range []textlist.FilterFunctions{
		versionfilterfunctions.VersionFilterFunctions,
		tarballnamefilterfunctions.TarballNameFilterFunctions,
	} {
		for k, v := range v1 {
			filter_functions[k] = v
		}
	}

	additional_data := make(map[string]interface{})
	additional_data["pkg_info"] = info

	ret, err := textlist.FilterList(
		data,
		textlist.ParseFilterTextLinesMust(info.TarballFilters),
		filter_functions,
		additional_data,
	)
	if err != nil {
		return []string{}, err
	}

	return ret, nil
}

func ListPackagesByGroups(groups []string) ([]string, error) {
	ret := make([]string, 0)

	found_groups := set.NewSetString()

	for k, v := range Index {
		found := false
	loop2:
		for _, v2 := range v.Groups {
			for _, v3 := range groups {
				if v2 == v3 {
					found = true
					found_groups.Add(v3)
					break loop2
				}
			}
		}
		if found {
			found2 := false
			for _, v2 := range ret {
				if v2 == k {
					found2 = true
					break
				}
			}
			if !found2 {
				ret = append(ret, k)
			}
		}
	}

	for _, i := range groups {
		found := false
		for _, j := range found_groups.ListStrings() {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			return []string{}, errors.New("some of named groups are not found")
		}
	}

	return ret, nil
}

func ListPackagesByCategories(categories []string) ([]string, error) {
	ret := make([]string, 0)

	found_categories := set.NewSetString()

	for k, v := range Index {
		found := false
		for _, v3 := range categories {
			if v.Category == v3 {
				found = true
				found_categories.Add(v3)
				break
			}
		}
		if found {
			found2 := false
			for _, v2 := range ret {
				if v2 == k {
					found2 = true
					break
				}
			}
			if !found2 {
				ret = append(ret, k)
			}
		}
	}

	for _, i := range categories {
		found := false
		for _, j := range found_categories.ListStrings() {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			return []string{}, errors.New("some of categories groups are not found")
		}
	}

	return ret, nil
}
