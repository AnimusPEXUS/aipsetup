package buildercollection

import (
	"bytes"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["mc"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_mc(bs)
	}
}

type Builder_mc struct {
	*Builder_std
}

func NewBuilder_mc(bs basictypes.BuildingSiteCtlI) (*Builder_mc, error) {
	self := new(Builder_mc)
	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_mc) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret, err := ret.AddActionsBeforeName(
		basictypes.BuilderActions{
			&basictypes.BuilderAction{
				Name:     "add_asp_open_support",
				Callable: self.BuilderActionAddASPSupport,
			},
		},
		"prepack",
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_mc) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	//	pkgconfig, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().GetPrefixPkgConfig()
	//	if err != nil {
	//		return nil, err
	//	}

	//	LIBSSH_CFLAGS, err := pkgconfig.CommandOutput("--cflags", "libssh2")
	//	if err != nil {
	//		return nil, err
	//	}

	//	LIBSSH_LIBS, err := pkgconfig.CommandOutput("--libs", "libssh2")
	//	if err != nil {
	//		return nil, err
	//	}

	//	LIBCRYPTO_CFLAGS, err := pkgconfig.CommandOutput("--cflags", "libcrypto")
	//	if err != nil {
	//		return nil, err
	//	}

	//	LIBCRYPTO_LIBS, err := pkgconfig.CommandOutput("--libs", "libcrypto")
	//	if err != nil {
	//		return nil, err
	//	}

	ret = append(
		ret,
		[]string{
			"--enable-charset",

			// NOTE: for some reason mc (4.8.21) can't be built with FLAGS
			//       (some scripting issues)
			//       and can't find values for lbicrypt manually, so
			//       sftp is disabled for now

			"--disable-vfs-sftp",
			//						"LIBSSH_CFLAGS=" + LIBSSH_CFLAGS,
			//						"LIBSSH_LIBS=" + LIBSSH_LIBS,
			//			"CFLAGS=" + LIBCRYPTO_CFLAGS,
			//			"LDFLAGS=" + LIBCRYPTO_LIBS,
		}...,
	)

	return ret, nil
}

func (self *Builder_mc) BuilderActionAddASPSupport(log *logger.Logger) error {

	exts_file := path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), "etc", "mc", "mc.ext")

	var lines []string

	{

		f, err := os.Open(exts_file)
		if err != nil {
			return err
		}

		defer f.Close()

		by := make([]byte, 0)

		_, err = f.Read(by)
		if err != nil {
			return err
		}

		b_str := string(by)

		strings.Split(b_str, "\n")

	}

	b := &bytes.Buffer{}

	{
		install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
		if err != nil {
			return err
		}

		for _, i := range lines {
			if i == "# asp" {
				return nil
			}
		}

		log.Info("adding ASP support")

		tpl, err := template.New("tpl").Parse(`
# asp
shell/i/.asp
` + "\t" + `Open=%cd %p/utar://
` + "\t" + `View=%view{ascii} {{.Prefix}}/libexec/mc/ext.d/archive.sh view tar

`)
		if err != nil {
			return err
		}

		err = tpl.Execute(
			b,
			struct{ Prefix string }{Prefix: install_prefix},
		)
	}

	f, err := os.Create(exts_file)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = b.WriteTo(f)
	if err != nil {
		return err
	}

	return nil
}
