package buildercollection

import (
	"os/exec"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["sdl"] = func(bs basictypes.BuildingSiteCtlI) (basictypes.BuilderI, error) {
		return NewBuilder_sdl(bs)
	}
}

type Builder_sdl struct {
	*Builder_std
}

func NewBuilder_sdl(bs basictypes.BuildingSiteCtlI) (*Builder_sdl, error) {
	self := new(Builder_sdl)

	self.Builder_std = NewBuilder_std(bs)

	self.EditActionsCB = self.EditActions

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_sdl) EditActions(ret basictypes.BuilderActions) (basictypes.BuilderActions, error) {

	ret, err := ret.AddActionAfterNameShort(
		"extract",
		"patch", self.Patch,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *Builder_sdl) Patch(log *logger.Logger) error {

	info, err := self.GetBuildingSiteCtl().ReadInfo()
	if err != nil {
		return err
	}

	if info.PackageVersion == "1.2.15" {

		log.Info("Patching 1.2.15")

		c := exec.Command(
			"sed",
			"-e",
			"/_XData32/s:register long:register _Xconst long:",
			"-i",
			"src/video/x11/SDL_x11sym.h",
		)
		c.Dir = self.GetBuildingSiteCtl().GetDIR_SOURCE()
		c.Stdout = log.StdoutLbl()
		c.Stderr = log.StderrLbl()
		err = c.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *Builder_sdl) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-audio",
			"--enable-video",
			"--enable-events",
			"--enable-libc",
			"--enable-loads",
			"--enable-file",
			"--disable-alsa",
			"--enable-pulseaudio",
			"--enable-pulseaudio-shared",
		}...,
	)

	return ret, nil
}
