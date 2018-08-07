package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["perl"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_perl(bs)
	}
}

type Builder_perl struct {
	*Builder_std
}

func NewBuilder_perl(bs basictypes.BuildingSiteCtlI) (*Builder_perl, error) {

	self := new(Builder_perl)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureScriptNameCB = self.EditConfigureScriptName
	self.EditConfigureArgsCB = self.EditConfigureArgs
	self.EditConfigureEnvCB = self.EditConfigureEnv

	return self, nil
}

func (self *Builder_perl) EditConfigureScriptName(log *logger.Logger, ret string) (string, error) {
	return "Configure", nil
}

func (self *Builder_perl) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = []string{
		"-Dprefix=" + install_prefix,
		"-Dcc=" + info.Host + "-gcc",
		"-Duseshrplib",
		"-d",
		"-e",
	}

	//	for i := len(ret) - 1; i != -1; i -= 1 {
	//		if strings.HasPrefix(ret[i], "CC=") ||
	//			strings.HasPrefix(ret[i], "CXX=") ||
	//			strings.HasPrefix(ret[i], "GCC=") {
	//			ret = append(ret[:i], ret[i+1:]...)
	//		}
	//	}

	return ret, nil
}

func (self *Builder_perl) EditConfigureEnv(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error) {
	for _, i := range []string{"CC", "CXX", "GCC"} {
		ret.Del(i)
	}
	return ret, nil
}
