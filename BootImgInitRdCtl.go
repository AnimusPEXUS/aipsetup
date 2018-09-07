package aipsetup

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

type BootImgInitRdCtl struct {
	src_root      string
	wd_path       string
	os_files      string
	initrd_tar_xz string
	log           *logger.Logger
}

func NewBootImgInitRdCtl(
	src_root string,
	wd_path string,
	log *logger.Logger,
) (*BootImgInitRdCtl, error) {

	self := new(BootImgInitRdCtl)

	self.src_root = src_root

	wd_path, err := filepath.Abs(wd_path)
	if err != nil {
		return nil, err
	}

	self.wd_path = wd_path
	self.os_files = path.Join(wd_path, "osfiles")
	self.initrd_tar_xz = path.Join(wd_path, "initrd.tar.xz")
	self.log = log

	return self, nil
}

func (self *BootImgInitRdCtl) CopyOSFiles() error {

	system := NewSystem(self.src_root, self.log)

	host, err := system.Host()
	if err != nil {
		return err
	}

	sys_packs := NewSystemPackages(system)

	pkg_names, err := pkginfodb.ListPackagesByGroups([]string{"fib"})
	if err != nil {
		return err
	}

	for _, i := range pkg_names {

		asps, err := sys_packs.ListInstalledPackageNameASPs(i, host, host)
		if err != nil {
			return err
		}

		for _, j := range asps {
			files, err := sys_packs.ListInstalledASPFiles(j)
			if err != nil {
				return err
			}

			self.log.Info("copying " + j.String() + "..")

			for _, k := range files {
				src := path.Join(system.Root(), k)

				if _, err := os.Lstat(src); err != nil {
					if !os.IsNotExist(err) {
						return err
					} else {
						continue
					}
				}

				dst := path.Join(self.os_files, k)

				err = os.MkdirAll(path.Dir(dst), 755)
				if err != nil {
					return err
				}

				err = filetools.CopyWithInfo(
					src,
					dst,
					nil,
				)
				if err != nil {
					return err
				}

			}
		}

	}

	for _, i := range []string{
		"mnt", "run", "tmp", "root", "dev",
		"proc", "sys", "root_new", "root_old",
	} {
		err := os.MkdirAll(path.Join(self.os_files, i), 0700)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BootImgInitRdCtl) MakeSymlinks() error {

	var err error

	names := []string{"bin", "sbin", "lib", "lib64"}

	for _, i := range names {
		// TODO: no hardcode
		err = os.Symlink(
			path.Join("multihost", "x86_64-pc-linux-gnu", i),
			path.Join(self.os_files, i),
		)
		if err != nil {
			return err
		}

	}

	return nil
}

func (self *BootImgInitRdCtl) WriteInit() error {
	init_file := path.Join(self.os_files, "sbin", "init")
	init := `#!/bin/bash

echo '+=(initrd)========================+'
echo '|                                 |'
echo '|  WELCOME TO HORIZON LIVE IMAGE  |'
echo '|                                 |'
echo '|=================================+'


export LD_LIBRARY_PATH=/lib:/lib64

# this should be already mounted by kernel
# mount -t devtmpfs devtmpfs /dev

mount -t proc proc /proc
mount -t sysfs sysfs /sys
mount -o ro PARTUUID=` + basictypes.BOOT_IMAGE_BOOT_PARTITION_UUID + ` /boot
mount /boot/squash.fs /root_new

# echo "testing overlayfs"
# /bin/bash

# echo "Ignore next 4 possible /root_new/* mount error messages"
# umount /boot

mount --move /boot /root_new/boot

umount /proc
umount /dev
umount /sys

cd /root_new

pivot_root . /root_new/root_old

exec chroot . /overlay_init.sh
`
	err := ioutil.WriteFile(
		init_file,
		[]byte(init),
		0700,
	)
	if err != nil {
		return err
	}

	err = os.Chmod(init_file, 0700)
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgInitRdCtl) DoEverythingBeforePack() error {
	for _, i := range [](func() error){
		self.CopyOSFiles,
		self.MakeSymlinks,
		self.WriteInit,
	} {
		err := i()
		if err != nil {
			return err
		}
	}
	return nil
}
