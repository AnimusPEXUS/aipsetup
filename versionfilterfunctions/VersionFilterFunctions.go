package versionfilterfunctions

import (
	"errors"
	"strconv"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/textlist"
	"github.com/AnimusPEXUS/utils/version/versioncomparators/types"
)

var VersionFilterFunctions textlist.FilterFunctions

func VersionCheck(
	function string,
	parameter string,
	value_to_match string,
	data map[string]interface{},
) (bool, error) {

	info, ok := data["pkg_info"].(*basictypes.PackageInfo)
	if !ok {
		panic("VersionCheck requires data[\"pkg_info\"]")
	}

	ver_comparator, ok := data["versioncomparator"].(types.VersionComparatorI)
	if !ok {
		panic("VersionCheck requires data[\"versioncomparator\"]")
	}

	p, err := tarballnameparsers.Get(info.TarballFileNameParser)
	if err != nil {
		return false, err
	}

	res, err := p.Parse(value_to_match)
	if err != nil {
		return false, err
	}

	// conterpart building means (and requires) value_to_match to have clear
	// version text (in form of \d+(\.\d+)*), so, for instance, infozip zip
	// version 300 should be represented as "3.00" in filter text
	res_conterpart := new(tarballname.ParsedTarballName)
	{
		std_parser := tarballnameparsers.Index["std"]
		p, err := std_parser.Parse("aaa-" + value_to_match + ".tar.xz")
		if err != nil {
			return false,
				errors.New("fileter text value for VersionCheck given is invalid")
		}

		res_conterpart.Name = res.Name
		res_conterpart.Version = p.Version
	}

	ret := false

	param_i_array := make([]int, 0)
	for _, i := range strings.Split(parameter, ".") {
		i_i, err := strconv.Atoi(i)
		if err != nil {
			return false, err
		}
		param_i_array = append(param_i_array, i_i)
	}

	switch function {
	default:
		return false, errors.New("invalid version comparison function")
	case "<":
		cres, err := ver_comparator.Compare(res, res_conterpart)
		if err != nil {
			return false, err
		}
		ret = cres == -1
	case "<=":
		cres, err := ver_comparator.Compare(res, res_conterpart)
		if err != nil {
			return false, err
		}
		ret = cres == -1 || cres == 0
	case "==":
		cres, err := ver_comparator.Compare(res, res_conterpart)
		if err != nil {
			return false, err
		}
		ret = cres == 0
	case ">=":
		cres, err := ver_comparator.Compare(res, res_conterpart)
		if err != nil {
			return false, err
		}
		ret = cres == 0 || cres == 1
	case ">":
		cres, err := ver_comparator.Compare(res, res_conterpart)
		if err != nil {
			return false, err
		}
		ret = cres == 1
	case "!=":
		cres, err := ver_comparator.Compare(res, res_conterpart)
		if err != nil {
			return false, err
		}
		ret = cres != 0
	}

	return ret, nil
}

func init() {
	VersionFilterFunctions = make(textlist.FilterFunctions)

	// for k, v := range VersionFilterFunctions {
	// 	VersionFilterFunctions[k] = v
	// }

	VersionFilterFunctions["version-<"] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
		data map[string]interface{},
	) (bool, error) {
		return VersionCheck("<", parameter, value_to_match, data)
	}
	VersionFilterFunctions["version-<="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
		data map[string]interface{},
	) (bool, error) {
		return VersionCheck("<=", parameter, value_to_match, data)
	}
	VersionFilterFunctions["version-=="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
		data map[string]interface{},
	) (bool, error) {
		return VersionCheck("==", parameter, value_to_match, data)
	}
	VersionFilterFunctions["version->="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
		data map[string]interface{},
	) (bool, error) {
		return VersionCheck(">=", parameter, value_to_match, data)
	}
	VersionFilterFunctions["version->"] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
		data map[string]interface{},
	) (bool, error) {
		return VersionCheck(">", parameter, value_to_match, data)
	}
	VersionFilterFunctions["version-!="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
		data map[string]interface{},
	) (bool, error) {
		return VersionCheck("!=", parameter, value_to_match, data)
	}
}
