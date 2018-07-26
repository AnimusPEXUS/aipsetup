package pkginfodb

import (
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"

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

func DetermineTarballPackageInfoSingle(filename string) (
	string,
	*basictypes.PackageInfo,
	error,
) {
	res, err := DetermineTarballPackageInfo(filename)
	if err != nil {
		return "", nil, err
	}
	if len(res) != 1 {

		keys := make([]string, 0)

		for k, _ := range res {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		return "", nil,
			errors.New(
				fmt.Sprintf(
					"couldn't determine single package info by tarball name. matches count %d: %v",
					len(keys),
					keys,
				),
			)

	}
	var name string
	var info *basictypes.PackageInfo
	for name, info = range res {
	}
	return name, info, nil
}

func DetermineTarballPackageInfo(filename string) (
	map[string]*basictypes.PackageInfo,
	error,
) {

	ret := make(map[string]*basictypes.PackageInfo)

	filename_s_base := path.Base(filename)

	keys := IndexKeysSorted()

searching:
	for _, key := range keys {

		value := Index[key]

		match, err := CheckTarballMatchesInfo(filename_s_base, key, value)
		if err != nil {
			return nil, err
		}

		if !match {
			continue searching
		}

		ret[key] = value

	}

	return ret, nil
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
			return nil, errors.New("some of named groups are not found")
		}
	}

	return ret, nil
}

func ListPackagesByCategories(categories []string, is_prefixes bool) ([]string, error) {

	ret := set.NewSetString()

	some_not_found := false

	for k, v := range Index {
		found := false
		for _, v3 := range categories {
			if (!is_prefixes && (v.Category == v3)) ||
				(is_prefixes && strings.HasPrefix(v.Category, v3)) {
				found = true
				ret.Add(k)
				break
			}
		}
		if !found {
			some_not_found = true
		}
	}

	if !some_not_found {
		return nil, errors.New("some of categories are not found")
	}

	return ret.ListStrings(), nil
}

func CheckTarballMatchesInfoByName(
	tarballfilename string,
	infoname string,
) (bool, error) {
	info, err := Get(infoname)
	if err != nil {
		return false, err
	}

	return CheckTarballMatchesInfo(tarballfilename, infoname, info)
}

func CheckTarballMatchesInfo(
	tarballfilename string,
	infoname string,
	info *basictypes.PackageInfo,
) (bool, error) {

	// TODO: looks like it would be good, if parse_result would also be returned

	if strings.Trim(info.TarballName, " \n") == "" {

		return false,
			errors.New(
				fmt.Sprintf("package %s have invalid tarball base name", infoname),
			)
	}

	tarballfilename = path.Base(tarballfilename)

	parser, err := tarballnameparsers.Get(info.TarballFileNameParser)
	if err != nil {
		return false, err
	}

	parse_result, err := parser.Parse(tarballfilename)
	if err != nil {
		return false, err
	}

	if parse_result.Name != info.TarballName {
		return false, nil // not an error
	}

	fres, err := ApplyInfoFilter(info, []string{tarballfilename})
	if err != nil {
		return false, err
	}

	if len(fres) != 1 {
		return false, nil // not an error
	}

	return true, nil
}
