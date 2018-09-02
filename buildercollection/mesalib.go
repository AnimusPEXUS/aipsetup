package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["mesalib"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_mesalib(bs)
	}
}

type Builder_mesalib struct {
	*Builder_std_meson
}

func NewBuilder_mesalib(bs basictypes.BuildingSiteCtlI) (*Builder_mesalib, error) {

	self := new(Builder_mesalib)

	if t, err := NewBuilder_std_meson(bs); err != nil {
		return nil, err
	} else {
		self.Builder_std_meson = t
	}

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_mesalib) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{

			//			//            # NOTE: on OpenGL version printed by glxinfo or qtdiag
			//			//            #       http://permalink.gmane.org/gmane.comp.video.mesa3d.user/3311

			//			//            # ------------------------
			//			//            # next block is copy&paste from
			//			//            # https://wiki.freedesktop.org/nouveau/InstallNouveau/
			//			//            # ->
			//			"--enable-texture-float",
			//			"--enable-gles1",
			//			"--enable-gles2",
			//			"--enable-glx",
			//			"--enable-egl",
			//			"--enable-gallium-egl",
			//			"--enable-llvm", // #'--enable-gallium-llvm",
			//			"--enable-shared-glapi",
			//			"--enable-gbm",
			//			"--enable-glx-tls", //  # undefined reference to `_glapi_tls_Dispatch'
			//			"--enable-dri",
			//			"--enable-osmesa",
			//			"--enable-vdpau",
			//			//            #'--with-platforms=x11,drm', '--with-egl-platforms=x11,drm',
			//			//            #'--with-gallium-drivers=nouveau',
			//			//            #'--with-dri-drivers=nouveau',
			//			//            # <-
			//			//            # end of https://wiki.freedesktop.org/nouveau/InstallNouveau/
			//			//            # block
			//			//            # ------------------------

			//			//            # ------------------------
			//			//            # https://pkg-xorg.alioth.debian.org/howto/build-mesa.html
			//			//            # ->

			//			//            #'--with-dri-driverdir={}/lib/dri'.format(
			//			//            #    self.calculate_install_prefix()
			//			//            #    ),
			//			"--enable-driglx-direct",

			//			//            # <-
			//			//            # ------------------------

			//			//            #'--libdir={}/lib'.format(
			//			//            #    self.calculate_install_prefix()
			//			//            #    ),

			//			"--enable-texture-float",

			//			"--enable-gles1",
			//			"--enable-gles2",

			//			"--enable-openvg=auto",

			//			"--enable-osmesa",       //  # -
			//			"--with-osmesa-bits=64", //  # -

			//			"--enable-xa",
			//			"--enable-gbm",

			//			//            #'--disable-gallium',
			//			//            #'--disable-llvm', #'--disable-gallium-llvm',

			//			"--enable-egl",
			//			"--enable-gallium-egl", // # -
			//			"--enable-gallium-gbm",

			//			"--enable-dri",       // # -
			//			"--enable-dri3=auto", // # -

			//			//            # '--enable-glx-tls",

			//			"--enable-xorg", //  # -

			//			// TODO: readd wayland
			//			"--with-platforms=x11,drm", //# '--with-egl-platforms=x11,drm,wayland',  # -

			//			"--with-gallium-drivers=nouveau,svga,swrast,virgl",        // # -
			//			"--with-dri-drivers=nouveau,i915,i965,r200,radeon,swrast", //  # -
			//			//            #'--without-gallium-drivers",
			//			//            #'--without-dri-drivers",

			//			//            # "--enable-d3d1x",
			//			//            # "--enable-opencl",

			//			//            #"--with-clang-libdir={}'.format(
			//			//            #    wayround_i2p.utils.path.join(
			//			//            #        self.get_host_dir(),
			//			//            #        'lib'
			//			//            #        )
			//			//            #    ),
			//			//            #'--with-llvm-prefix={}'.format(self.get_host_dir()),

			//			//            #'PYTHON2={}'.format(
			//			//            #    wayround_i2p.utils.file.which(
			//			//            #        'python2',
			//			//            #        self.get_host_dir()
			//			//            #        )
			//			//            #    )

			//			//            # NOTE: llvm is installed into 'lib' dir and
			//			//            #       trying to use 32-bit glibc libs, while it must use
			//			//            #       64-bit. so here is the hack to point it to right
			//			//            #       'lib64' dir
			//			//            'LLVM_LDFLAGS=-L{}'.format(
			//			//                wayround_i2p.utils.path.join(
			//			//                    self.calculate_install_libdir(),
			//			//                    #'lib'
			//			//                    )
			//			//                ),
		}...,
	)

	return ret, nil
}
