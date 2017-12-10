package main

import (
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

	STD_BUILDER_HOST_OPTION = &cliapp.GetOptCheckListItem{
		Name: "--host",
		Description: "System which will run. " +
			"If omitted - current system's host value is used",
		HaveDefault:   false,
		Default:       "x86_64-pc-linux-gnu",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_BUILDER_ARCH_OPTION = &cliapp.GetOptCheckListItem{
		Name: "--arch",
		Description: "System's subarch which will run. " +
			"If omitted - host value is used",
		HaveDefault:   false,
		Default:       "x86_64-pc-linux-gnu",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_BUILDER_BUILD_OPTION = &cliapp.GetOptCheckListItem{
		Name: "--build",
		Description: "System which will build. " +
			"If omitted - host value is used",
		HaveDefault:   false,
		Default:       "x86_64-pc-linux-gnu",
		IsRequired:    false,
		MustHaveValue: true,
	}

	STD_BUILDER_TARGET_OPTION = &cliapp.GetOptCheckListItem{
		Name:          "--target",
		Description:   "This option is for configuring crosscompiler",
		HaveDefault:   false,
		Default:       "x86_64-pc-linux-gnu",
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

func StdRoutineGetHostArchOptions(getopt_result *cliapp.GetOptResult) (
	string, bool,
	string, bool,
) {

	host := ""
	host_ok := false
	arch := ""
	arch_ok := false

	{
		res := getopt_result.GetLastNamedRetOptItem("--host")
		if res != nil {
			host = res.Value
			host_ok = true
		}
	}

	{
		res := getopt_result.GetLastNamedRetOptItem("--arch")
		if res != nil {
			arch = res.Value
			arch_ok = true
		}
	}

	return host, host_ok, arch, arch_ok
}

func StdRoutineRootHostArchSys(getopt_result *cliapp.GetOptResult) (
	root string,
	host string,
	arch string,
	sys *aipsetup.System,
	ret *cliapp.AppResult,
) {

	root = "/"
	host = ""
	arch = ""
	ret = &cliapp.AppResult{Code: 0}

	root_opt, ok := StdRoutineGetRootOption(getopt_result)
	if ok {
		root = root_opt
	}

	sys = aipsetup.NewSystem(root)

	{
		host_opt, host_ok, arch_opt, arch_ok :=
			StdRoutineGetHostArchOptions(getopt_result)

		if host_ok {
			host = host_opt
		}

		if arch_ok {
			arch = arch_opt
		}
	}

	if arch != "" && host == "" {
		ret = &cliapp.AppResult{
			Code:    10,
			Message: "if host is empty, arch must be empty too",
		}
	}
	return
}

func StdRoutineRootSys(getopt_result *cliapp.GetOptResult) (
	root string,
	sys *aipsetup.System,
	ret *cliapp.AppResult,
) {

	root = "/"
	ret = &cliapp.AppResult{Code: 0}

	if root_opt, ok := StdRoutineGetRootOption(getopt_result); ok {
		root = root_opt
	}

	sys = aipsetup.NewSystem(root)

	return
}

func StdRoutineMustGetOneArg(getopt_result *cliapp.GetOptResult) (
	res string,
	err *cliapp.AppResult,
) {

	if len(getopt_result.Args) != 1 {
		res = ""
		err = &cliapp.AppResult{
			Code:    1,
			Message: "exactly one argument required",
		}
	} else {
		res = getopt_result.Args[0]
		err = &cliapp.AppResult{Code: 0}
	}

	return
}

func StdRoutineMustGetASPName(getopt_result *cliapp.GetOptResult) (
	*aipsetup.ASPName,
	*cliapp.AppResult,
) {

	res, res_err := StdRoutineMustGetOneArg(getopt_result)
	if res_err.Code != 0 {
		return nil, res_err
	}

	name, err := aipsetup.NewASPNameFromString(res)
	if err != nil {
		return nil, &cliapp.AppResult{
			Code:    11,
			Message: "Can't parse given string as ASP name",
		}
	}

	return name, &cliapp.AppResult{Code: 0}
}

func StdRoutineHostArchBuildTarget(
	getopt_result *cliapp.GetOptResult,
	system *aipsetup.System,
) (
	host string,
	arch string,
	build string,
	target string,
) {

	host = system.Host()

	host_o := getopt_result.GetLastNamedRetOptItem("--host")
	if host_o != nil {
		host = host_o.Value
	}

	arch = host

	arch_o := getopt_result.GetLastNamedRetOptItem("--arch")
	if arch_o != nil {
		arch = arch_o.Value
	}

	build = host

	build_o := getopt_result.GetLastNamedRetOptItem("--build")
	if build_o != nil {
		build = build_o.Value
	}

	target = host

	target_o := getopt_result.GetLastNamedRetOptItem("--target")
	if target_o != nil {
		target = target_o.Value
	}

	return
}
