package main

import (
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
)

var (
	STD_ROOT_OPTION = &cliapp.GetOptCheckListItem{
		Name:          "--root",
		Description:   "System root to work with",
		HaveDefault:   true,
		Default:       "/",
		IsRequired:    false,
		MustHaveValue: true,
	}

	// TODO: hot sure about this option
	STD_OPTION_BUILD_USING = &cliapp.GetOptCheckListItem{
		Name: "--build-using",
		Description: "Select system which' builder should be used as builder" +
			"Default is hosting system, but this may change " +
			"in future releases of aipsetup. Value passed here should be one " +
			"or two system tripletts separated with comma.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_FOR_HOST = &cliapp.GetOptCheckListItem{
		Name:          "--build-for-host",
		Description:   "Select which system will run. Default is current host",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_FOR_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--build-for-hostarch",
		Description: "Select which system arch will run. " +
			"Default is equal to value calculated (or given to) with " +
			"--build-for-host",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_TO_TARGET = &cliapp.GetOptCheckListItem{
		Name: "--build-to-target",
		Description: "This option is for building crosscompiler, " +
			"so crosscompiler will be built to compile for named system. " +
			"Default value for this option is empty value, what disables this option",
		HaveDefault:   false,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_NAMED_INSTALLATION_FOR_HOST = &cliapp.GetOptCheckListItem{
		Name: "--install-for-host",
		Description: "Select hosting system. " +
			"Default is to get value from aipsetup5.system.ini",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_NAMED_INSTALLATION_FOR_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--install-for-hostarch",
		Description: "Select hosting subarch system. " +
			"Default value is 'config', which means to read aipsetup5.system.ini " +
			"and get subarch values for hosting system selected with " +
			"--install-for-host",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_NAMED_INSTALLATION_TO_TARGET = &cliapp.GetOptCheckListItem{
		Name: "--install-to-target",
		Description: "This option is for selecting crosscompiler, " +
			"which can build for named target. No default or automatic value - " +
			"select one",
		HaveDefault:   false,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_HOST = &cliapp.GetOptCheckListItem{
		Name: "--show-only-host",
		Description: "Value is regexp. Show only selected by regexp. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--show-only-hostarchs",
		Description: "Value is regexp. Show only selected by regexp. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_TARGET = &cliapp.GetOptCheckListItem{
		Name: "--show-only-target",
		Description: "Value is regexp. Show only selected by regexp. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}
)

func StdRoutineGetRootOption(getopt_result *cliapp.GetOptResult) (
	string,
	bool,
) {

	res := getopt_result.GetLastNamedRetOptItem("--root")
	if res != nil {
		return res.Value, true
	}

	return "", false
}

// func StdRoutineGetHostArchOptionsDeprecated(getopt_result *cliapp.GetOptResult) (
// 	string, bool,
// 	string, bool,
// ) {
//
// 	host := ""
// 	host_ok := false
// 	arch := ""
// 	arch_ok := false
//
// 	{
// 		res := getopt_result.GetLastNamedRetOptItem("--host")
// 		if res != nil {
// 			host = res.Value
// 			host_ok = true
// 		}
// 	}
//
// 	{
// 		res := getopt_result.GetLastNamedRetOptItem("--arch")
// 		if res != nil {
// 			arch = res.Value
// 			arch_ok = true
// 		}
// 	}
//
// 	return host, host_ok, arch, arch_ok
// }

// func StdRoutineRootHostArchSysDeprecated(getopt_result *cliapp.GetOptResult) (
// 	root string,
// 	host string,
// 	arch string,
// 	sys *aipsetup.System,
// 	ret *cliapp.AppResult,
// ) {
//
// 	root = "/"
// 	host = ""
// 	arch = ""
// 	ret = &cliapp.AppResult{Code: 0}
//
// 	root_opt, ok := StdRoutineGetRootOption(getopt_result)
// 	if ok {
// 		root = root_opt
// 	}
//
// 	sys = aipsetup.NewSystem(root)
//
// 	{
// 		host_opt, host_ok, arch_opt, arch_ok :=
// 			StdRoutineGetHostArchOptions(getopt_result)
//
// 		if host_ok {
// 			host = host_opt
// 		}
//
// 		if arch_ok {
// 			arch = arch_opt
// 		}
// 	}
//
// 	if arch != "" && host == "" {
// 		ret = &cliapp.AppResult{
// 			Code:    10,
// 			Message: "if host is empty, arch must be empty too",
// 		}
// 	}
// 	return
// }
//
func StdRoutineGetRootOptionAndSystemObject(getopt_result *cliapp.GetOptResult) (
	root string,
	sys *aipsetup.System,
	ret *cliapp.AppResult,
) {

	// TODO: cleanup required
	// TODO: move this function closer to StdRoutineGetRootOption()
	// TODO: or, maybe, even, StdRoutineGetRootOption() may be removed, if it's
	//       newer used alone.

	root = "/"
	ret = &cliapp.AppResult{Code: 0}

	if root_opt, ok := StdRoutineGetRootOption(getopt_result); ok {
		root = root_opt
	}

	sys = aipsetup.NewSystem(root)

	return
}

// func StdRoutineMustGetOneArgDeprecated(getopt_result *cliapp.GetOptResult) (
// 	res string,
// 	err *cliapp.AppResult,
// ) {
//
// 	if len(getopt_result.Args) != 1 {
// 		res = ""
// 		err = &cliapp.AppResult{
// 			Code:    1,
// 			Message: "exactly one argument required",
// 		}
// 	} else {
// 		res = getopt_result.Args[0]
// 		err = &cliapp.AppResult{Code: 0}
// 	}
//
// 	return
// }

// func StdRoutineMustGetASPNameDeprecated(getopt_result *cliapp.GetOptResult) (
// 	*basictypes.ASPName,
// 	*cliapp.AppResult,
// ) {
//
// 	res, res_err := StdRoutineMustGetOneArg(getopt_result)
// 	if res_err.Code != 0 {
// 		return nil, res_err
// 	}
//
// 	name, err := basictypes.NewASPNameFromString(res)
// 	if err != nil {
// 		return nil, &cliapp.AppResult{
// 			Code:    11,
// 			Message: "Can't parse given string as ASP name",
// 		}
// 	}
//
// 	return name, &cliapp.AppResult{Code: 0}
// }

// ----v-------v-------v---- rework 20 march 2018

func StdRoutineGetBuildingHHaT(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (string, string, string, *cliapp.AppResult) {

	var (
		host     = ""
		hostarch = ""
		target   = ""
	)

	if t := getopt_result.GetLastNamedRetOptItem("--build-for-host"); t != nil {
		if t.Value != "" {
			host = t.Value
		} else {
			// TODO: smarter decigen needed for current_host value, or change description
			//       of --build-for-host option
			host = system.Host()
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--build-for-hostarch"); t != nil {
		if t.Value != "" {
			hostarch = t.Value
		} else {
			hostarch = host
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--build-to-target"); t != nil {
		if t.Value != "" {
			target = t.Value
		}
	}

	return host, hostarch, target, nil
}

func StdRoutineGetInstallationHHaT(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (string, []string, string, *cliapp.AppResult) {

	var (
		host     = ""
		hostarch = []string{}
		target   = ""
	)

	if t := getopt_result.GetLastNamedRetOptItem("--install-for-host"); t != nil {
		if t.Value != "" {
			host = t.Value
		} else {
			host = system.Host()
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--install-for-hostarch"); t != nil {
		if t.Value != "" {
			hostarch = strings.Split(t.Value, ",")
		} else {
			hostarch = system.Archs()
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--install-to-target"); t != nil {
		if t.Value != "" {
			target = t.Value
		}
	}

	return host, hostarch, target, nil
}

func StdRoutineGetASPListFiltersHHaT(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (string, string, string, *cliapp.AppResult) {

	var (
		host     = ""
		hostarch = ""
		target   = ""
	)

	if t := getopt_result.GetLastNamedRetOptItem("--show-only-host"); t != nil {
		if t.Value != "" {
			host = t.Value
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--show-only-hostarchs"); t != nil {
		if t.Value != "" {
			hostarch = t.Value
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--show-only-target"); t != nil {
		if t.Value != "" {
			target = t.Value
		}
	}

	return host, hostarch, target, nil
}
