package main

import (
	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupBootImgInitRd() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "boot-img-initrd",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:             "copy",
				Callable:         CmdAipsetupBootImgInitRdCopyOSFiles,
				ShortDescription: "copy os files to ./osfiles",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:             "mk-links",
				Callable:         CmdAipsetupBootImgInitRdCopyOSFiles,
				ShortDescription: "copy os files to ./osfiles",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:             "write-init",
				Callable:         CmdAipsetupBootImgInitRdCopyOSFiles,
				ShortDescription: "copy os files to ./osfiles",
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "all-above",
				Callable: CmdAipsetupBootImgInitRdDoEverythingBeforeCompress,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "compress",
				Callable: CmdAipsetupBootImgInitRdCompressOSFiles,
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

func CmdAipsetupBootImgInitRdCopyOSFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgInitRdCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.CopyOSFiles()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgInitRdMakeSymlinks(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgInitRdCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.MakeSymlinks()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgInitRdWriteInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgInitRdCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.WriteInit()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgInitRdDoEverythingBeforeCompress(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgInitRdCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.DoEverythingBeforePack()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgInitRdCompressOSFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgInitRdCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.CompressOSFiles()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}
