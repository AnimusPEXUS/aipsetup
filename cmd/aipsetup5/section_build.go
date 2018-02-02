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
	"github.com/AnimusPEXUS/aipsetup/tarballrepository"
	"github.com/AnimusPEXUS/aipsetup/versionstabilityclassifiers"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers/types"
	"github.com/AnimusPEXUS/utils/version"
)

func SectionAipsetupBuild() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{

		Name: "build",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:      "get-src",
				Callable:  CmdAipsetupBuildGetSrc,
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

				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
					STD_BUILDER_HOST_OPTION,
					STD_BUILDER_ARCH_OPTION,
					STD_BUILDER_BUILD_OPTION,
					STD_BUILDER_TARGET_OPTION,
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

			&cliapp.AppCmdNode{
				Name: "full",

				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   -1,

				Callable: CmdAipsetupBuildFull,
			},

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
	main_tarball string,
	addittional_tarballs []string,
	host, arch, build, target string,
) error {

	target_tarball := main_tarball

	buildinfo, err := pkginfodb.DetermineTarballsBuildInfo(target_tarball)
	if err != nil {
		return errors.New("error searching matching info record: " + err.Error())
	}

	var buildinfoname string
	var buildinfo0 *basictypes.PackageInfo = nil

	for n, v := range buildinfo {
		buildinfoname = n
		buildinfo0 = v
		break
	}

	var version string
	{
		var parser types.TarballNameParserI

		{
			parser_c, ok :=
				tarballnameparsers.Index[buildinfo0.TarballFileNameParser]
			if !ok {
				return errors.New(
					"can't find tarball name parser pointed by info file: " + err.Error(),
				)
			}

			parser = parser_c()
		}

		err := tarballname.IsPossibleTarballNameErr(target_tarball)
		if err != nil {
			return err
		}

		parsed, err := parser.Parse(target_tarball)
		if err != nil {
			return err
		}

		version = parsed.Version.Str

	}

	bs_ctl, err := aipsetup.NewBuildingSiteCtl(
		fmt.Sprintf("build/%s-%s", buildinfoname, version),
	)

	err = bs_ctl.Init()
	if err != nil {
		return errors.New("can't init new building site: " + err.Error())
	}

	err = bs_ctl.ApplyInitialInfo(buildinfoname, buildinfo0)
	if err != nil {
		return errors.New("can't apply initial info to building site: " + err.Error())
	}

	err = bs_ctl.ApplyHostArchBuildTarget(host, arch, build, target)
	if err != nil {
		return errors.New("can't apply habt info to building site: " + err.Error())
	}

	all_tarballs := make([]string, 0)
	all_tarballs = append(all_tarballs, main_tarball)
	all_tarballs = append(all_tarballs, addittional_tarballs...)

	err = bs_ctl.CopyInTarballs(all_tarballs)
	if err != nil {
		return errors.New("can't copy tarballs into building site: " + err.Error())
	}

	err = bs_ctl.ApplyTarballs(all_tarballs[0])
	if err != nil {
		return errors.New("can't apply tarballs to building site: " + err.Error())
	}

	return nil
}

func CmdAipsetupBuildInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	sys := aipsetup.NewSystem(
		getopt_result.GetLastNamedRetOptItem("--root").Value,
	)

	host, arch, build, target := StdRoutineHostArchBuildTarget(getopt_result, sys)

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
			getopt_result.Args[0],
			getopt_result.Args[1:],
			host, arch, build, target,
		)
	} else {

		for _, i := range getopt_result.Args {

			err = CmdAipsetupBuildInitSub01(
				i,
				[]string{},
				host, arch, build, target,
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

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupBuildListActions(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	bs_ctl, err := aipsetup.NewBuildingSiteCtl(".")
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
	bs_ctl, err := aipsetup.NewBuildingSiteCtl(".")
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

	return &cliapp.AppResult{Code: 0}
}

// func CmdAipsetupBuildPack(
// 	getopt_result *cliapp.GetOptResult,
// 	adds *cliapp.AdditionalInfo,
// ) *cliapp.AppResult {
// 	bs_ctl, err := aipsetup.NewBuildingSiteCtl(".")
// 	if err != nil {
// 		return &cliapp.AppResult{
// 			Code:    10,
// 			Message: "Can't create building site object",
// 		}
// 	}
//
// 	log, err := bs_ctl.CreateLogger("packaging", true)
// 	if err != nil {
// 		return &cliapp.AppResult{
// 			Code:    11,
// 			Message: "Can't create logger",
// 		}
// 	}
//
// 	err = bs_ctl.Packager().Run(log)
// 	if err != nil {
// 		log.Error(err.Error())
// 		return &cliapp.AppResult{
// 			Code:    12,
// 			Message: "Packaging error",
// 		}
// 	}
//
// 	log.Info("Finished")
// 	log.Close()
//
// 	return new(cliapp.AppResult)
// }

func CmdAipsetupBuildGetSrc(
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

	work_on_groups := getopt_result.DoesHaveNamedRetOptItem("-g")
	work_on_categories := getopt_result.DoesHaveNamedRetOptItem("-c")

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

		version_tool, err := versionstabilityclassifiers.Get(name_info.TarballVersionTool)
		if err != nil {
			return err
		}

		err = version.SortByVersion(tarballs, p)
		if err != nil {
			return err
		}

		{
			tarballs2 := make([]string, 0)
			for _, i := range tarballs {

				isstable, err := version_tool.IsStable(p, i)
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

	if work_on_groups && work_on_categories {
		return &cliapp.AppResult{
			Code:    12,
			Message: "mutual exclusive options given",
		}
	} else if !work_on_groups && !work_on_categories {
		for _, i := range getopt_result.Args {
			err := get_by_name_func(i)

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

		for _, i := range pkgs {
			err := get_by_name_func(i)

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

		for _, i := range pkgs {
			err := get_by_name_func(i)

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

	return &cliapp.AppResult{Code: 0}
}
