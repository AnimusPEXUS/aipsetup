package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/buildingtools"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["ffmpeg"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_ffmpeg(bs)
	}
}

type Builder_ffmpeg struct {
	*Builder_std
}

func NewBuilder_ffmpeg(bs basictypes.BuildingSiteCtlI) (*Builder_ffmpeg, error) {

	self := new(Builder_ffmpeg)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_ffmpeg) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	// TODO: host compiler specification resolving needed.

	ret, err := buildingtools.FilterAutotoolsConfigOptions(
		ret,
		[]string{},
		[]string{
			"--includedir=",
			"--sysconfdir=",
			"--localstatedir=",
			"--sbindir=",
			"--bindir=",
			"--libexecdir=",
			"--datarootdir=",
			"--exec-prefix=",
			"--host=",
			"--build=",
			"--target=",
			"CC=",
			"GCC=",
			"CXX=",
		},
	)
	if err != nil {
		return nil, err
	}

	ret = append(
		ret,
		[]string{
			"--enable-shared",
			"--enable-gpl",
			"--enable-libtheora",
			"--enable-libvorbis",
			//			"--enable-x11grab",
			"--enable-libmp3lame",
			"--enable-libx264",
			"--enable-libxvid",
			"--enable-runtime-cpudetect",
			"--enable-doc",
			"--enable-avresample",
		}...,
	)

	return ret, nil
}
