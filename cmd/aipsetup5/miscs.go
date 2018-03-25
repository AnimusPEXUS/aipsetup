package main

import (
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/cliapp"
)

func MiscDoSomethingForGroupsCategoriesOrLists(
	sys *aipsetup.System,
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
	action func(name string) error,
) *cliapp.AppResult {

	work_on_groups := getopt_result.DoesHaveNamedRetOptItem("-g")
	work_on_categories := getopt_result.DoesHaveNamedRetOptItem("-c")

	if work_on_groups && work_on_categories {
		return &cliapp.AppResult{
			Code:    12,
			Message: "mutual exclusive options given",
		}
	} else if !work_on_groups && !work_on_categories {
		for _, i := range getopt_result.Args {
			err := action(i)

			if err != nil {
				return &cliapp.AppResult{
					Code:    10,
					Message: err.Error(),
				}
			}
		}
	} else if work_on_groups {
		pkgs, err := pkginfodb.ListPackagesByGroups(getopt_result.Args)
		if err != nil {
			return &cliapp.AppResult{
				Code:    11,
				Message: err.Error(),
			}
		}

		sort.Strings(pkgs)

		for _, i := range pkgs {
			err := action(i)

			if err != nil {
				return &cliapp.AppResult{
					Code:    10,
					Message: err.Error(),
				}
			}
		}
	} else if work_on_categories {
		pkgs, err := pkginfodb.ListPackagesByCategories(getopt_result.Args)
		if err != nil {
			return &cliapp.AppResult{
				Code:    11,
				Message: err.Error(),
			}
		}

		sort.Strings(pkgs)

		for _, i := range pkgs {
			err := action(i)

			if err != nil {
				return &cliapp.AppResult{
					Code:    10,
					Message: err.Error(),
				}
			}
		}
	} else {
		panic("programming error")
	}
	return &cliapp.AppResult{}
}
