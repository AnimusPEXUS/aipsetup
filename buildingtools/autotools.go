package buildingtools

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/AnimusPEXUS/gologger"
)

type Autotools struct {
}

func (self *Autotools) Extract(
	filename string,
	srcdir string,
	tempdir string,
	unwrap bool,
	rename_dir bool,
	new_name string,
	is_more_than_one_extracted_ok bool,
	cleanup_srcdir bool,
	log *gologger.Logger,
) error {

	os.RemoveAll(tempdir)
	os.MkdirAll(tempdir, 0700)
	defer os.RemoveAll(tempdir)

	if cleanup_srcdir {
		os.RemoveAll(srcdir)
	}

	os.MkdirAll(srcdir, 0700)

	// TODO: mega fast decision. this should be replaced by internal functionality
	c := exec.Command("tar", "-xf", filename, "-C", tempdir)
	c_res := c.Run()

	if c_res != nil {
		return c_res
	}

	info, err := ioutil.ReadDir(tempdir)
	if err == nil {
		return err
	}

	if unwrap {

		if len(info) != 1 && !is_more_than_one_extracted_ok {
			return errors.New("extracted more than one item")
		}

		extracted_dir := ""

		for _, i := range info {
			if i.IsDir() {
				extracted_dir = i.Name()
				break
			}
		}

		if extracted_dir == "" {
			return errors.New("no directories extracted")
		}

		extracted_dir = path.Join(tempdir, extracted_dir)

		info, err = ioutil.ReadDir(extracted_dir)
		if err == nil {
			return err
		}

		for _, i := range info {
			os.Rename(
				path.Join(extracted_dir, i.Name()),
				path.Join(srcdir, i.Name()),
			)
		}

	} else {
		for _, i := range info {
			os.Rename(
				path.Join(tempdir, i.Name()),
				path.Join(srcdir, i.Name()),
			)
		}
	}

	return nil
}

func (self *Autotools) GenerateConfigureIfNeeded(
	directory string,
	force bool,
) error {

	presumed_configure_full_filename := path.Join(directory, "configure")

	is_presumed_configure_full_filename_exists := false

	if _, err := os.Stat(presumed_configure_full_filename); err == nil {
		is_presumed_configure_full_filename_exists = true
	}

	if !is_presumed_configure_full_filename_exists || force {
		// ('makeconf.sh', ['./makeconf.sh']),
		// ('autogen.sh', ['./autogen.sh']),
		// ('bootstrap.sh', ['./bootstrap.sh']),
		// ('bootstrap', ['./bootstrap']),
		// ('genconfig.sh', ['./genconfig.sh']),
		// ('configure.ac', ['autoreconf', '-i']),
		// ('configure.in', ['autoreconf', '-i']),

		{
			checked_file := path.Join(directory, "makeconf.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./makeconf.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "autogen.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./autogen.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "bootstrap.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./bootstrap.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "bootstrap")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./bootstrap")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "genconfig.sh")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("./genconfig.sh")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "configure.ac")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("autoconf", "-i")
				c.Dir = directory
				return c.Run()
			}
		}

		{
			checked_file := path.Join(directory, "configure.in")

			if _, err := os.Stat(checked_file); err == nil {
				c := exec.Command("autoconf", "-i")
				c.Dir = directory
				return c.Run()
			}
		}

		return errors.New(
			"no acceptable configure script creating tools found",
		)

	}

	return nil
}
