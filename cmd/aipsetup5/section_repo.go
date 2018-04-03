package main

import (
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/providers"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
)

func SectionAipsetupRepo() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{

		Name: "repo",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name: "providers",

				SubCmds: []*cliapp.AppCmdNode{

					&cliapp.AppCmdNode{
						Name:     "list",
						Callable: CmdAipsetupRepoProvidersList,
					},
				},
			},

			&cliapp.AppCmdNode{
				Name:      "get-src",
				Callable:  CmdAipsetupRepoGetSrc,
				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   -1,

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_NAMES_ARE_CATEGORIES,
					STD_NAMES_ARE_GROUPS,
				},
			},

			&cliapp.AppCmdNode{
				Name:        "up",
				Callable:    CmdAipsetupRepoUp,
				Description: "update sources of named package or packages' names by group or category",
				CheckArgs:   true,
				MinArgs:     1,
				MaxArgs:     -1,

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_NAMES_ARE_CATEGORIES,
					STD_NAMES_ARE_GROUPS,
				},
			},

			&cliapp.AppCmdNode{
				Name:        "put",
				Callable:    CmdAipsetupRepoPut,
				Description: "copy selected files to repository",
				CheckArgs:   true,
				MinArgs:     0,
				MaxArgs:     -1,

				AvailableOptions: cliapp.GetOptCheckList{
					&cliapp.GetOptCheckListItem{
						Name:        "-m",
						Description: "move. not copy.",
					},
				},
			},
		},
	}

}

func CmdAipsetupRepoProvidersList(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	names := make([]string, 0)

	for k, _ := range providers.Index {
		names = append(names, k)
	}

	sort.Strings(names)

	for _, i := range names {
		fmt.Println(i)
	}

	return new(cliapp.AppResult)
}

func CmdAipsetupRepoUp(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	repo, err := tarballrepository.NewRepository(sys, log)
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: err.Error(),
		}
	}

	failed_list := make([]string, 0)

	get_by_name_func := func(name string) error {

		fmt.Println("--------------------------------------")
		fmt.Println("   updating " + name)
		fmt.Println("--------------------------------------")

		err := repo.PerformPackageSourcesUpdate(name)
		if err != nil {
			failed_list = append(failed_list, name)
			fmt.Println("error")
			fmt.Println(err)
		}

		return nil
	}

	err2 := MiscDoSomethingForGroupsCategoriesOrLists(
		sys,
		getopt_result,
		adds,
		get_by_name_func,
	)
	if err2.Code != 0 {
		return err2
	}

	if len(failed_list) != 0 {
		sort.Strings(failed_list)
		fmt.Println("failed to update:")
		for _, i := range failed_list {
			fmt.Println("   " + i)
		}
		return &cliapp.AppResult{
			Code:    20,
			Message: "Some packages src repo update failed",
		}
	}

	return &cliapp.AppResult{}

}

func CmdAipsetupRepoPut(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	repo, err := tarballrepository.NewRepository(sys, log)
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: err.Error(),
		}
	}

	copy := true

	if getopt_result.DoesHaveNamedRetOptItem("-m") {
		copy = false
	}

	word := "moving"
	if copy {
		word = "copying"
	}

	was_errors := false

	for _, i := range getopt_result.Args {
		fmt.Print(word, " ", i)
		if _, err := basictypes.NewASPNameFromString(i); err == nil {
			err = repo.MoveInASP(i, copy)
			if err != nil {
				fmt.Println(" - error:", err)
				was_errors = true
				continue
			}
			fmt.Println(" - ok")
			continue
		}
		if tarballname.IsPossibleTarballName(i) {
			err = repo.MoveInTarball(i, copy)
			if err != nil {
				fmt.Println(" - error:", err)
				was_errors = true
				continue
			}
			fmt.Println(" - ok")
			continue
		}
		fmt.Println("- error: unknown file")
		was_errors = true
	}

	if was_errors {
		return &cliapp.AppResult{Code: 12}
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupRepoGetSrc(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	repo, err := tarballrepository.NewRepository(sys, log)
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

		err = repo.CopyTarballToDir(name, t, ".")
		if err != nil {
			return err
		}
		return nil
	}

	ret := MiscDoSomethingForGroupsCategoriesOrLists(
		sys,
		getopt_result,
		adds,
		get_by_name_func,
	)

	return ret
}
