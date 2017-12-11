package main

import (
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/providers"
	"github.com/AnimusPEXUS/utils/cliapp"
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

					&cliapp.AppCmdNode{
						Name: "tarballs",
						//Callable: CmdAipsetupTarGetProvidersTarballs,
					},
				},
			},

			// &cliapp.AppCmdNode{
			//
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
				MinArgs:   0,
				MaxArgs:   1,
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

	name := getopt_result.Args[0]

	err = repo.PerformPackageTarballsUpdate(name)
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{Code: 0}
}
