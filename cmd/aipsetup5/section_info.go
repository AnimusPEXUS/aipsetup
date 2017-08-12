package main

import (
	"fmt"

	"github.com/AnimusPEXUS/cliapp"
)

func SectionAipsetupInfo() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "info",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:     "serve",
				Callable: CmdAipsetupInfoServe,
			},

			&cliapp.AppCmdNode{
				Name:     "serve-cfg",
				Callable: CmdAipsetupInfoServePrintExampleConfig,
			},
		},
	}

}

func CmdAipsetupInfoServe(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {

	srv, err := NewInfoServer()
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}
	fmt.Println("Host", srv.host)
	fmt.Println("Port", srv.port)
	fmt.Println("Prefix", srv.prefix)
	fmt.Println("Infodir", srv.infodir)
	srv.Run()

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupInfoServePrintExampleConfig(
	getopt_result *cliapp.GetOptResult,
	adds *cliapp.AdditionalInfo,
) *cliapp.AppResult {
	fmt.Println("aipsetup5.info_server.ini")
	fmt.Println(string(INFO_SERVER_CONFIG))
	return &cliapp.AppResult{Code: 0, DoNotPrintResult: true}
}
