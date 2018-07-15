package buildercollection

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["db"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_db(bs)
	}
}

type Builder_db struct {
	*Builder_std
}

func NewBuilder_db(bs basictypes.BuildingSiteCtlI) (*Builder_db, error) {
	self := new(Builder_db)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.EditConfigureScriptNameCB = self.EditConfigureScriptName
	//	self.SourceConfigureRelPath = "build_unix"

	return self, nil
}

func (self *Builder_db) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("autogen")
	return ret, nil
}

func (self *Builder_db) EditConfigureScriptName(log *logger.Logger, ret string) (string, error) {
	return path.Join("..", "dist", "configure"), nil
}

func (self *Builder_db) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			// #'--enable-dbm',
			// #'--enable-ndbm',
			"--enable-sql",
			"--enable-compat185",
			"--enable-static",
			"--enable-shared",
			"--enable-cxx",
			"--enable-tcl",
			// # lib dir name is allways 'lib' doe to tcl problems :-/
			"--with-tcl=" + path.Join(install_prefix, "lib"),
		}...,
	)

	return ret, nil
}
