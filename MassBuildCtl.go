package aipsetup

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
)

// Used us and upper interface for packages mass building
type MassBuildCtl struct {
	path string
	sys  *System
	log  *logger.Logger
	info *basictypes.MassBuilderInfo
}

func NewMassBuilder(
	path string,
	sys *System,
	log *logger.Logger,
) (*MassBuildCtl, error) {
	self := new(MassBuildCtl)

	if path, err := filepath.Abs(path); err != nil {
		return nil, err
	} else {
		self.path = path
	}

	self.sys = sys

	self.log = logger.New()
	self.log.AddOutput(log)

	return self, nil
}

func (self *MassBuildCtl) ReadInfo() (*basictypes.MassBuilderInfo, error) {

	if self.info == nil {
		fullpath := path.Join(self.path, basictypes.MASSBUILDER_INFO_FILENAME)

		res, err := ioutil.ReadFile(fullpath)
		if err != nil {
			return nil, err
		}

		j_res := new(basictypes.MassBuilderInfo)

		err = json.Unmarshal(res, j_res)
		if err != nil {
			return nil, err
		}

		self.info = j_res
	}

	return self.info, nil
}

func (self *MassBuildCtl) WriteInfo(info *basictypes.MassBuilderInfo) error {

	fullpath := path.Join(self.path, basictypes.MASSBUILDER_INFO_FILENAME)

	res, err := json.Marshal(info)
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)

	err = json.Indent(b, res, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fullpath, b.Bytes(), 0700)
	if err != nil {
		return err
	}

	self.info = info

	return nil
}

func (self *MassBuildCtl) GetTarballsPath() string {
	//	return path.Join(self.path, basictypes.MASSBUILDER_TARBALLS_DIR)
	// NOTE: using self.path for tarballs storaging is better/handier
	//       for fast mass building init creation
	return self.path
}

func (self *MassBuildCtl) GetAspsPath() string {
	return path.Join(self.path, basictypes.MASSBUILDER_ASPS_DIR)
}

func (self *MassBuildCtl) GetDoneTarballsPath() string {
	return path.Join(self.path, basictypes.MASSBUILDER_DONE_TARBALLS)
}

