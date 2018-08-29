package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ntfs3gntfsprogs"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_ntfs3gntfsprogs(bs)
	}
}

type Builder_ntfs3gntfsprogs struct {
	*Builder_std
}

func NewBuilder_ntfs3gntfsprogs(bs basictypes.BuildingSiteCtlI) (*Builder_ntfs3gntfsprogs, error) {

	self := new(Builder_ntfs3gntfsprogs)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_ntfs3gntfsprogs) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	ret = append(
		ret,
		[]string{
			"--disable-ldconfig",
		}...,
	)
	return ret, nil
}
