package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["zlib"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_zlib(bs)
	}
}

type Builder_zlib struct {
	*Builder_std
}

func NewBuilder_zlib(bs basictypes.BuildingSiteCtlI) (*Builder_zlib, error) {

	self := new(Builder_zlib)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditDistributeArgsCB = self.EditDistributeArgs

	//	self.EditDistributeDESTDIRCB = func(log *logger.Logger, ret string) (string, error) {
	//		return "prefix", nil
	//	}

	return self, nil
}

func (self *Builder_zlib) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("build")

	//	ret, err := ret.AddActionAfterNameShort(
	//		"distribute",
	//		"after-distribute", self.BuilderActionAfterDistribute,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}

	return ret, nil
}

func (self *Builder_zlib) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	variant, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	ret = []string{
		"--prefix=" + install_prefix,
		"--shared",
	}

	if variant == "64" {
		ret = append(ret, "--64")
	}

	return ret, nil
}

func (self *Builder_zlib) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{
			"prefix=",
			"DESTDIR=",
		},
	)
	if err != nil {
		return nil, err
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"prefix=" + dst_install_prefix,
		}...,
	)

	return ret, nil
}

//func (self *Builder_zlib) BuilderActionAfterDistribute(log *logger.Logger) error {

//	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
//		CalculateDstInstallPrefix()
//	if err != nil {
//		return err
//	}

//	dst_share := path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), "share")
//	dst_ip_share := path.Join(dst_install_prefix, "share")

//	dst_include := path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), "include")
//	dst_ip_include := path.Join(dst_install_prefix, "include")

//	for _, i := range [][2]string{
//		[2]string{dst_share, dst_ip_share},
//		[2]string{dst_include, dst_ip_include},
//	} {
//		err = os.Rename(i[0], i[1])
//		if err != nil {
//			return err
//		}
//	}

//	return nil
//}
