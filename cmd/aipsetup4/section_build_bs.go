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
			},
		},
	}

}

func CmdAipsetupBuildBSInit(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {

	target_dir := "."

	switch len(getopt_result.Args) {
	case 0:
	case 1:
		target_dir = getopt_result.Args[0]
	default:
		return &cliapp.AppResult{Code: 10, Message: "too many arguments"}
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
