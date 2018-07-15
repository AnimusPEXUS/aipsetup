package buildercollection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/archive"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ncurses"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_ncurses(bs)
	}
}

var (
	ROLLING_PATCH_RE_C  = regexp.MustCompile(`ncurses-(\d+.\d+)-(\d+)-patch\.sh(.*)`)
	ORDINARY_PATCH_RE_C = regexp.MustCompile(`ncurses-(\d+.\d+)-(\d+).patch(.*)`)
)

type Builder_ncurses struct {
	*Builder_std
}

func NewBuilder_ncurses(bs basictypes.BuildingSiteCtlI) (*Builder_ncurses, error) {
	self := new(Builder_ncurses)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions
	self.PatchCB = self.Patch
	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_ncurses) EditActions(ret basictypes.BuilderActions) (
	basictypes.BuilderActions,
	error,
) {
	var err error

	ret, err = ret.AddActionsAfterName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "make_lib_symlinks",
				Callable: self.AfterDistributeLinks,
			},
		},
		"distribute",
	)
	if err != nil {
		return nil, err
	}

	ret, err = ret.AddActionsAfterName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "make_pkgconfig_symlinks",
				Callable: self.AfterDistributePkgConfig,
			},
		},
		"make_lib_symlinks",
	)
	if err != nil {
		return nil, err
	}

	ret, err = ret.AddActionsAfterName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "make_ln_ncurses_to_ncursesw",
				Callable: self.AfterDistributeNcursesLnNcursesw,
			},
		},
		"make_pkgconfig_symlinks",
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_ncurses) Patch(log *logger.Logger) error {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	ver_str := info.PackageVersion

	patches_dir := path.Join(self.bs.GetDIR_PATCHES(), ver_str)

	{
		_, err := os.Stat(patches_dir)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			} else {
				// considering user does not want patching, or there is no patches
				// yet
				return nil
			}
		}
	}

	file_list := make([]string, 0)

	{
		stat_lst, err := ioutil.ReadDir(patches_dir)
		if err != nil {
			return err
		}

		for _, i := range stat_lst {
			if i.IsDir() {
				continue
			}

			if !strings.HasSuffix(i.Name(), ".patch.gz") {
				continue
			}

			file_list = append(file_list, i.Name())
		}
	}

	sort.Strings(file_list)

	rolling_index := -1
	rolling_name := ""

	for i := len(file_list) - 1; i != -1; i = i - 1 {
		filename := file_list[i]

		if ROLLING_PATCH_RE_C.MatchString(filename) {
			sub := ROLLING_PATCH_RE_C.FindStringSubmatch(filename)
			if sub[1] == ver_str {
				rolling_index = i
				rolling_name = filename
				break
			}
		}
	}

	if rolling_index != -1 {
		file_list = file_list[rolling_index+1:]
	}

	for i := len(file_list) - 1; i != -1; i = i - 1 {
		filename := file_list[i]
		ok := true

		if !ORDINARY_PATCH_RE_C.MatchString(filename) {
			ok = false
		} else {
			sub := ORDINARY_PATCH_RE_C.FindStringSubmatch(filename)
			if sub[1] != ver_str {
				ok = false
			}
		}

		if !ok {
			file_list = append(file_list[:i], file_list[i+1:]...)
		}
	}

	log.Info("Patches to apply:")
	if rolling_name == "" {
		log.Info("  No rolling patch")
	} else {
		log.Info("  Rolling:  " + rolling_name)
	}

	for _, i := range file_list {
		log.Info("  Ordinary: " + i)
	}

	if rolling_name != "" {
		log.Info("Patching with " + rolling_name)

		ok, compressor := archive.DetermineCompressorByFilename(rolling_name)
		if !ok {
			return errors.New("not supported compressor of rolling patch")
		}

		// TODO: use native code
		c := exec.Command(compressor, "-kfd", rolling_name)
		c.Dir = patches_dir

		err = c.Run()
		if err != nil {
			return err
		}

		ok, ext := archive.DetermineExtensionByFilename(rolling_name)
		if !ok {
			return errors.New("error determining filename extension")
		}

		rolling_name = rolling_name[:len(rolling_name)-len(ext)]

		rolling_full_path := path.Join(patches_dir, rolling_name)

		c = exec.Command("bash", rolling_full_path)
		c.Dir = self.bs.GetDIR_SOURCE()
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()

		err = c.Run()
		if err != nil {
			return err
		}

		log.Info("")
	}

	for _, i := range file_list {
		log.Info("Patching with " + i)

		ok, compressor := archive.DetermineCompressorByFilename(i)
		if !ok {
			return errors.New("not supported compressor of ordinary patch")
		}

		// TODO: use native code
		c := exec.Command(compressor, "-kfd", i)
		c.Dir = patches_dir

		err = c.Run()
		if err != nil {
			return err
		}

		ok, ext := archive.DetermineExtensionByFilename(i)
		if !ok {
			return errors.New("error determining filename extension")
		}

		i = i[:len(i)-len(ext)]

		i_full_path := path.Join(patches_dir, i)

		c = exec.Command("patch", "-p1", "-i", i_full_path)
		c.Dir = self.bs.GetDIR_SOURCE()
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()

		err = c.Run()
		if err != nil {
			return err
		}

		log.Info("")
	}

	return nil
}

