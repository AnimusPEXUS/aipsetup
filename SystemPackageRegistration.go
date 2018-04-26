package aipsetup

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

type SystemPackageRegistration struct {
	ASPName *basictypes.ASPName
	sys     *System

	Pkg       []string
	Sums      []string
	Buildlogs []string
	Deps      []string
}

func FindSystemPackageRegistrationByName(
	aspname *basictypes.ASPName,
	sys *System,
) (*SystemPackageRegistration, error) {
	self := new(SystemPackageRegistration)
	self.ASPName = aspname
	self.sys = sys

	search := func(aspname *basictypes.ASPName, dir string) ([]string, error) {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}

		ret := make([]string, 0)

		for _, i := range files {
			if i.IsDir() {
				continue
			}

			fullpath := path.Join(dir, path.Base(i.Name()))

			parsed, err := basictypes.NewASPNameFromString(fullpath)
			if err != nil {
				self.sys.log.Warning("couldn't parse " + fullpath)
				continue
			}

			if parsed.IsEqual(aspname) {
				ret = append(ret, path.Base(i.Name()))
			}

		}
		return ret, nil
	}

	was_errors := false

	if t, err := search(aspname, self.sys.GetInstalledASPDir()); err != nil {
		self.sys.log.Error(
			"error searching package " +
				aspname.String() +
				" registration: " +
				err.Error(),
		)
		was_errors = true
	} else {
		self.Pkg = t
	}

	if t, err := search(aspname, self.sys.GetInstalledASPDepsDir()); err == nil {
		self.Deps = t
	}

	if t, err := search(aspname, self.sys.GetInstalledASPSumsDir()); err == nil {
		self.Sums = t
	}

	if t, err := search(aspname, self.sys.GetInstalledASPBuildLogsDir()); err == nil {
		self.Buildlogs = t
	}

	if was_errors {
		return nil, errors.New(
			"there was errors searching package's registration files",
		)
	}

	return self, nil
}

func (self *SystemPackageRegistration) Found() bool {
	return len(self.Pkg) != 0
}

func (self *SystemPackageRegistration) DeleteAll() error {
	var err error

	d := self.sys.GetInstalledASPDepsDir()

	for _, i := range self.Deps {
		j := path.Join(d, i)
		err = os.Remove(j)
		if err != nil {
			return err
		}
	}

	d = self.sys.GetInstalledASPBuildLogsDir()

	for _, i := range self.Buildlogs {
		j := path.Join(d, i)
		err = os.Remove(j)
		if err != nil {
			return err
		}
	}

	d = self.sys.GetInstalledASPSumsDir()

	for _, i := range self.Sums {
		j := path.Join(d, i)
		err = os.Remove(j)
		if err != nil {
			return err
		}
	}

	d = self.sys.GetInstalledASPDir()

	for _, i := range self.Pkg {
		j := path.Join(d, i)
		err = os.Remove(j)
		if err != nil {
			return err
		}
	}

	return nil
}
