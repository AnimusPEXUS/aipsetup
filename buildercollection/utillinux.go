package buildercollection

import (
	"io/ioutil"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["utillinux"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_utillinux(bs)
	}
}

type Builder_utillinux struct {
	*Builder_std
}

func NewBuilder_utillinux(bs basictypes.BuildingSiteCtlI) (*Builder_utillinux, error) {

	self := new(Builder_utillinux)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_utillinux) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret, err := ret.AddActionBeforeNameShort(
		"autogen",
		"place_version", self.BuilderActionPlaceVersion,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_utillinux) BuilderActionPlaceVersion(log *logger.Logger) error {

	// TODO: promote this to std.go?

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	for _, i := range []string{".version", ".tarball-version"} {
		fn := path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), i)

		err = ioutil.WriteFile(fn, []byte(info.PackageVersion), 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_utillinux) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--with-python=3",
			"--disable-makeinstall-chown",
		}...,
	)

	return ret, nil
}
