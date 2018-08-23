package buildercollection

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["glibc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_glibc(bs)
	}
}

type Builder_glibc struct {
	*Builder_std

	slibdir string
}

func NewBuilder_glibc(bs basictypes.BuildingSiteCtlI) (*Builder_glibc, error) {

	self := new(Builder_glibc)

	self.Builder_std = NewBuilder_std(bs)

	//	self.SeparateBuildDir = true
	//	self.ForcedTarget = true
	//	self.ApplyHostSpecCompilerOptions = true

	self.EditConfigureArgsCB = self.EditConfigureArgs

	// calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	// if t, err := calc.CalculateInstallLibDir(); err != nil {
	// 	return nil, err
	// } else {
	// 	self.slibdir = t
	// }

	if t, err := self._CalculateSlibdir(); err != nil {
		return nil, err
	} else {
		self.slibdir = t
	}

	self.EditBuildArgsCB = self.EditBuildArgs
	self.EditDistributeArgsCB = self.EditDistributeArgs
	self.EditConfigureWorkingDirCB = self.EditConfigureWorkingDir

	return self, nil
}

func (self *Builder_glibc) EditConfigureWorkingDir(log *logger.Logger, ret string) (string, error) {
	return self.GetBuildingSiteCtl().GetDIR_BUILDING(), nil
}

func (self *Builder_glibc) DefineActions() (basictypes.BuilderActions, error) {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	ret, err := self.Builder_std.DefineActions()
	if err != nil {
		return nil, err
	}

	if info.ThisIsCrossbuilder() {
		ret = ret.Remove("build")
		ret = ret.Remove("distribute")

		ret = ret.AddActionsAfter(
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

func (self *Builder_glibc) _CalculateSlibdir() (string, error) {
	// # NOTE: on multilib installations glibc libraries shuld be allways
	// #       installed in host_lib_dir. Else - multilib GCC could
	// #       not be built

	// NOTE 2 and BIG FAT WARNING: if ix86-pc-linux-gnu's lib dir will go as
	//      /multihost/x86_64-pc-linux-gnu/multiarch/ix86-pc-linux-gnu/lib, then
	//      compilation of 32bit programms will stop working. this is a huge
	//      issue on which I dont have time right now

	// TODO: do something smart about this problem

	// === garbage code start ===
	// info,err := self.GetBuildingSiteCtl().ReadInfo()
	// if err != nil {
	// 	return "", nil
	// }
	//
	// host_dir := calc.CalculateHostDir()
	//
	// if info.Host == "x86_64-pc-linux-gnu"  {
	// 	if  info.HostArch == info.Host{
	// 		ret = ""
	// 	}
	//
	// }
	// === garbage code end ===

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	ret, err := calc.CalculateHostLibDir()
	if err != nil {
		return "", err
	}

	return ret, nil
}

func (self *Builder_glibc) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	var with_headers string

	if t, err := calc.CalculateHostDir(); err != nil {
		return nil, err
	} else {
		with_headers = path.Join(t, basictypes.DIRNAME_INCLUDE)
	}

	if info.ThisIsCrossbuilder() {

		host_builders_dir, err := calc.CalculateHostCrossbuildersDir()
		if err != nil {
			return nil, err
		}

		prefix := path.Join(host_builders_dir, info.CrossbuilderTarget)

		with_headers = path.Join(prefix, basictypes.DIRNAME_INCLUDE)

		hbt_opts, err := calc.CalculateAutotoolsHBTOptions()
		if err != nil {
			return nil, err
		}

		ret = make([]string, 0)
		ret = append(
			ret,
			[]string{
				"--prefix=" + prefix,
				"--mandir=" + path.Join(prefix, basictypes.DIRNAME_SHARE, "man"),
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

	slibdir, err := self._CalculateSlibdir()
	if err != nil {
		return nil, err
	}

	for i := range ret {
		if strings.HasPrefix(ret[i], "--libdir=") {
			ret[i] = "--libdir=" + slibdir
		}
	}

	ENABLE_KERNEL := "4.17"

	ret = append(
		ret,
		[]string{

			"--enable-obsolete-rpc",
			"--enable-kernel=" + ENABLE_KERNEL,
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
	// # NOTE: don't remove this block. it's for informational reason
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

	if info.ThisIsCrossbuilder() {
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

func (self *Builder_glibc) EditBuildArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(ret, "slibdir="+self.slibdir)
	return ret, nil
}

func (self *Builder_glibc) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(ret, "slibdir="+self.slibdir)
	return ret, nil
}

func (self *Builder_glibc) BuilderActionDistribute_01(
	log *logger.Logger,
) error {

	self.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{
			"install-bootstrap-headers=yes",
			"install-headers",
			"DESTDIR=" + self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
		}, nil
	}

	return self.BuilderActionBuild(log)
}

func (self *Builder_glibc) BuilderActionDistribute_01_2(
	log *logger.Logger,
) error {

	self.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{
			"csu/subdir_lib",
		}, nil
	}

	return self.BuilderActionBuild(log)
}

func (self *Builder_glibc) BuilderActionDistribute_01_3(
	log *logger.Logger,
) error {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	dhcd, err := calc.CalculateDstHostCrossbuildersDir()
	if err != nil {
		return err
	}

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	mmldn, err := calc.CalculateMainMultiarchLibDirName()
	if err != nil {
		return err
	}

	gres, err := filepath.Glob(
		path.Join(self.GetBuildingSiteCtl().GetDIR_BUILDING(), "csu", "*crt*.o"),
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

		err = filetools.CopyWithInfo(i, path.Join(dest_lib_dir, path.Base(i)), nil)
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_glibc) BuilderActionDistribute_01_4(
	log *logger.Logger,
) error {
	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	dhcd, err := calc.CalculateDstHostCrossbuildersDir()
	if err != nil {
		return err
	}

	info, err := self.GetBuildingSiteCtl().ReadInfo()
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

func (self *Builder_glibc) BuilderActionDistribute_01_5(
	log *logger.Logger,
) error {

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	dhcd, err := calc.CalculateDstHostCrossbuildersDir()
	if err != nil {
		return err
	}

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	cwd := path.Join(dhcd, info.CrossbuilderTarget, basictypes.DIRNAME_INCLUDE, "gnu")

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

func (self *Builder_glibc) BuilderActionIntermediateInstruction(
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

func (self *Builder_glibc) BuilderActionBuild_02(
	log *logger.Logger,
) error {

	self.EditBuildArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{}, nil
	}

	return self.BuilderActionBuild(log)
}

func (self *Builder_glibc) BuilderActionDistribute_02(
	log *logger.Logger,
) error {

	self.EditDistributeArgsCB = func(
		log *logger.Logger,
		ret []string,
	) ([]string, error) {
		return []string{"install", "DESTDIR=" + self.GetBuildingSiteCtl().GetDIR_DESTDIR()}, nil
	}

	return self.BuilderActionDistribute(log)
}
