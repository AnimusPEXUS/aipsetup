package aipsetup

import (
	"errors"
	"strconv"
	"strings"

	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/textlist"
	"github.com/AnimusPEXUS/utils/version"
)

var StdVersionFilterFunctions textlist.FilterFunctions

func VersionCheck(
	function string,
	parameter string,
	value_to_match string,
) (bool, error) {
	res, err := tarballname.Parse(value_to_match)
	if err != nil {
		return false, err
	}

	// fmt.Println("VersionCheck", function, parameter, value_to_match)

	ret := false

	vtm_i_array, err := res.Version.ArrInt()
	if err != nil {
		return false, err
	}

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
		ret = version.Compare(vtm_i_array, param_i_array) == -1
	case "<=":
		r := version.Compare(vtm_i_array, param_i_array)
		ret = r == -1 || r == 0
	case "==":
		ret = version.Compare(vtm_i_array, param_i_array) == 0
	case ">=":
		r := version.Compare(vtm_i_array, param_i_array)
		ret = r == 0 || r == 1
	case ">":
		ret = version.Compare(vtm_i_array, param_i_array) == 1
	case "!=":
		ret = version.Compare(vtm_i_array, param_i_array) != 0
	}

	return ret, nil
}

func init() {

	StdVersionFilterFunctions = make(textlist.FilterFunctions)

	for k, v := range StdVersionFilterFunctions {
		StdVersionFilterFunctions[k] = v
	}

	StdVersionFilterFunctions["version-<"] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
	) (bool, error) {
		return VersionCheck("<", parameter, value_to_match)
	}
	StdVersionFilterFunctions["version-<="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
	) (bool, error) {
		return VersionCheck("<=", parameter, value_to_match)
	}
	StdVersionFilterFunctions["version-=="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
	) (bool, error) {
		return VersionCheck("==", parameter, value_to_match)
	}
	StdVersionFilterFunctions["version->="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
	) (bool, error) {
		return VersionCheck(">=", parameter, value_to_match)
	}
	StdVersionFilterFunctions["version->"] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
	) (bool, error) {
		return VersionCheck(">", parameter, value_to_match)
	}
	StdVersionFilterFunctions["version-!="] = func(
		parameter string,
		case_sensitive bool,
		value_to_match string,
	) (bool, error) {
		return VersionCheck("!=", parameter, value_to_match)
	}
}
