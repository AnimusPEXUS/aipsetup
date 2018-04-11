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

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
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

func (self *MassBuildCtl) TarballsPath() string {
	return path.Join(self.path, "01.tarballs")
}

func (self *MassBuildCtl) PerformMassBuilding(pth string) (
	[]string, []string, error,
) {

	var err error

	files := make([]string, 0)

	pth, err = filepath.Abs(pth)
	if err != nil {
		return nil, nil, err
	}

	tarball_dir := path.Join(pth, basictypes.DIR_TARBALL)

	stats, err := ioutil.ReadDir(tarball_dir)
	if err != nil {
		return nil, nil, err
	}

	for _, i := range stats {
		if i.IsDir() {
			continue
		}
		if !tarballname.IsPossibleTarballName(i.Name()) {
			continue
		}
		files = append(files, i.Name())
	}

	built := make([]string, 0)
	notbuilt := make([]string, 0)

	for _, i := range files {

		istat, err := os.Stat(i)
		if err != nil {
			notbuilt = append(notbuilt, i)
			continue
		}

		if istat.IsDir() {
			notbuilt = append(notbuilt, i)
			continue
		}

		if !tarballname.IsPossibleTarballName(i) {
			notbuilt = append(notbuilt, i)
			continue
		}

		infoname, _, err := pkginfodb.DetermineTarballPackageInfoSingle(i)
		if err != nil {
			notbuilt = append(notbuilt, i)
			continue
		}

		new_timestamp := basictypes.NewASPTimeStampFromCurrentTime()

		building_site_dir := path.Join(
			self.path,
			fmt.Sprintf("%s-%s", infoname, new_timestamp.String()),
		)

		err = os.Mkdir(building_site_dir, 0700)
		if err != nil {
			notbuilt = append(notbuilt, i)
			continue
		}

		bsctl, err := NewBuildingSiteCtl(self.sys, building_site_dir, self.log)
		if err != nil {
			notbuilt = append(notbuilt, i)
			continue
		}

		new_bs_info := &basictypes.BuildingSiteInfo{
			PackageName:      infoname,
			PackageTimeStamp: new_timestamp.String(),
		}

		err = bsctl.WriteInfo(new_bs_info)
		if err != nil {
			notbuilt = append(notbuilt, i)
			continue
		}

		err = bsctl.PrepareToRun()
		if err != nil {
			notbuilt = append(notbuilt, i)
			continue
		}

	}

	if len(notbuilt) != 0 {
		err = errors.New("some packages not built")
	}

	return built, notbuilt, err
}
