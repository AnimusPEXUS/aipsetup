package main

import (
	"fmt"
	"path"
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/repository"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupMBuild() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{

		Name: "mbuild",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Callable: CmdAipsetupMassBuildInit,
				Name:     "init",

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_OPTION_MASS_BUILD_CURRENT_HOST,
					STD_OPTION_MASS_BUILD_FOR_HOST,
					STD_OPTION_MASS_BUILD_FOR_HOSTARCHS,
					STD_OPTION_MASS_BUILD_CROSSBUILDER,
					STD_OPTION_MASS_BUILD_CROSSBUILDING,
				},

				MaxArgs:   0,
				MinArgs:   0,
				CheckArgs: true,
			},

			&cliapp.AppCmdNode{
				Callable: CmdAipsetupMassBuildGetSrc,
				Name:     "get-src",

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_NAMES_ARE_CATEGORIES,
					STD_NAMES_ARE_GROUPS,
				},

				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Callable: CmdAipsetupMassBuildPerform,
				Name:     "full",

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},

				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},
		},
	}
}

func CmdAipsetupMassBuildInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(
		getopt_result,
		log,
	)
	if res != nil && res.Code != 0 {
		return res
	}

	mbuild, err := aipsetup.NewMassBuilder(".", sys, log)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	current_host,
		for_host, for_hostarchs,
		crossbuilder, crossbuilding,
		res := StdRoutineGetMassBuildOptions(getopt_result, sys)

	mbuild_info := &basictypes.MassBuilderInfo{
		Host:               for_host,
		HostArchs:          for_hostarchs,
		CrossbuilderTarget: crossbuilder,
		CrossbuildersHost:  crossbuilding,
		InitiatedByHost:    current_host,
	}

	err = mbuild.WriteInfo(mbuild_info)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupMassBuildGetSrc(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	mbuild, err := aipsetup.NewMassBuilder(".", sys, log)
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: err.Error(),
		}
	}

	repo, err := repository.NewRepository(sys, log)
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: err.Error(),
		}
	}

	get_by_name_func := func(name string) error {

		t, err := repo.DetermineNewestStableTarball(name)
		if err != nil {
			return err
		}

		log.Info(path.Base(t))

		err = repo.CopyTarballToDir(name, t, path.Join(mbuild.GetTarballsPath()))
		if err != nil {
			return err
		}
		return nil
	}

	res = MiscDoSomethingForGroupsCategoriesOrLists(
		sys,
		getopt_result,
		adds,
		get_by_name_func,
	)
	if res != nil && res.Code != 0 {
		return res
	}

	return nil

}

func CmdAipsetupMassBuildPerform(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(
		getopt_result,
		log,
	)
	if res != nil && res.Code != 0 {
		return res
	}

	mbuild, err := aipsetup.NewMassBuilder(".", sys, log)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	_, f, err := mbuild.PerformMassBuilding()
	if err != nil {
		return &cliapp.AppResult{Code: 11, Message: err.Error()}
	}

	keys := make([]string, 0)
	for k, _ := range f {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	failed := false

	for _, i := range keys {
		fmt.Println("arch", i)
		sort.Strings(f[i])
		for _, j := range f[i] {
			fmt.Println("   ", j)
			failed = true
		}
	}

	if failed {
		return &cliapp.AppResult{
			Code:    12,
			Message: "some packages building have failed",
		}
	}

	return nil
}
