package aipsetup

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

// Some packages, like linux, and many others, have no features to precicely
// point where to install what, so, without additional file movements,
// Packager.DestDirCheckCorrectness() will return error.

// PrePackager role is to be called by builder at the final stage of package
// distribution in case if builder can't (or see it irrational to do it on
// it own) performe such movements.

type PrePackager struct {
	site *BuildingSiteCtl
}

func NewPrePackager(site *BuildingSiteCtl) *PrePackager {
	ret := new(PrePackager)
	ret.site = site
	return ret
}

func (self *PrePackager) Run(log *logger.Logger) error {
	for _, i := range [](func(log *logger.Logger) error){
		self.DestDirMoveRootToUsr,
		self.DestDirMoveUsrToPrefix,
	} {
		err := i(log)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *PrePackager) DestDirMoveRootToUsr(log *logger.Logger) error {
	dst_dir := self.site.GetDIR_DESTDIR()
	for _, i := range []string{"bin", "sbin", "lib", "lib64"} {
		i_j := path.Join(dst_dir, i)
		stat, err := os.Lstat(i_j)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			} else {
				continue
			}
		}
		if filetools.Is(stat.Mode()).Symlink() {
			return errors.New("src must be directory")
		}
		if !stat.IsDir() {
			return errors.New("src must be directory")
		}

		d_j := path.Join(dst_dir, "usr", i)

		err = filetools.CopyTree(
			i_j, d_j,
			false,
			false,
			true,
			false,
			log,
			filetools.CopyWithInfo,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *PrePackager) DestDirMoveUsrToPrefix(log *logger.Logger) error {

	usr_dir := path.Join(self.site.GetDIR_DESTDIR(), "usr")

	usr_dir_stat, err := os.Stat(usr_dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			return nil
		}
	}

	if !usr_dir_stat.IsDir() {
		return errors.New("invalid type of /usr in DESTDIR")
	}

	new_usr_dir, err := self.site.ValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		return err
	}

	err = os.MkdirAll(new_usr_dir, 0700)
	if err != nil {
		return err
	}

	usr_dir_lst, err := ioutil.ReadDir(usr_dir)
	if err != nil {
		return err
	}

	for _, i := range usr_dir_lst {
		new_dst_name := path.Join(new_usr_dir, i.Name())

		_, err := os.Lstat(new_dst_name)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			// already exists = error
			return errors.New("can't move item from usr to new place. already exists")
		}

		i_joined := path.Join(usr_dir, i.Name())
		i_joined_stat, err := os.Lstat(i_joined)
		if err != nil {
			return err
		}

		if i_joined_stat.IsDir() {
			err = filetools.CopyTree(
				i_joined,
				new_dst_name,
				false,
				false,
				true,
				false,
				log,
				filetools.CopyWithInfo,
			)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
