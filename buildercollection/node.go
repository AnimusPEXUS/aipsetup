package buildercollection

import (
	"fmt"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["node"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_node(bs)
	}
}

type Builder_node struct {
	*Builder_std
}

func NewBuilder_node(bs basictypes.BuildingSiteCtlI) (*Builder_node, error) {
	self := new(Builder_node)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_node) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	if !info.ThisIsSubarchBuilding() {

		ret, err = ret.AddActionAfterNameShort(
			"distribute",
			"env_config", self.BuilderActionEnvConfig,
		)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_node) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{
			"--enable-shared",
		},
		[]string{
			"--libdir=",
			"--docdir=",
			"--sysconfdir",
			"--localstatedir",
			"--host",
			"--build",
			"CC=",
			"CXX=",
		},
	)
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{}...,
	)

	return ret, nil
}

func (self *Builder_node) BuilderActionEnvConfig(log *logger.Logger) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	node_version_string := info.PackageVersion
	nodenode_version_string := "nodejs" + node_version_string

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
			nodenode_version_string,
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
		`#!/bin/bash

export NPM_CONFIG_PREFIX="$HOME/.nodejs_npm_prefix"

export PATH+=":$NPM_CONFIG_PREFIX/bin"

`,
	)
	if err != nil {
		return err
	}

	return nil
}
