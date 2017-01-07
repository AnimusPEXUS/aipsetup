package main

import (
	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/cliapp"
)

func ReturnStdOptRoot(getopt_result *cliapp.GetOptResult) *string {
	var ret *string = nil

	res := getopt_result.GetLastNamedRetOptItem("--root")
	if res != nil {
		ret = &res.Value
	}

	return ret
}

func ReturnStdOptHostArch(
	getopt_result *cliapp.GetOptResult,
	host_if_empty string,
) (*string, *string) {

	var host *string = nil
	var arch *string = nil

	{
		res := getopt_result.GetLastNamedRetOptItem("--host")
		if res != nil {
			host = &res.Value
		}
	}

	{
		res := getopt_result.GetLastNamedRetOptItem("--arch")
		if res != nil {
			arch = &res.Value
		}
	}

	if arch != nil && host == nil {
		host = &host_if_empty
	}

	return host, arch
}

func CmdRoutineRootHostArchSys(getopt_result *cliapp.GetOptResult) (
	root string,
	host string,
	arch string,
	sys *aipsetup.System,
	ret *cliapp.AppResult,
) {

	root = "/"
	host = ""
	arch = ""
	ret = nil

	root_opt := ReturnStdOptRoot(getopt_result)

	if root_opt != nil {
		root = *root_opt
	}

	sys = aipsetup.NewSystem(root)

	host_opt, arch_opt := ReturnStdOptHostArch(getopt_result, sys.Host())

	if host_opt != nil {
		host = *host_opt
	}

	if arch_opt != nil {
		arch = *arch_opt
	}

	if arch != "" && host == "" {
		ret = &cliapp.AppResult{
			Code:    10,
			Message: "if host is empty, arch must be empty also",
		}
	}
	return
}

func CmdRoutineRootSys(getopt_result *cliapp.GetOptResult) (
	root string,
	sys *aipsetup.System,
	ret *cliapp.AppResult,
) {

	root = "/"
	ret = nil

	root_opt := ReturnStdOptRoot(getopt_result)

	if root_opt != nil {
		root = *root_opt
	}

	sys = aipsetup.NewSystem(root)

	return
}

func CmdRoutineASPName(getopt_result *cliapp.GetOptResult) (
	name string,
	ret *cliapp.AppResult,
) {

	name = ""
	ret = nil

	if len(getopt_result.Args) != 1 {
		ret = &cliapp.AppResult{Code: 10, Message: "one argument is required"}
	} else {
		name = getopt_result.Args[0]
	}

	return
}
