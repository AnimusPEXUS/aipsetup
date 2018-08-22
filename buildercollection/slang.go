package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["slang"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_slang(bs)
	}
}

type Builder_slang struct {
	*Builder_std
}

func NewBuilder_slang(bs basictypes.BuildingSiteCtlI) (*Builder_slang, error) {

	self := new(Builder_slang)

	self.Builder_std = NewBuilder_std(bs)

	self.EditBuildConcurentJobsCountCB = self.EditBuildConcurentJobsCount

	return self, nil
}

func (self *Builder_slang) EditBuildConcurentJobsCount(log *logger.Logger, ret int) int {
	return 1
}
