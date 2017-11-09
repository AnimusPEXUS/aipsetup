package main

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
	"github.com/AnimusPEXUS/cliapp"
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
