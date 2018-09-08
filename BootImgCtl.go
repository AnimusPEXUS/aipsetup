package aipsetup

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

type BootImgCtl struct {
	src_root string
	wd_path  string

	initrd_ctl *BootImgInitRdCtl
	squash_ctl *BootImgSquashCtl

	img_file   string
	mnt_dir    string
	kernel_dir string

	loop_dev   string
	loop_devp1 string
	loop_devp2 string

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

	initrd_ctl, err := NewBootImgInitRdCtl(src_root, wd_path, log)
	if err != nil {
		return nil, err
	}

	squash_ctl, err := NewBootImgSquashCtl(src_root, wd_path, log)
	if err != nil {
		return nil, err
	}

	self.initrd_ctl = initrd_ctl
	self.squash_ctl = squash_ctl

	self.img_file = path.Join(wd_path, "liveflash.img")
	self.mnt_dir = path.Join(wd_path, "mnt")
	self.kernel_dir = path.Join(wd_path, "kernel")

	self.loop_dev = path.Join("/dev", "loop0")
	self.loop_devp1 = self.loop_dev + "p1"
	self.loop_devp2 = self.loop_dev + "p2"

	self.log = log

	return self, nil
}

func (self *BootImgCtl) CheckFiles() error {
	files, err := ioutil.ReadDir(self.kernel_dir)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("no files inside " + self.kernel_dir + " dir")
	}

	if _, err := os.Stat(self.initrd_ctl.initrd_tar_xz); err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			return errors.New("make initrd")
		}
	}

	if _, err := os.Stat(self.squash_ctl.squashed_fs); err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			return errors.New("make squash.fs")
		}
	}
	return nil
}

func (self *BootImgCtl) FallocateImage() error {
	c := exec.Command("fallocate", "-l10G", self.img_file)
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) LoSetup() error {
	c := exec.Command("losetup", "-P", self.loop_dev, self.img_file)
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) EditParts() error {

	script := `
label: gpt
device: ` + self.loop_dev + `
unit: sectors

` + self.loop_devp1 + ` : type=21686148-6449-6E6F-744E-656564454649, attrs="LegacyBIOSBootable", size=10MiB
` + self.loop_devp2 + ` : type=0FC63DAF-8483-4772-8E79-3D69D8477DE4, uuid=` +
		basictypes.BOOT_IMAGE_BOOT_PARTITION_UUID + `
`

	b := bytes.NewReader([]byte(script))

	c := exec.Command("sfdisk", self.loop_dev)
	c.Stdin = b
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) FormatBoot() error {
	c := exec.Command(
		"mke2fs",
		"-U", basictypes.BOOT_IMAGE_BOOT_PARTITION_FS_UUID,
		self.loop_devp2,
	)
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) MountBoot() error {
	err := os.MkdirAll(self.mnt_dir, 0700)
	if err != nil {
		return err
	}
	c := exec.Command("mount", self.loop_devp2, self.mnt_dir)
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err = c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) InstallGrub() error {
	c := exec.Command("grub-install", "--boot-directory="+self.mnt_dir, self.loop_dev)
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) CopyFiles() error {

	self.log.Info("copy " + self.initrd_ctl.initrd_tar_xz)
	err := filetools.CopyWithInfo(
		self.initrd_ctl.initrd_tar_xz,
		path.Join(self.mnt_dir, path.Base(self.initrd_ctl.initrd_tar_xz)),
		self.log,
	)
	if err != nil {
		return err
	}

	self.log.Info("copy " + self.squash_ctl.squashed_fs)
	err = filetools.CopyWithInfo(
		self.squash_ctl.squashed_fs,
		path.Join(self.mnt_dir, path.Base(self.squash_ctl.squashed_fs)),
		self.log,
	)
	if err != nil {
		return err
	}

	k_files, err := ioutil.ReadDir(self.kernel_dir)
	if err != nil {
		return err
	}

	for _, i := range k_files {
		err := filetools.CopyWithInfo(
			path.Join(self.kernel_dir, i.Name()),
			path.Join(self.mnt_dir, i.Name()),
			self.log,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BootImgCtl) CreateGrubCfg() error {
	kernel := ""
	{
		k_files, err := ioutil.ReadDir(self.kernel_dir)
		if err != nil {
			return err
		}

		for _, i := range k_files {
			if strings.HasPrefix(i.Name(), "vmlinuz-") {
				kernel = i.Name()
				break
			}
		}
	}

	txt := `
menuentry start {
	search --fs-uuid --set=root ` + basictypes.BOOT_IMAGE_BOOT_PARTITION_FS_UUID + `
	linux /` + kernel + `
	initrd /` + path.Base(self.initrd_ctl.initrd_tar_xz) + `
}
`
	cfg_path := path.Join(self.mnt_dir, "grub", "grub.cfg")

	err := ioutil.WriteFile(cfg_path, []byte(txt), 0755)
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgCtl) UMountBoot() error {
	c := exec.Command("umount", self.loop_devp2)
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) LoSetupD() error {
	c := exec.Command("losetup", "-d", self.loop_dev)
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) DoEverything() error {
	for _, i := range [](func() error){
		self.CheckFiles,
		self.FallocateImage,
		self.LoSetup,
		self.EditParts,
		self.FormatBoot,
		self.MountBoot,
		self.InstallGrub,
		self.CopyFiles,
		self.CreateGrubCfg,
		self.UMountBoot,
		self.LoSetupD,
	} {
		err := i()
		if err != nil {
			return err
		}
	}

	return nil
}
