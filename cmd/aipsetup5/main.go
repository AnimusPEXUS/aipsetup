package main

import (
	"os"

	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func main() {

	app := cliapp.AppCmdNode{
		Name:        "aipsetup",
		Description: "Lilith GNU/Linux system's package manager",
		DevStatus:   "dev",
		License:     "GPLv3+",
		Version:     "5.0",
		SubCmds: []*cliapp.AppCmdNode{
			SectionAipsetupSys(),
			SectionAipsetupSysConfig(),
			SectionAipsetupBuild(),
			SectionAipsetupMBuild(),
			SectionAipsetupRepo(),
		},
	}

	global_logger := logger.New()
	global_logger.AddOutput(os.Stdout)

	cliapp.RunApp(&app, global_logger)

}
