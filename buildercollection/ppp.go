package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ppp"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_ppp(bs)
	}
}

type Builder_ppp struct {
	*Builder_std
}

func NewBuilder_ppp(bs basictypes.BuildingSiteCtlI) (*Builder_ppp, error) {

	self := new(Builder_ppp)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_ppp) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("build")

	return ret, nil
}

func (self *Builder_ppp) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{"--with-pydebug"}...,
	)
	return ret, nil
}

func (self *Builder_ppp) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	var err error

	//	pkgconfig, err := self.GetBuildingSiteCtl().
	//		GetBuildingSiteValuesCalculator().GetPrefixPkgConfig()
	//	if err != nil {
	//		return nil, err
	//	}

	//	openssl_cflags, err := pkgconfig.CommandOutput("--cflags", "openssl")
	//	if err != nil {
	//		return nil, err
	//	}

	//	openssl_libs, err := pkgconfig.CommandOutput("--libs", "openssl")
	//	if err != nil {
	//		return nil, err
	//	}

	ret, err = buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{"DESTDIR="},
	)
	if err != nil {
		return nil, err
	}

	args := []string{
		"all",
		"install",
		"CHAPMS=1",

		// TODO: have to restore crypt
		// "USE_CRYPT=1",

		//"INCLUDE_DIRS+= -I../include ",
		"INSTROOT=" + path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR()),
		//		"CFLAGS+=" + openssl_cflags,
		//		"LDFLAGS+=" + openssl_libs,
	}

	ret = append(ret, args...)

	return ret, nil
}
