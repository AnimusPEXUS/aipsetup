package main

import (
	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func SectionAipsetupBootImgSquash() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "boot-img-squash",
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

			&cliapp.AppCmdNode{
				Name:     "reset-root-passwd",
				Callable: CmdAipsetupBootImgResetRootPasswd,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "cleanup-fs",
				Callable: CmdAipsetupBootImgCleanupOSFS,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "cleanup-linux",
				Callable: CmdAipsetupBootImgCleanupLinuxSrc,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "overlay-init",
				Callable: CmdAipsetupBootImgInstallOverlayInit,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "all-above",
				Callable: CmdAipsetupBootImgDoEverythingBeforeSquash,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},

			&cliapp.AppCmdNode{
				Name:     "mksquashfs",
				Callable: CmdAipsetupBootImgSquashOSFiles,
				AvailableOptions: cliapp.GetOptCheckList{
					STD_ROOT_OPTION,
				},
				CheckArgs: true,
				MinArgs:   0,
				MaxArgs:   0,
			},
		},

		Description: "Do 'all-above' action and before mksquashfs" +
			" chroot into osfiles and 'aipsetup5 sys-setup make-good'",
	}

}

func CmdAipsetupBootImgCopyOSFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
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

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
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

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.RemoveUsers()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgResetRootPasswd(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.ResetRootPasswd()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgCleanupOSFS(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.CleanupOSFS()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgCleanupLinuxSrc(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.CleanupLinuxSrc()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgInstallOverlayInit(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.InstallOverlayInit()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgDoEverythingBeforeSquash(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.DoEverythingBeforeSquash()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}

func CmdAipsetupBootImgSquashOSFiles(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	log := adds.PassData.(*logger.Logger)

	bic, err := aipsetup.NewBootImgSquashCtl("/", ".", log)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	err = bic.SquashOSFiles()
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}
