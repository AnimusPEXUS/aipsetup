package main

import (
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupSysSetup() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "sys-setup",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:     "gen-locale",
				Callable: CmdAipsetupSysSetupGenLocale,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name: "regen-userdb",
				Description: "new (based on existing) userdb will be rendered and" +
					" placed inside /root/tmp/newuserdb. \n" +
					" You'll have to review it and move files to /etc manually",
				Callable: CmdAipsetupSysSetupResetUsers,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name: "regen-daemon-homes",
				Description: "creates, chowns and chmods home dirs of users 1-999" +
					" inside /daemons",
				Callable: CmdAipsetupSysSetupResetDaemonHomes,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "reset-permissions",
				Callable: CmdAipsetupSysSetupResetPermissions,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "reinstall-etc",
				Callable: CmdAipsetupSysSetupReinstallEtc,
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

func CmdAipsetupSysSetupGenLocale(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	err := sys.GenLocale()
	if err != nil {
		return &cliapp.AppResult{10, "error generating locales", false}
	}

	return nil
}

func CmdAipsetupSysSetupResetUsers(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	user_ctl := sys.GetUserCtl()

	err := user_ctl.RecreateUserDB()
	if err != nil {
		return &cliapp.AppResult{
			Code:             10,
			Message:          err.Error(),
			DoNotPrintResult: false,
		}
	}

	return nil
}

func CmdAipsetupSysSetupResetDaemonHomes(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	user_ctl := sys.GetUserCtl()

	err := user_ctl.RecreateDaemonHomes()
	if err != nil {
		return &cliapp.AppResult{
			Code:             10,
			Message:          err.Error(),
			DoNotPrintResult: false,
		}
	}

	return nil
}

func CmdAipsetupSysSetupResetPermissions(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	user_ctl := sys.GetUserCtl()

	err := user_ctl.ResetSystemPermissions(log)
	if err != nil {
		return &cliapp.AppResult{
			Code:             10,
			Message:          err.Error(),
			DoNotPrintResult: false,
		}
	}

	return nil
}

func CmdAipsetupSysSetupReinstallEtc(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	_, sys, res := StdRoutineGetRootOptionAndSystemObject(getopt_result, log)
	if res != nil && res.Code != 0 {
		return res
	}

	err := sys.GetSystemUpdates().InstallEtc(log)
	if err != nil {
		return &cliapp.AppResult{
			Code:             10,
			Message:          err.Error(),
			DoNotPrintResult: false,
		}
	}

	return nil
}
