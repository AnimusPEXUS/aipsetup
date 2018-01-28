package aipsetup

import (
	"errors"
	"fmt"
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
	log.Info("checking and moving *bin/lib* dirs to usr")
	dst_dir := self.site.GetDIR_DESTDIR()
	for _, i := range []string{"bin", "sbin", "lib", "lib64"} {
		log.Info(" searching for " + i + " directory")
		i_j := path.Join(dst_dir, i)
		stat, err := os.Lstat(i_j)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Error("  error lstating " + i + " directory:" + err.Error())
				return err
			} else {
				log.Info("  " + i + " not found = ok - continuing..")
				continue
			}
		}

		log.Info("   item " + i + " is found. checking it's properties..")

		if filetools.Is(stat.Mode()).Symlink() || !stat.IsDir() {
			log.Error(
				fmt.Sprintf(
					"    item has unacceptable properties: not directory == %v, symlink == %v",
					!stat.IsDir(),
					filetools.Is(stat.Mode()).Symlink()),
			)
			return errors.New("src must be directory")
		}

		log.Info("    properties acceptable")

		d_j := path.Join(dst_dir, "usr", i)

		log.Info("     starting copy process of " + i)
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
			log.Error("      copy process ended with error: " + err.Error())
			return err
		}
		log.Info("      copy succeeded")
		log.Info("       removing " + i)
		err = os.RemoveAll(i_j)
		if err != nil {
			log.Error("       error: " + err.Error())
			return err
		}
	}
	return nil
}

func (self *PrePackager) DestDirMoveUsrToPrefix(log *logger.Logger) error {
	log.Info("checking and moving usr to install prefix")

	usr_dir := path.Join(self.site.GetDIR_DESTDIR(), "usr")

	log.Info(" checking usr dir existance")

	usr_dir_stat, err := os.Stat(usr_dir)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Info("  error " + err.Error())
			return err
		} else {
			log.Info("  usr does not exists")
			return nil
		}
	}

	log.Info("   usr found. checking it's properties")

	if !usr_dir_stat.IsDir() {
		err = errors.New("invalid type of /usr in DESTDIR")
		log.Error("    error. " + err.Error())
		return err
	}

	log.Info("    usr parameters looks acceptable")

	log.Info("calculating new path for usr contents")

	new_usr_dir, err := self.site.ValuesCalculator().CalculateDstInstallPrefix()
	if err != nil {
		log.Info(" error: " + err.Error())
		return err
	}

	log.Info(" " + new_usr_dir)

	log.Info("making new dir")

	err = os.MkdirAll(new_usr_dir, 0700)
	if err != nil {
		return err
	}

	log.Info("listing usr")

	usr_dir_lst, err := ioutil.ReadDir(usr_dir)
	if err != nil {
		return err
	}

	log.Info("copying files..")
	for _, i := range usr_dir_lst {
		i_joined := path.Join(usr_dir, i.Name())

		new_dst_name := path.Join(new_usr_dir, i.Name())

		log.Info(" from " + i_joined)
		log.Info("   to " + new_dst_name)

		new_dst_name_stat, err := os.Lstat(new_dst_name)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Error("  error stating " + new_dst_name)
				log.Error("   " + err.Error())
				return err
			}
		} else {

			if !new_dst_name_stat.IsDir() {
				log.Error("dst " + new_dst_name + " already exists and it's not a directory")
				return errors.New("dst already exists and it's not a directory")
			}

		}

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
				log.Error("copy error: " + err.Error())
				return err
			}
			err = os.RemoveAll(i_joined)
		} else {
			log.Error(i_joined + " is not the directory. don't know what to do")
			return errors.New(
				"src is not the directory. don't know what to do",
			)
		}

	}

	log.Info("Removing usr")
	err = os.RemoveAll(usr_dir)
	if err != nil {
		log.Error(" error: " + err.Error())
		return err
	}

	log.Info("usr moving complete")

	return nil
}
