package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/cliapp"
)

func SectionAipsetupSysConfig() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "sys-config",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:     "write-example",
				Callable: CmdAipsetupSysConfigWriteExample,
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

func CmdAipsetupSysConfigWriteExample(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	root := "/"

	if root_opt, ok := StdRoutineGetRootOption(getopt_result); ok {
		root = root_opt
	}

	root, err := filepath.Abs(root)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	file_dir := path.Join(root, "/etc")
	file_path :=
		path.Join(file_dir, basictypes.AIPSETUP_SYSTEM_CONFIG_FILENAME+".example")

	err = os.MkdirAll(file_dir, 0755)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	f, err := os.Create(file_path)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}

	defer f.Close()

	_, err = f.Write(aipsetup.DEFAULT_AIPSETUP_SYSTEM_CONFIG)
	if err != nil {
		return &cliapp.AppResult{Code: 20, Message: err.Error()}
	}

	return nil
}
