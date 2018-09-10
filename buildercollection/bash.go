package buildercollection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/versionorstatus"
)

func init() {
	Index["bash"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_bash(bs)
	}
}

type Builder_bash struct {
	*Builder_std
}

func NewBuilder_bash(bs basictypes.BuildingSiteCtlI) (*Builder_bash, error) {
	self := new(Builder_bash)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_bash) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

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

func (self *Builder_bash) BuilderActionPatch(log *logger.Logger) error {

	bash_patch_name_regexp, err := regexp.Compile(`^bash(\d)(\d)-(\d+)$`)
	if err != nil {
		return err
	}

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	version_parsed := versionorstatus.NewParsedVersionOrStatusFromString(
		info.PackageVersion,
		".",
	)

	version_parsed_i_slice, err := version_parsed.IntSlice()
	if err != nil {
		return err
	}

	current_patch := 0

	if len(version_parsed_i_slice) > 2 {
		current_patch = version_parsed_i_slice[2]
	}

	current_patch++

	patches := make([]string, 0)

	all_patches_files, err := ioutil.ReadDir(self.GetBuildingSiteCtl().GetDIR_PATCHES())
	if err != nil {
		return err
	}

	for {
		found := false
		for _, i := range all_patches_files {
			if !bash_patch_name_regexp.MatchString(i.Name()) {
				continue
			}

			subs := bash_patch_name_regexp.FindStringSubmatch(i.Name())

			if len(subs) != 4 {
				continue
			}

			v1s := subs[1]
			v2s := subs[2]
			v3s := subs[3]

			v1, err := strconv.Atoi(v1s)
			if err != nil {
				return err
			}

			v2, err := strconv.Atoi(v2s)
			if err != nil {
				return err
			}

			v3, err := strconv.Atoi(v3s)
			if err != nil {
				return err
			}

			if v1 != version_parsed_i_slice[0] || v2 != version_parsed_i_slice[1] {
				continue
			}

			if v3 == current_patch {
				found = true
				patches = append(patches, i.Name())
			}
		}

		if !found {
			break
		}

		current_patch++
	}

	log.Info(fmt.Sprintf("Found applicabale patch names (%d)", len(patches)))
	for _, i := range patches {
		log.Info("  " + i)
	}

	for _, i := range patches {

		log.Info("Going to apply patch " + i)

		subs := bash_patch_name_regexp.FindStringSubmatch(i)

		if len(subs) != 4 {
			panic("programming error")
		}

		v3s := subs[3]

		v3, err := strconv.Atoi(v3s)
		if err != nil {
			return err
		}

		full_patch_file_path := path.Join(self.GetBuildingSiteCtl().GetDIR_PATCHES(), i)

		c := exec.Command("patch", "-i", full_patch_file_path, "-p0")
		c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
		c.Stderr = log.StderrLbl()
		c.Stdout = log.StdoutLbl()
		res := c.Run()
		if res != nil {
			return errors.New("patching error")
		}

		info.ModifyVersionBeforePack = true
		if t, err := versionorstatus.NewParsedVersionOrStatusFromIntSlice(
			[]int{
				version_parsed_i_slice[0],
				version_parsed_i_slice[1],
				v3,
			},
			".",
		).IntSliceString("."); err != nil {
			return err
		} else {
			info.NewVersion = t
		}

	}

	err = self.GetBuildingSiteCtl().WriteInfo(info)
	if err != nil {
		return err
	}

	if info.ModifyVersionBeforePack {
		log.Info("resulting package version will be modified to " + info.NewVersion)
	}

	return nil
}

func (self *Builder_bash) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = append(ret, "--enable-multibyte")

	if info.ThisIsCrossbuilder() || info.ThisIsCrossbuilding() {
		ret = append(ret, "--without-curses")
	} else {
		ret = append(ret, "--with-curses")
	}

	return ret, nil
}

func (self *Builder_bash) BuilderActionAfterDistribute(log *logger.Logger) error {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	tsl := path.Join(dst_install_prefix, "bin", "sh")

	err = os.Symlink("bash", tsl)
	if err != nil {
		return err
	}

	return nil
}
