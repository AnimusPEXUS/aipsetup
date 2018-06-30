package main

import (
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/repository"
	"github.com/AnimusPEXUS/aipsetup/repository/providers"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
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
					STD_NAMES_ARE_CATEGORIES_PRESERVE_NESTING,
					STD_NAMES_ARE_CATEGORIES_IS_PREFIXES,
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
					STD_NAMES_ARE_CATEGORIES_IS_PREFIXES,
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

			&cliapp.AppCmdNode{
				Name:        "parse",
				Callable:    CmdAipsetupRepoTarballParse,
				Description: "parse tarball name and print results",
				CheckArgs:   true,
				MinArgs:     0,
				MaxArgs:     -1,
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

	repo, err := repository.NewRepository(sys, log)
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
	if err2 != nil && err2.Code != 0 {
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

	return nil

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

	repo, err := repository.NewRepository(sys, log)
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

	ok_asp_names := make([]string, 0)
	ok_src_names := make([]string, 0)
	err_names := make([]string, 0)
	was_errors := false

	for _, i := range getopt_result.Args {

		_, err := basictypes.NewASPNameFromString(i)
		if err == nil {
			ok_asp_names = append(ok_asp_names, i)
			continue
		}

		if tarballname.IsPossibleTarballName(i) {
			ok_src_names = append(ok_src_names, i)
			continue
		}

		err_names = append(err_names, i)
		was_errors = true
	}

	for _, i := range ok_asp_names {
		fmt.Print(word, " ", i)

		ok, err := aipsetup.CheckAspPackageByFilename(i)
		if err != nil {
			fmt.Println(" - error:", err)
			was_errors = true
			continue

		}

		if !ok {
			err := errors.New("given file didn't passed package check")
			fmt.Println(" - error:", err)
			was_errors = true
			continue
		}

		err = repo.MoveInASP(i, copy)
		if err != nil {
			fmt.Println(" - error:", err)
			was_errors = true
			continue
		}

		fmt.Println(" - ok")
	}

	for _, i := range ok_src_names {
		fmt.Print(word, " ", i)

		err = repo.MoveInTarball(i, copy)
		if err != nil {
			fmt.Println(" - error:", err)
			was_errors = true
			continue
		}

		fmt.Println(" - ok")
	}

	if len(err_names) != 0 {
		log.Error(fmt.Sprintf("There was %d error(s):"))
		for _, i := range err_names {
			fmt.Println("   ", i)
		}
	}

	if was_errors {
		return &cliapp.AppResult{Code: 12}
	}

	return nil
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

func CmdAipsetupRepoTarballParse(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	for _, i := range getopt_result.Args {

		fmt.Println(i)

		res, err := pkginfodb.DetermineTarballPackageInfo(i)
		if err != nil {
			fmt.Println("error:", err.Error())
			continue
		}

		res_names := make([]string, 0)

		for k, _ := range res {
			res_names = append(res_names, k)
		}

		sort.Strings(res_names)

		fmt.Printf(
			"  results (count %d): %s.\n",
			len(res_names),
			strings.Join(res_names, ", "),
		)

		for _, j := range res_names {

			fmt.Println("  parsing and classification for result:", j)

			//			var name string
			var info *basictypes.PackageInfo

			for _, info = range res {
			}

			parser, err := tarballnameparsers.Get(info.TarballFileNameParser)
			if err != nil {
				fmt.Println("   error:", err.Error())
				continue
			}

			result, err := parser.Parse(i)
			if err != nil {
				fmt.Println("   tarballname parsing error:", err.Error())
				continue
			}

			fmt.Println("parse result")
			fmt.Println(result.InfoText())

			clssfier, err := tarballstabilityclassification.Get(info.TarballStabilityClassifier)
			if err != nil {
				fmt.Println("   error:", err.Error())
				continue
			}

			clssfier_ret, err := clssfier.Check(result)
			if err != nil {
				fmt.Println("   classifier returned classification error:", err.Error())
				continue
			}

			fmt.Println("classification result")
			fmt.Println(clssfier_ret.String())

			fmt.Println("======")

		}

		fmt.Println("-------")
	}

	return nil
}
