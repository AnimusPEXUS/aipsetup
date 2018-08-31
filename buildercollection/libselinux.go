package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libselinux"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_libselinux(bs), nil
	}
}

type Builder_libselinux struct {
	*Builder_std
}

func NewBuilder_libselinux(bs basictypes.BuildingSiteCtlI) *Builder_libselinux {

	self := new(Builder_libselinux)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureDirCB = func(log *logger.Logger, ret string) (string, error) {
		info, err := self.GetBuildingSiteCtl().ReadInfo()
		if err != nil {
			return "", err
		}
		return path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), info.PackageName), nil
	}

	self.EditConfigureWorkingDirCB = self.EditConfigureDirCB

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self
}

func (self *Builder_libselinux) EditActions(ret basictypes.BuilderActions) (
	basictypes.BuilderActions,
	error,
) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	//	ret, err := ret.AddActionAfterNameShort(
	//		"distribute",
	//		"fix-symlinks", self.BuilderActionFixSymlinks,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}

	return ret, nil
}

func (self *Builder_libselinux) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	libdir := path.Join(install_prefix, "lib")

	ret = append(
		ret,
		[]string{
			"PREFIX=" + install_prefix,
			"LIBDIR=" + libdir,
			"SHLIBDIR=" + libdir,
		}...,
	)
	return ret, nil
}

//func (self *Builder_libselinux) BuilderActionFixSymlinks(log *logger.Logger) error {

//	dst_install_prefix, err := self.GetBuildingSiteCtl().
//		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
//	if err != nil {
//		return err
//	}

//	for _, i := range []string{"lib", "lib64"} {

//		dp := path.Join(dst_install_prefix, i)

//		if _, err := os.Stat(dp); err != nil {
//			if !os.IsNotExist(err) {
//				return err
//			} else {
//				log.Info("doesn't exists " + dp)
//				continue
//			}
//		}

//		log.Info("searching symlinks inside " + dp)

//		dp_files, err := ioutil.ReadDir(dp)
//		if err != nil {
//			return err
//		}

//		for _, i := range dp_files {
//			if filetools.Is(i.Mode()).Symlink() {

//				sl_fn := i.Name()
//				sl_fp := path.Join(dp, sl_fn)

//				log.Info("fixing " + sl_fp)

//				sl_value, err := os.Readlink(sl_fp)
//				if err != nil {
//					return nil
//				}

//				sl_value_base := path.Base(sl_value)

//				err = os.Remove(sl_fp)
//				if err != nil {
//					return nil
//				}

//				err = os.Symlink(sl_value_base, sl_fp)
//				if err != nil {
//					return nil
//				}
//			}
//		}

//	}

//	return nil
//}
