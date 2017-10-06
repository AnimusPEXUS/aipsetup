package main

import (
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/cliapp"
)

func SectionAipsetupSys() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "sys",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:             "names",
				Callable:         CmdAipsetupSysAllNames,
				ShortDescription: "list installed package names",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:             "name-asps",
				Callable:         CmdAipsetupSysNameASPs,
				ShortDescription: "list asps with given package name",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   1,
			},

			// &cliapp.AppCmdNode{
			// 	Name:    "pkg",
			// 	SubCmds: []*cliapp.AppCmdNode{},
			// },

			&cliapp.AppCmdNode{
				Name: "asp",
				SubCmds: []*cliapp.AppCmdNode{

					&cliapp.AppCmdNode{
						Name:             "list",
						ShortDescription: "list installed asps",
						Callable:         CmdAipsetupSysAllAsps,
						AvailableOptions: cliapp.GetOptCheckList{
							STD_ROOT_OPTION,
						},
						CheckArgs: true,
						MinArgs:   0,
						MaxArgs:   0,
					},

					&cliapp.AppCmdNode{
						Name:             "files",
						ShortDescription: "list files installed by named asp",
						Callable:         CmdAipsetupSysASPFiles,
						AvailableOptions: cliapp.GetOptCheckList{
							STD_ROOT_OPTION,
						},
						CheckArgs: true,
						MaxArgs:   1,
						MinArgs:   1,
					},

					&cliapp.AppCmdNode{
						Name:             "remove",
						ShortDescription: "remove asp package",
						Callable:         CmdAipsetupSysRemoveASP,
						AvailableOptions: cliapp.GetOptCheckList{
							STD_ROOT_OPTION,
						},
						CheckArgs: true,
						MinArgs:   1,
						MaxArgs:   1,
					},

					&cliapp.AppCmdNode{
						Name: "install",
						ShortDescription: "install asp package.\n" +
							"install asps via 'sys pkg install' for automatic " +
							"uninstallation of old asps",
						Callable: CmdAipsetupSysInstallASP,
						AvailableOptions: cliapp.GetOptCheckList{
							STD_ROOT_OPTION,
						},
						CheckArgs: true,
						MinArgs:   1,
						MaxArgs:   1,
					},
				},
			},
		},
	}

}

func CmdAipsetupSysAllAsps(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, host, arch, sys, res := StdRoutineRootHostArchSys(getopt_result)
	if res.Code != 0 {
		return res
	}

	installed_asps, err := sys.ASPs.ListInstalledASPs(host, arch)
	if err != nil {
		return &cliapp.AppResult{Code: 1}
	}

	num_len := len(fmt.Sprintf("%d", len(installed_asps)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range installed_asps {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysAllNames(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, host, arch, sys, res := StdRoutineRootHostArchSys(getopt_result)
	if res.Code != 0 {
		return res
	}

	installed_names, err := sys.ASPs.ListInstalledPackageNames(host, arch)
	if err != nil {
		return &cliapp.AppResult{Code: 1}
	}

	sort.Strings(installed_names)

	num_len := len(fmt.Sprintf("%d", len(installed_names)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range installed_names {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysNameASPs(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, host, arch, sys, res := StdRoutineRootHostArchSys(getopt_result)
	if res.Code != 0 {
		return res
	}

	name, res := StdRoutineMustGetOneArg(getopt_result)
	if res.Code != 0 {
		return res
	}

	installed_archs, err :=
		sys.ASPs.ListInstalledPackageNameASPs(name, host, arch)
	if err != nil {
		return &cliapp.AppResult{Code: 1}
	}

	sort.Strings(installed_archs)

	num_len := len(fmt.Sprintf("%d", len(installed_archs)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range installed_archs {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysASPFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, sys, res := StdRoutineRootSys(getopt_result)
	if res.Code != 0 {
		return res
	}

	name, res := StdRoutineMustGetASPName(getopt_result)
	if res.Code != 0 {
		return res
	}

	files, err := sys.ASPs.ListInstalledASPFiles(name)
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: "can't get file list for named ASP",
		}
	}

	sort.Strings(files)

	num_len := len(fmt.Sprintf("%d", len(files)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range files {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysRemoveASP(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, sys, res := StdRoutineRootSys(getopt_result)
	if res.Code != 0 {
		return res
	}

	name, res := StdRoutineMustGetASPName(getopt_result)
	if res.Code != 0 {
		return res
	}

	err := sys.ASPs.RemoveASP(name, false, []string{})
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysInstallASP(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	fmt.Println(getopt_result.String())

	_, sys, res := StdRoutineRootSys(getopt_result)
	if res.Code != 0 {
		return res
	}

	name, res := StdRoutineMustGetOneArg(getopt_result)
	if res.Code != 0 {
		return res
	}

	err := sys.ASPs.InstallASP(name)

	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{Code: 0}
}
