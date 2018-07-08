package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["domterm"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_domterm(bs)
	}
}

type Builder_domterm struct {
	*Builder_std
}

func NewBuilder_domterm(bs basictypes.BuildingSiteCtlI) (*Builder_domterm, error) {
	self := new(Builder_domterm)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs
	return self, nil
}

func (self *Builder_domterm) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {
	return append(
		ret,
		[]string{
			"--without-asciidoctor",
			"--without-java",
			"--without-javafx",
			//			"--disable-doc",
			//			"--without-doc",
		}...,
	), nil
}
