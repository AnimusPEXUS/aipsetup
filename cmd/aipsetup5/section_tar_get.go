package main

import (
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballdownloader/providers"
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

	names := []string{"make"}

	sort.Strings(names)

	for _, i := range names {
		_, ok := distropkginfodb.Index[i]
		if !ok {
			return &cliapp.AppResult{
				Code:    10,
				Message: "info for " + i + " not found",
			}
		}
	}

	for _, i := range names {
		info := distropkginfodb.Index[i]

		provider_name := info.TarballProvider

		prov_c, ok := providers.Index[provider_name]
		if !ok {
			return &cliapp.AppResult{
				Code:    11,
				Message: "provider " + provider_name + " not found",
			}
		}

		prov, err := prov_c(info, sys, info.TarballProviderArguments)
		if err != nil {
			return &cliapp.AppResult{
				Code:    12,
				Message: "error instantinating provider " + provider_name + ": " + err.Error(),
			}
		}

		tarballs
	}

}
