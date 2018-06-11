package buildercollection

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["bzip2"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_bzip2(bs)
	}
}

type Builder_bzip2 struct {
	Builder_std

	fixed_CC     string
	fixed_AR     string
	fixed_RANLIB string
}

func NewBuilder_bzip2(bs basictypes.BuildingSiteCtlI) (*Builder_bzip2, error) {
	//        thr['CC'] = 'CC={}-gcc -m{}'.format(
	//            self.get_host_from_pkgi(),
	//            self.get_multilib_variant_int()
	//            )
	//        thr['AR'] = 'AR={}-gcc-ar'.format(self.get_host_from_pkgi())
	//        thr['RANLIB'] = 'RANLIB={}-gcc-ranlib'.format(
	//            self.get_host_from_pkgi()
	//            )

	self := new(Builder_bzip2)

	self.Builder_std = *NewBuilder_std(bs)

	calc := bs.GetBuildingSiteValuesCalculator()

	info, err := bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	mlv, err := calc.CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	self.fixed_CC = fmt.Sprintf("CC=%s-gcc -m%s", info.Host, mlv)
	self.fixed_AR = fmt.Sprintf("AR=%s-gcc-ar", info.Host)
	self.fixed_RANLIB = fmt.Sprintf("RANLIB=%s-gcc-ranlib", info.Host)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_bzip2) EditActions(ret basictypes.BuilderActions) (
	basictypes.BuilderActions,
	error,
) {

	//            ('so', self.builder_action_so),
	//            ('copy_so', self.builder_action_copy_so),
	//            ('fix_links', self.builder_action_fix_links),
	//            ('fix_libdir_positions', self.builder_action_fix_libdir_positions)

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")
	ret = ret.Remove("distribute")

	new_actions := basictypes.BuilderActions{
		&basictypes.BuilderAction{"build", self.BuilderActionBuild},
		&basictypes.BuilderAction{"distribute", self.BuilderActionDistribute},
		&basictypes.BuilderAction{"so", self.BuilderActionSO},
		&basictypes.BuilderAction{"copy_so", self.BuilderActionCopySo},
		&basictypes.BuilderAction{"fix_links", self.BuilderActionFixLinks},
		&basictypes.BuilderAction{"fix_libdir_name", self.BuilderActionFixLibdirName},
		&basictypes.BuilderAction{"fix_mandir_position", self.BuilderActionFixMandirPosition},
	}

	ret, err := ret.AddActionsBeforeName(new_actions, "prepack")
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_bzip2) BuilderActionBuild(log *logger.Logger) error {

	//        if self.get_host_from_pkgi() != 'x86_64-pc-linux-gnu':
	//            raise Exception("fix for others is required")

	//        ret = autotools.make_high(
	//            self.buildingsite_path,
	//            log=log,
	//            options=[],
	//            arguments=[
	//                'PREFIX={}'.format(self.calculate_install_prefix()),
	//                'CFLAGS=  -fpic -fPIC -Wall -Winline -O2 -g '
	//                '-D_FILE_OFFSET_BITS=64',
	//                'libbz2.a',
	//                'bzip2',
	//                'bzip2recover'
	//                ] + [self.custom_data['thr']['CC']] +
	//            [self.custom_data['thr']['AR']] +
	//            [self.custom_data['thr']['RANLIB']],
	//            environment={},
	//            environment_mode='copy',
	//            use_separate_buildding_dir=self.separate_build_dir,
	//            source_configure_reldir=self.source_configure_reldir
	//            )

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return err
	}

	args := []string{
		"PREFIX=" + install_prefix,
	}

	args = append(args, self.flags()...)

	args = append(
		args,
		[]string{
			"libbz2.a",
			"bzip2",
			"bzip2recover",
		}...,
	)

	err = buildingtools.Autotools{}.Make(
		args,
		[]string{},
		buildingtools.Copy,
		"",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_bzip2) BuilderActionDistribute(log *logger.Logger) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	att := &buildingtools.Autotools{}

	err = att.Make(
		[]string{
			"PREFIX=" + install_prefix,
			"install",
		},
		[]string{},
		buildingtools.Copy,
		"",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_bzip2) flags() []string {
	ret := []string{
		"CFLAGS= -fpic -fPIC -Wall -Winline -O2 -g " +
			"-D_FILE_OFFSET_BITS=64",
		self.fixed_CC,
		self.fixed_AR,
		self.fixed_RANLIB,
	}
	return ret
}

