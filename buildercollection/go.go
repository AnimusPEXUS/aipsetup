package buildercollection

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["go"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_go(bs)
	}
}

var _ basictypes.BuilderI = &Builder_go{}

type Builder_go struct {
	*Builder_std

	os_name string
	arch    string
}

func NewBuilder_go(bs basictypes.BuildingSiteCtlI) (*Builder_go, error) {
	self := new(Builder_go)

	self.Builder_std = NewBuilder_std(bs)

	// TODO: dehardcode
	self.os_name = "linux"
	self.arch = "amd64"

	return self, nil
}

func (self *Builder_go) DefineActions() (basictypes.BuilderActions, error) {

	ret, err := self.Builder_std.DefineActions()
	if err != nil {
		return nil, err
	}

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	err = ret.Replace(
		"build",
		&basictypes.BuilderAction{
			"build",
			self.BuilderActionBuild,
		},
	)
	if err != nil {
		return nil, err
	}

	// ret, err = ret.AddAfterName(
	// 	basictypes.BuilderActions{
	// 		&basictypes.BuilderAction{
	// 			"stop",
	// 			func(log *logger.Logger) error { return errors.New("todo") },
	// 		},
	// 	},
	// 	"build",
	// )
	// if err != nil {
	// 	return nil, err
	// }

	err = ret.Replace(
		"distribute",
		&basictypes.BuilderAction{
			"distribute",
			self.BuilderActionDistribute,
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_go) BuilderActionBuild(
	log *logger.Logger,
) error {
	cwd := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "src")

	env := environ.NewFromStrings(os.Environ())

	env_new := environ.New()
	env_new.Set("GOROOT_BOOTSTRAP", env.Get("GOROOT", "/"))
	env_new.Set("GOOS", self.os_name)
	env_new.Set("GOARCH", self.arch)

	log.Info("Environment Edits:")
	for _, i := range env_new.Strings() {
		log.Info(" " + i)
	}

	env.UpdateWith(env_new)

	c := exec.Command("bash", "./bootstrap.bash")
	c.Dir = cwd
	c.Env = env.Strings()

	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_go) BuilderActionDistribute(
	log *logger.Logger,
) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	host_lib_dir, err := calc.CalculateHostLibDir()
	if err != nil {
		return err
	}

	os_name := self.os_name
	arch := self.arch

	go_version_string := info.PackageVersion
	gogo_version_string := "go" + go_version_string

	go_dir := fmt.Sprintf("go-%s-%s-bootstrap", os_name, arch)

	godir_path := path.Join(self.GetBuildingSiteCtl().GetPath(), go_dir)

	dir_path := path.Join(host_lib_dir, gogo_version_string)

	dst_dir_path := path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), dir_path)

	err = os.MkdirAll(dst_dir_path, 0755)
	if err != nil {
		return err
	}

	err = filetools.CopyTree(
		godir_path,
		dst_dir_path,
		false,
		false,
		false,
		true,
		log,
		false,
		true,
		filetools.CopyWithInfo,
	)
	if err != nil {
		return err
	}

	dst_etc_dir := path.Join(
		self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
		"etc",
		"profile.d",
		"SET",
	)

	etc_file_path := path.Join(
		dst_etc_dir,
		fmt.Sprintf(
			"009.%s.%s.%s.sh",
			gogo_version_string,
			info.Host,
			info.HostArch,
		),
	)

	err = os.MkdirAll(dst_etc_dir, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(etc_file_path)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(
		fmt.Sprintf(
			`#!/bin/bash

export GOROOT="%s"

export PATH+=":$GOROOT/bin"

TEMP_PATH="$HOME/gopath_clean"

export GOPATH="$TEMP_PATH"
export PATH+=":$TEMP_PATH/bin"

TEMP_PATH="$HOME/gopath_work"

export GOPATH+=":$TEMP_PATH"
export PATH+=":$TEMP_PATH/bin"

unset TEMP_GOPATH

`,
			dir_path,
		),
	)
	if err != nil {
		return err
	}

	return nil
}
