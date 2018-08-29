package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["mupdf"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_mupdf(bs)
	}
}

type Builder_mupdf struct {
	*Builder_std
}

func NewBuilder_mupdf(bs basictypes.BuildingSiteCtlI) (*Builder_mupdf, error) {

	self := new(Builder_mupdf)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditDistributeArgsCB = self.EditDistributeArgs

	return self, nil
}

func (self *Builder_mupdf) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")
	ret = ret.Remove("configure")

	return ret, nil
}

func (self *Builder_mupdf) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"prefix=/usr",
		}...,
	)

	return ret, nil
}
