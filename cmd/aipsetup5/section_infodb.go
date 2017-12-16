package main

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
	"github.com/AnimusPEXUS/aipsetup/infoeditor"
	"github.com/AnimusPEXUS/utils/cliapp"
)

func SectionAipsetupInfoDB() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "info-db",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:      "write",
				Callable:  CmdAipsetupInfoWrite,
				CheckArgs: true,
				MinArgs:   1,
				MaxArgs:   1,
			},

			&cliapp.AppCmdNode{
				Name:             "code",
				ShortDescription: "Generate new distropkginfodb editing it with InfoEditor.go",
				Callable:         CmdAipsetupInfoCode,
				CheckArgs:        true,
				MinArgs:          0,
				MaxArgs:          0,
			},
		},
	}

}

func CmdAipsetupInfoWrite(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	filebasename := getopt_result.Args[0]

	f, err := os.Create(filebasename)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: "can't create named file"}
	}

	defer f.Close()

	j_result, err := json.Marshal(distropkginfodb.Index)
	if err != nil {
		return &cliapp.AppResult{
			Code:    11,
			Message: "can't convert internal DB into JSON",
		}
	}

	var out bytes.Buffer

	json.Indent(&out, j_result, "", "\t")

	out.WriteTo(f)

	return &cliapp.AppResult{
		Code:    0,
		Message: "looks like no errors",
	}

}

func CmdAipsetupInfoCode(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	ie := &infoeditor.InfoEditor{}
	err := ie.Run()
	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: err.Error(),
		}
	}
	return &cliapp.AppResult{
		Code: 0,
	}
}