// if len(tarballs) == 0  - all will be done
func (self *MassBuildCtl) PerformMassBuilding(tarballs []string) (
	map[string][]string,
	map[string][]string,
	error,
) {
	var err error

	info, err := self.ReadInfo()
	if err != nil {
		return nil, nil, err
	}

	err = os.MkdirAll(self.GetAspsPath(), 0700)
	if err != nil {
		return nil, nil, err
	}

	err = os.MkdirAll(self.GetDoneTarballsPath(), 0700)
	if err != nil {
		return nil, nil, err
	}

	host := info.Host
	archs := info.HostArchs

	if len(tarballs) == 0 {
		tarballs, err = self.tarballsList()
		if err != nil {
			return nil, nil, err
		}
	}

	{
		self.log.Info("searching for package determination problems...")
		package_determination_errors := make(map[string]error, 0)

		counter := 0

		for _, i := range tarballs {

			counter++

			fmt.Printf(" %.2f%%   \r", 100.0/(float64(len(tarballs))/float64(counter)))

			bi := path.Base(i)

			_, _, err := pkginfodb.DetermineTarballPackageInfoSingle(bi)
			if err != nil {
				package_determination_errors[bi] = err
			}

		}

		if len(package_determination_errors) != 0 {
			self.log.Error("package determination problems has been found")
			for k, v := range package_determination_errors {
				self.log.Error(fmt.Sprintf("   %s: %s\n", k, v))
			}
			return nil, nil, errors.New("discovered package determination problems")
		} else {
			self.log.Info("   not found")
		}
	}

	sret := make(map[string][]string)
	fret := make(map[string][]string)

	for _, i := range tarballs {

		bi := path.Base(i)

		_, pkginfo, err := pkginfodb.DetermineTarballPackageInfoSingle(bi)
		if err != nil {
			return nil, nil, err
		}

		if pkginfo.NonBuildable {
			continue
		}

		all_archs_succeeded := true

		for _, arch := range archs {

			if pkginfo.PrimaryInstallOnly && host != arch {
				continue
			}

			self.log.Info("-----//=********************--")
			self.log.Info("---<{[ building " + i + " for " + host + "-" + arch)
			self.log.Info(`-----\\=********************--`)
			res := self.fullBuildTarball(bi, host, arch)
			if res != nil {
				self.log.Error("building failed: " + res.Error())
				all_archs_succeeded = false
			}

			var vret map[string][]string

			if res == nil {
				vret = sret
			} else {
				vret = fret
			}

			if _, ok := vret[arch]; !ok {
				vret[arch] = make([]string, 0)
			}
			vret[arch] = append(vret[arch], bi)
			self.log.Info("")
		}

		if all_archs_succeeded {
			self.log.Info("moving succeeded tarball to separate dir")
			dtbp := self.GetDoneTarballsPath()
			tbm := path.Base(bi)
			err = os.Rename(bi, path.Join(dtbp, tbm))
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return sret, fret, nil
}

func (self *MassBuildCtl) checkAlreadyReady(
	pkgname, version,
	host, hostarch string,
) (bool, error) {

	files, err := ioutil.ReadDir(self.GetAspsPath())
	if err != nil {
		return false, err
	}

	for _, i := range files {
		if p, err := basictypes.NewASPNameFromString(i.Name()); err != nil {
			continue
		} else {
			if p.Name == pkgname &&
				p.Version == version &&
				p.Host == host &&
				p.HostArch == hostarch {
				return true, nil
			}
		}
	}

	return false, nil
}

func (self *MassBuildCtl) findBuildingSite(
	pkgname, version,
	host, hostarch string,
) (*BuildingSiteCtl, bool, error) {
	files, err := ioutil.ReadDir(self.path)
	if err != nil {
		return nil, false, err
	}

dirs_search:
	for _, i := range files {
		if i.IsDir() {
			{
				ib := path.Base(i.Name())
				for _, i := range []string{
					//					basictypes.MASSBUILDER_TARBALLS_DIR,
					basictypes.MASSBUILDER_ASPS_DIR,
				} {
					if ib == i {
						continue dirs_search
					}
				}
			}

			pth := path.Join(self.path, path.Base(i.Name()))
			nbs, err := NewBuildingSiteCtl(pth, self.sys, self.log)
			if err != nil {
				self.log.Error(err.Error())
				continue
			}

			if !nbs.IsBuildingSite() {
				continue
			}

			nbs_info, err := nbs.ReadInfo()
			if err != nil {
				self.log.Error(err.Error())
				continue
			}

			if nbs_info.PackageName == pkgname &&
				nbs_info.PackageVersion == version &&
				nbs_info.Host == host &&
				nbs_info.HostArch == hostarch {
				return nbs, true, nil
			}
		}
	}
	return nil, false, nil
}

func (self *MassBuildCtl) createBuildingSite(
	pkgname string,
	host, hostarch string,
	tarball_parsed *tarballname.ParsedTarballName,
) (*BuildingSiteCtl, error) {
	bs_ts := basictypes.NewASPTimeStampFromCurrentTime().String()

	vstr, err := tarball_parsed.Version.IntSliceString(".")
	if err != nil {
		return nil, err
	}

	bs_name := fmt.Sprintf(
		"%s-%s-%s",
		pkgname,
		vstr,
		bs_ts,
	)

	ret, err := NewBuildingSiteCtl(
		path.Join(self.path, bs_name),
		self.sys,
		self.log,
	)
	if err != nil {
		return nil, err
	}

	ret.InitDirs()
	if err != nil {
		return nil, err
	}

	mb_info, err := self.ReadInfo()
	if err != nil {
		return nil, err
	}

	new_bs_info := &basictypes.BuildingSiteInfo{
		Host:               host,
		HostArch:           hostarch,
		InitiatedByHost:    mb_info.InitiatedByHost,
		PackageName:        pkgname,
		PackageVersion:     vstr,
		PackageStatus:      tarball_parsed.Status.StrSliceString(""),
		PackageTimeStamp:   bs_ts,
		CrossbuilderTarget: mb_info.CrossbuilderTarget,
		CrossbuildersHost:  mb_info.CrossbuildersHost,
		TarballsDir:        self.GetTarballsPath(),
		AspsDir:            self.GetAspsPath(),
	}

	new_bs_info.SetInfoHorizon50()

	err = ret.WriteInfo(new_bs_info)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *MassBuildCtl) fullBuildTarball(tarballname, host, hostarch string) error {

	pkgname, pkginfo, err :=
		pkginfodb.DetermineTarballPackageInfoSingle(tarballname)
	if err != nil {
		return err
	}

	parser, err := tarballnameparsers.Get(pkginfo.TarballFileNameParser)
	if err != nil {
		return err
	}

	tarball_parsed, err := parser.Parse(tarballname)
	if err != nil {
		return err
	}

	vstr, err := tarball_parsed.Version.IntSliceString(".")
	if err != nil {
		return err
	}

	already_done := false

	if ok, err := self.checkAlreadyReady(
		pkgname,
		vstr,
		host, hostarch,
	); err != nil {
		return err
	} else {
		if ok {
			self.log.Info("  already done")
			already_done = true
		}
	}

	var bs *BuildingSiteCtl

	self.log.Info(
		"trying to find " + pkgname + "-" + vstr + "-" + host + "-" + hostarch,
	)

	if tbs, found, err := self.findBuildingSite(pkgname, vstr, host, hostarch); err != nil {
		self.log.Error("  finding error: " + err.Error())
		return err
	} else {

		if found {
			bs = tbs
			self.log.Info("  found existing bs: " + bs.path)
		} else {

			if !already_done {
				self.log.Info("  creating new bs")
				tbs, err := self.createBuildingSite(pkgname, host, hostarch, tarball_parsed)
				if err != nil {
					return err
				}
				self.log.Info("     " + tbs.GetPath())
				bs = tbs
			}
		}
	}

	if already_done {
		self.log.Info(
			"this bs's package already complete, so this bs going to be removed",
		)
	}

	if !already_done {

		self.log.Info("getting sources and patches..")

		err = bs.GetSources()
		if err != nil {
			return err
		}

		self.log.Info("getting action list..")

		bs_actions, err := bs.ListActions()
		if err != nil {
			return err
		}

		self.log.Info("running actions: " + strings.Join(bs_actions, ", ") + "..")

		err = bs.Run(bs_actions)
		if err != nil {
			return err
		}

		self.log.Info("run complete")

		already_done = true

	}

	if ok, err := self.checkAlreadyReady(pkgname, vstr, host, hostarch); err != nil {
		return err
	} else {
		if !ok {
			return errors.New("not built yet. but we need it to be. so this is error")
		}
	}

	if bs != nil && already_done {
		self.log.Info("removing bs which already done")
		err = os.RemoveAll(bs.GetPath())
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *MassBuildCtl) tarballsList() ([]string, error) {
	pth := self.GetTarballsPath()

	files, err := ioutil.ReadDir(pth)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0)

	for _, i := range files {
		if i.IsDir() {
			continue
		}

		if !tarballname.IsPossibleTarballName(i.Name()) {
			continue
		}

		ret = append(ret, path.Base(i.Name()))
	}

	return ret, nil
}
