package main

import (
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/providers"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/tarballname"
)

func SectionAipsetupTarGet() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{

		Name: "tar-get",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name: "providers",

				SubCmds: []*cliapp.AppCmdNode{

					&cliapp.AppCmdNode{
						Name:     "list",
						Callable: CmdAipsetupTarGetProvidersList,
					},

					// &cliapp.AppCmdNode{
					// 	Name:     "tarballs",
					// 	Callable: CmdAipsetupTarGetProvidersTarballs,
					// },
				},
			},

			// &cliapp.AppCmdNode{
			// 	Name:      "init",
			// 	Callable:  CmdAipsetupTarGetInit,
			// 	CheckArgs: true,
			// 	MinArgs:   0,
			// 	MaxArgs:   0,
			// },

			&cliapp.AppCmdNode{
				Name:      "for",
				Callable:  CmdAipsetupTarGetFor,
				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   -1,

				AvailableOptions: cliapp.GetOptCheckList{
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
				Name:      "move-in",
				Callable:  CmdAipsetupTarGetMoveIn,
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   2,
			},
		},
	}

}

func CmdAipsetupTarGetProvidersList(
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

func CmdAipsetupTarGetFor(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	// TODO: add root parameter to command
	sys := aipsetup.NewSystem("/")

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

		err := repo.PerformPackageTarballsUpdate(name)
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

	return &cliapp.AppResult{Code: 0}

}

func CmdAipsetupTarGetMoveIn(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	arg := "."
	arg2 := "../rejected"

	if len(getopt_result.Args) > 0 {
		arg = getopt_result.Args[0]
	}

	if len(getopt_result.Args) > 1 {
		arg2 = getopt_result.Args[1]
	}

	err := os.MkdirAll(arg2, 0700)
	if err != nil {
		return &cliapp.AppResult{Code: 14, Message: err.Error()}
	}

	s, err := os.Stat(arg)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	// TODO: add root parameter to command
	sys := aipsetup.NewSystem("/")

	repo, err := tarballrepository.NewRepository(sys)
	if err != nil {
		return &cliapp.AppResult{Code: 9, Message: err.Error()}
	}

	if !s.IsDir() {
		err = repo.MoveInTarball(arg)
		if err != nil {
			return &cliapp.AppResult{Code: 11, Message: err.Error()}
		}
	} else {
		filetools.Walk(
			arg,
			func(
				dir string,
				dirs []os.FileInfo,
				files []os.FileInfo,
			) error {
				for _, i := range files {
					if tarballname.IsPossibleTarballName(i.Name()) {
						fp := path.Join(dir, i.Name())
						fmt.Printf("trying to movein %s:", fp)
						res := repo.MoveInTarball(fp)
						if res == nil {
							fmt.Println("ok")
						} else {
							fmt.Println("error:", res.Error())
							err := os.Rename(fp, path.Join(arg2, i.Name()))
							if err != nil {
								return err
							}
						}
					}
				}
				return nil
			},
		)
	}
	return nil
}
