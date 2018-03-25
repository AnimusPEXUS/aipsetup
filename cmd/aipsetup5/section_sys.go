package main

import (
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/cliapp"
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

			&cliapp.AppCmdNode{
				Name:     "install",
				Callable: CmdAipsetupSysInstall,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_OPTION_NAMED_INSTALLATION_FOR_HOST,
					STD_OPTION_NAMED_INSTALLATION_FOR_HOSTARCH,
					// STD_OPTION_NAMED_INSTALLATION_TO_TARGET,
				},
				CheckArgs: true,
				MinArgs:   -1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Name:     "remove",
				Callable: CmdAipsetupSysRemove,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_OPTION_NAMED_INSTALLATION_FOR_HOST,
					STD_OPTION_NAMED_INSTALLATION_FOR_HOSTARCH,
					// STD_OPTION_NAMED_INSTALLATION_TO_TARGET,
				},
				CheckArgs: true,
				MinArgs:   -1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Name:     "reduce-to",
				Callable: CmdAipsetupSysReduceTo,
				Description: "Reduces named installed asp to named installed asp. " +
					"Second argument may be a 'leatest' keyword. " +
					"If named asp is last left in system or is already latest - " +
					"nothing is done and 0 code returned.",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   2,
				MaxArgs:   2,
			},

			&cliapp.AppCmdNode{
				Name:             "list-asps",
				ShortDescription: "list installed asps",
				Description: "argumenth have to be package name. " +
					"Currently installed asps in pointed root will be listed.",
				Callable: CmdAipsetupSysListAsps,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_OPTION_ASP_LIST_FILTER_HOST,
					STD_OPTION_ASP_LIST_FILTER_HOSTARCH,
					// STD_OPTION_ASP_LIST_FILTER_TARGET,
				},
				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   1,
			},

			&cliapp.AppCmdNode{
				Name:             "list-files",
				ShortDescription: "list files installed by named asp",
				Callable:         CmdAipsetupSysASPFiles,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MaxArgs:   1,
				MinArgs:   1,
			},
		},
	}

}

// ----v-------v-------v---- rework 21 march 2018

func CmdAipsetupSysListAsps(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	host, hostarch, res := StdRoutineGetASPListFiltersHostHostArch(getopt_result, sys)
	if res != nil && res.Code != 0 {
		return res
	}

	res_lst, err := sys.ASPs.ListFilteredInstalledASPs(host, hostarch)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	sort.Strings(res_lst)

	for _, i := range res_lst {
		fmt.Println(i)
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupSysAllNames(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	host, hostarch, res := StdRoutineGetASPListFiltersHostHostArch(getopt_result, sys)
	if res != nil && res.Code != 0 {
		return res
	}

	res_lst, err := sys.ASPs.ListInstalledPackageNames(host, hostarch)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	sort.Strings(res_lst)

	for _, i := range res_lst {
		fmt.Println(i)
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupSysNameASPs(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	host, hostarch, res := StdRoutineGetASPListFiltersHostHostArch(getopt_result, sys)
	if res != nil && res.Code != 0 {
		return res
	}

	asp_name := getopt_result.Args[0]

	res_lst, err := sys.ASPs.ListInstalledPackageNameASPs(asp_name, host, hostarch)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	sort.Strings(res_lst)

	for _, i := range res_lst {
		fmt.Println(i)
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupSysASPFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	asp_name, err := basictypes.NewASPNameFromString(getopt_result.Args[0])
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: err.Error(),
		}
	}

	files, err := sys.ASPs.ListInstalledASPFiles(asp_name)
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

	return &cliapp.AppResult{}
}

func CmdAipsetupSysInstall(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	fmt.Println(getopt_result.String())

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	// host, hostarch, target, res := StdRoutineGetInstallationHHaT(getopt_result, sys)
	// if res != nil && res.Code != 0 {
	// 	return res
	// }
	//
	// names := getopt_result.Args
	panic("TODO")

	err := sys.ASPs.InstallASP("dummy")

	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupSysRemove(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	// host, hostarch, target, res := StdRoutineGetInstallationHHaT(getopt_result, sys)
	// if res != nil && res.Code != 0 {
	// 	return res
	// }
	//
	// names := getopt_result.Args
	panic("TODO")

	err := sys.ASPs.RemoveASP(nil, false, []string{})
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupSysReduceTo(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	return nil
}
