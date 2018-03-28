package main

import (
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/systemtriplet"
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

	STD_OPTION_BUILD_CURRENT_HOST = &cliapp.GetOptCheckListItem{
		Name: "--build-current-host",
		Description: "Override current host. " +
			"Default is current host.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_FOR_HOST = &cliapp.GetOptCheckListItem{
		Name: "--build-for-host",
		Description: "Select main name of system which will run package. " +
			"Default is current host.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_FOR_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--build-for-hostarch",
		Description: "Select subsystem (subarch) name which will run package. " +
			"Default is equal to value calculated (or given to) with " +
			"--build-for-host.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_CROSSBUILDER = &cliapp.GetOptCheckListItem{
		Name: "--build-crossbuilder",
		Description: "Configure and build this package to be a " +
			"crossbuilder for named system. " +
			"Default is empty - disabling this mode",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_CROSSBUILDING = &cliapp.GetOptCheckListItem{
		Name: "--build-crossbuilding",
		Description: "Configure this package to be built with crosscompiler " +
			"and use crosscompiler to build it. " +
			"Default is empty - disabling this mode",
		HaveDefault:   true,
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
			"Default is to use subarch values for named host from aipsetup5.system.ini. ",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_NAMED_INSTALLATION_CROSSBUILDER = &cliapp.GetOptCheckListItem{
		Name: "--install-crossbuilder",
		Description: "Select crossbuilder for which system you wish to install. " +
			"No default value. You must select one manually.",
		HaveDefault:   false,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_HOST = &cliapp.GetOptCheckListItem{
		Name: "--show-only-host",
		Description: "Value is exact name. Show only selected. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--show-only-hostarchs",
		Description: "Value is exact name. Show only selected. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_CROSSBUILDER = &cliapp.GetOptCheckListItem{
		Name: "--show-only-crossbuilder",
		Description: "Value is exact name. Show only selected. " +
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

func StdRoutineGetRootOptionAndSystemObject(getopt_result *cliapp.GetOptResult) (
	root string,
	sys *aipsetup.System,
	ret *cliapp.AppResult,
) {

	root = "/"
	ret = &cliapp.AppResult{}

	if root_opt, ok := StdRoutineGetRootOption(getopt_result); ok {
		root = root_opt
	}

	sys = aipsetup.NewSystem(root)

	return
}

func StdRoutineGetBuildingHostHostArch(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (string, string, *cliapp.AppResult) {

	var (
		err      error
		host     = ""
		hostarch = ""
	)

	if t := getopt_result.GetLastNamedRetOptItem("--build-for-host"); t != nil {
		if t.Value != "" {
			host = t.Value
		} else {
			host, err = system.Host()
			if err != nil {
				return "", "", &cliapp.AppResult{Code: 22, Message: err.Error()}
			}
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--build-for-hostarch"); t != nil {
		if t.Value != "" {
			hostarch = t.Value
		} else {
			hostarch = host
		}
	}

	if _, err := systemtriplet.NewFromString(host); err != nil {
		return "", "", &cliapp.AppResult{
			Code:    21,
			Message: "Can't parse --build-for-host as system name triplet",
		}
	}

	if _, err := systemtriplet.NewFromString(hostarch); err != nil {
		return "", "", &cliapp.AppResult{
			Code:    21,
			Message: "Can't parse --build-for-hostarch as system name triplet",
		}
	}

	return host, hostarch, nil
}

func StdRoutineGetInstallationHostHostArch(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (string, []string, *cliapp.AppResult) {

	var (
		err      error
		host     = ""
		hostarch = []string{}
	)

	if t := getopt_result.GetLastNamedRetOptItem("--install-for-host"); t != nil {
		if t.Value != "" {
			host = t.Value
		} else {
			host, err = system.Host()
			if err != nil {
				return "", nil, &cliapp.AppResult{Code: 22, Message: err.Error()}
			}
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem("--install-for-hostarch"); t != nil {
		if t.Value != "" {
			hostarch = strings.Split(t.Value, ",")
		} else {
			hostarch, err = system.Archs()
			if err != nil {
				return "", nil, &cliapp.AppResult{Code: 22, Message: err.Error()}
			}
		}
	}

	if _, err := systemtriplet.NewFromString(host); err != nil {
		return "", nil, &cliapp.AppResult{
			Code:    21,
			Message: "Can't parse --install-for-host as system name triplet",
		}
	}

	for _, i := range hostarch {
		if _, err := systemtriplet.NewFromString(i); err != nil {
			return "", nil, &cliapp.AppResult{
				Code:    21,
				Message: "Can't parse --install-for-hostarch as system name triplet",
			}
		}
	}

	return host, hostarch, nil
}

func StdRoutineGetASPListFiltersHostHostArch(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (string, string, *cliapp.AppResult) {

	var (
		host     = ""
		hostarch = ""
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

	if _, err := systemtriplet.NewFromString(host); err != nil {
		return "", "", &cliapp.AppResult{
			Code:    21,
			Message: "Can't parse --show-only-host as system name triplet",
		}
	}

	if _, err := systemtriplet.NewFromString(hostarch); err != nil {
		return "", "", &cliapp.AppResult{
			Code:    21,
			Message: "Can't parse --show-only-hostarchs as system name triplet",
		}
	}

	return host, hostarch, nil
}
