package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"sort"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/repository"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupMBuild() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{

		Name: "mbuild",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Callable: CmdAipsetupMassGetDistroTarballs,
				Name:     "get-all",

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
					STD_NAMES_ARE_CATEGORIES_PRESERVE_NESTING,
					STD_NAMES_ARE_CATEGORIES_IS_PREFIXES,
					STD_NAMES_ARE_GROUPS,
				},

				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Callable: CmdAipsetupMassBuildPerform,
				Name:     "run",

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},

				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   -1,

				Description: "if args count == 0, all tarballs wil be tried.",
			},
		},
	}
}

func CmdAipsetupMassGetDistroTarballs(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	// TODO: fix options

	log := adds.PassData.(*logger.Logger)

	cwd, err := os.Getwd()
	if err != nil {
		return &cliapp.AppResult{Code: 10}
	}

	aipsetup_exec := os.Args[0]

	err = MiscDoSomethingForGroupsCategoriesOrListsAllAtOnce(
		cwd,
		func(
			name string,
			wd string,
		) error {

			log.Info("doing " + wd)

			err := MiscDoSomethingForGroupsCategoriesOrListsAllAtOnce_Sub(
				wd,
				"00.BuildGetAll.log",
				func(log *logger.Logger) error {
					c := exec.Command(aipsetup_exec, "mbuild", "init")
					c.Dir = wd
					err = c.Run()
					if err != nil {
						return err
					}

					c = exec.Command(aipsetup_exec, "mbuild", "get-src", "-g", name)
					c.Dir = wd
					c.Stdout = log.StdoutLbl()
					c.Stderr = log.StderrLbl()
					err = c.Run()
					if err != nil {
						return err
					}

					return nil

				},
			)
			if err != nil {
				return err
			}

			return nil
		},
		func(
			name string,
			wd string,
		) error {

			log.Info("doing " + wd)

			err := MiscDoSomethingForGroupsCategoriesOrListsAllAtOnce_Sub(
				wd,
				"00.BuildGetAll.log",
				func(log *logger.Logger) error {
					c := exec.Command(aipsetup_exec, "mbuild", "init")
					c.Dir = wd
					err = c.Run()
					if err != nil {
						return err
					}

					c = exec.Command(
						aipsetup_exec,
						"mbuild", "get-src",
						"-c", "--cpn", "--cip",
						name+"/",
					)
					c.Dir = wd
					c.Stdout = log.StdoutLbl()
					c.Stderr = log.StderrLbl()
					err = c.Run()
					if err != nil {
						return err
					}

					return nil

				},
			)
			if err != nil {
				return err
			}

			return nil
		},
		log,
	)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	return nil
}

func subCmdAipsetupMassBuildInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
	wd string,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(
		getopt_result,
		log,
	)
	if res != nil && res.Code != 0 {
		return res
	}

	mbuild, err := aipsetup.NewMassBuilder(wd, sys, log)
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

func subCmdAipsetupMassBuildGetSrc(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
	wd string,
) *cliapp.AppResult {
	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	work_on_categories_preserve := getopt_result.DoesHaveNamedRetOptItem(
		STD_NAMES_ARE_CATEGORIES_PRESERVE_NESTING.Name,
	)

	mbuild, err := aipsetup.NewMassBuilder(wd, sys, log)
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

	tarballs_not_found := make([]string, 0)

	get_by_name_func := func(name string) error {

		t, err := repo.DetermineNewestStableTarball(name)
		if err != nil {
			tarballs_not_found = append(tarballs_not_found, name)
			return nil
		}

		log.Info(path.Base(t))

		target_dir := mbuild.GetTarballsPath()

		if work_on_categories_preserve {
			pkginfo, err := pkginfodb.Get(name)
			if err != nil {
				return err
			}
			target_dir = path.Join(target_dir, pkginfo.Category)
		}

		err = repo.CopyTarballToDir(name, t, target_dir)
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

	if len(tarballs_not_found) != 0 {
		sort.Strings(tarballs_not_found)
		log.Error("Couldn't find tarballs for package(s):")
		for _, i := range tarballs_not_found {
			log.Error(fmt.Sprintf("  %s", i))
		}
		return &cliapp.AppResult{
			Code:    12,
			Message: "couldn't get needed tarballs",
		}
	}

	return nil
}

func CmdAipsetupMassBuildInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	return subCmdAipsetupMassBuildInit(getopt_result, adds, ".")
}

func CmdAipsetupMassBuildGetSrc(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	return subCmdAipsetupMassBuildGetSrc(getopt_result, adds, ".")
}

func CmdAipsetupMassBuildPerform(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	type ResStru struct {
		name string
		lst  map[string][]string
	}

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

	s, f, err := mbuild.PerformMassBuilding(getopt_result.Args)
	if err != nil {
		return &cliapp.AppResult{Code: 11, Message: err.Error()}
	}

	t := []*ResStru{
		&ResStru{
			name: "success",
			lst:  s,
		},
		&ResStru{
			name: "fail",
			lst:  f,
		},
	}

	ts := basictypes.NewASPTimeStampFromCurrentTime()

	failed := false

	for i := 0; i != len(t); i++ {

		f, err := os.Create("02." + ts.String() + "." + t[i].name + ".txt")

		if err != nil {
			return &cliapp.AppResult{
				Code:    13,
				Message: "error creating " + t[i].name + " build report",
			}
		}

		keys := make([]string, 0)

		for k, _ := range t[i].lst {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, j := range keys {
			t2 := fmt.Sprintf("arch %s", j)
			fmt.Println(t2)
			f.WriteString(t2 + "\n")
			sort.Strings(t[i].lst[j])
			for _, k := range t[i].lst[j] {
				t2 := fmt.Sprintf("  "+t[i].name+" - %s", k)
				fmt.Println(t2)
				f.WriteString(t2 + "\n")
				if !failed && i == 1 {
					failed = true
				}
			}
		}

		f.Close()

	}

	if failed {
		return &cliapp.AppResult{
			Code:    12,
			Message: "some packages building have failed",
		}
	}

	return nil
}
