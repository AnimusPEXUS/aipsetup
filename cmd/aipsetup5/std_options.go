package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
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
		Name: "--current-host",
		Description: "Override current host. " +
			"Default is current host.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_FOR_HOST = &cliapp.GetOptCheckListItem{
		Name: "--host",
		Description: "Select main name of system which will run package. " +
			"Default is current host.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_FOR_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--hostarch",
		Description: "Select subsystem (subarch) name which will run package. " +
			"Default is equal to value calculated (or given to) with " +
			"--build-for-host.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_CROSSBUILDER = &cliapp.GetOptCheckListItem{
		Name: "--crossbuilder",
		Description: "Configure and build this package to be a " +
			"crossbuilder for named system. " +
			"Default is empty - disabling this mode",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_BUILD_CROSSBUILDING = &cliapp.GetOptCheckListItem{
		Name: "--crossbuilding",
		Description: "Configure this package to be built with crosscompiler " +
			"and use crosscompiler to build it. " +
			"Default is empty - disabling this mode",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	// -------------------------

	STD_OPTION_MASS_BUILD_CURRENT_HOST = STD_OPTION_BUILD_CURRENT_HOST

	STD_OPTION_MASS_BUILD_FOR_HOST = STD_OPTION_BUILD_FOR_HOST

	STD_OPTION_MASS_BUILD_FOR_HOSTARCHS = &cliapp.GetOptCheckListItem{
		Name: "--hostarchs",
		Description: "Select subsystems (subarchs) names which will run packages. " +
			"Default is use aipsetup.system.ini to take all subarchs for host pointed " +
			"or calculated for --host.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_MASS_BUILD_CROSSBUILDER = STD_OPTION_BUILD_CROSSBUILDER

	STD_OPTION_MASS_BUILD_CROSSBUILDING = STD_OPTION_BUILD_CROSSBUILDING

	// -------------------------

	STD_OPTION_NAMED_GET_ASP_FOR_HOST = &cliapp.GetOptCheckListItem{
		Name: "--host",
		Description: "Select hosting system. " +
			"Default is current host",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_NAMED_GET_ASP_FOR_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--hostarch",
		Description: "Select hosting subarch system. " +
			"Default is equal to current host + " +
			"all it's subarchs found already installed in system.",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_NAMED_GET_ASP_CROSSBUILDER = &cliapp.GetOptCheckListItem{
		Name: "--crossbuilder",
		Description: "Select crossbuilder for which system you wish to install. " +
			"No default value. You must select one manually.",
		HaveDefault:   false,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,

		// -------------------------
	}

	STD_OPTION_ASP_LIST_FILTER_HOST = &cliapp.GetOptCheckListItem{
		Name: "--host",
		Description: "Value is exact name. Show only selected. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_HOSTARCH = &cliapp.GetOptCheckListItem{
		Name: "--hostarch",
		Description: "Value is exact name. Show only selected. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_OPTION_ASP_LIST_FILTER_CROSSBUILDER = &cliapp.GetOptCheckListItem{
		Name: "--crossbuilder",
		Description: "Value is exact name. Show only selected. " +
			"default is empty value, which shows all",
		HaveDefault:   true,
		Default:       "",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_NAMES_ARE_CATEGORIES = &cliapp.GetOptCheckListItem{
		Name:        "-c",
		Description: "named names are categories",
	}

	STD_NAMES_ARE_CATEGORIES_PRESERVE_NESTING = &cliapp.GetOptCheckListItem{
		Name: "--cpn",
		Description: "category preserve nesting. " +
			"use with -c. creates subdirs and " +
			"preserves categorization",
	}

	STD_NAMES_ARE_CATEGORIES_IS_PREFIXES = &cliapp.GetOptCheckListItem{
		Name: "--cip",
		Description: "category is prefix. " +
			"use with -c. all packages " +
			"category of which starts with value will match",
	}

	STD_NAMES_ARE_GROUPS = &cliapp.GetOptCheckListItem{
		Name:        "-g",
		Description: "named names are groups",
	}

	STD_FORCE = &cliapp.GetOptCheckListItem{
		Name:        "-f",
		Description: "force",
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

func StdRoutineGetRootOptionAndSystemObject(
	getopt_result *cliapp.GetOptResult,
	log *logger.Logger,
) (
	root string,
	sys *aipsetup.System,
	ret *cliapp.AppResult,
) {

	root = "/"
	ret = nil

	if root_opt, ok := StdRoutineGetRootOption(getopt_result); ok {
		root = root_opt
	}

	sys = aipsetup.NewSystem(root, log)

	return
}

func StdRoutineGetASPListFiltersHostHostArch(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (string, string, *cliapp.AppResult) {

	var (
		host     = ""
		hostarch = ""
	)

	if t := getopt_result.GetLastNamedRetOptItem(STD_OPTION_ASP_LIST_FILTER_HOST.Name); t != nil {
		if t.Value != "" {
			host = t.Value
		}
	}

	if t := getopt_result.GetLastNamedRetOptItem(STD_OPTION_ASP_LIST_FILTER_HOSTARCH.Name); t != nil {
		if t.Value != "" {
			hostarch = t.Value
		}
	}

	if hostarch != "" && host == "" {
		return "", "", &cliapp.AppResult{
			Code: 22,
			Message: STD_OPTION_ASP_LIST_FILTER_HOSTARCH.Name +
				" can't be specified without specifiyng " +
				STD_OPTION_ASP_LIST_FILTER_HOST.Name,
		}
	}

	if host != "" {
		if _, err := systemtriplet.NewFromString(host); err != nil {
			return "", "", &cliapp.AppResult{
				Code: 21,
				Message: "Can't parse " +
					STD_OPTION_ASP_LIST_FILTER_HOST.Name +
					" as system name triplet",
			}
		}
	}

	if hostarch != "" {
		if _, err := systemtriplet.NewFromString(hostarch); err != nil {
			return "", "", &cliapp.AppResult{
				Code: 21,
				Message: "Can't parse " +
					STD_OPTION_ASP_LIST_FILTER_HOSTARCH.Name +
					"as system name triplet",
			}
		}
	}

	return host, hostarch, nil
}

func StdRoutineGetMassBuildOptions(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (
	current_host, for_host string,
	for_hostarchs []string,
	crossbuilder, crossbuilding string,
	ret *cliapp.AppResult,
) {

	var err error

	if o := getopt_result.GetLastNamedRetOptItem("--build-current-host"); o != nil {
		fmt.Println("--build-current-host")
		current_host = o.Value
	} else {
		current_host, err = system.Host()
		if err != nil {
			ret = &cliapp.AppResult{Code: 200, Message: err.Error()}
			return
		}
	}

	if o := getopt_result.GetLastNamedRetOptItem("--build-for-host"); o != nil {
		for_host = o.Value
	} else {
		for_host = current_host
	}

	if o := getopt_result.GetLastNamedRetOptItem("--build-for-hostarchs"); o != nil {
		for_hostarchs = strings.Split(o.Value, ",")
	} else {
		cfg := system.Cfg()
		if cfg == nil {
			ret = &cliapp.AppResult{Code: 201, Message: "aipsetup configuration problems"}
		}
		s := cfg.Section(for_host)
		// if err != nil {
		// 	ret = &cliapp.AppResult{Code: 202, Message: err.Error()}
		// }
		a := s.Key("archs")
		for_hostarchs = a.Strings(",")
		for_hostarchs = append(for_hostarchs, for_host)
		sort.Strings(for_hostarchs)
	}

	if o := getopt_result.GetLastNamedRetOptItem("--build-crossbuilder"); o != nil {
		crossbuilder = o.Value
	} else {
		crossbuilder = ""
	}

	if o := getopt_result.GetLastNamedRetOptItem("--build-crossbuilding"); o != nil {
		crossbuilding = o.Value
	} else {
		crossbuilding = ""
	}

	return
}
