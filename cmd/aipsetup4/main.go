package main

import (
	"github.com/AnimusPEXUS/cliapp"
)

func main() {

	app := cliapp.AppCmdNode{
		Name: "aipsetup",
		SubCmds: []*cliapp.AppCmdNode{
			SectionAipsetupSys(),
		},
	}

	cliapp.RunApp(&app, nil)

}
