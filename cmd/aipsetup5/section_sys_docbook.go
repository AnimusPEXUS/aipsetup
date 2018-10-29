package main

import (
	"fmt"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupSysDocBook() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "sys-docbook",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:      "instruction",
				Callable:  CmdAipsetupSysDocBookInstruction,
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "install",
				Callable: CmdAipsetupSysDocBookInstall,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},
		},
	}
}

func CmdAipsetupSysDocBookInstruction(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	fmt.Println(aipsetup.DOCBOOK_INSTRUCTION)
	return nil
}

func CmdAipsetupSysDocBookInstall(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	host, err := sys.Host()
	if err != nil {
		return &cliapp.AppResult{
			Code:    15,
			Message: "host triplet determination problem",
		}
	}

	dbs := &aipsetup.InstallDockBookSettings{}
	dbs.SetDefaults(host)
	dbs.BaseDir = sys.Root()
	dbs.Log = log

	db_ctl := aipsetup.NewDocBookCtl(dbs)

	err = db_ctl.InstallDockBook()
	if err != nil {
		return &cliapp.AppResult{
			Code:    16,
			Message: err.Error(),
		}
	}

	return nil
}
