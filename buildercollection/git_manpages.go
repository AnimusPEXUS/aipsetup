package buildercollection

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["git_manpages"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_git_manpages(bs)
	}
}

type Builder_git_manpages struct {
	*Builder_std
}

func NewBuilder_git_manpages(bs basictypes.BuildingSiteCtlI) (*Builder_git_manpages, error) {

	self := new(Builder_git_manpages)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_git_manpages) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")
	ret = ret.Remove("build")

	err := ret.Replace(
		"extract",
		&basictypes.BuilderAction{
			Name:     "extract",
			Callable: self.BuilderActionExtract,
		},
	)
	if err != nil {
		return nil, err
	}

	err = ret.Replace(
		"distribute",
		&basictypes.BuilderAction{
			Name:     "distribute",
			Callable: self.BuilderActionDistribute,
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_git_manpages) BuilderActionExtract(log *logger.Logger) error {

	a_tools := new(buildingtools.Autotools)

	main_tarball, err := self.bs.DetermineMainTarrball()
	if err != nil {
		return err
	}

	err = a_tools.Extract(
		path.Join(self.bs.GetDIR_TARBALL(), main_tarball),
		self.bs.GetDIR_SOURCE(),
		path.Join(self.bs.GetDIR_TEMP(), "primary_tarball"),
		false,
		true,
		false,
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_git_manpages) BuilderActionDistribute(log *logger.Logger) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	src_dir_dirs, err := ioutil.ReadDir(self.bs.GetDIR_SOURCE())
	if err != nil {
		return err
	}

	dst_install_prefix, err := calc.CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_prefix_man := path.Join(dst_install_prefix, "share", "man")

	os.MkdirAll(dst_install_prefix_man, 0700)

	for _, i := range src_dir_dirs {
		if i.IsDir() && strings.HasPrefix(i.Name(), "man") {
			err = os.Rename(
				path.Join(self.bs.GetDIR_SOURCE(), i.Name()),
				path.Join(dst_install_prefix_man, i.Name()),
			)
		}
	}

	return nil
}
