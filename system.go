package aipsetup

import (
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"
)

type AIPSETUP_SYSTEM_CFG_YAML struct {
	Host  string   `yaml:"host"`
	Archs []string `yaml:"archs"`
}

type System struct {
	root string

	host  string
	archs []string

	asps *SystemPackages
}

func (self *System) Root() string {
	return self.root
}

func (self *System) Host() string {
	return self.host
}

func (self *System) Asps() *SystemPackages {
	return self.asps
}

func (self *System) GetSystemConfigFileName() string {
	return path.Join(self.root, "etc", "aipsetup.system.cfg.yaml")
}

func (self *System) GetInstalledASPDir() string {
	return path.Join(self.root, "var", "log", "packages")
}

func (self *System) GetInstalledASPSumsDir() string {
	return path.Join(self.GetInstalledASPDir(), "sums")
}

func (self *System) GetInstalledASPBuildLogsDir() string {
	return path.Join(self.GetInstalledASPDir(), "buildlogs")
}

func (self *System) GetInstalledASPDepsDir() string {
	return path.Join(self.GetInstalledASPDir(), "deps")
}

func NewSystem(root string) *System {
	ret := new(System)
	ret.root = root
	ret.host = "x86_64-pc-linux-gnu"
	ret.archs = append(ret.archs, ret.host, "i686-pc-linux-gnu")

	yaml_file, err := ioutil.ReadFile(ret.GetSystemConfigFileName())

	if err == nil && yaml_file != nil {
		t := AIPSETUP_SYSTEM_CFG_YAML{}
		err := yaml.Unmarshal(yaml_file, &t)
		if err == nil {
		}
	}

	ret.asps = NewSystemPackages(ret)

	return ret
}
