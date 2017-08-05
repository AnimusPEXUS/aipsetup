package main

import (
	"fmt"
	//"github.com/AnimusPEXUS/aipsetup"
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
				Name:     "print-server-example-config",
				Callable: CmdAipsetupInfoServePrintExampleConfig,
			},
		},
	}

}

func CmdAipsetupInfoServe(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {

	var (
		target_dir string = "."
		err        error
	)

	switch len(getopt_result.Args) {
	case 0:
	case 1:
		target_dir = getopt_result.Args[0]
	default:
		return &cliapp.AppResult{Code: 10, Message: "too many arguments"}
	}

	srv, err := NewInfoServer(target_dir)
	if err != nil {
		return &cliapp.AppResult{Code: 10, Message: err.Error()}
	}
	srv.Run()

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupInfoServePrintExampleConfig(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {
	t := NewInfoServerConfig()
	fmt.Println(t.YAMLString())
	return &cliapp.AppResult{Code: 0, DoNotPrintResult: true}
}
