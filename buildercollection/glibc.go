package buildercollection

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["glibc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilderGlibc(bs)
	}
}

type BuilderGlibc struct {
	bs basictypes.BuildingSiteCtlI

	std_builder *BuilderStdAutotools

	slibdir string
}

func NewBuilderGlibc(bs basictypes.BuildingSiteCtlI) (*BuilderGlibc, error) {

	self := new(BuilderGlibc)
	self.bs = bs

	self.std_builder = NewBuilderStdAutotools(bs)

	self.std_builder.SeparateBuildDir = true
	self.std_builder.ForcedTarget = true
	self.std_builder.ApplyHostSpecCompilerOptions = true

	self.std_builder.EditConfigureArgsCB = self.EditConfigureArgs

	calc := self.bs.GetBuildingSiteValuesCalculator()

	if t, err := calc.CalculateInstallLibDir(); err != nil {
		return nil, err
	} else {
		self.slibdir = t
	}

	self.std_builder.EditBuildArgsCB = self.EditBuildArgs
	self.std_builder.EditDistributeArgsCB = self.EditDistributeArgs

	return self, nil
}

func (self *BuilderGlibc) DefineActions() (basictypes.BuilderActions, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret, err := self.std_builder.DefineActions()
	if err != nil {
		return nil, err
	}

	if info.ThisIsCrossbuilder {
		ret = ret.Remove("build")
		ret = ret.Remove("distribute")

		ret = ret.AddAfter(
			basictypes.BuilderActions{
				&basictypes.BuilderAction{"distribute_01", self.BuilderActionDistribute_01},
				&basictypes.BuilderAction{"distribute_01_2", self.BuilderActionDistribute_01_2},
				&basictypes.BuilderAction{"distribute_01_3", self.BuilderActionDistribute_01_3},
				&basictypes.BuilderAction{"distribute_01_4", self.BuilderActionDistribute_01_4},
				&basictypes.BuilderAction{"distribute_01_5", self.BuilderActionDistribute_01_5},

				&basictypes.BuilderAction{"intermediate_instruction", self.BuilderActionIntermediateInstruction},

				&basictypes.BuilderAction{"build_02", self.BuilderActionBuild_02},
				&basictypes.BuilderAction{"distribute_02", self.BuilderActionDistribute_02},
			},
			len(ret)-1,
		)
	}

	return ret, nil
}

func (self *BuilderGlibc) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	var with_headers string

	if t, err := calc.CalculateHostDir(); err != nil {
		return nil, err
	} else {
		with_headers = path.Join(t, "include")
	}

	if info.ThisIsCrossbuilder {

		host_builders_dir, err := calc.CalculateHostCrossbuildersDir()
		if err != nil {
			return nil, err
		}

		prefix := path.Join(host_builders_dir, info.CrossbuilderTarget)

		with_headers = path.Join(prefix, "include")

		hbt_opts, err := calc.CalculateAutotoolsHBTOptions()
		if err != nil {
			return nil, err
		}

		ret = make([]string, 0)
		ret = append(
			ret,
			[]string{
				"--prefix=" + prefix,
				"--mandir=" + path.Join(prefix, "share", "man"),
				"--sysconfdir=/etc",
				"--localstatedir=/var",
				"--enable-shared",
			}...,
		)
		ret = append(
			ret,
			hbt_opts...,
		)

	}

	ret = append(
		ret,
		[]string{

			"--enable-obsolete-rpc",
			"--enable-kernel=4.9",
			"--enable-tls",
			"--with-elf",
			// # disabled those 3 items on 2 jul 2015
			// # reenabled those 3 items on 11 aug 2015: sims I need it
			"--enable-multi-arch",
			"--enable-multiarch",
			"--enable-multilib",

			// # this is from configure --help. configure looking for
			// # linux/version.h file
			// #"--with-headers=/usr/src/linux/include",
			"--with-headers=" + with_headers,
			"--enable-shared",

			// # temp
			// #"libc_cv_forced_unwind=yes"
		}...,
	)

	// """
	// # NOTE: it's not working
	// # NOTE: don"t remove this block. it"s for informational reason
	// if self.get_arch_from_pkgi().startswith("x86_64"):
	//     ret += ["slibdir=lib64"]
	// else:
	//     ret += ["slibdir=lib"]
	// """

	// NOTE: this block is not found in pythonish aipsetup clibc builder, and
	//       is only binutils.go builder copy result. probably can be removed
	//       without any consiquences.
	// if cb {
	// 	ret = append(ret, "--with-sysroot")
	// }

	if info.ThisIsCrossbuilder {
		ret = append(
			ret,
			[]string{
			// TODO: 2018 mar 14. it was easy to comment/uncomment when it was
			//       python script, but from now on this should be done somehow
			//       wiser. to be done..

			// # this can be commented whan gcc fully built and installed
			// #"libc_cv_forced_unwind=yes",
			//
			// # this parameter is required to build `build_02+"
			// # stage.  comment and completle rebuild this glibc
			// # after rebuilding gcc without --without-headers and
			// # with --with-sysroot parameter.
			// #
			// # "libc_cv_ssp=no"
			// #
			// # else You will see errors like this:
			// #     gethnamaddr.c:185: undefined reference to
			// #     `__stack_chk_guard"

			}...,
		)
	}

	return ret, nil
}