func (self *Builder_ncurses) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	gcc_line, err := calc.CalculateAutotoolsCCParameterValue()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--enable-shared",
			"--enable-widec",
			"--enable-const",
			"--enable-ext-colors",
			"--enable-pc-files",
			"--with-shared",
			"--with-gpm",
			"--with-ticlib",
			"--with-termlib",
			"--with-pkg-config=" + path.Join(install_prefix, "share", "pkgconfig"),

			"--with-ada",
			// TODO: NOTE: building with ada fails on new installations
			// "--without-ada",

			fmt.Sprintf("ADAFLAGS= --GCC=\"%s\" ", gcc_line),
		}...,
	)

	if info.ThisIsCrossbuilder() || info.ThisIsCrossbuilding() {
		ret = append(ret, "--without-ada")
	}

	return ret, nil
}

func (self *Builder_ncurses) _AfterDistributeLinksCreator(
	log *logger.Logger,
	pth string,
	what, from, to string,
) error {
	files, err := filepath.Glob(path.Join(pth, what))
	if err != nil {
		return nil
	}

	for _, i := range files {
		o_name := path.Base(i)
		l_name := strings.Replace(o_name, from, to, -1)

		rrr := path.Join(pth, l_name)

		_, err := os.Lstat(rrr)
		if err == nil {
			continue
		} else {
			if !os.IsNotExist(err) {
				return err
			}
		}

		log.Info("  " + o_name + " -> " + l_name)
		err = os.Symlink(o_name, rrr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *Builder_ncurses) AfterDistributeLinks(log *logger.Logger) error {

	log.Info("Going to make lib symlinks")

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dst_lib_dir, err := calc.CalculateDstInstallLibDir()
	if err != nil {
		return err
	}

	for _, s := range [][3]string{
		[3]string{"*w.so*", "w.so", ".so"},
		[3]string{"*w_g.so*", "w_g.so", "_g.so"},
		[3]string{"*w.a*", "w.a", ".a"},
		[3]string{"*w_g.a*", "w_g.a", "_g.a"},
	} {

		err = self._AfterDistributeLinksCreator(
			log,
			dst_lib_dir,
			s[0], s[1], s[2],
		)
		if err != nil {
			return err
		}

	}

	return nil

}

func (self *Builder_ncurses) AfterDistributePkgConfig(log *logger.Logger) error {

	log.Info("Going to make pkgconfig symlinks")

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dst_lib_dir, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_share_dir := path.Join(dst_lib_dir, "share")

	dst_pc_lib_dir := path.Join(dst_share_dir, "pkgconfig")

	err = self._AfterDistributeLinksCreator(
		log,
		dst_pc_lib_dir,
		"*w.pc", "w.pc", ".pc",
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_ncurses) AfterDistributeNcursesLnNcursesw(log *logger.Logger) error {

	log.Info("Making headers ncurses symlink to ncursesw")

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	ncurses := path.Join(dst_install_prefix, "include", "ncurses")
	ncursesw := path.Join(dst_install_prefix, "include", "ncursesw")

	err = os.Symlink(ncursesw, ncurses)
	if err != nil {
		return err
	}

	return nil
}
