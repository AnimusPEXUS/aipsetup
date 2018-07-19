package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["rustc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_rustc(bs)
	}
}

type Builder_rustc struct {
	*Builder_std
}

func NewBuilder_rustc(bs basictypes.BuildingSiteCtlI) (*Builder_rustc, error) {

	self := new(Builder_rustc)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	return self, nil
}

func (self *Builder_rustc) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("patch")
	ret = ret.Remove("autogen")

	err := ret.Replace(
		"configure",
		&basictypes.BuilderAction{
			Name:     "configure",
			Callable: self.BuilderActionConfigure,
		},
	)
	if err != nil {
		return nil, err
	}

	err = ret.Replace(
		"build",
		&basictypes.BuilderAction{
			Name:     "build",
			Callable: self.BuilderActionBuild,
		},
	)
	if err != nil {
		return nil, err
	}

	err = ret.Replace(
		"distribute",
		&basictypes.BuilderAction{
			Name:     "distribute",
			Callable: self.BuilderActionDistribute,
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_rustc) BuilderActionConfigure(
	log *logger.Logger,
) error {
	src_config_toml := path.Join(self.bs.GetDIR_SOURCE(), "config.toml")

	prefix = install_prefix
	sysconfdir = "/etc"

	cfg_txt := `
[llvm]
[build]
[install]
` +
		"prefix = '" + prefix + "'\n" +
		"sysconfdir = '" + etc + "'\n" +
		"docdir = '" + docdir + "'\n" +
		"libdir = '" + libdir + "'\n" +
		"localstatedir = '" + localstatedir + "'\n" +
		`
[rust]
[dist]
`

	return nil
}

func (self *Builder_rustc) BuilderActionBuild(
	log *logger.Logger,
) error {
	return nil
}

func (self *Builder_rustc) BuilderActionDistribute(
	log *logger.Logger,
) error {
	return nil
}
