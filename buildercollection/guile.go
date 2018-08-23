package buildercollection

import (
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["guile"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_guile(bs)
	}
}

type Builder_guile struct {
	*Builder_std
}

func NewBuilder_guile(bs basictypes.BuildingSiteCtlI) (*Builder_guile, error) {

	self := new(Builder_guile)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_guile) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	pv := strings.Split(info.PackageVersion, ".")

	if pv[0] == "2" && pv[1] == "0" {
		ret = append(
			ret,
			[]string{
				"--program-suffix=-2.0",
			}...,
		)
	}

	return ret, nil
}
