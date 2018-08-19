package buildercollection

import (
	"bytes"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["libcap2"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_libcap2(bs)
	}
}

type Builder_libcap2 struct {
	*Builder_std
}

func NewBuilder_libcap2(bs basictypes.BuildingSiteCtlI) (*Builder_libcap2, error) {

	self := new(Builder_libcap2)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditDistributeArgsCB = self.EditDistributeArgs

	self.EditBuildConcurentJobsCountCB = self.EditBuildConcurentJobsCount

	return self, nil
}

func (self *Builder_libcap2) EditBuildConcurentJobsCount(log *logger.Logger, ret int) int {
	return 1
}

func (self *Builder_libcap2) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = ret.Remove("configure")
	ret = ret.Remove("autogen")
	ret = ret.Remove("build")

	if info.PackageName == "libcap2" && info.PackageVersion == "2.25" {

		err := ret.ReplaceShort("patch", self.BuilderActionPatch)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *Builder_libcap2) BuilderActionPatch(
	log *logger.Logger,
) error {

	patch := bytes.NewBuffer([]byte(`44c
	echo "#include <stddef.h>" > $@
	perl -e 'print "struct __cap_token_s { const char *name; int index; };\n%{\nconst struct __cap_token_s *__cap_lookup_name(register const char *str, register size_t len);\n%}\n%%\n"; while ($$l = <>) { $$l =~ s/[\{\"]//g; $$l =~ s/\}.*// ; print $$l; }' < $< | gperf --ignore-case --language=ANSI-C --readonly --null-strings --global-table --hash-function-name=__cap_hash_name --lookup-function-name="__cap_lookup_name" -c -t -m20 $(INDENT) >> $@
.
w
q
`))

	c := exec.Command("ed", "./Makefile")
	c.Dir = path.Join(self.bs.GetDIR_SOURCE(), "libcap")
	c.Stdout = log.StdoutLbl()
	c.Stderr = log.StderrLbl()
	c.Stdin = patch

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

func (self *Builder_libcap2) EditDistributeArgs(log *logger.Logger, ret []string) ([]string, error) {

	install_prefix, err := self.bs.GetBuildingSiteValuesCalculator().CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	ret = []string{
		"-j1",
		"all",
		"install",
		"prefix=" + install_prefix,
		//		"lib=" + main_multiarch_libdir_name,
		"lib=lib",
		"DESTDIR=" + self.bs.GetDIR_DESTDIR(),
		"RAISE_SETFCAP=no",
		"PAM_CAP=yes",
	}

	return ret, nil
}
