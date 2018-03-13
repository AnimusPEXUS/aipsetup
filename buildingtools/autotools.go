package buildingtools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/utils/logger"
)

type (
	EnvironmentOperationMode uint
)

const (
	Copy EnvironmentOperationMode = iota
	Clean
)

type Autotools struct {
}

func (self *Autotools) Extract(
	filename string,
	outputdir string,
	tempdir string,
	unwrap bool,
	is_more_than_one_extracted_ok bool,
	cleanup_outputdir bool,
	log *logger.Logger,
) error {

	os.RemoveAll(tempdir)
	os.MkdirAll(tempdir, 0700)
	defer os.RemoveAll(tempdir)

	if cleanup_outputdir {
		os.RemoveAll(outputdir)
	}

	os.MkdirAll(outputdir, 0700)

	// TODO: mega fast decision. this should be replaced by internal functionality
	log.Info("starting tar utility")
	c := exec.Command("tar", "-vxf", filename, "-C", tempdir)

	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()

	c_res := c.Run()

	if c_res != nil {
		log.Error("error running tar utility: " + c_res.Error())
		return c_res
	}

	log.Info("tar exited ok")

	info, err := ioutil.ReadDir(tempdir)
	if err != nil {
		return err
	}

	log.Info(
		fmt.Sprintf("tar work resulted in %d files", len(info)),
	)

	directory_to_work_with := tempdir

	if unwrap {
		is_more_than_one_extracted_ok = false
	}

	if !is_more_than_one_extracted_ok {
		if len(info) != 1 {
			return errors.New("can't unwrap: extracted more than one item")
		}
	}

	if unwrap {

		extracted_dir := ""

		for _, i := range info {
			if i.IsDir() {
				extracted_dir = i.Name()
				break
			}
		}

		if extracted_dir == "" {
			return errors.New("unwrap failed no directories extracted")
		}

		directory_to_work_with = path.Join(directory_to_work_with, extracted_dir)
	}

	info, err = ioutil.ReadDir(directory_to_work_with)
	if err != nil {
		return err
	}

	for _, i := range info {
		os.Rename(
			path.Join(directory_to_work_with, i.Name()),
			path.Join(outputdir, i.Name()),
		)
	}

	log.Info("extraction procedure ended without errors")

	return nil
}

func (self *Autotools) GenerateConfigureIfNeeded(
	directory string,
	force bool,
) error {

	presumed_configure_full_filename := path.Join(directory, "configure")

	is_presumed_configure_full_filename_exists := false

	if _, err := os.Stat(presumed_configure_full_filename); err == nil {
		is_presumed_configure_full_filename_exists = true
	}

	if !is_presumed_configure_full_filename_exists || force {
		// ('makeconf.sh', ['./makeconf.sh']),
		// ('autogen.sh', ['./autogen.sh']),
		// ('bootstrap.sh', ['./bootstrap.sh']),
		// ('bootstrap', ['./bootstrap']),
		// ('genconfig.sh', ['./genconfig.sh']),
		// ('configure.ac', ['autoreconf', '-i']),
		// ('configure.in', ['autoreconf', '-i']),

		{
			checked_file := path.Join(directory, "makeconf.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./makeconf.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "autogen.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./autogen.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "bootstrap.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./bootstrap.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "bootstrap")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./bootstrap")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "genconfig.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./genconfig.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "configure.ac")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("autoconf", "-i")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "configure.in")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("autoconf", "-i")
				c.Dir = directory
				return c.Run()
			}
		}

		return errors.New(
			"no acceptable configure script creating tools found",
		)

	}

	return nil
}

