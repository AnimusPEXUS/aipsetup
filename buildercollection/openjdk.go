package buildercollection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openjdk"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_openjdk(bs)
	}
}

type Builder_openjdk struct {
	*Builder_std
}

func NewBuilder_openjdk(bs basictypes.BuildingSiteCtlI) (*Builder_openjdk, error) {
	self := new(Builder_openjdk)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.EditConfigureEnvCB = self.EditConfigureEnv
	self.EditBuildConcurentJobsCountCB = self.EditBuildConcurentJobsCount
	self.EditDistributeDESTDIRCB = self.EditDistributeDESTDIR

	return self, nil
}

func (self *Builder_openjdk) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret, err := ret.AddActionsAfterName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "after_distribute",
				Callable: self.AfterOpenJDKDistribution,
			},
		},
		"distribute",
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_openjdk) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	for i := len(ret) - 1; i != -1; i -= 1 {
		if ret[i] == "--enable-shared" ||
			strings.HasPrefix(ret[i], "CC=") ||
			strings.HasPrefix(ret[i], "CXX=") ||
			strings.HasPrefix(ret[i], "GCC=") {
			ret = append(ret[:i], ret[i+1:]...)
		}
	}

	ret = append(
		ret,
		[]string{"--disable-warnings-as-errors"}...,
	)

	return ret, nil
}

func (self *Builder_openjdk) EditConfigureEnv(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error) {
	for _, i := range []string{"CC", "CXX", "GCC"} {
		ret.Del(i)
	}
	return ret, nil
}

func (self *Builder_openjdk) EditBuildConcurentJobsCount(log *logger.Logger, ret int) int {
	return -1
}

func (self *Builder_openjdk) EditDistributeDESTDIR(log *logger.Logger, ret string) (string, error) {
	return "INSTALL_PREFIX", nil
}

func (self *Builder_openjdk) AfterOpenJDKDistribution(log *logger.Logger) error {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return err
	}

	java_dir := path.Join(install_prefix, "opt", "java")

	dst_dir := self.bs.GetDIR_DESTDIR()

	dst_java_dir := path.Join(dst_dir, java_dir)

	dst_etc_dir := path.Join(dst_dir, "etc", "profile.d", "SET")

	java09 := path.Join(
		dst_etc_dir,
		fmt.Sprintf(
			"009.java.%s.%s.sh",
			info.Host,
			info.HostArch,
		),
	)

	os.RemoveAll(path.Join(dst_dir, "bin"))

	dst_jvm := path.Join(dst_dir, "jvm")
	dst_jvm_jdk := ""
	jdk_basename := ""

	{

		_, err := os.Stat(dst_jvm)
		if err != nil {
			return err
		}

		files, err := ioutil.ReadDir(dst_jvm)

		for _, i := range files {
			if strings.HasPrefix(i.Name(), "openjdk") {
				jdk_basename = i.Name()
				dst_jvm_jdk = path.Join(dst_jvm, jdk_basename)
				break
			}
		}

	}

	if dst_jvm_jdk == "" {
		return errors.New("couldn't find openjdk dir inside DESTDIR")
	}

	dst_java_dir_jdk := path.Join(dst_java_dir, jdk_basename)

	err = os.MkdirAll(dst_java_dir, 0700)
	if err != nil {
		return err
	}

	err = os.Rename(dst_jvm_jdk, dst_java_dir_jdk)
	if err != nil {
		return err
	}

	err = os.Remove(dst_jvm)
	if err != nil {
		return err
	}

	for _, i := range []string{"jre", "jdk", "java"} {

		nfn := path.Join(dst_java_dir, i)

		err = os.Remove(nfn)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}

		err = os.Symlink(jdk_basename, nfn)
		if err != nil {
			return err
		}

	}

	err = os.MkdirAll(dst_etc_dir, 0700)
	if err != nil {
		return err
	}

	// NOTE: this is old variant. looks like JKD10 doesnt have jre subdir
	//	script := ""
	//	script += `#!/bin/bash` + "\n"
	//	script += `export JAVA_HOME=` + java_dir + `/jdk` + "\n"
	//	script += `export PATH=$PATH:$JAVA_HOME/bin:$JAVA_HOME/jre/bin` + "\n"
	//	script += `export MANPATH=$MANPATH:$JAVA_HOME/man` + "\n"
	//	script += `if [ "${#LD_LIBRARY_PATH}" -ne "0" ]; then` + "\n"
	//	script += `    LD_LIBRARY_PATH+=":"` + "\n"
	//	script += `fi` + "\n"
	//	// TODO: following paths are hardcoded, should not be.
	//	script += `export LD_LIBRARY_PATH+="$JAVA_HOME/jre/lib/i386:$JAVA_HOME/jre/lib/i386/client"` + "\n"
	//	script += `export LD_LIBRARY_PATH+=":$JAVA_HOME/jre/lib/amd64:$JAVA_HOME/jre/lib/amd64/client"` + "\n"

	script := ""
	script += `#!/bin/bash` + "\n"
	script += `export JAVA_HOME=` + java_dir + `/jdk` + "\n"
	script += `export PATH=$PATH:$JAVA_HOME/bin` + "\n"
	script += `export MANPATH=$MANPATH:$JAVA_HOME/man` + "\n"
	script += `if [ "${#LD_LIBRARY_PATH}" -ne "0" ]; then` + "\n"
	script += `    LD_LIBRARY_PATH+=":"` + "\n"
	script += `fi` + "\n"
	script += `export LD_LIBRARY_PATH+="$JAVA_HOME/lib"` + "\n"

	f, err := os.Create(java09)
	if err != nil {
		return err
	}

	f.WriteString(script)

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
