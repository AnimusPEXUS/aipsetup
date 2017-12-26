package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/cliapp"
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
				Name: "pack",

				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,

				Callable: CmdAipsetupBuildPack,
			},
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

		parsed, err := parser.ParseName(target_tarball)
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

	arg0 := actions[0] + "+"
	if len(getopt_result.Args) != 0 {
		arg0 = getopt_result.Args[0]
	}

	plus := false
	if strings.HasSuffix(arg0, "+") {
		arg0 = arg0[:len(arg0)-1]
	}

	append_act := false
	actions2 := make([]string, 0)
	for _, i := range actions {
		if i == arg0 {
			append_act = true
		}
		if append_act {
			actions2 = append(actions2, i)
			if !plus {
				break
			}
		}
	}

	err = bs_ctl.Run(actions2)
	if err != nil {
		return &cliapp.AppResult{
			Code:    14,
			Message: err.Error(),
		}
	}

	return new(cliapp.AppResult)
}

func CmdAipsetupBuildPack(
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

	log, err := bs_ctl.CreateLogger("packaging", true)
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: "Can't create logger",
		}
	}

	err = bs_ctl.Packager().Run(log)
	if err != nil {
		log.Error(err.Error())
		return &cliapp.AppResult{
			Code:    12,
			Message: "Packaging error",
		}
	}

	log.Info("Finished")
	log.Close()

	return new(cliapp.AppResult)
}
