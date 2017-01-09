package main

import (
	"fmt"
	"sort"

	"github.com/AnimusPEXUS/cliapp"

	"github.com/AnimusPEXUS/aipsetup"
)

func SectionAipsetupSys() *cliapp.AppCmdNode {

	return &cliapp.AppCmdNode{
		Name: "sys",
		SubCmds: []*cliapp.AppCmdNode{

			&cliapp.AppCmdNode{
				Name:     "asps",
				Callable: CmdAipsetupSysAllAsps,
			},

			&cliapp.AppCmdNode{
				Name:     "names",
				Callable: CmdAipsetupSysAllNames,
			},

			&cliapp.AppCmdNode{
				Name:     "name-asps",
				Callable: CmdAipsetupSysNameASPs,
			},

			&cliapp.AppCmdNode{
				Name:     "asp-files",
				Callable: CmdAipsetupSysASPFiles,
			},

			&cliapp.AppCmdNode{
				Name:     "asp-rm",
				Callable: CmdAipsetupSysRemoveASP,
			},
		},
	}

}

func CmdAipsetupSysAllAsps(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {

	var (
		host string
		arch string
		sys  *aipsetup.System
	)

	{
		var fast_ret *cliapp.AppResult

		_, host, arch, sys, fast_ret = CmdRoutineRootHostArchSys(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	installed_asps := sys.Asps().ListInstalledASPs(host, arch)

	num_len := len(fmt.Sprintf("%d", len(installed_asps)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range installed_asps {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysAllNames(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {

	var (
		host string
		arch string
		sys  *aipsetup.System
	)

	{
		var fast_ret *cliapp.AppResult

		_, host, arch, sys, fast_ret = CmdRoutineRootHostArchSys(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	installed_names := sys.Asps().ListInstalledPackageNames(host, arch)

	sort.Strings(installed_names)

	num_len := len(fmt.Sprintf("%d", len(installed_names)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range installed_names {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysNameASPs(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {

	var (
		// root string = "/"
		host string
		arch string
		name string

		sys *aipsetup.System
	)

	{
		var fast_ret *cliapp.AppResult

		_, host, arch, sys, fast_ret = CmdRoutineRootHostArchSys(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	{
		var fast_ret *cliapp.AppResult

		name, fast_ret = CmdRoutineASPName(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	installed_archs := sys.Asps().ListInstalledPackageNameASPs(name, host, arch)

	sort.Strings(installed_archs)

	num_len := len(fmt.Sprintf("%d", len(installed_archs)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range installed_archs {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysASPFiles(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {

	var (
		//root string
		name string
		sys  *aipsetup.System
	)

	{
		var fast_ret *cliapp.AppResult

		_, sys, fast_ret = CmdRoutineRootSys(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	{
		var fast_ret *cliapp.AppResult

		name, fast_ret = CmdRoutineASPName(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	liaf_res, err := sys.Asps().ListInstalledASPFiles(name)

	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: "can't get file list for named ASP",
		}
	}

	files := liaf_res.FileList

	sort.Strings(files)

	num_len := len(fmt.Sprintf("%d", len(files)))

	print_fmt := "#%0" + fmt.Sprintf("%d", num_len) + "d %s\n"

	for ii, i := range files {
		fmt.Printf(print_fmt, ii, i)
	}

	return &cliapp.AppResult{Code: 0}
}

func CmdAipsetupSysRemoveASP(
	getopt_result *cliapp.GetOptResult,
	available_options cliapp.GetOptCheckList,
	depth_level []string,
	subnode *cliapp.AppCmdNode,
	rootnode *cliapp.AppCmdNode,
	arg0 string,
	pass_data *interface{},
) *cliapp.AppResult {

	var (
		//	root string
		asp string
		sys *aipsetup.System
	)

	{
		var fast_ret *cliapp.AppResult

		_, sys, fast_ret = CmdRoutineRootSys(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	{
		var fast_ret *cliapp.AppResult

		asp, fast_ret = CmdRoutineASPName(getopt_result)

		if fast_ret != nil {
			return fast_ret
		}
	}

	err := sys.Asps().RemoveASP(asp, true, true)

	if err != nil {
		return &cliapp.AppResult{
			Code:    10,
			Message: err.Error(),
		}
	}

	return &cliapp.AppResult{Code: 0}
}
