package buildercollection

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
)

func init() {
	Index["expect"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_expect(bs), nil
	}
}

type Builder_expect struct {
	*Builder_std
}

func NewBuilder_expect(bs basictypes.BuildingSiteCtlI) *Builder_expect {

	self := new(Builder_expect)

	self.Builder_std = NewBuilder_std(bs)
	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.AfterExtractCB = self.AfterExtract

	return self
}

func (self *Builder_expect) AfterExtract(log *logger.Logger, ret error) error {
	log.Info("Looking for Tcl and Tk too extract them...")
	tcl_found := ""
	tk_found := ""

	tarb_dir := self.bs.GetDIR_TARBALL()

	tarb_dir_files, err := ioutil.ReadDir(tarb_dir)
	if err != nil {
		return err
	}

	for _, i := range tarb_dir_files {
		i_name := i.Name()
		if tarballname.IsPossibleTarballName(i_name) {
			if strings.HasPrefix(i_name, "tcl") {
				tcl_found = i_name
			}
			if strings.HasPrefix(i_name, "tk") {
				tk_found = i_name
			}

			if tcl_found != "" && tk_found != "" {
				break
			}
		}
	}

	a_tools := new(buildingtools.Autotools)

	for _, i := range [][2]string{
		[2]string{tcl_found, "tcl_tarball"},
		[2]string{tk_found, "tk_tarball"},
	} {
		err = a_tools.Extract(
			path.Join(tarb_dir, i[0]),
			self.bs.GetPath(),
			path.Join(self.bs.GetDIR_TEMP(), i[1]),
			false,
			false,
			false,
			log,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_expect) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--enable-threads",
			"--enable-wince",
		}...,
	)

	if strings.HasPrefix(info.HostArch, "x86_64") {
		ret = append(
			ret,
			[]string{
				"--enable-64bit",
				"--enable-64bit-vis",
			}...,
		)
	}

	return ret, nil
}
