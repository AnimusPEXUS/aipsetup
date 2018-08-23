package buildercollection

import (
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["python"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_python(bs), nil
	}
}

type Builder_python struct {
	*Builder_std
}

func NewBuilder_python(bs basictypes.BuildingSiteCtlI) *Builder_python {
	self := new(Builder_python)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self
}

func (self *Builder_python) EditConfigureArgs(
	log *logger.Logger,
	ret []string,
) (
	[]string,
	error,
) {
	// DISCLAIMER //
	// Python 2.7.10 and 3.4.3 have have Me a lot of butt heart.
	//  When i've naively gave them --libdir=.../lib64, they'r resulting
	//  installations have spleatup between lib and lib64. There were ways
	//  to force installation to lib64 using Makefile's internal variables,
	//  but this where leading to unworking python installations.
	// Are you agree with me?: this is not good for so called "stable"
	//  versions. This was the last drop, which lead me to thinking about
	//  changing prograaming language of aipsetup.
	// This module will contain commentaries left from pythonish aipsetup.

	calc := self.GetBuildingSiteCtl().GetBuildingSiteValuesCalculator()

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return nil, err
	}

	install_prefix, err := calc.CalculateInstallPrefix()
	if err != nil {
		return nil, err
	}

	cb_opts := make([]string, 0)

	if info.ThisIsCrossbuilding() {
		cb_opts = append(
			cb_opts,
			[]string{
				"--disable-ipv6",
				"--without-ensurepip",
				"ac_cv_file__dev_ptmx=no",
				"ac_cv_file__dev_ptc=no",
			}...,
		)
	}

	//            # TODO: need to figure out what is this for
	//            '--with-libc={}'.format(
	//                wayround_i2p.utils.path.join(
	//                    self.target_host_root,
	//                    '/usr'
	//                    )
	//                )

	ret = append(
		ret,
		[]string{
			"--without-ensurepip",
			// # '--with-pydebug' # NOTE: enabling may cause problems to Cython
		}...,
	)

	// # NOTE: python shuld be ALLWAYS be installed in 'lib' dir. be it i?86
	// #       or x86_64 build, else *.so modules will go into lib64 and
	// #       python modules will remain in lib and Your system
	// #       will crush because of this

	for i := len(ret) - 1; i != -1; i -= 1 {
		if strings.HasPrefix(ret[i], "--libdir=") {
			ret = append(ret[:i], ret[i+1:]...)
		}
	}

	// # NOTE: at least python 2.7.10 and 3.4.3 are hard coded and can't
	// #       be configured to install scripts into lib64 dir

	ret = append(
		ret,
		"--libdir="+path.Join(install_prefix, "lib"),
	)

	//        ret += [
	//            # 'SCRIPTDIR={}'.format(
	//            #     wayround_i2p.utils.path.join(
	//            #         self.get_host_dir(),
	//            #         'lib'
	//            #         )
	//            #     ),
	//            ] + cb_opts

	ret = append(ret, cb_opts...)

	//        '''
	//            'DESTLIB={}'.format(
	//                self.get_host_lib_dir()
	//                ),
	//            'LIBDEST={}'.format(
	//                self.get_host_lib_dir()
	//                ),
	//        '''

	return ret, nil
}
