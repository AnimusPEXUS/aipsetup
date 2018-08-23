package buildercollection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["openssl"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_openssl(bs)
	}
}

type Builder_openssl struct {
	*Builder_std
}

func NewBuilder_openssl(bs basictypes.BuildingSiteCtlI) (*Builder_openssl, error) {

	self := new(Builder_openssl)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	self.EditActionsCB = self.EditActions

	self.EditConfigureScriptNameCB = self.EditConfigureScriptName
	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditConfigureIsArgToShellCB = self.EditConfigureIsArgToShell

	self.EditConfigureShellCB = self.EditConfigureShell
	//	self.AfterDistributeCB = self.AfterDistribute

	return self, nil
}

func (self *Builder_openssl) veredName() (string, error) {
	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return "", err
	}

	spl := strings.Split(info.PackageVersion, ".")

	ret := fmt.Sprintf(
		"openssl-%s.%s",
		spl[0],
		spl[1],
	)

	return ret, nil
}

func (self *Builder_openssl) EditConfigureShell(log *logger.Logger, ret string) (string, error) {
	return "perl", nil
}

func (self *Builder_openssl) EditConfigureIsArgToShell(log *logger.Logger, ret bool) (bool, error) {
	return true, nil
}

func (self *Builder_openssl) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {
	ret = ret.Remove("build")
	ret, err := ret.AddActionAfterNameShort(
		"distribute",
		"after-distribute", self.BuilderActionAfterDistribute,
	)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (self *Builder_openssl) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	platform := "linux-generic32"
	if strings.HasPrefix(info.HostArch, "x86_64") {
		platform = "linux-x86_64"
	}

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = make([]string, 0)

	ret = append(
		ret,
		[]string{
			"--prefix=" + install_prefix,
			"--openssldir=/etc/ssl",
		}...,
	)

	libdir_slice := []string{"lib"}

	{

		ver_name, err := self.veredName()
		if err != nil {
			return nil, err
		}

		// TODO: I don't like this. Need to add some better comparators
		spl := strings.Split(info.PackageVersion, ".")
		if !(spl[0] == "1" && spl[1] == "1") {
			libdir_slice = append(libdir_slice, ver_name)
		}
	}

	ret = append(
		ret,
		[]string{
			"--libdir=" + path.Join(libdir_slice...),
		}...,
	)

	ret = append(
		ret,
		[]string{
			"shared",
			"zlib-dynamic",
			platform,
		}...,
	)

	return ret, nil
}

func (self *Builder_openssl) EditConfigureScriptName(log *logger.Logger, ret string) (string, error) {
	return "Configure", nil
}

func (self *Builder_openssl) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	//	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	//	if err != nil {
	//		return nil, err
	//	}

	ret = []string{
		"depend",
		"install",
		// FIXME: fix path join
		"MANDIR=" + path.Join(install_prefix, "share", "man"),
		// "MANSUFFIX=ssl",
		"INSTALL_PREFIX=" + self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
		"DESTDIR=" + self.GetBuildingSiteCtl().GetDIR_DESTDIR(),
	}

	return ret, nil
}

func (self *Builder_openssl) BuilderActionAfterDistribute(log *logger.Logger) error {
	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	spl := strings.Split(info.PackageVersion, ".")

	// TODO: I don't like this. Need to add some better version comparators

	if spl[0] == "1" && spl[1] == "1" {
		return nil
	}

	dst_install_prefix, err := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	ver_name, err := self.veredName()
	if err != nil {
		return err
	}

	lib_dir := path.Join(dst_install_prefix, "lib")
	lib_ossl_dir := ""
	lib_dir_files, err := ioutil.ReadDir(lib_dir)
	if err != nil {
		return err
	}

	include_openssl := path.Join(dst_install_prefix, "include", "openssl")

	include_openssl_ver := path.Join(dst_install_prefix, "include", ver_name)

	err = os.Rename(include_openssl, include_openssl_ver)
	if err != nil {
		return err
	}

	for _, i := range lib_dir_files {
		if strings.HasPrefix(i.Name(), "open") {
			lib_ossl_dir = path.Join(lib_dir, i.Name())
			break
		}
	}

	if lib_ossl_dir == "" {
		return errors.New("openssl* dir not found inside lib dir")
	}

	lib_ossl_dir_files, err := ioutil.ReadDir(lib_ossl_dir)
	if err != nil {
		return err
	}

	for _, i := range lib_ossl_dir_files {
		fn := path.Join(lib_ossl_dir, i.Name())
		if yes, err := regexp.MatchString(`lib.*\.so.*`, i.Name()); err != nil {
			return err
		} else {
			if yes {
				if s, err := os.Lstat(fn); err != nil {
					//				if !os.IsNotExist(err ) {
					return err
					//				}
				} else {
					if !filetools.Is(s.Mode()).Symlink() {
						nname := path.Join(lib_dir, i.Name())
						// yes - hardlink
						err = os.Link(fn, nname)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	lib_ossl_pkgconfig_dir := path.Join(lib_ossl_dir, "pkgconfig")

	if _, err := os.Stat(lib_ossl_pkgconfig_dir); err != nil {
		return err
	}

	lib_ossl_pkgconfig_dir_files, err := ioutil.ReadDir(lib_ossl_pkgconfig_dir)
	if err != nil {
		return err
	}

	for _, i := range lib_ossl_pkgconfig_dir_files {
		fn := path.Join(lib_ossl_pkgconfig_dir, i.Name())
		if strings.HasSuffix(i.Name(), ".pc") {

			ft, err := ioutil.ReadFile(fn)
			if err != nil {
				return err
			}

			ft_lines := strings.Split(string(ft), "\n")

			for i := len(ft_lines) - 1; i != -1; i-- {
				if ft_lines[i] == "includedir=${prefix}/include" {
					ft_lines[i] = "includedir=${prefix}/include/" + ver_name
				}

				if strings.HasPrefix(ft_lines[i], "Requires:") {
					t := strings.SplitN(ft_lines[i], ":", 2)[1]
					tss := strings.Split(t, " ")
					for k, v := range tss {
						fmt.Println(" tss", k, v)
					}
					for j := len(tss) - 1; j != -1; j-- {
						if len(tss[j]) == 0 {
							tss = append(tss[:j], tss[j+1:]...)
						} else {
							if ok, err := regexp.MatchString(`\s+`, tss[j]); err != nil {
								return err
							} else {
								if ok {
									tss = append(tss[:j], tss[j+1:]...)
								}
							}
						}
					}

					for j := 0; j != len(tss); j++ {
						tss[j] = tss[j] + "-1.0"
					}

					ft_lines[i] = "Requires: " + strings.Join(tss, " ")
				}
			}

			err = ioutil.WriteFile(
				fn[:len(fn)-3]+ver_name[len(ver_name)-4:]+".pc",
				[]byte(strings.Join(ft_lines, "\n")),
				0700,
			)
			if err != nil {
				return err
			}

			err = os.Remove(fn)
			if err != nil {
				return err
			}

		}
	}

	err = os.Rename(lib_ossl_pkgconfig_dir, path.Join(lib_dir, "pkgconfig"))
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(dst_install_prefix, "bin"))
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(dst_install_prefix, "share"))
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(self.GetBuildingSiteCtl().GetDIR_DESTDIR(), "etc"))
	if err != nil {
		return err
	}

	return nil
}
