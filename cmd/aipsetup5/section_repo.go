package main

import (
	"errors"
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/providers"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
	"github.com/AnimusPEXUS/utils/version/versioncomparators"
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
					&cliapp.GetOptCheckListItem{
						Name:        "-c",
						Description: "named names are categories, from for which tarballs to get",
					},
					&cliapp.GetOptCheckListItem{
						Name:        "-g",
						Description: "named names are groups, from for which tarballs to get",
					},
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
					&cliapp.GetOptCheckListItem{
						Name:        "-c",
						Description: "named names are categories, from for which tarballs to get",
					},
					&cliapp.GetOptCheckListItem{
						Name:        "-g",
						Description: "named names are groups, from for which tarballs to get",
					},
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

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	repo, err := tarballrepository.NewRepository(sys)
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

	sys := aipsetup.NewSystem("/")

	repo, err := tarballrepository.NewRepository(sys)
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

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result)
	if res != nil && res.Code != 0 {
		return res
	}

	repo, err := tarballrepository.NewRepository(sys)
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: err.Error(),
		}
	}

	get_by_name_func := func(name string) error {

		name_info, err := pkginfodb.Get(name)
		if err != nil {
			return err
		}

		tarballs, err := repo.ListLocalTarballs(name, true)
		if err != nil {
			return err
		}

		if len(tarballs) == 0 {
			return errors.New("repository does not have tarballs for this package")
		}

		p, err := tarballnameparsers.Get(name_info.TarballFileNameParser)
		if err != nil {
			return err
		}

		c, err := versioncomparators.Get(name_info.TarballVersionComparator)
		if err != nil {
			return err
		}

		version_tool, err := tarballstabilityclassification.Get(name_info.TarballStabilityClassifier)
		if err != nil {
			return err
		}

		err = c.SortStrings(tarballs, p)
		if err != nil {
			return err
		}

		{
			tarballs2 := make([]string, 0)
			for _, i := range tarballs {

				parsed, err := p.Parse(i)
				if err != nil {
					return err
				}

				isstable, err := version_tool.IsStable(parsed)
				if err != nil {
					return err
				}
				if isstable {
					tarballs2 = append(tarballs2, i)
				}
			}
			tarballs = tarballs2
		}

		t := tarballs[len(tarballs)-1]
		fmt.Println(t)
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
