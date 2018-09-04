package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func MiscDoSomethingForGroupsCategoriesOrListsAllAtOnce_Sub(
	dir string,
	log_file_name string,
	cb func(log *logger.Logger) error,
) error {

	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return err
	}

	log_f, err := os.Create(path.Join(dir, log_file_name))
	if err != nil {
		return err
	}

	log := logger.New()
	log.AddOutput(log_f)
	log.AddOutput(os.Stdout)

	return cb(log)
}

func MiscDoSomethingForGroupsCategoriesOrListsAllAtOnce(
	wd string,
	cb_for_group func(
		name string,
		wd string,
	) error,
	cb_for_category func(
		name string,
		wd string,
	) error,
	log *logger.Logger,
) error {

	group_errors := make(map[string]error)

	for _, i := range basictypes.HORIZON_GROUPS {
		dir := path.Join(wd, i)

		res := cb_for_group(i, dir)
		if res != nil {
			group_errors[i] = res
			continue
		}
	}

	cat_errors := make(map[string]error)

	for _, i := range basictypes.HORIZON_CATEGORIES {
		cn := "cat_" + i

		dir := path.Join(wd, cn)

		res := cb_for_category(i, dir)
		if res != nil {
			group_errors[i] = res
			continue
		}
	}

	if log != nil {
		if len(group_errors) != 0 {
			log.Error("group errors:")
			for k, v := range group_errors {
				log.Error("   " + k + ": " + v.Error())
			}
		}

		if len(cat_errors) != 0 {
			log.Error("category errors:")
			for k, v := range cat_errors {
				log.Error("   " + k + ": " + v.Error())
			}
		}
	}

	if len(group_errors) != 0 && len(cat_errors) != 0 {
		return errors.New("there was errors")
	}

	return nil
}

func MiscDoSomethingForGroupsCategoriesOrLists(
	sys *aipsetup.System,
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
	action func(name string) error,
) *cliapp.AppResult {

	// TODO: think about removing getopt_result and adds parameters

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
