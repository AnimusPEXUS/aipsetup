package main

import (
	"fmt"
	"strings"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupBuild() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{

		Name: "build",
		SubCmds: []*cliapp.AppCmdNode{

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
		},
	}

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

	if actions == nil {
		return &cliapp.AppResult{
			Code:    20,
			Message: "builder returned nil in place of actions",
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

	return nil
}
