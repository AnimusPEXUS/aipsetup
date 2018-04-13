package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers/types"
)

func SectionAipsetupBuild() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{

		Name: "build",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_OPTION_BUILD_CURRENT_HOST,
					STD_OPTION_BUILD_FOR_HOST,
					STD_OPTION_BUILD_FOR_HOSTARCH,
					STD_OPTION_BUILD_CROSSBUILDER,
					STD_OPTION_BUILD_CROSSBUILDING,
					&cliapp.GetOptCheckListItem{
						Name: "-o",
						Description: "" +
							"when more than one tarball defined, this option " +
							"makes init command treat named tarballs as for same single " +
							"building site, also the first defined tarball will be the main " +
							"one",
					},
				},
				Name:      "init",
				Callable:  CmdAipsetupBuildInit,
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   -1,
			},

			&cliapp.AppCmdNode{
				Name: "info",

				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,

				Callable: CmdAipsetupBuildPrintInfo,
			},

			&cliapp.AppCmdNode{
				Name: "list",

				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,

				Callable: CmdAipsetupBuildListActions,
			},

			&cliapp.AppCmdNode{
				Name: "run",

				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   1,

				Callable: CmdAipsetupBuildRun,
			},

			// &cliapp.AppCmdNode{
			// 	Name: "full",
			//
			// 	CheckArgs: true,
			// 	MinArgs:   0,
			// 	MaxArgs:   -1,
			//
			// 	Callable: CmdAipsetupBuildFull,
			// },

			// &cliapp.AppCmdNode{
			// 	Name: "pack",
			//
			// 	CheckArgs: true,
			// 	MinArgs:   0,
			// 	MaxArgs:   0,
			//
			// 	Callable: CmdAipsetupBuildPack,
			// },
		},
	}

}

func CmdAipsetupBuildInitSub01(
	sys *aipsetup.System,
	main_tarball string,
	addittional_tarballs []string,
	host, hostarch string,
	log *logger.Logger,
) error {

	target_tarball := main_tarball

	buildinfoname, buildinfo0, err :=
		pkginfodb.DetermineTarballPackageInfoSingle(target_tarball)
	if err != nil {
		return err
	}

	var version string
	{
		var parser types.TarballNameParserI

		parser, err := tarballnameparsers.Get(buildinfo0.TarballFileNameParser)
		if err != nil {
			return err
		}

		err = tarballname.IsPossibleTarballNameErr(target_tarball)
		if err != nil {
			return err
		}

		parsed, err := parser.Parse(target_tarball)
		if err != nil {
			return err
		}

		version = parsed.Version.Str

	}

	new_timestamp := basictypes.NewASPTimeStampFromCurrentTime()

	bs_ctl, err := aipsetup.NewBuildingSiteCtl(
		fmt.Sprintf("build/%s-%s-%s", buildinfoname, version, new_timestamp.String()),
		sys,
		log,
	)

	err = bs_ctl.Init()
	if err != nil {
		return errors.New("can't init new building site: " + err.Error())
	}

	new_bs_info := &basictypes.BuildingSiteInfo{
		PackageName:      buildinfoname,
		Host:             host,
		HostArch:         hostarch,
		PackageVersion:   version,
		PackageTimeStamp: new_timestamp.String(),
	}
	new_bs_info.SetInfoLailalo50()

	err = bs_ctl.WriteInfo(new_bs_info)
	if err != nil {
		return err
	}

	// TODO: here was source application. but should be PrepareToRun()

	return nil
}

func CmdAipsetupBuildInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	sys := aipsetup.NewSystem(
		getopt_result.GetLastNamedRetOptItem("--root").Value,
		log,
	)

	host, hostarch, res := StdRoutineGetBuildingHostHostArch(getopt_result, sys)
	if res != nil && res.Code != 0 {
		return res
	}

	if len(getopt_result.Args) == 0 {
		return &cliapp.AppResult{Code: 13, Message: "no tarballs defined"}
	}

	all_files_found := true
	for _, i := range getopt_result.Args {
		if _, err := os.Stat(i); err != nil {
			all_files_found = false
			fmt.Println("file", i, "not found")
		}
	}

	if !all_files_found {
		return &cliapp.AppResult{
			Code:    14,
			Message: "error checking input files existance",
		}
	}

	var err error

	if getopt_result.DoesHaveNamedRetOptItem("-o") {
		err = CmdAipsetupBuildInitSub01(
			sys,
			getopt_result.Args[0],
			getopt_result.Args[1:],
			host, hostarch,
			log,
		)
	} else {

		for _, i := range getopt_result.Args {

			err = CmdAipsetupBuildInitSub01(
				sys,
				i,
				[]string{},
				host, hostarch,
				log,
			)

			if err != nil {
				break
			}

		}

	}

	if err != nil {
		return &cliapp.AppResult{
			Code:    15,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupBuildListActions(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	bs_ctl, err := aipsetup.NewBuildingSiteCtl(".", sys, log)
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: "Can't create building site object",
		}
	}

	actions, err := bs_ctl.ListActions()
	if err != nil {
		return &cliapp.AppResult{
			Code:    13,
			Message: err.Error(),
		}
	}

	for _, i := range actions {
		fmt.Println(i)
	}

	return new(cliapp.AppResult)
}

func CmdAipsetupBuildRun(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	bs_ctl, err := aipsetup.NewBuildingSiteCtl(".", sys, log)
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: "Can't create building site object",
		}
	}

	actions, err := bs_ctl.ListActions()
	if err != nil {
		return &cliapp.AppResult{
			Code:    13,
			Message: err.Error(),
		}
	}

	// copy(targets, actions)

	action := actions[0] + "+"

	if len(getopt_result.Args) != 0 {
		action = getopt_result.Args[0]
	}

	plus := false
	if strings.HasSuffix(action, "+") {
		action = action[:len(action)-1]
		plus = true
	}

	{
		actions2 := make([]string, 0)
		found := false

		for ii, i := range actions {
			if i == action {
				found = true
				actions2 = actions[ii:]
				if !plus {
					actions2 = actions2[:1]
				}
				break
			}
		}
		if !found {
			return &cliapp.AppResult{
				Code:    15,
				Message: "asked action not found",
			}
		}
		actions = actions2
	}

	err = bs_ctl.Run(actions)
	if err != nil {
		return &cliapp.AppResult{
			Code:    14,
			Message: err.Error(),
		}
	}

	return new(cliapp.AppResult)
}

func CmdAipsetupBuildFull(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	app := adds.Rootnode

	for _, i := range getopt_result.Args {

		s, err := os.Stat(i)
		if err == nil && s.IsDir() {
			continue
		}

		appres := cliapp.RunCmd(
			adds.Arg0,
			[]string{"build", "init", i},
			app,
			adds.PassData,
		)
		if appres.Code != 0 {
			return &cliapp.AppResult{
				Code:    10,
				Message: "error initiating building site for " + i,
			}
		}

	}

	build_dir_files, err := ioutil.ReadDir("build")
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: "error listing 'build' dir",
		}
	}

	bsites := make([]string, 0)

	for _, i := range build_dir_files {
		if i.IsDir() {
			bsites = append(bsites, i.Name())
		}
	}

	sort.Strings(bsites)

	init_wd, err := os.Getwd()
	if err != nil {
		return &cliapp.AppResult{
			Code:    12,
			Message: "error treating current directory",
		}
	}

	was_build_errors := false

	for _, i := range bsites {
		joined := path.Join(init_wd, "build", i)

		err := os.Chdir(joined)
		if err != nil {
			return &cliapp.AppResult{
				Code:    14,
				Message: "error cd into " + joined,
			}
		}

		appres := cliapp.RunCmd(
			adds.Arg0,
			[]string{"build", "run"},
			app,
			adds.PassData,
		)
		if appres.Code != 0 {
			was_build_errors = true
		}
	}

	err = os.Chdir(init_wd)
	if err != nil {
		return &cliapp.AppResult{
			Code:    16,
			Message: "error cd into " + init_wd,
		}
	}

	if was_build_errors {
		return &cliapp.AppResult{
			Code:    17,
			Message: "some packages failed to build",
		}
	}

	return &cliapp.AppResult{}
}

func CmdAipsetupBuildPrintInfo(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	bs_ctl, err := aipsetup.NewBuildingSiteCtl(".", sys, log)
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: "Can't create building site object",
		}
	}

	err = bs_ctl.PrintCalculations()
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{}
}
