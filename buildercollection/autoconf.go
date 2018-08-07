package buildercollection

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["autoconf"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_autoconf(bs)
	}
}

type Builder_autoconf struct {
	*Builder_std
}

func NewBuilder_autoconf(bs basictypes.BuildingSiteCtlI) (*Builder_autoconf, error) {

	self := new(Builder_autoconf)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditDistributeDESTDIRCB = self.EditDistributeDESTDIR

	return self, nil
}

func (self *Builder_autoconf) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	err := ret.Replace(
		"patch",
		&basictypes.BuilderAction{
			Name:     "patch",
			Callable: self.BuilderActionPatch,
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_autoconf) BuilderActionPatch(
	log *logger.Logger,
) error {
	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageName == "autoconf" && info.PackageVersion == "2.13" {
		ptch_dir := self.bs.GetDIR_PATCHES()
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

		cmd := exec.Command("patch", "-p1", path.Join(ptch_dir, pth_name))
		cmd.Dir = self.bs.GetDIR_SOURCE()
		cmd.Stdout = log.StdoutLbl()
		cmd.Stderr = log.StderrLbl()

		err = cmd.Run()
		if err != nil {
			return err
		}

	}

	return nil
}

func (self *Builder_autoconf) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	dst_dir := self.bs.GetDIR_DESTDIR()

	ret = append(
		ret,
		[]string{
			"--program-suffix=2.13",
			"--datadir=" + path.Join(dst_dir, install_prefix, "share"),
			"--infodir=" + path.Join(dst_dir, install_prefix, "share", "info"),
		}...,
	)

	for i := len(ret) - 1; i != -1; i -= 1 {
		if strings.HasPrefix(ret[i], "--docdir=") ||
			strings.HasPrefix(ret[i], "CC=") ||
			strings.HasPrefix(ret[i], "GCC=") ||
			strings.HasPrefix(ret[i], "CXX=") {
			ret = append(ret[:i], ret[i+1:]...)
		}
	}

	return ret, nil
}

func (self *Builder_autoconf) EditDistributeDESTDIR(log *logger.Logger, ret string) (string, error) {
	return "prefix", nil
}
