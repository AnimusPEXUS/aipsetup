package main

import (
	"os"

	"github.com/AnimusPEXUS/utils/cliapp"
	"github.com/AnimusPEXUS/utils/logger"
)

func main() {

	app := cliapp.AppCmdNode{
		Name:        "aipsetup",
		Description: "Horizon GNU/Linux system's package manager",
		DevStatus:   "development",
		License:     "GPLv3+",
		Version:     "5.1",
		Developers:  []string{"AnimusPEXUS"},
		URIs:        []string{"https://github.com/AnimusPEXUS/aipsetup"},
		SubCmds: []*cliapp.AppCmdNode{
			SectionAipsetupSys(),
			SectionAipsetupSysSetup(),
			SectionAipsetupSysDocBook(),
			SectionAipsetupBuild(),
			SectionAipsetupMBuild(),
			SectionAipsetupBootImg(),
			SectionAipsetupBootImgSquash(),
			SectionAipsetupBootImgInitRd(),
			SectionAipsetupRepo(),
			SectionAipsetupConfig(),
		},
	}

	global_logger := logger.New()
	global_logger.AddOutput(os.Stdout)

	cliapp.RunApp(&app, global_logger)

}
