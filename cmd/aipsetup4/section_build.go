package main

import (
	"github.com/AnimusPEXUS/cliapp"
)

func SectionAipsetupBuild() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "build",
		SubCmds: []*cliapp.AppCmdNode{
			SectionAipsetupBuildBS(),
		},
	}

}
