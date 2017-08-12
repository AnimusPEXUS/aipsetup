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
				Name:     "asps",
				Callable: CmdAipsetupSysAllAsps,
			},

			&cliapp.AppCmdNode{
				Name:     "names",
				Callable: CmdAipsetupSysAllNames,
			},

			&cliapp.AppCmdNode{
				Name:     "name-asps",
				Callable: CmdAipsetupSysNameASPs,
			},

			&cliapp.AppCmdNode{
				Name:     "asp-files",
				Callable: CmdAipsetupSysASPFiles,
			},

			&cliapp.AppCmdNode{
				Name:     "asp-rm",
				Callable: CmdAipsetupSysRemoveASP,
			},

			&cliapp.AppCmdNode{
				Name:     "asp-install",
				Callable: CmdAipsetupSysInstallASP,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
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

	name, res := StdRoutineMustGetASPName(getopt_result)
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

	name, res := StdRoutineMustGetOneArg(getopt_result)
	if res.Code != 0 {
		return res
	}

	liaf_res, err := sys.ASPs.ListInstalledASPFiles(name)
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: "can't get file list for named ASP",
		}
	}

	files := liaf_res.FileList

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

	err := sys.ASPs.RemoveASP(name, true, true)

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

	name, res := StdRoutineMustGetASPName(getopt_result)
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
