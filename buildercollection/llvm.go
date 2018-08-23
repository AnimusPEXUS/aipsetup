package buildercollection

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/pkgconfig"
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
	self.EditConfigureEnvCB = self.EditConfigureEnv

	return self, nil
}

func (self *Builder_llvm) AfterExtract(log *logger.Logger, ret error) error {

	if ret != nil {
		return ret
	}

	// NOTE: see https://llvm.org/docs/GettingStarted.html#getting-started-quickly-a-summary

	//Remember that you were warned twice about reading the documentation.

	//    In particular, the relative paths specified are important.

	//Checkout LLVM:

	//    cd where-you-want-llvm-to-live
	//    svn co http://llvm.org/svn/llvm-project/llvm/trunk llvm

	//Checkout Clang:

	//    cd where-you-want-llvm-to-live
	//    cd llvm/tools
	//    svn co http://llvm.org/svn/llvm-project/cfe/trunk clang

	//Checkout Extra Clang Tools [Optional]:

	//    cd where-you-want-llvm-to-live
	//    cd llvm/tools/clang/tools
	//    svn co http://llvm.org/svn/llvm-project/clang-tools-extra/trunk extra

	//Checkout LLD linker [Optional]:

	//    cd where-you-want-llvm-to-live
	//    cd llvm/tools
	//    svn co http://llvm.org/svn/llvm-project/lld/trunk lld

	//Checkout Polly Loop Optimizer [Optional]:

	//    cd where-you-want-llvm-to-live
	//    cd llvm/tools
	//    svn co http://llvm.org/svn/llvm-project/polly/trunk polly

	//Checkout Compiler-RT (required to build the sanitizers) [Optional]:

	//    cd where-you-want-llvm-to-live
	//    cd llvm/projects
	//    svn co http://llvm.org/svn/llvm-project/compiler-rt/trunk compiler-rt

	//Checkout Libomp (required for OpenMP support) [Optional]:

	//    cd where-you-want-llvm-to-live
	//    cd llvm/projects
	//    svn co http://llvm.org/svn/llvm-project/openmp/trunk openmp

	//Checkout libcxx and libcxxabi [Optional]:

	//    cd where-you-want-llvm-to-live
	//    cd llvm/projects
	//    svn co http://llvm.org/svn/llvm-project/libcxx/trunk libcxx
	//    svn co http://llvm.org/svn/llvm-project/libcxxabi/trunk libcxxabi

	//Get the Test Suite Source Code [Optional]

	//    cd where-you-want-llvm-to-live
	//    cd llvm/projects
	//    svn co http://llvm.org/svn/llvm-project/test-suite/trunk test-suite

	findings_table := [][3]string{

		[3]string{"cfe", "tools/clang", ""},
		[3]string{"clang-tools-extra", "tools/clang/tools/extra", ""},
		//		[3]string{"lld", "tools/lld", ""},
		[3]string{"polly", "tools/polly", ""},

		[3]string{"compiler-rt", "projects/compiler-rt", ""},
		[3]string{"openmp", "projects/openmp", ""},

		//		[3]string{"libcxx", "projects/libcxx", ""},
		//		[3]string{"libcxxabi", "projects/libcxxabi", ""},
		//		[3]string{"libunwind", "projects/libunwind", ""},
	}

	log.Info("Looking for additional packages to extract them...")

	tarb_dir := self.GetBuildingSiteCtl().GetDIR_TARBALL()

	tarb_dir_files, err := ioutil.ReadDir(tarb_dir)
	if err != nil {
		return err
	}

	for _, i := range tarb_dir_files {
		i_name := i.Name()
		if tarballname.IsPossibleTarballName(i_name) {

			//			for _, j := range findings_table {

			//				if strings.HasPrefix(i_name, j[0]) {
			//					j[2] = i_name
			//					log.Info(fmt.Sprintf("%s found to be %s", j[0], j[2]))
			//				}

			//			}

			for j := 0; j != len(findings_table); j++ {

				if strings.HasPrefix(i_name, findings_table[j][0]) {
					findings_table[j][2] = i_name
					log.Info(
						fmt.Sprintf(
							"%s found to be %s",
							findings_table[j][0],
							findings_table[j][2],
						),
					)
				}

			}

			all_found := true

			for _, j := range findings_table {
				if j[2] == "" {
					all_found = false
					break
				}
			}

			if all_found {
				break
			}
		}
	}

	a_tools := new(buildingtools.Autotools)

	for _, i := range findings_table {

		if i[2] == "" {
			return errors.New(i[0] + " tarball not found")
		}

		log.Info("Extracting " + i[0])
		err = a_tools.Extract(
			path.Join(tarb_dir, i[2]),
			path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), i[1]),
			path.Join(self.GetBuildingSiteCtl().GetDIR_TEMP(), i[0]),
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

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	PATH, err := calc.Calculate_PATH()
	if err != nil {
		return nil, err
	}

	p, err := pkgconfig.NewPkgConfig(PATH, []string{})
	if err != nil {
		return nil, err
	}

	lst, err := p.GetIncludeDirs("libffi")
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

			//			"-DLLVM_ENABLE_LIBCXX=yes",
			//			"-DLLVM_ENABLE_LIBCXXABI=yes",
			//			"-DLIBCXXABI_LIBCXX_INCLUDES=" +
			//				path.Join(self.GetBuildingSiteCtl().GetDIR_SOURCE(), "projects", "libcxx", "include"),

			//			"-DLLVM_ENABLE_MODULES=yes",

			"-DLLVM_ENABLE_FFI=yes",
			"-DFFI_INCLUDE_DIR=" + lst[0],
		}...,
	)

	//	for i := len(ret) - 1; i != -1; i -= 1 {
	//		for _, j := range []string{"-DCC=", "-DCXX="} {
	//			if strings.HasPrefix(ret[i], j) {
	//				ret = append(ret[:i], ret[i+1:]...)
	//			}
	//		}
	//	}

	return ret, nil
}

func (self *Builder_llvm) EditConfigureEnv(log *logger.Logger, ret environ.EnvVarEd) (environ.EnvVarEd, error) {
	//	for _, i := range []string{"CC", "CXX"} {
	//		ret.Del(i)
	//	}
	return ret, nil
}
