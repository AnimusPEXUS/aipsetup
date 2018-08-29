package buildercollection

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

// BIG FAT WARNING: on x86_64 with multilib configured GCC boost libs
//                  MUST be installed into lib64 subdir of sysroot,
//                  else you will get some packages tending to be
//                  linked ageinst 32-bit libs under 32-bit lib dir: at
//                  least libtool does so for source-highlite package

func init() {
	Index["boost"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_boost(bs)
	}
}

type Builder_boost struct {
	*Builder_std

	python           string
	user_config_path string
}

func NewBuilder_boost(bs basictypes.BuildingSiteCtlI) (*Builder_boost, error) {

	self := new(Builder_boost)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	if t, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().
		CalculateInstallPrefixExecutable("python2"); err != nil {
		return nil, err
	} else {
		self.python = t
	}

	self.user_config_path =
		path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "user-config.jam")

	return self, nil
}

func (self *Builder_boost) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret = ret.Remove("autogen")

	err := ret.ReplaceShort("configure", self.BuilderActionConfigure)
	if err != nil {
		return nil, err
	}

	ret, err = ret.AddActionAfterNameShort(
		"configure",
		"bootstrap", self.BuilderActionBootstrap,
	)
	if err != nil {
		return nil, err
	}

	ret = ret.Remove("build")

	//	err = ret.ReplaceShort("build", self.BuilderActionBuild)
	//	if err != nil {
	//		return nil, err
	//	}

	err = ret.ReplaceShort("distribute", self.BuilderActionDistribute)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_boost) BuilderActionConfigure(log *logger.Logger) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	variant, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateMultilibVariant()
	if err != nil {
		return nil
	}

	//	cfg_data, err := ioutil.ReadFile(self.user_config_path)
	//	if err != nil {
	//		return err
	//	}

	compiler, err := self.GetBuildingSiteCtl().
		GetBuildingSiteValuesCalculator().CalculateInstallPrefixExecutable(
		fmt.Sprintf("%s-g++", info.Host),
	)
	if err != nil {
		return err
	}

	log.Info("Configuring")
	log.Info("   compiler: " + compiler)
	log.Info("   bitness: " + variant)

	appendix_tpl, err := template.New("pkg_config").Parse(
		"\n\nusing gcc : : {{.Compiler}} : <compileflags>-m{{.Bitness}} <linkflags>-m{{.Bitness}} ;\n\n",
	)
	if err != nil {
		return err
	}

	b := &bytes.Buffer{}

	err = appendix_tpl.Execute(
		b,
		struct {
			Compiler string
			Bitness  string
		}{
			Compiler: compiler,
			Bitness:  variant,
		},
	)
	if err != nil {
		return err
	}

	cfg_data_str := b.String()

	err = ioutil.WriteFile(self.user_config_path, []byte(cfg_data_str), 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_boost) BuilderActionBootstrap(log *logger.Logger) error {

	install_prefix, err :=
		self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return err
	}

	install_libdir, err :=
		self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallLibDir()
	if err != nil {
		return err
	}

	args := []string{
		"./bootstrap.sh",
		"--prefix=" + install_prefix,
		"--libdir=" + install_libdir,
		"--with-python=" + self.python,
	}

	log.Info("bootstrap arguments: " + strings.Join(args[1:], " "))

	//	env := environ.NewFromStrings(os.Environ())
	//	env

	c := exec.Command("bash", args...)
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_boost) BuilderActionDistribute(log *logger.Logger) error {

	dst_install_prefix, err :=
		self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	dst_install_libdir, err :=
		self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallLibDir()
	if err != nil {
		return err
	}

	args := []string{
		// NOTE: this is not an error:
		//       prefix = self.calculate_dst_install_prefix()
		"--prefix=" + dst_install_prefix,
		"--libdir=" + dst_install_libdir,
		// '--build-type=complete',
		// '--layout=versioned',
		// '--build-dir={}'.format(self.bld_dir),

		// NOTE: boost configurer and it's docs is crappy shit..
		//       thanks to Sergey Popov from Gentoo for pointing
		//       on --user-config= option
		"--user-config=" + self.user_config_path,
		// '--with-python={}'.format(self.custom_data['python']),
		"threading=multi",
		"link=shared",
		"stage",
		"install",
	}

	log.Info("b2 arguments: " + strings.Join(args, " "))

	c := exec.Command(path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "b2"), args...)
	c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	err = c.Run()
	if err != nil {
		return err
	}

	return nil
}
