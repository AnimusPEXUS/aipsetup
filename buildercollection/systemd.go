package buildercollection

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"path"
	"sort"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["systemd"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_systemd(bs)
	}
}

type Builder_systemd struct {
	*Builder_std_meson
}

func NewBuilder_systemd(bs basictypes.BuildingSiteCtlI) (*Builder_systemd, error) {

	self := new(Builder_systemd)

	if t, err := NewBuilder_std_meson(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_meson = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_systemd) BuilderActionPatch(
	log *logger.Logger,
) error {
	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageName != "systemd" || info.PackageVersion != "239" {
		return errors.New("this builder is for systemd-239 only")
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
		return errors.New("systemd-239 requires patches")
	} else {
		// TODO: have to do this smarter
		sort.Strings(ptch_dir_files_ls)
	}

	pth_name := ptch_dir_files_ls[len(ptch_dir_files_ls)-1]

	cmd := exec.Command("patch", "-p1", path.Join(ptch_dir, pth_name))
	cmd.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_systemd) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-Ddefault_library=both", // TODO: promote this option to std_meson?

		}...,
	)

	return ret, nil
}