func (self *BuilderGlibc) EditBuildArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(ret, "slibdir="+self.slibdir)
	return ret, nil
}

func (self *BuilderGlibc) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(ret, "slibdir="+self.slibdir)
	return ret, nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01(
	log *logger.Logger,
) error {

	self.std_builder.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{
			"install-bootstrap-headers=yes",
			"install-headers",
			"DESTDIR=" + self.bs.GetDIR_DESTDIR(),
		}, nil
	}

	return self.std_builder.BuilderActionBuild(log)
}

func (self *BuilderGlibc) BuilderActionDistribute_01_2(
	log *logger.Logger,
) error {

	self.std_builder.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{
			"csu/subdir_lib",
		}, nil
	}

	return self.std_builder.BuilderActionBuild(log)
}

func (self *BuilderGlibc) BuilderActionDistribute_01_3(
	log *logger.Logger,
) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dhcd, err := calc.CalculateDstHostCrossbuildersDir()
	if err != nil {
		return err
	}

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	mmldn, err := calc.CalculateMainMultiarchLibDirName()
	if err != nil {
		return err
	}

	gres, err := filepath.Glob(
		path.Join(self.bs.GetDIR_BUILDING(), "csu", "*crt*.o"),
	)
	if err != nil {
		return err
	}

	dest_lib_dir := path.Join(dhcd, info.CrossbuilderTarget, mmldn)

	err = os.MkdirAll(dest_lib_dir, 0755)
	if err != nil {
		return err
	}

	for _, i := range gres {
		fromf, err := os.Open(i)
		if err != nil {
			return err
		}
		tof, err := os.Create(path.Join(dest_lib_dir, path.Base(i)))
		if err != nil {
			fromf.Close()
			return err
		}
		_, err = io.Copy(tof, fromf)

		fromf.Close()
		tof.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01_4(
	log *logger.Logger,
) error {
	calc := self.bs.GetBuildingSiteValuesCalculator()

	dhcd, err := calc.CalculateDstHostCrossbuildersDir()
	if err != nil {
		return err
	}

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	mmldn, err := calc.CalculateMainMultiarchLibDirName()
	if err != nil {
		return err
	}

	cwd := path.Join(dhcd, info.CrossbuilderTarget, mmldn)

	cmd := []string{
		"-nostdlib",
		"-nostartfiles",
		"-shared",
		"-x",
		"c",
		"/dev/null",
		"-o",
		"libc.so",
	}

	log.Info("directory: " + cwd)
	log.Info("cmd: " + strings.Join(cmd, "/"))

	c := exec.Command(info.CrossbuilderTarget+"-gcc", cmd...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Dir = cwd

	err = c.Run()

	return nil
}

func (self *BuilderGlibc) BuilderActionDistribute_01_5(
	log *logger.Logger,
) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	dhcd, err := calc.CalculateDstHostCrossbuildersDir()
	if err != nil {
		return err
	}

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	cwd := path.Join(dhcd, info.CrossbuilderTarget, "include", "gnu")

	cwdf := path.Join(cwd, "stubs.h")

	err = os.MkdirAll(cwd, 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(cwdf)
	if err != nil {
		return err
	}

	f.Close()

	return nil
}

func (self *BuilderGlibc) BuilderActionIntermediateInstruction(
	log *logger.Logger,
) error {
	for _, i := range []string{
		"---------------",
		"pack and install this glibc build.",
		"then continue with gcc build_02+",
		"---------------",
	} {
		log.Info(i)
	}
	return errors.New("user action required")
}

func (self *BuilderGlibc) BuilderActionBuild_02(
	log *logger.Logger,
) error {

	self.std_builder.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{}, nil
	}

	return self.std_builder.BuilderActionBuild(log)
}

func (self *BuilderGlibc) BuilderActionDistribute_02(
	log *logger.Logger,
) error {

	self.std_builder.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install", "DESTDIR=" + self.bs.GetDIR_DESTDIR()}, nil
	}

	return self.std_builder.BuilderActionDistribute(log)
}
