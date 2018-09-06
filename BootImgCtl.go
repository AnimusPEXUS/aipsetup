package aipsetup

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/shadowusers"
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

	for _, i := range []string{"mnt", "run", "tmp", "root", "dev"} {
		err := os.MkdirAll(path.Join(self.os_files, i), 0700)
		if err != nil {
			return err
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

func (self *BootImgCtl) RemoveUsers() error {

	susers := shadowusers.NewCtl(path.Join(self.os_files, "etc"))

	err := susers.ReadAll()
	if err != nil {
		return err
	}

	simple_users := make([]int, 0)
	simple_users_names := make([]string, 0)

	for _, i := range susers.Passwds.Passwds {
		if i.UserId > basictypes.SYS_UID_MAX {
			simple_users = append(simple_users, i.UserId)
			simple_users_names = append(simple_users_names, i.Login)
		}
	}

	for i := len(susers.Passwds.Passwds) - 1; i != -1; i-- {
		if susers.Passwds.Passwds[i].UserId > basictypes.SYS_UID_MAX {
			susers.Passwds.Passwds = append(
				susers.Passwds.Passwds[:i],
				susers.Passwds.Passwds[i+1:]...,
			)
		}
	}

	for i := len(susers.Groups.Groups) - 1; i != -1; i-- {
		del := false

		for _, j := range simple_users {
			if susers.Groups.Groups[i].GID == j {
				del = true
				break
			}
		}

		if del {
			susers.Groups.Groups = append(
				susers.Groups.Groups[:i],
				susers.Groups.Groups[i+1:]...,
			)
		}
	}

	for i := len(susers.Shadows.Shadows) - 1; i != -1; i-- {
		del := false

		for _, j := range simple_users_names {
			if susers.Shadows.Shadows[i].Login == j {
				del = true
				break
			}
		}

		if del {
			susers.Shadows.Shadows = append(
				susers.Shadows.Shadows[:i],
				susers.Shadows.Shadows[i+1:]...,
			)
		}
	}

	for i := len(susers.GShadows.GShadows) - 1; i != -1; i-- {
		del := false

		for _, j := range simple_users_names {
			if susers.GShadows.GShadows[i].Name == j {
				del = true
				break
			}
		}

		if del {
			susers.GShadows.GShadows = append(
				susers.GShadows.GShadows[:i],
				susers.GShadows.GShadows[i+1:]...,
			)
		}
	}

	err = susers.WriteAll()
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgCtl) ResetRootPasswd() error {
	susers := shadowusers.NewCtl(path.Join(self.os_files, "etc"))

	err := susers.ReadAll()
	if err != nil {
		return err
	}

	r, err := susers.Shadows.GetByLogin("root")
	if err != nil {
		return err
	}

	r.Password =
		`$6$cLewkeecW2A4.SvS$t0ckgHfOu8jPtPPDelZwH6binT7sLigIhyz0Pou6Kz9lb.Xy/qMWA6bNvWTOfSGwNsssTTWiKc2bmo10GEjVP.`

	err = susers.WriteAll()
	if err != nil {
		return err
	}

	return nil
}

func (self *BootImgCtl) CleanupOSFS() error {

	targets_to_remove := []string{
		"/etc/passwd-",
		"/etc/shadow-",
		"/etc/groups-",
		"/etc/gshadow-",
		"/etc/machine-id",
		"/etc/dhcpcd.secret",
		"/etc/dhcpcd.secret",
		"/etc/tor",
	}

	for _, i := range targets_to_remove {
		err := os.RemoveAll(path.Join(self.os_files, i))
		if err != nil {
			return err
		}
	}

	sd_journals_pth := path.Join(self.os_files, "var", "log", "journal")

	err := os.MkdirAll(sd_journals_pth, 0700)
	if err != nil {
		return err
	}

	sd_journals_pth_files, err := ioutil.ReadDir(sd_journals_pth)
	if err != nil {
		return err
	}

	for _, i := range sd_journals_pth_files {
		err = os.RemoveAll(i.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *BootImgCtl) CleanupLinuxSrc() error {
	pth := path.Join(self.os_files, "usr", "src", "linux")
	c := exec.Command("make", "clean")
	c.Stdout = self.log.StdoutLbl()
	c.Stderr = self.log.StderrLbl()
	c.Dir = pth
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (self *BootImgCtl) SquashOSFiles() error {
	c := exec.Command("mksquashfs", self.os_files, self.squashed_fs, "-comp", "xz")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Dir = self.wd_path
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
