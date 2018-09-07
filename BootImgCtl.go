package aipsetup

import (
	"os/exec"
	"path"
	"path/filepath"

	"github.com/AnimusPEXUS/utils/logger"
)

type BootImgCtl struct {
	src_root string
	wd_path  string

	initrd_ctl *BootImgInitRdCtl
	squash_ctl *BootImgSquashCtl

	img_file string

	log *logger.Logger
}

func NewBootImgCtl(
	src_root string,
	wd_path string,
	log *logger.Logger,
) (*BootImgCtl, error) {

	self := new(BootImgCtl)

	self.src_root = src_root

	wd_path, err := filepath.Abs(wd_path)
	if err != nil {
		return nil, err
	}

	self.wd_path = wd_path

	initrd_ctl, err = NewBootImgInitRdCtl(src_root, wd_path, log)
	if err != nil {
		return nil, err
	}

	squash_ctl, err = NewBootImgSquashCtl(src_root, wd_path, log)
	if err != nil {
		return nil, err
	}

	self.initrd_ctl = initrd_ctl
	self.squash_ctl = squash_ctl

	self.img_file = path.Join(wd_path, "liveflash.img")

	self.log = log

	return self, nil
}

func (self *BootImgCtl) FallocateImage() error {
	c := exec.Command("fallocate", "-l10G", self.img_file)
	return nil
}
