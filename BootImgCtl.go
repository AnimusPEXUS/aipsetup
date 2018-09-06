package aipsetup

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

type BootImgCtl struct {
	src_root    string
	wd_path     string
	os_files    string
	squashed_fs string
	log         *logger.Logger
}

func NewBootImgCtl(src_root string, wd_path string, log *logger.Logger) (*BootImgCtl, error) {

	self := new(BootImgCtl)

	self.src_root = src_root

	wd_path, err := filepath.Abs(wd_path)
	if err != nil {
		return nil, err
	}

	self.wd_path = wd_path
	self.os_files = path.Join(wd_path, "osfiles")
	self.squashed_fs = path.Join(wd_path, "squash.fs")
	self.log = log

	return self, nil
}

func (self *BootImgCtl) CopyOSFiles() error {

	{
		root_files_to_copy := []string{
			"bin", "sbin", "lib", "lib64", "usr",
			"var", "etc", "daemons", "multihost",
		}

		{
			root_files, err := ioutil.ReadDir(self.src_root)
			if err != nil {
				return err
			}
			for _, i := range root_files {
				for _, j := range []string{"etc.", "var."} {
					if strings.HasPrefix(i.Name(), j) {
						root_files_to_copy = append(root_files_to_copy, i.Name())
					}
				}
			}
		}

		for _, i := range root_files_to_copy {
			err := filetools.CopyTree(
				path.Join(self.src_root, i),
				path.Join(self.os_files, i),
				false,
				true,
				true,
				true,
				self.log,
				true,
				true,
				func(f, t string, log logger.LoggerI) error {
					fstat, err := os.Lstat(f)
					if err != nil {
						return err
					}

					if !fstat.Mode().IsRegular() && !filetools.Is(fstat.Mode()).Symlink() {
						log.Error("skipping irregular file " + f)
						return nil
					}

					err = filetools.CopyWithInfo(f, t, log)
					if err != nil {
						return err
					}

					return nil
				},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (self *BootImgCtl) InstallAipSetup() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	err = filetools.CopyWithInfo(
		exe,
		path.Join(self.os_files, "bin", "aipsetup"),
		self.log,
	)
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgCtl) CleanupOSFS() error {
	return nil
}

func (self *BootImgCtl) SquashOSFS() error {
	return nil
}

func (self *BootImgCtl) CreateDiskImage() error {
	return nil
}
