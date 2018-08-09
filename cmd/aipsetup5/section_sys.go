package main

import (
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
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
				Name:             "get-asp",
				Callable:         CmdAipsetupSysGetASP,
				ShortDescription: "get asps for named package from repository",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_OPTION_NAMED_GET_ASP_FOR_HOST,
					STD_OPTION_NAMED_GET_ASP_FOR_HOSTARCH,
					STD_OPTION_NAMED_GET_ASP_CROSSBUILDER,
				},
				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Name:     "install-asp",
				Callable: CmdAipsetupSysInstall,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					// STD_OPTION_NAMED_INSTALLATION_TO_TARGET,
				},
				CheckArgs: true,
				MinArgs:   -1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Name:     "remove-asp",
				Callable: CmdAipsetupSysRemove,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_FORCE,
					// STD_OPTION_NAMED_INSTALLATION_TO_TARGET,
				},
				CheckArgs: true,
				MinArgs:   -1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Name:     "reduce-to",
				Callable: CmdAipsetupSysReduceTo,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					// STD_OPTION_NAMED_INSTALLATION_TO_TARGET,
				},
				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Name:             "list-asps",
				ShortDescription: "list installed asps",
				Description: "list all installed asps. " +
					"can be filtered with --host and --hostarch options",
				Callable: CmdAipsetupSysListAsps,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_OPTION_ASP_LIST_FILTER_HOST,
					STD_OPTION_ASP_LIST_FILTER_HOSTARCH,
				},
				CheckArgs: true,
				MinArgs:   -1,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:             "list-pkg-asps",
				ShortDescription: "list installed asps",
				Description:      "list installed asps of given package name.",
				Callable:         CmdAipsetupSysNameASPs,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
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

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	host, hostarch, res := StdRoutineGetASPListFiltersHostHostArch(getopt_result, sys)
	if res != nil && res.Code != 0 {
		return res
	}

	res_lst, err := sys.ASPs.ListFilteredInstalledASPs(
		host,
		hostarch,
	)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	res_lst_str := make([]string, 0)
	for _, i := range res_lst {
		res_lst_str = append(res_lst_str, i.String())
	}

	sort.Strings(res_lst_str)

	for _, i := range res_lst_str {
		fmt.Println(i)
	}

	return nil
}

func CmdAipsetupSysAllNames(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
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

	return nil
}

func CmdAipsetupSysNameASPs(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	asp_name := getopt_result.Args[0]

	res_lst, err := sys.ASPs.ListInstalledPackageNameASPs(asp_name, "", "")
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	res_lst_str := make([]string, 0)
	for _, i := range res_lst {
		res_lst_str = append(res_lst_str, i.String())
	}

	sort.Strings(res_lst_str)

	for _, i := range res_lst_str {
		fmt.Println(i)
	}

	return nil
}

func CmdAipsetupSysASPFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
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

	return nil
}

func CmdAipsetupSysInstall(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	errors_lst := make([][2]string, 0)

	for _, i := range getopt_result.Args {
		log.Info("Installing " + i)

		res := sys.ASPs.InstallASP(i)

		if res != nil {
			log.Error(res.Error())
			errors_lst = append(errors_lst, [2]string{i, res.Error()})
		}
	}

	if len(errors_lst) != 0 {

		for _, i := range errors_lst {
			fmt.Println("  ", i[0], ":", i[1])
		}

		return &cliapp.AppResult{
			Code:    10,
			Message: "some packages was installed with errors",
		}
	}

	sys.GetSystemUpdates().UpdatesAfterPkgInstall()

	return nil
}

func CmdAipsetupSysReduceTo(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	errors := false

	for _, i := range getopt_result.Args {
		res := sys.ASPs.InstallASPReduceToSubRoutine(i, true)

		if res != nil {
			log.Error(res.Error())
			errors = true
		}
	}

	if errors {
		return &cliapp.AppResult{
			Code:    10,
			Message: "some packages was reduced with errors",
		}
	}

	sys.GetSystemUpdates().UpdatesAfterPkgInstall()

	return nil
}

func CmdAipsetupSysRemove(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	force_if_reduce_possible := getopt_result.DoesHaveNamedRetOptItem(STD_FORCE.Name)

	// host, hostarch, target, res := StdRoutineGetInstallationHHaT(getopt_result, sys)
	// if res != nil && res.Code != 0 {
	// 	return res
	// }
	//
	// names := getopt_result.Args

	for _, i := range getopt_result.Args {
		log.Info("Removing" + i)

		iname, err := basictypes.NewASPNameFromString(i)
		if err != nil {
			return &cliapp.AppResult{
				Code:    11,
				Message: err.Error(),
			}
		}

		err = sys.ASPs.RemoveASP(iname, false, false, nil, force_if_reduce_possible)
		if err != nil {
			return &cliapp.AppResult{
				Code:    10,
				Message: err.Error(),
			}
		}

	}

	sys.GetSystemUpdates().UpdatesAfterPkgInstall()

	return nil
}

func CmdAipsetupSysGetASP(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	// TODO
	return nil
}
