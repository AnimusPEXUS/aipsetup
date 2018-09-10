package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["sshfs_fuse"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_sshfs_fuse(bs)
	}
}

type Builder_sshfs_fuse struct {
	*Builder_std_meson
}

func NewBuilder_sshfs_fuse(bs basictypes.BuildingSiteCtlI) (*Builder_sshfs_fuse, error) {

	self := new(Builder_sshfs_fuse)

	if t, err := NewBuilder_std_meson(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_meson = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_sshfs_fuse) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"-Ddocs=false",
		}...,
	)

	return ret, nil
}
