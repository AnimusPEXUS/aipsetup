package buildingtools

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/AnimusPEXUS/utils/logger"
)

type CMake struct {
}

func (self *CMake) CMake(
	args []string,
	env []string,
	env_mode EnvironmentOperationMode,
	cmakelist_filename string,
	cmakelist_dirpath string,
	working_dirpath string,
	cmake_program string,
	log *logger.Logger,
) error {

	if cmakelist_filename == "" {
		cmakelist_filename = "CMakeList.txt"
	}

	cmakelist_filename = path.Base(cmakelist_filename)

	if cmake_program == "" {
		cmake_program = "cmake"
	}

	cmakelist_dirpath, err := filepath.Abs(cmakelist_dirpath)
	if err != nil {
		return err
	}

	working_dirpath, err = filepath.Abs(working_dirpath)
	if err != nil {
		return err
	}

	int_env := make([]string, 0)
	if env_mode == Copy {
		int_env = append(int_env, os.Environ()...)
	}
	int_env = append(int_env, env...)

	//   cmd = ['cmake'] + opts + ['--build=' + build_dir] + args + [src_dir]
	int_args := make([]string, 0)

	int_args = append(
		int_args,
		"--build="+working_dirpath,
	)

	int_args = append(
		int_args,
		args...,
	)

	int_args = append(
		int_args,
		cmakelist_dirpath,
	)

	log.Info("CMake Parameters:")
	log.Info("  executable: " + cmake_program)
	log.Info("  arguments:")
	for _, i := range int_args {
		log.Info(fmt.Sprintf("    %s", i))
	}
	log.Info("  environment:")
	for _, i := range env {
		log.Info(fmt.Sprintf("    %s", i))
	}
	log.Info("  CMakeList.txt dir: " + cmakelist_dirpath)
	log.Info("  working dir: " + working_dirpath)

	c := exec.Command(cmake_program, int_args...)
	c.Dir = working_dirpath
	c.Env = int_env
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()

	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}
