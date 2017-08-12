package main

import (
	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/cliapp"
)

func SectionAipsetupBuildBS() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "bs",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:     "init",
				Callable: CmdAipsetupBuildBSInit,
				MinArgs:  0,
				MaxArgs:  1,
			},
		},
	}

}

func CmdAipsetupBuildBSInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	target_dir := "."

	if len(getopt_result.Args) != 0 {
		target_dir = getopt_result.Args[0]
	}

	bs_ctl, err := aipsetup.NewBuildingSiteCtl(target_dir)

	if err != nil {
		return &cliapp.AppResult{Code: 12, Message: err.Error()}
	}

	res := bs_ctl.Init()

	if res != nil {
		return &cliapp.AppResult{Code: 11, Message: res.Error()}
	}

	return &cliapp.AppResult{Code: 0}
}
