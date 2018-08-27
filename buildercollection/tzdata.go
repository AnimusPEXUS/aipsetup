package buildercollection

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["tzdata"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_tzdata(bs)
	}
}

type Builder_tzdata struct {
	*Builder_std
}

func NewBuilder_tzdata(bs basictypes.BuildingSiteCtlI) (*Builder_tzdata, error) {

	self := new(Builder_tzdata)

	self.Builder_std = NewBuilder_std(bs)

	self.EditExtractMoreThanOneExtractedOkCB = func(log *logger.Logger, ret bool) (bool, error) {
		return true, nil
	}

	self.EditExtractUnwrapCB = func(log *logger.Logger, ret bool) (bool, error) {
		return false, nil
	}

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_tzdata) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("build")

	err := ret.ReplaceShort("configure", self.BuilderActionConfigure)
	if err != nil {
		return nil, err
	}

	err = ret.ReplaceShort("distribute", self.BuilderActionDistribute)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_tzdata) BuilderActionConfigure(log *logger.Logger) error {

	makefile := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "Makefile")

	data, err := ioutil.ReadFile(makefile)
	if err != nil {
		return err
	}

	data_s := string(data)

	data_s += "\nprinttdata:\n\t\t@echo \"$(TDATA)\"\n"

	err = ioutil.WriteFile(makefile, []byte(data_s), 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_tzdata) BuilderActionDistribute(log *logger.Logger) error {

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	zoneinfo_dir := path.Join(dst_install_prefix, "share", "zoneinfo")
	zoneinfo_dir_posix := path.Join(zoneinfo_dir, "posix")
	zoneinfo_dir_right := path.Join(zoneinfo_dir, "right")

	for _, i := range []string{zoneinfo_dir, zoneinfo_dir_posix, zoneinfo_dir_right} {
		err := os.MkdirAll(i, 0700)
		if err != nil {
			return err
		}
	}

	c := exec.Command("make", "printtdata")
	c.Stderr = log.StderrLbl()
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()

	data, err := c.Output()
	if err != nil {
		return err
	}

	printed_data := string(data)

	zone_files := strings.Split(printed_data, " ")

	for i := len(zone_files) - 1; i != -1; i-- {
		zone_files[i] = strings.Trim(zone_files[i], "\n\t\x00\r")
	}

	sort.Strings(zone_files)

	log.Info("ZF: " + strings.Join(zone_files, " "))

	for _, tz := range zone_files {

		log.Info(fmt.Sprintf("   '%s'", tz))

		c := exec.Command(
			"zic",
			"-L", "/dev/null",
			"-d", zoneinfo_dir,
			"-y", "sh yearistype.sh", tz,
		)
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()
		c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
		err := c.Run()
		if err != nil {
			return err
		}

		c = exec.Command(
			"zic",
			"-L", "/dev/null",
			"-d", zoneinfo_dir_posix,
			"-y", "sh yearistype.sh", tz,
		)
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()
		c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
		err = c.Run()
		if err != nil {
			return err
		}

		c = exec.Command(
			"zic",
			"-L", "leapseconds",
			"-d", zoneinfo_dir_right,
			"-y", "sh yearistype.sh", tz,
		)
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()
		c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
		err = c.Run()
		if err != nil {
			return err
		}

	}

	return nil

}
