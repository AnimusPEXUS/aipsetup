package main

import (
	"github.com/AnimusPEXUS/aipsetup"

	"github.com/AnimusPEXUS/cliapp"
)

var (
	STD_ROOT_OPTION = &cliapp.GetOptCheckListItem{
		Name:          "--root",
		Description:   "System root to work with",
		Default:       "/",
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
	res string,
	err *cliapp.AppResult,
) {

	res, err = StdRoutineMustGetOneArg(getopt_result)
	if err.Code != 0 {
		return
	}

	// TODO: add ASP name validity check

	return
}
