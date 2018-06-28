package aipsetup

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"syscall"

	"github.com/AnimusPEXUS/aipsetup/etcfiles"
	"github.com/AnimusPEXUS/utils/environ"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/set"
)

// TODO: this functionality is outdated and requires some improvements and
//			 cleanups

type SystemUpdates struct {
	sys *System
}

func NewSystemUpdates(sys *System) *SystemUpdates {
	self := new(SystemUpdates)
	self.sys = sys
	return self
}

func (self *SystemUpdates) UpdatesAfterPkgInstall() error {
	uid := os.Getuid()

	was_errors := false

	separator_line := "---------------------------------------------------"

	if uid != 0 {

		self.sys.log.Info(separator_line)
		self.sys.log.Info("You are not a root, so no updates (except sync)")
		self.sys.log.Info(separator_line)
		self.Sync()

	}

	if uid == 0 {

		self.sys.log.Info(separator_line)
		self.sys.log.Info(
			"System updates: Performing System Updates After Installing/Removing " +
				"packages",
		)
		self.sys.log.Info(separator_line)

		for _, i := range []func() error{
			self.Sync,
			self.LDConfig,
			self.UpdateMimeDatabase,
			self.GDKPixbuffQueryLoaders,
			self.GlibCompileSchemas,
			self.GtkQueryImmodules20,
			self.GtkQueryImmodules30,
			self.Sync,
		} {
			res := i()
			if res != nil {
				was_errors = true
			}
		}
	}
	if was_errors {
		return errors.New("system updates errors")
	}
	return nil
}

func (self *SystemUpdates) Sync() error {
	self.sys.log.Info("Running sync")
	syscall.Sync()
	return nil
}

func (self *SystemUpdates) LDConfig() error {

	self.sys.log.Info("Running ldconfig")

	c := exec.Command("ldconfig")
	err := c.Run()
	if err != nil {
		self.sys.log.Error("  " + err.Error())
		return err
	}
	return nil
}

func (self *SystemUpdates) _Roots() ([]string, error) {
	ret := make([]string, 0)
	calc := self.sys.GetSystemValuesCalculator()

	host, err := self.sys.Host()
	if err != nil {
		return nil, err
	}

	archs, err := self.sys.Archs()
	if err != nil {
		return nil, err
	}

	for _, i := range archs {
		pth := calc.CalculateHostArchDir(host, i)
		ret = append(ret, pth)
	}

	return ret, nil
}

func (self *SystemUpdates) _UpdateMimeDatabase_Check(
	mime_dir,
	check_file string,
) (bool, error) {

	check_file = path.Base(check_file)

	mime_dir_check_file := path.Join(mime_dir, check_file)

	err := os.MkdirAll(mime_dir, 0755)
	if err != nil {
		return false, err
	}

	mime_dir_check_file_s, err := os.Stat(mime_dir_check_file)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		} else {
			return true, err
		}
	}
	mime_dir_check_file_s_t := mime_dir_check_file_s.ModTime()

	ret := false

	err = filetools.Walk(
		mime_dir,
		func(
			dir string,
			dirs []os.FileInfo,
			files []os.FileInfo,
		) error {
			for _, i := range files {

				if i.ModTime().After(mime_dir_check_file_s_t) {
					ret = true
					return errors.New("ok")
				}
			}
			return nil
		},
	)
	if err != nil {
		if err.Error() != "ok" {
			return true, err
		}
	}

	return ret, nil
}

func (self *SystemUpdates) _UpdateMimeDatabase_Run(
	root,
	mime_dir,
	check_file string,
) error {
	check_file = path.Base(check_file)

	mime_dir_check_file := path.Join(mime_dir, check_file)

	err := os.MkdirAll(mime_dir, 0755)
	if err != nil {
		return err
	}

	c := exec.Command(
		path.Join(root, "/bin/update-mime-database"),
		mime_dir,
	)

	err = c.Run()
	if err != nil {
		return err
	}

	if o, err := os.Create(mime_dir_check_file); err != nil {
		return err
	} else {
		o.Close()
	}

	return nil
}

func (self *SystemUpdates) UpdateMimeDatabase() error {

	self.sys.log.Info("Update Mime Database")

	roots, err := self._Roots()
	if err != nil {
		return err
	}

	was_errors := false

	for _, i := range roots {
		check_file := path.Join(i, "share", "mime", ".modcheck")

		file_missing := false
		if _, err := os.Stat(check_file); os.IsNotExist(err) {
			file_missing = true
		}

		up_required, err := self._UpdateMimeDatabase_Check(
			path.Dir(check_file),
			path.Base(check_file),
		)
		if err != nil {
			was_errors = true
			self.sys.log.Error(
				"  error checking mime database at " + path.Dir(check_file) + " " + err.Error(),
			)
			continue
		}

		if file_missing || up_required {
			err := self._UpdateMimeDatabase_Run(
				i,
				path.Dir(check_file),
				path.Base(check_file),
			)
			if err != nil {
				was_errors = true
				self.sys.log.Error(
					"  error updating mime database at " + path.Dir(check_file) + " " + err.Error(),
				)
				continue
			}
		}
	}

	if was_errors {
		return errors.New("errors while running Mime Database updates")
	}

	return nil
}

