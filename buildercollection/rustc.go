package buildercollection

import (
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/environ"
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

	info, err := self.bs.ReadInfo()
	if err != nil {
		return err
	}

	calc := self.bs.GetBuildingSiteValuesCalculator()

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return err
	}

	src_config_toml := path.Join(self.bs.GetDIR_SOURCE(), "config.toml")

	prefix := install_prefix
	sysconfdir := "/etc"
	docdir := path.Join(prefix, "share", "doc", "rust-"+info.PackageVersion)
	//	libdir, err := calc.CalculateInstallLibDir()
	//	if err != nil {
	//		return err
	//	}
	localstatedir := "/var"

	llvmconfig, err := calc.CalculateInstallPrefixExecutable("llvm-config")
	if err != nil {
		return err
	}

	rustc, err := calc.CalculateInstallPrefixExecutable("rustc")
	if err != nil {
		return err
	}

	cargo, err := calc.CalculateInstallPrefixExecutable("cargo")
	if err != nil {
		return err
	}

	//	prefix := path.Join(self.bs.GetDIR_DESTDIR(), install_prefix)
	//	sysconfdir := path.Join(self.bs.GetDIR_DESTDIR(), "/etc")
	//	docdir := path.Join(prefix, "share", "doc", "rust-"+info.PackageVersion)

	//	//	libdir, err := calc.CalculateInstallLibDir()
	//	//	if err != nil {
	//	//		return err
	//	//	}
	//	//	libdir = path.Join(self.bs.GetDIR_DESTDIR(), libdir)
	//	localstatedir := path.Join(self.bs.GetDIR_DESTDIR(), "/var")

	err = os.MkdirAll(prefix, 0700)
	if err != nil {
		return err
	}

	//	bits, err := calc.CalculateMultilibVariant()
	//	if err != nil {
	//		return err
	//	}

	cfg_txt := `
[llvm]
` +
		"link-shared = true\n" +
		`
[build]
` +
		"rustc = '" + rustc + "'\n" +
		"cargo = '" + cargo + "'\n" +
		"build = 'x86_64-unknown-linux-gnu'" + "\n" +
		"host = ['x86_64-unknown-linux-gnu']" + "\n" +
		"target = ['x86_64-unknown-linux-gnu']" + "\n" +
		//		"target = ['x86_64-pc-linux-gnu', 'i686-pc-linux-gnu']" + "\n" +
		//		"target = ['" + info.HostArch + "']" + "\n" +
		//		"local-rebuild = true\n" +
		"verbose = 1\n" +
		//		"vendor = true\n" +
		//		"extended = true\n" +
		`
[install]
` +
		"prefix = '" + prefix + "'\n" +
		"sysconfdir = '" + sysconfdir + "'\n" +
		"docdir = '" + docdir + "'\n" +
		//		"libdir = '" + libdir + "'\n" +
		"localstatedir = '" + localstatedir + "'\n" +
		`
[rust]
[dist]
[target.` + "x86_64-unknown-linux-gnu" + `]
` +
		//		"cc = '" + info.Host + "-gcc -m" + bits + "'\n" +
		//		"cxx = '" + info.Host + "-g++ -m" + bits + "'\n" +
		//		"ar = '" + info.Host + "-ar'\n" +
		//		"linker = '" + info.Host + "-gcc -m" + bits + "'\n" +

		"llvm-config = '" + llvmconfig + "'\n" +
		""

	f, err := os.Create(src_config_toml)
	if err != nil {
		return err
	}

	f.WriteString(cfg_txt)

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_rustc) BuilderActionBuild(
	log *logger.Logger,
) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	python2, err := calc.CalculateInstallPrefixExecutable("python2")
	if err != nil {
		return err
	}

	cmd := exec.Command(python2, "./x.py", "build")
	cmd.Dir = self.bs.GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_rustc) BuilderActionDistribute(
	log *logger.Logger,
) error {

	calc := self.bs.GetBuildingSiteValuesCalculator()

	python2, err := calc.CalculateInstallPrefixExecutable("python2")
	if err != nil {
		return err
	}

	e := environ.NewFromStrings(os.Environ())
	e.Set("DESTDIR", self.bs.GetDIR_DESTDIR())

	cmd := exec.Command(python2, "./x.py", "install")
	cmd.Dir = self.bs.GetDIR_SOURCE()
	cmd.Stdout = log.StdoutLbl()
	cmd.Stderr = log.StderrLbl()
	cmd.Env = e.Strings()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