func (self *Builder_bzip2) BuilderActionSO(log *logger.Logger) error {
	//        ret = autotools.make_high(
	//            self.buildingsite_path,
	//            log=log,
	//            options=[],
	//            arguments=[
	//                'CFLAGS= -fpic -fPIC -Wall -Winline -O2 -g '
	//                '-D_FILE_OFFSET_BITS=64',
	//                'PREFIX={}'.format(self.calculate_dst_install_prefix())
	//                ] + [self.custom_data['thr']['CC']] +
	//            [self.custom_data['thr']['AR']] +
	//            [self.custom_data['thr']['RANLIB']],
	//            environment={},
	//            environment_mode='copy',
	//            use_separate_buildding_dir=self.separate_build_dir,
	//            source_configure_reldir=self.source_configure_reldir,
	//            make_filename='Makefile-libbz2_so'
	//            )

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return err
	}

	args := []string{
		"PREFIX=" + install_prefix,
	}

	args = append(args, self.flags()...)

	err = buildingtools.Autotools{}.Make(
		args,
		[]string{},
		buildingtools.Copy,
		"Makefile-libbz2_so",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"",
		log,
	)
	if err != nil {
		return err
	}
	return nil
}

func (self *Builder_bzip2) BuilderActionCopySo(log *logger.Logger) error {
	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	di := path.Join(install_prefix, "lib")

	sos, err := filepath.Glob(
		strings.Join(
			[]string{self.bs.GetDIR_SOURCE(), "*.so*"},
			"/",
		),
	)
	if err != nil {
		return err
	}

	for _, i := range sos {

		base := path.Base(i)

		j := path.Join(self.bs.GetDIR_SOURCE(), base)
		j2 := path.Join(di, base)

		err = filetools.CopyWithInfo(j, j2, log)
		if err != nil {
			return err
		}

	}

	return nil
}

func (self *Builder_bzip2) BuilderActionFixLinks(log *logger.Logger) error {
	calc := self.bs.GetBuildingSiteValuesCalculator()

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_bin := path.Join(dst_install_prefix, "bin")

	dst_install_prefix_bin_files, err := ioutil.ReadDir(dst_install_prefix_bin)
	if err != nil {
		return err
	}

	for _, i := range dst_install_prefix_bin_files {
		fp := path.Join(dst_install_prefix_bin, i.Name())
		fp_s, err := os.Lstat(fp)
		if err != nil {
			return err
		}

		if !filetools.Is(fp_s.Mode()).Symlink() {
			continue
		}

		fp_val, err := os.Readlink(fp)
		if err != nil {
			return err
		}

		err = os.Remove(fp)
		if err != nil {
			return err
		}

		err = os.Symlink(path.Base(fp_val), fp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_bzip2) BuilderActionFixLibdirName(log *logger.Logger) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_lib := path.Join(dst_install_prefix, "lib")
	dst_install_prefix_lib64 := path.Join(dst_install_prefix, "lib64")

	//	dst_install_prefix_lib_s, err:= os.Stat(dst_install_prefix_lib )
	//	if err != nil {
	//		return err
	//	}

	//	dst_install_prefix_lib64_s, err:= os.Stat(dst_install_prefix_lib64 )
	//	if err != nil {
	//		return err
	//	}

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	if info.HostArch == "x86_64-pc-linux-gnu" {
		log.Info(
			fmt.Sprintf(
				"renaming %s to %s",
				dst_install_prefix_lib,
				dst_install_prefix_lib64,
			),
		)
		err = os.Rename(dst_install_prefix_lib, dst_install_prefix_lib64)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_bzip2) BuilderActionFixMandirPosition(log *logger.Logger) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_man := path.Join(dst_install_prefix, "man")
	dst_install_prefix_share := path.Join(dst_install_prefix, "share")
	dst_install_prefix_share_man := path.Join(dst_install_prefix_share, "man")

	log.Info(
		fmt.Sprintf(
			"renaming %s to %s",
			dst_install_prefix_man,
			dst_install_prefix_share_man,
		),
	)

	err = os.Mkdir(dst_install_prefix_share, 0700)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	err = os.Rename(dst_install_prefix_man, dst_install_prefix_share_man)
	if err != nil {
		return err
	}

	return nil
}
