package buildercollection

import (
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["xfs"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_xfs(bs), nil
	}
}

type Builder_xfs struct {
	*Builder_std
}

func NewBuilder_xfs(bs basictypes.BuildingSiteCtlI) *Builder_xfs {
	self := new(Builder_xfs)

	self.Builder_std = NewBuilder_std(bs)

	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditActionsCB = self.EditActions

	return self
}

func (self *Builder_xfs) EditActions(
	ret basictypes.BuilderActions,
) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionsBeforeName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "preconfiguration",
				Callable: self.ActionCompletePreconfiguration,
			},
		},
		"configure",
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_xfs) ActionCompletePreconfiguration(
	log *logger.Logger,
) error {
	//	info, err := self.bs.ReadInfo()
	//	if err != nil {
	//		return err
	//	}

	//	calc := self.bs.GetBuildingSiteValuesCalculator()

	//	dst_install_prefix,err:=calc .CalculateDstInstallPrefix()
	//	if err != nil {
	//		return err
	//	}

	//	install_sh_path := path.Join(self.bs.GetDIR_SOURCE(), "include", "install-sh")
	install_sh_path2 := path.Join(self.bs.GetDIR_SOURCE(), "install-sh")
	configure_path := path.Join(self.bs.GetDIR_SOURCE(), "configure")

	//	if _, err := os.Stat(install_sh_path); err == nil {

	//		if _, err := os.Stat(install_sh_path2); err == nil {
	//			return nil
	//		}

	//		err = os.Link(install_sh_path, install_sh_path2)
	//		if err != nil {
	//			return err
	//		}
	//	}

	if _, err := os.Stat(install_sh_path2); err == nil {
		return nil
	}

	os.Remove(configure_path)

	a_tools := new(buildingtools.Autotools)

	err := a_tools.Make(
		[]string{"configure"},
		[]string{},
		buildingtools.Copy,
		"Makefile",
		self.bs.GetDIR_SOURCE(),
		self.bs.GetDIR_SOURCE(),
		"make",
		log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_xfs) EditDistributeArgs(
	log *logger.Logger,
	ret []string,
) (
	[]string,
	error,
) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = make([]string, 0)

	ret = append(ret, "install")

	// NOTE: looks like install-lib and install-dev has been deprecated

	{
		add_install_dev := true
		for _, i := range []string{"acl", "attr"} {
			if info.PackageName == i {
				add_install_dev = false
				break
			}
		}
		if add_install_dev {
			ret = append(ret, "install-dev")
		}
	}

	//	{
	//		add_install_lib := true
	//		for _, i := range []string{"acl", "attr", "xfsprogs", "xfsdump", "dmapi"} {
	//			if info.PackageName == i {
	//				add_install_lib = false
	//				break
	//			}
	//		}
	//		if add_install_lib {
	//			ret = append(ret, "install-lib")
	//		}
	//	}

	ret = append(ret, "DESTDIR="+self.bs.GetDIR_DESTDIR())

	return ret, nil
}
