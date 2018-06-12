package main

import (
	"fmt"
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

	work_on_categories := getopt_result.DoesHaveNamedRetOptItem(
		STD_NAMES_ARE_CATEGORIES.Name,
	)

	//	work_on_categories_preserve := getopt_result.DoesHaveNamedRetOptItem(
	//		STD_NAMES_ARE_CATEGORIES_PRESERVE_NESTING.Name,
	//	)

	work_on_categories_prefixes := getopt_result.DoesHaveNamedRetOptItem(
		STD_NAMES_ARE_CATEGORIES_IS_PREFIXES.Name,
	)

	work_on_groups := getopt_result.DoesHaveNamedRetOptItem(
		STD_NAMES_ARE_GROUPS.Name,
	)

	//					STD_NAMES_ARE_CATEGORIES,
	//					STD_NAMES_ARE_CATEGORIES_PRESERVE_NESTING,
	//					STD_NAMES_ARE_CATEGORIES_IS_PREFIXES,
	//					STD_NAMES_ARE_GROUPS,

	if !work_on_categories && work_on_categories_prefixes {
		return &cliapp.AppResult{
			Code: 13,
			Message: fmt.Sprintf(
				"%s option can be used only with %s option",
				STD_NAMES_ARE_CATEGORIES_IS_PREFIXES.Name,
				STD_NAMES_ARE_CATEGORIES.Name,
			),
		}
	}

	if work_on_groups && work_on_categories {
		return &cliapp.AppResult{
			Code:    12,
			Message: "mutualy exclusive options given",
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

		//		var pkgs []string

		pkgs, err := pkginfodb.ListPackagesByCategories(
			getopt_result.Args,
			work_on_categories_prefixes,
		)
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
	return nil
}
