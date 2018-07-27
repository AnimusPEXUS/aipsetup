package buildercollection

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
)

func init() {
	Index["llvm"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_llvm(bs)
	}
}

type Builder_llvm struct {
	*Builder_std_cmake
}

func NewBuilder_llvm(bs basictypes.BuildingSiteCtlI) (*Builder_llvm, error) {

	self := new(Builder_llvm)

	if t, err := NewBuilder_std_cmake(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_cmake = t
	}

	self.AfterExtractCB = self.AfterExtract
	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_llvm) AfterExtract(log *logger.Logger, ret error) error {

	if ret != nil {
		return ret
	}

	log.Info("Looking for cfe and clang-tools-extra too extract them...")
	cfe_found := ""
	cet_found := ""

	tarb_dir := self.bs.GetDIR_TARBALL()

	tarb_dir_files, err := ioutil.ReadDir(tarb_dir)
	if err != nil {
		return err
	}

	for _, i := range tarb_dir_files {
		i_name := i.Name()
		if tarballname.IsPossibleTarballName(i_name) {
			if strings.HasPrefix(i_name, "cfe") {
				cfe_found = i_name
			}
			if strings.HasPrefix(i_name, "clang-tools-extra") {
				cet_found = i_name
			}

			if cfe_found != "" && cet_found != "" {
				break
			}
		}
	}

	a_tools := new(buildingtools.Autotools)

	for _, i := range [][3]string{
		[3]string{cfe_found, "cfe", "tools/clang"},
		[3]string{cet_found, "clang-tools-extra", "tools/clang/tools/extra"},
	} {
		if i[0] == "" {
			continue
		}

		log.Info("Extracting " + i[1])
		err = a_tools.Extract(
			path.Join(tarb_dir, i[0]),
			path.Join(self.bs.GetDIR_SOURCE(), i[2]),
			path.Join(self.bs.GetDIR_TEMP(), i[1]),
			true,
			false,
			false,
			log,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_llvm) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	info, err := self.bs.ReadInfo()
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"-DLLVM_EXPERIMENTAL_TARGETS_TO_BUILD=WebAssembly",
			"-DLLVM_INSTALL_UTILS=on",

			"-DBUILD_SHARED_LIBS=on",
			"-DCMAKE_BUILD_TYPE=Release",

			"-DLLVM_BUILD_DOCS=on",
			"-DLLVM_DEFAULT_TARGET_TRIPLE=" + info.HostArch,
			// #"-DLLVM_ENABLE_FFI=yes",
			"-DLLVM_ENABLE_LIBCXX=yes",
			"-DLLVM_ENABLE_LIBCXXABI=yes",
			// #"-DLLVM_ENABLE_MODULES=yes",
		}...,
	)

	return ret, nil
}
