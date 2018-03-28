package aipsetup

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/set"
	"github.com/AnimusPEXUS/utils/systemtriplet"
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

	host  string
	archs []string

	root_paths *SystemMovedRoot

	valuescalculator *SystemValuesCalculator
}

func NewSystem(root string) *System {
	self := new(System)

	self.valuescalculator = NewSystemValuesCalculator(self)

	if root, err := filepath.Abs(root); err != nil {
		panic(err)
	} else {
		self.root = root
	}

	if cfg, err := ini.Load(
		DEFAULT_AIPSETUP_SYSTEM_CONFIG,
	); err != nil {
		panic(err)
	} else {
		self.cfg = cfg
	}

	if res, err :=
		ioutil.ReadFile(
			path.Join(self.root, "/etc", AIPSETUP_SYSTEM_CONFIG_FILENAME),
		); err != nil {
		panic(err)
	} else {
		if err := self.cfg.Append(res); err != nil {
			panic(err)
		}
	}

	self.ASPs = NewSystemPackages(self)
	self.root_paths = NewSystemMovedRoot(root, self)

	return self
}

func (self *System) Root() string {
	return self.root
}

func (self *System) Host() (string, error) {

	if self.host == "" {

		st := &systemtriplet.SystemTriplet{}

		switch runtime.GOARCH {
		case "i386":
			st.CPU = "i686"
		case "amd64":
			st.CPU = "x86_64"
		default:
			return "", errors.New("can't determine current cpu")
		}

		st.Company = "pc"

		switch runtime.GOOS {
		case "linux":
			st.Kernel = "linux"
			st.OS = "gnu"
		default:
			return "", errors.New("can't determine current kernel and os")
		}

		self.host = st.String()
	}

	return self.host, nil
}

func (self *System) Archs() ([]string, error) {
	if self.archs == nil {
		host, err := self.Host()
		if err != nil {
			return nil, err
		}

		sect, err := self.cfg.GetSection("archs")
		if err != nil {
			return nil, err
		}

		k := sect.Key(host)

		s := set.NewSetString()
		s.Add(host)

		for _, i := range strings.Split(k.String(), " ") {
			s.Add(i)
		}

		self.archs = s.ListStrings()

	}

	return self.archs, nil
}

func (self *System) GetInstalledASPDir() string {
	return self.root_paths.GetInstalledASPDir()
}

func (self *System) GetInstalledASPSumsDir() string {
	return self.root_paths.GetInstalledASPSumsDir()
}

func (self *System) GetInstalledASPBuildLogsDir() string {
	return self.root_paths.GetInstalledASPBuildLogsDir()
}

func (self *System) GetInstalledASPDepsDir() string {
	return self.root_paths.GetInstalledASPDepsDir()
}

func (self *System) GetMovedRootForCrossbuilder(host, target string) (*SystemMovedRoot, error) {
	calc := self.GetSystemValuesCalculator()
	d, err := calc.CalculateHostCrossbuilderDir(host, target)
	if err != nil {
		return nil, err
	}
	return NewSystemMovedRoot(d, self), nil
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

func (self *System) GetSystemValuesCalculator() basictypes.SystemValuesCalculatorI {
	return self.valuescalculator
}
