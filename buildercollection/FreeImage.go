package buildercollection

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["FreeImage"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_FreeImage(bs)
	}
}

type Builder_FreeImage struct {
	*Builder_std
}

func NewBuilder_FreeImage(bs basictypes.BuildingSiteCtlI) (*Builder_FreeImage, error) {

	self := new(Builder_FreeImage)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditBuildArgsCB = self.EditBuildArgs

	return self, nil
}

func (self *Builder_FreeImage) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	err := ret.ReplaceShort("distribute", self.BuilderActionDistribute)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_FreeImage) EditBuildArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"DESTDIR=" + install_prefix,
		}...,
	)

	return ret, nil
}

func (self *Builder_FreeImage) BuilderActionDistribute(log *logger.Logger) error {

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_include := path.Join(dst_install_prefix, "include")

	dst_install_libdir, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallLibDir()
	if err != nil {
		return err
	}

	src_dst := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "Dist")

	src_dst_files, err := ioutil.ReadDir(src_dst)
	if err != nil {
		return err
	}

	for _, i := range []string{dst_install_libdir, dst_install_prefix_include} {
		err = os.MkdirAll(i, 0700)
		if err != nil {
			return err
		}
	}

	for _, i := range src_dst_files {

		i_name := i.Name()

		if strings.HasSuffix(i_name, ".a") || strings.HasSuffix(i_name, ".so") {
			err = filetools.CopyWithInfo(
				path.Join(src_dst, i_name),
				path.Join(dst_install_libdir, i_name),
				log,
			)
			if err != nil {
				return err
			}
		}

		if strings.HasSuffix(i_name, ".h") {
			err = filetools.CopyWithInfo(
				path.Join(src_dst, i_name),
				path.Join(dst_install_prefix_include, i_name),
				log,
			)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
