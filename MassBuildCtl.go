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
	return path.Join(self.path, basictypes.MASSBUILDER_TARBALLS_DIR)
}

func (self *MassBuildCtl) GetAspsPath() string {
	return path.Join(self.path, basictypes.MASSBUILDER_ASPS_DIR)
}

func (self *MassBuildCtl) PerformMassBuilding() (
	map[string][]string,
	map[string][]string,
	error,
) {
	info, err := self.ReadInfo()
	if err != nil {
		return nil, nil, err
	}

	err = os.MkdirAll(self.GetAspsPath(), 0700)
	if err != nil {
		return nil, nil, err
	}

	host := info.Host
	archs := info.HostArchs

	tarballs, err := self.tarballsList()
	if err != nil {
		return nil, nil, err
	}

	sret := make(map[string][]string)
	fret := make(map[string][]string)

	for _, i := range tarballs {
		bi := path.Base(i)
		for _, arch := range archs {
			self.sys.log.Info("--//=********************--")
			self.log.Info("---<{[ building " + i + " for " + host + "-" + arch)
			self.sys.log.Info(`--\\=********************--`)
			res := self.fullBuildTarball(bi, host, arch)
			if res != nil {
				self.sys.log.Error("building failed: " + res.Error())
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
			self.sys.log.Info("")
		}
	}

	return sret, fret, nil
}

func (self *MassBuildCtl) checkAlreadyReady(
	pkgname, version,
	host, hostarch string,
) error {

	files, err := ioutil.ReadDir(self.GetAspsPath())
	if err != nil {
		return err
	}

	for _, i := range files {
		if p, err := basictypes.NewASPNameFromString(i.Name()); err != nil {
			if p.Name == pkgname &&
				p.Version == version &&
				p.Host == host &&
				p.HostArch == hostarch {
				return nil
			}
		}
	}

	return errors.New("not built yet")
}

func (self *MassBuildCtl) findBuildingSite(
	pkgname, version,
	host, hostarch string,
) (*BuildingSiteCtl, error) {
	files, err := ioutil.ReadDir(self.path)
	if err != nil {
		return nil, err
	}

dirs_search:
	for _, i := range files {
		if i.IsDir() {
			{
				ib := path.Base(i.Name())
				for _, i := range []string{
					basictypes.MASSBUILDER_TARBALLS_DIR,
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
				self.sys.log.Error(err.Error())
				continue
			}

			nbs_info, err := nbs.ReadInfo()
			if err != nil {
				self.sys.log.Error(err.Error())
				continue
			}

			if nbs_info.PackageName == pkgname &&
				nbs_info.PackageVersion == version &&
				nbs_info.Host == host &&
				nbs_info.HostArch == hostarch {
				return nbs, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (self *MassBuildCtl) createBuildingSite(
	pkgname string,
	host, hostarch string,
	tarball_parsed *tarballname.ParsedTarballName,
) (*BuildingSiteCtl, error) {
	bs_ts := basictypes.NewASPTimeStampFromCurrentTime().String()

	bs_name := fmt.Sprintf(
		"%s-%s-%s",
		pkgname,
		tarball_parsed.Version.Str,
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

	ret.Init()
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
		PackageStatus:      tarball_parsed.Status.Str,
		PackageVersion:     tarball_parsed.Version.Str,
		PackageTimeStamp:   bs_ts,
		CrossbuilderTarget: mb_info.CrossbuilderTarget,
		CrossbuildersHost:  mb_info.CrossbuildersHost,
		TarballsDir:        self.GetTarballsPath(),
		AspsDir:            self.GetAspsPath(),
	}

	new_bs_info.SetInfoLilith50()

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

	if self.checkAlreadyReady(pkgname, tarball_parsed.Version.Str, host, hostarch) == nil {
		return nil
	}

	var bs *BuildingSiteCtl

	self.log.Info(
		"trying to find " + pkgname + "-" + tarball_parsed.Version.Str + "-" +
			host + "-" + hostarch,
	)

	if tbs, err := self.findBuildingSite(
		pkgname, tarball_parsed.Version.Str, host, hostarch,
	); err != nil {
		self.log.Info("  finding error: " + err.Error())
		tbs, err := self.createBuildingSite(
			pkgname,
			host, hostarch,
			tarball_parsed,
		)
		if err != nil {
			return err
		}
		bs = tbs
	} else {
		bs = tbs
		self.log.Info("  using existing: " + bs.path)
	}

	self.log.Info("  getting sources and patches..")

	err = bs.GetSources()
	if err != nil {
		return err
	}

	self.log.Info("  getting action list..")

	bs_actions, err := bs.ListActions()
	if err != nil {
		return err
	}

	self.log.Info("  running actions: " + strings.Join(bs_actions, ", ") + "..")

	err = bs.Run(bs_actions)
	if err != nil {
		return err
	}

	self.log.Info("  run complete")

	if err := self.checkAlreadyReady(
		pkgname, tarball_parsed.Version.Str, host, hostarch,
	); err != nil {
		return err
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
		if tarballname.IsPossibleTarballName(i.Name()) {
			ret = append(ret, path.Base(i.Name()))
		}
	}

	return ret, nil
}
