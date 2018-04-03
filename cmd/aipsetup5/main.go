package main

import (
	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func main() {

	app := cliapp.AppCmdNode{
		Name:        "aipsetup",
		Description: "LAILALO GNU/Linux system's package manager",
		DevStatus:   "dev",
		License:     "GPLv3+",
		Version:     "5.0",
		SubCmds: []*cliapp.AppCmdNode{
			SectionAipsetupSys(),
			SectionAipsetupSysConfig(),
			SectionAipsetupBuild(),
			SectionAipsetupRepo(),
		},
	}

	global_logger := logger.New()

	cliapp.RunApp(&app, global_logger)

}