func (self *Autotools) Configure(
	args []string,
	env []string,
	env_mode EnvironmentOperationMode,
	configure_filename string,
	configure_dirpath string,
	working_dirpath string,
	// 1) calculates absolute path to configure and uses it as
	// run-path or 2) calculates relative path from working dir to configure file
	// and runs it relatively
	run_relative bool,
	// whatever to start script it self or to
	// execute shell programm with configure's path as parameter
	run_as_argument_to_shell bool,
	shell_program string,
	log *logger.Logger,
) error {

	configure_dirpath, err := filepath.Abs(configure_dirpath)
	if err != nil {
		return err
	}

	working_dirpath, err = filepath.Abs(working_dirpath)
	if err != nil {
		return err
	}

	var dirpath string

	if run_relative {
		dirpath, err = filepath.Rel(working_dirpath, configure_dirpath)
		if err != nil {
			return err
		}
	} else {
		dirpath = configure_dirpath
	}

	executable := configure_filename

	int_env := make([]string, 0)
	if env_mode == Copy {
		int_env = append(int_env, os.Environ()...)
	}
	int_env = append(int_env, env...)

	int_args := make([]string, 0)

	dirpath_script_to_run := strings.Join(
		[]string{dirpath, configure_filename},
		string(os.PathSeparator),
	)

	if run_as_argument_to_shell {
		executable = shell_program
		int_args = append(int_args, dirpath_script_to_run)
		int_args = append(int_args, args...)
	} else {
		executable = dirpath_script_to_run
		// int_args = append(int_args, path.Join(dirpath, script_to_run))
		int_args = append(int_args, args...)
	}

	cmd := exec.Command(executable, int_args...)
	cmd.Env = int_env
	cmd.Dir = working_dirpath

	log.Info("Configure Parameters:")
	log.Info("  executable: " + executable)
	log.Info("  arguments:")
	for _, i := range int_args {
		log.Info(fmt.Sprintf("    %s", i))
	}
	log.Info("  environment:")
	for _, i := range env {
		log.Info(fmt.Sprintf("    %s", i))
	}
	log.Info("  configure dir: " + configure_dirpath)
	log.Info("  working dir: " + working_dirpath)

	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	ret := cmd.Run()

	log.Info(cmd.ProcessState.String())

	return ret
}

/*

env - additional environment variables

env_mode - use empty variables or copy of variables passed to aipsetup5

makefile_filename - based automatically. name of Makefile. usually "Makefile"

makefile_dirpath - absoluted automatically. directory in which named makefile
	contained

working_dirpath  - absoluted automatically. working dir which shold be current
	for startend make utility

*/
func (self *Autotools) Make(
	args []string,
	env []string,
	env_mode EnvironmentOperationMode,
	makefile_filename string,
	makefile_dirpath string,
	working_dirpath string,
	make_program string,
	log *logger.Logger,
) error {

	makefile_filename = path.Base(makefile_filename)

	makefile_dirpath, err := filepath.Abs(makefile_dirpath)
	if err != nil {
		return err
	}

	working_dirpath, err = filepath.Abs(working_dirpath)
	if err != nil {
		return err
	}

	executable := make_program

	int_env := make([]string, 0)
	if env_mode == Copy {
		int_env = append(int_env, os.Environ()...)
	}
	int_env = append(int_env, env...)

	int_args := make([]string, 0)

	dirpath, err := filepath.Rel(working_dirpath, makefile_dirpath)
	if err != nil {
		return err
	}

	dirpath_script_to_run := strings.Join(
		[]string{dirpath, makefile_filename},
		string(os.PathSeparator),
	)

	int_args = append(int_args, []string{"-f", dirpath_script_to_run}...)

	int_args = append(int_args, args...)

	cmd := exec.Command(executable, int_args...)
	cmd.Env = int_env
	cmd.Dir = working_dirpath

	log.Info("Make Parameters:")
	log.Info("  executable: " + executable)
	log.Info("  arguments:")
	for _, i := range int_args {
		log.Info(fmt.Sprintf("    %s", i))
	}
	log.Info("  environment:")
	for _, i := range env {
		log.Info(fmt.Sprintf("    %s", i))
	}
	log.Info("  Makefile dir: " + makefile_dirpath)
	log.Info("  working dir: " + working_dirpath)

	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	ret := cmd.Run()

	log.Info(cmd.ProcessState.String())

	return ret
}
