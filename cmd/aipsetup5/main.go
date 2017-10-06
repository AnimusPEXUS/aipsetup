package main

import "github.com/AnimusPEXUS/cliapp"

func main() {

	app := cliapp.AppCmdNode{
		Name:        "aipsetup",
		Description: "LAILALO GNU/Linux system's package manager",
		DevStatus:   "pre-alpha",
		License:     "GPLv3+",
		Version:     "5.0",
		SubCmds: []*cliapp.AppCmdNode{
			SectionAipsetupSys(),
			SectionAipsetupSysConfig(),
			SectionAipsetupBuild(),
			SectionAipsetupInfo(),
		},
	}

	cliapp.RunApp(&app, nil)

}
