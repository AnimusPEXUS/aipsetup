package aipsetup

import (
	"path"
)

type SystemMovedRoot struct {
	system   *System
	new_root string
}

func NewSystemMovedRoot(new_root string, sys *System) *SystemMovedRoot {
	return &SystemMovedRoot{
		system:   sys,
		new_root: new_root,
	}
}

func (self *SystemMovedRoot) GetInstalledASPDir() string {
	return path.Join(self.new_root, "/var", "aipsetup5", "packages")
}

func (self *SystemMovedRoot) GetInstalledASPSumsDir() string {
	return path.Join(self.GetInstalledASPDir(), "sums")
}

func (self *SystemMovedRoot) GetInstalledASPBuildLogsDir() string {
	return path.Join(self.GetInstalledASPDir(), "buildlogs")
}

func (self *SystemMovedRoot) GetInstalledASPDepsDir() string {
	return path.Join(self.GetInstalledASPDir(), "deps")
}
