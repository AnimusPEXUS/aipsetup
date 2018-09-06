package main

import (
	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupBootImg() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "boot-img",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:             "copy",
				Callable:         CmdAipsetupBootImgCopyOSFiles,
				ShortDescription: "copy os files to ./osfiles",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "install-aipsetup",
				Callable: CmdAipsetupBootImgInstallAipSetup,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "rm-users",
				Callable: CmdAipsetupBootImgRemoveUsers,
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

func CmdAipsetupBootImgCopyOSFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.CopyOSFiles()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgInstallAipSetup(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.InstallAipSetup()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgRemoveUsers(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.RemoveUsers()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}
