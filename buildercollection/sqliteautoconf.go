package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["sqliteautoconf"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_sqliteautoconf(bs)
	}
}

type Builder_sqliteautoconf struct {
	*Builder_std
}

func NewBuilder_sqliteautoconf(bs basictypes.BuildingSiteCtlI) (*Builder_sqliteautoconf, error) {

	self := new(Builder_sqliteautoconf)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_sqliteautoconf) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	// TODO: build for 32 bit

	ret = append(
		ret,
		[]string{
			"CFLAGS=" +
				"-DSQLITE_ENABLE_FTS3=1 " +
				"-DSQLITE_ENABLE_COLUMN_METADATA=1 " +
				"-DSQLITE_ENABLE_UNLOCK_NOTIFY=1 " +
				"-DSQLITE_SECURE_DELETE=1 ",
		}...,
	)

	return ret, nil
}
