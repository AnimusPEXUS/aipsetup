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
				Name:     "make",
				Callable: CmdAipsetupBootImgMake,
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

func CmdAipsetupBootImgMake(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.DoEverything()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}
