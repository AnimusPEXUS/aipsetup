package main

import (
	"github.com/AnimusPEXUS/cliapp"
)

func main() {

	app := cliapp.AppCmdNode{
		Name: "aipsetup",
		SubCmds: []*cliapp.AppCmdNode{
			SectionAipsetupBuild(),
			SectionAipsetupSys(),
			SectionAipsetupInfo(),
		},
	}

	cliapp.RunApp(&app, nil)

}
