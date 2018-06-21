package aipsetup

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
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
[x86_64-pc-linux-gnu]
archs = i686-pc-linux-gnu

[tarball_downloading]
enabled = false
repository = path
`)

var _ basictypes.SystemI = &System{}

type System struct {
	root string

	log *logger.Logger

	ASPs *SystemPackages

	cfg *ini.File

	host  string
	archs []string

	root_paths *SystemMovedRoot

	valuescalculator *SystemValuesCalculator
	sysupdates       *SystemUpdates
}

func NewSystem(root string, log *logger.Logger) *System {
	self := new(System)
	self.log = log

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
			path.Join(
				self.root,
				"/etc",
				basictypes.AIPSETUP_SYSTEM_CONFIG_FILENAME,
			),
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

func (self *System) Cfg() *ini.File {
	return self.cfg
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

func (self *System) GetMovedRootForCrossbuilder(host, target string) *SystemMovedRoot {
	d := self.GetSystemValuesCalculator().CalculateHostCrossbuilderDir(host, target)
	return NewSystemMovedRoot(d, self)
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

func (self *System) GetSystemUpdates() *SystemUpdates {
	if self.sysupdates == nil {
		self.sysupdates = NewSystemUpdates(self)
	}
	return self.sysupdates
}

func (self *System) GenLocale() error {

	sys_calc := self.GetSystemValuesCalculator()

	host, err := self.Host()
	if err != nil {
		return err
	}

	archs, err := self.Archs()
	if err != nil {
		return err
	}

	for _, arch := range archs {
		prefix := sys_calc.CalculateHostArchDir(host, arch)
		for _, libdirname := range basictypes.POSSIBLE_LIBDIR_NAMES {
			prefix_lib := path.Join(prefix, libdirname)
			prefix_lib_locale := path.Join(prefix_lib, "locale")
			if _, err := os.Stat(prefix_lib); err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return err
			}

			err := os.RemoveAll(prefix_lib_locale)
			if err != nil {
				if !os.IsNotExist(err) {
					return err
				}
			}

			err = os.MkdirAll(prefix_lib_locale, 0755)
			if err != nil {
				return err
			}

			for _, locale_name := range []string{
				"en_US.UTF-8",
			} {
				locale_name_spl := strings.SplitN(locale_name, ".", 2)
				prefix_lib_locale_name := path.Join(prefix_lib_locale, locale_name)
				rel_locale_name_dir, err := filepath.Rel(self.Root(), prefix_lib_locale_name)
				if err != nil {
					return err
				}

				cmd_args := []string{
					self.Root(),
					"localedef",
					"-f", locale_name_spl[1],
					"-i", locale_name_spl[0],
					rel_locale_name_dir,
				}
				fmt.Printf("Running %v inside %s\n", cmd_args[1:], self.Root())
				c := exec.Command("chroot", cmd_args...)
				c.Dir = self.Root()
				err = c.Run()
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
