package buildingtools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/AnimusPEXUS/gologger"
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
	srcdir string,
	tempdir string,
	unwrap bool,
	rename_dir bool,
	new_name string,
	is_more_than_one_extracted_ok bool,
	cleanup_srcdir bool,
	log *gologger.Logger,
) error {

	os.RemoveAll(tempdir)
	os.MkdirAll(tempdir, 0700)
	defer os.RemoveAll(tempdir)

	if cleanup_srcdir {
		os.RemoveAll(srcdir)
	}

	os.MkdirAll(srcdir, 0700)

	// TODO: mega fast decision. this should be replaced by internal functionality
	log.Info("starting tar utility")
	c := exec.Command("tar", "-xf", filename, "-C", tempdir)
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
		fmt.Sprintf("tar work resulted in %d files and/or directories", len(info)),
	)

	if unwrap {

		log.Info("unwrapping and moving files to source directory")

		if len(info) != 1 && !is_more_than_one_extracted_ok {
			return errors.New("extracted more than one item")
		}

		extracted_dir := ""

		for _, i := range info {
			if i.IsDir() {
				extracted_dir = i.Name()
				break
			}
		}

		if extracted_dir == "" {
			return errors.New("no directories extracted")
		}

		extracted_dir = path.Join(tempdir, extracted_dir)

		info, err = ioutil.ReadDir(extracted_dir)
		if err != nil {
			return err
		}

		for _, i := range info {
			os.Rename(
				path.Join(extracted_dir, i.Name()),
				path.Join(srcdir, i.Name()),
			)
		}

	} else {
		log.Info("moving files to source directory")
		for _, i := range info {
			os.Rename(
				path.Join(tempdir, i.Name()),
				path.Join(srcdir, i.Name()),
			)
		}
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
	log *gologger.Logger,
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

	dirpath_script_to_run := path.Join(dirpath, configure_filename)

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
	cmd.Env = env
	cmd.Dir = working_dirpath
	ret := cmd.Run()

	return ret
}
