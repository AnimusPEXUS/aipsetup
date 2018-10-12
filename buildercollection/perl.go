package buildercollection

import (
	"fmt"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/systemtriplet"
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

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	st, err := systemtriplet.NewFromString(info.HostArch)
	if err != nil {
		return nil, err
	}

	//	if st.Kernel != "linux" {
	//		return nil, errors.New("unsupported kernel")
	//	}

	//	archname := ""

	//	switch st.CPU {
	//	case "i486", "i586", "i686":
	//		archname = "i386-linux"
	//	case "x86_64":
	//		archname = "x86_64-linux"
	//	}

	//	if archname == "" {
	//		return nil, errors.New("unsupported cpu")
	//	}

	archname := fmt.Sprintf("%s-%s", st.CPU, st.Kernel)

	variant, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateMultilibVariant()
	if err != nil {
		return nil, err
	}

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = []string{
		"-Dprefix=" + install_prefix,
		"-Dcc=" + info.Host + "-gcc -m" + variant,
		"-Dld=" + info.Host + "-gcc -m" + variant,
		"-Duseshrplib",
		"-Dtargetarch=" + archname,
		"-Darchname=" + archname,
		//		"-Dtargethost=" + archname,
		//		"-Accflags=-m" + variant,
		//		"-Aldflags=-m" + variant,
		"-d",
		"-e",
	}

	//	ret, err = buildingtools.FilterAutotoolsConfigOptions(
	//		ret,
	//		[]string{},
	//		[]string{ "CC=", "CXX=", "GCC="},
	//	)
	//	if err != nil {
	//		return nil, err
	//	}

	return ret, nil
}

func (self *Builder_perl) EditConfigureEnv(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error) {
	//	for _, i := range []string{"CC", "CXX", "GCC"} {
	//		ret.Del(i)
	//	}
	return ret, nil
}