func (self *SystemUpdates) GDKPixbuffQueryLoaders() error {

	self.sys.log.Info("Qurying Pixbuf loaders")

	roots, err := self._Roots()
	if err != nil {
		return err
	}

	was_errors := false

	paths := make([][3]string, 0)

	for _, i := range roots {
		paths2, err := filepath.Glob(
			path.Join(i, "*", "gdk-pixbuf-2.0", "*", "loaders"),
		)
		if err != nil {
			was_errors = true
			self.sys.log.Error(
				"  error while searching for loaders dir " + i + " " + err.Error(),
			)
			continue
		}

		for _, j := range paths2 {
			paths = append(
				paths,
				[3]string{
					path.Join(i, "bin", "gdk-pixbuf-query-loaders"),
					j,
					path.Join(j, "..", "loaders.cache"),
				},
			)
		}
	}

	for _, i := range paths {
		_, err := os.Stat(i[1])
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				was_errors = true
				self.sys.log.Error(
					"  error at " + i[1] + " " + err.Error(),
				)
				continue
			}
		}

		c := exec.Command(i[0], "--update-cache")

		env := environ.NewFromStrings(os.Environ())
		env.Set("GDK_PIXBUF_MODULEDIR", i[1])
		env.Set("GDK_PIXBUF_MODULE_FILE", i[2])

		c.Env = env.Strings()
		err = c.Run()
		if err != nil {
			was_errors = true
			self.sys.log.Error(
				"  error at " + i[1] + " " + err.Error(),
			)
			continue
		}

	}

	if was_errors {
		return errors.New("errors while querrying pixbuf loaders")
	}

	return nil
}

// def pango_querymodules():
//     if not os.path.exists('/etc/pango'):
//         os.mkdir('/etc/pango')
//         logging.info('Created /etc/pango')
//     logging.info('pango-querymodules')
//     f = open('/etc/pango/pango.modules', 'wb')
//     r = subprocess.Popen(
//         ['pango-querymodules'], stdout=f
//         ).wait()
//     f.close()
//     return r

func (self *SystemUpdates) GlibCompileSchemas() error {
	roots, err := self._Roots()
	if err != nil {
		return err
	}

	self.sys.log.Info("Compile Glib Schemas")

	was_errors := false

	for _, i := range roots {
		pth := path.Join(i, "share", "glib-2.0", "schemas")
		pth_s, err := os.Stat(pth)
		if err == nil && pth_s.IsDir() {
			c := exec.Command("glib-compile-schemas", pth)
			res := c.Run()
			if res != nil {
				was_errors = true
				self.sys.log.Error(
					"  error at " + i + " " + err.Error(),
				)
			}
		}
	}

	if was_errors {
		return errors.New("was errors")
	}

	return nil
}

func (self *SystemUpdates) GtkQueryImmodules20() error {

	self.sys.log.Info("Querry Gtk 2.0 Immodules")

	// TODO: fix hardcoded path

	f, err := os.Create("/etc/gtk-2.0/gtk.immodules")
	if err != nil {
		return err
	}

	defer f.Close()

	c := exec.Command("gtk-query-immodules-2.0")
	c.Stdout = f

	r := c.Run()

	return r
}

func (self *SystemUpdates) GtkQueryImmodules30() error {

	self.sys.log.Info("Querry Gtk 3.0 Immodules")

	// TODO: fix uncerteinties

	c := exec.Command("gtk-query-immodules-3.0", "--update-cache")

	err := c.Run()

	if err != nil {
		self.sys.log.Error("  " + err.Error())
		return err
	}
	return nil
}

func (self *SystemUpdates) InstallEtc(log *logger.Logger) error {

	loginfo := func(txt string) {
		if log != nil {
			log.Info(txt)
		}
	}

	//	logerr := func(err error) {
	//		if log != nil {
	//			log.Error(err)
	//		}
	//	}

	r := self.sys.Root()

	dirs := set.NewSetString()

	an := etcfiles.AssetNames()

	sort.Strings(an)

	for _, i := range an {

		dir := path.Dir(i)

		dirs.Add(dir)

		dir_pth := path.Join(r, "/", dir)
		file_pth := path.Join(r, "/", i)

		loginfo("replacing " + file_pth)

		err := os.MkdirAll(dir_pth, 0755)
		if err != nil {
			return err
		}

		ass, err := etcfiles.Asset(i)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(file_pth, ass, 0755)
		if err != nil {
			return err
		}

		err = os.Chown(file_pth, 0, 0)
		if err != nil {
			return err
		}

		err = os.Chmod(file_pth, 0755)
		if err != nil {
			return err
		}

	}

	{
		dirs2 := set.NewSetString()

		for _, i := range dirs2.ListStrings() {
			t := i
			for {
				dirs2.Add(t)
				t = path.Dir(t)
				if t == "." {
					break
				}
			}
		}

		dirs = dirs2
	}

	ss := dirs.ListStrings()
	dirs = nil
	sort.Strings(ss)

	for _, i := range ss {
		dir_pth := path.Join(r, "/", i)

		err := os.Chown(dir_pth, 0, 0)
		if err != nil {
			return err
		}

		err = os.Chmod(dir_pth, 0755)
		if err != nil {
			return err
		}
	}

	loginfo("replacing /etc/mtab")
	p := path.Join(r, "/", "etc", "mtab")
	err := os.Remove(p)
	if err != nil {
		return err
	}

	err = os.Symlink("../proc/self/mounts", p)
	if err != nil {
		return err
	}

	return nil
}
