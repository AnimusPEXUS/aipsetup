package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/cliapp"
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
		},
	}

}

func CmdAipsetupBuildInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	// TODO: this function requires optimization

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

	if getopt_result.DoesHaveNamedRetOptItem("-o") {

		target_tarball := getopt_result.Args[0]

		buildinfo, err := aipsetup.DetermineTarballsBuildinfo(target_tarball)
		if err != nil {
			if len(buildinfo) > 1 {
				fmt.Println("error: too many info records match this tarball")
				for i, _ := range buildinfo {
					fmt.Println("  ", i)
				}
			}
			return &cliapp.AppResult{
				Code:    15,
				Message: "error getting buildinfo for tarball: " + err.Error(),
			}
		}

		var buildinfoname string
		var buildinfo0 *basictypes.PackageInfo = nil

		for n, v := range buildinfo {
			buildinfoname = n
			buildinfo0 = v
		}

		bs_ctl, err := aipsetup.NewBuildingSiteCtl(fmt.Sprintf("build/%s", buildinfoname))

		err = bs_ctl.Init()
		if err != nil {
			return &cliapp.AppResult{
				Code:    16,
				Message: "can't init new building site",
			}
		}

		err = bs_ctl.ApplyInitialInfo(buildinfoname, buildinfo0)
		if err != nil {
			return &cliapp.AppResult{
				Code:    16,
				Message: "can't apply initial info to building site",
			}
		}

		err = bs_ctl.ApplyHostArchBuildTarget(host, arch, build, target)
		if err != nil {
			return &cliapp.AppResult{
				Code:    16,
				Message: "can't apply habt info to building site",
			}
		}

		err = bs_ctl.CopyInTarballs(getopt_result.Args)
		if err != nil {
			return &cliapp.AppResult{
				Code:    16,
				Message: "can't copy tarballs into building site",
			}
		}

		err = bs_ctl.ApplyTarballs(getopt_result.Args[0])
		if err != nil {
			return &cliapp.AppResult{
				Code:    16,
				Message: "can't apply tarballs to building site",
			}
		}

	} else {

		counter := 0

		for _, i := range getopt_result.Args {

			buildinfo, err := aipsetup.DetermineTarballsBuildinfo(i)
			if err != nil {
				if len(buildinfo) > 1 {
					fmt.Println("error: too many info records match this tarball")
					for i, _ := range buildinfo {
						fmt.Println("  ", i)
					}
				}
				return &cliapp.AppResult{
					Code:    15,
					Message: "error getting buildinfo for tarball: " + err.Error(),
				}
			}

			var buildinfoname string
			var buildinfo0 *basictypes.PackageInfo = nil

			for n, v := range buildinfo {
				buildinfoname = n
				buildinfo0 = v
			}

			bs_ctl, err := aipsetup.NewBuildingSiteCtl(
				fmt.Sprintf("build/%s-%d", buildinfoname, counter),
			)

			err = bs_ctl.Init()
			if err != nil {
				return &cliapp.AppResult{
					Code:    16,
					Message: "can't init new building site",
				}
			}

			err = bs_ctl.ApplyInitialInfo(buildinfoname, buildinfo0)
			if err != nil {
				return &cliapp.AppResult{
					Code:    16,
					Message: "can't apply initial info to building site",
				}
			}

			err = bs_ctl.ApplyHostArchBuildTarget(host, arch, build, target)
			if err != nil {
				return &cliapp.AppResult{
					Code:    16,
					Message: "can't apply habt info to building site",
				}
			}

			err = bs_ctl.CopyInTarballs([]string{i})
			if err != nil {
				return &cliapp.AppResult{
					Code:    16,
					Message: "can't copy tarball into building site",
				}
			}

			err = bs_ctl.ApplyTarballs(i)
			if err != nil {
				return &cliapp.AppResult{
					Code:    16,
					Message: "can't apply tarballs to building site",
				}
			}

			counter++
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
