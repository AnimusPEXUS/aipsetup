package aipsetup

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/set"
	"github.com/go-ini/ini"
)

var (
	DEFAULT_HOST         = "x86_64-pc-linux-gnu"
	DEFAULT_ARCHS_STRING = "i686-pc-linux-gnu"
)

var DEFAULT_AIPSETUP_SYSTEM_CONFIG = []byte("" +
	`
[main]
host = x86_64-pc-linux-gnu
archs = i686-pc-linux-gnu

[tarball_downloading]
enabled = false
repository = path
`)

var _ basictypes.SystemI = &System{}

type System struct {
	root string

	ASPs *SystemPackages

	cfg *ini.File
}

func NewSystem(root string) *System {
	ret := new(System)

	if root, err := filepath.Abs(root); err != nil {
		panic(err)
	} else {
		ret.root = root
	}

	if cfg, err := ini.Load(
		DEFAULT_AIPSETUP_SYSTEM_CONFIG,
	); err != nil {
		panic(err)
	} else {
		ret.cfg = cfg
	}

	if res, err :=
		ioutil.ReadFile(
			path.Join(ret.root, "/etc", AIPSETUP_SYSTEM_CONFIG_FILENAME),
		); err != nil {
		panic(err)
	} else {
		if err := ret.cfg.Append(res); err != nil {
			panic(err)
		}
	}

	ret.ASPs = NewSystemPackages(ret)

	return ret
}

func (self *System) Root() string {
	return self.root
}

func (self *System) Host() string {
	sect, err := self.cfg.GetSection("main")
	if err != nil {
		return DEFAULT_HOST
	}
	ret := sect.Key("host").MustString(DEFAULT_HOST)
	return ret
}

func (self *System) Archs() []string {
	archs_strings := DEFAULT_ARCHS_STRING
	sect, err := self.cfg.GetSection("main")
	if err == nil {
		archs_strings = sect.Key("archs").MustString(DEFAULT_ARCHS_STRING)
	}

	res := strings.Split(archs_strings, " ")

	lst := make([]string, 0)
	lst = append(lst, self.Host())
	lst = append(lst, res...)

	s := set.NewSetString()
	for _, val := range lst {
		s.Add(val)
	}

	return s.ListStrings()
}

// func (self *System) GetSystemConfigFileName() string {
// 	return path.Join(self.Root(), "etc", "aipsetup5.ini")
// }

func (self *System) GetInstalledASPDir() string {
	return path.Join(self.root, "/var", "log", "packages")
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

func (self *System) GetTarballRepoRootDir() string {
	ret := self.cfg.Section("tarball_downloading").Key("path").MustString("")
	if ret == "" ||
		!strings.HasPrefix(ret, "/") ||
		(func() error { _, err := os.Stat(ret); return err }() != nil) {
		panic("invalid value for tarball_downloading -> path")
	}
	return ret
}

// func (self *System) GetTarballsRepository() *tarballrepository.Repository {
// 	return tarballrepository.NewRepository(self)
// }
