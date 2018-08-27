package buildercollection

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["autoconf2.13"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_autoconf213(bs)
	}
}

type Builder_autoconf213 struct {
	*Builder_std
}

func NewBuilder_autoconf213(bs basictypes.BuildingSiteCtlI) (*Builder_autoconf213, error) {

	self := new(Builder_autoconf213)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	//	self.EditDistributeDESTDIRCB = self.EditDistributeDESTDIR

	self.EditDistributeArgsCB = self.EditDistributeArgs

	//	self.AfterDistributeCB = self.AfterDistribute

	return self, nil
}

func (self *Builder_autoconf213) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"extract",
		"patch", self.BuilderActionPatch,
	)
	if err != nil {
		return nil, err
	}

	ret, err = ret.AddActionAfterNameShort(
		"distribute",
		"after-distribute", self.BuilderActionAfterDistribute,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_autoconf213) BuilderActionPatch(
	log *logger.Logger,
) error {
	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageName != "autoconf2.13" {
		return errors.New("this builder is for autoconf-2.13 only")
	}

	ptch_dir := self.GetBuildingSiteCtl().GetDIR_PATCHES()
	ptch_dir_files, err := ioutil.ReadDir(ptch_dir)
	if err != nil {
		return err
	}

	ptch_dir_files_ls := make([]string, 0)

	for _, i := range ptch_dir_files {
		ptch_dir_files_ls = append(ptch_dir_files_ls, i.Name())
	}

	if len(ptch_dir_files_ls) == 0 {
		return errors.New("autoconf 2.13 requires patches")
	} else {
		// TODO: have to do this smarter
		sort.Strings(ptch_dir_files_ls)
	}

	pth_name := ptch_dir_files_ls[len(ptch_dir_files_ls)-1]

	cmd := exec.Command("patch", "-p1", "-i", path.Join(ptch_dir, pth_name))
	cmd.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_autoconf213) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--program-suffix=2.13",
			"--datadir=" + path.Join(install_prefix, "share"),
			"--infodir=" + path.Join(install_prefix, "share", "info"),
		}...,
	)

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{"--docdir=", "CC=", "CXX=", "GCC="},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_autoconf213) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	//	ret = []string{}

	for i := len(ret) - 1; i != -1; i -= 1 {
		if strings.HasPrefix(ret[i], "prefix=") ||
			strings.HasPrefix(ret[i], "DESTDIR=") {
			ret = append(ret[:i], ret[i+1:]...)
		}
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"prefix=" + dst_install_prefix,
			"bindir=" + path.Join(dst_install_prefix, "bin"),
			"infodir=" + path.Join(dst_install_prefix, "share", "info"),
			"acdatadir=" + path.Join(dst_install_prefix, "share", "autoconf2.13"),
		}...,
	)

	return ret, nil
}

func (self *Builder_autoconf213) BuilderActionAfterDistribute(log *logger.Logger) error {

	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	share_info := path.Join(dst_install_prefix, "share", "info")

	err = os.RemoveAll(share_info)
	if err != nil {
		return err
	}

	return nil
}
