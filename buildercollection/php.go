package buildercollection

import (
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/utils/logger"
)

func init() {
	Index["php"] = func(bs basictypes.BuildingSiteCtlI) (
		basictypes.BuilderI,
		error,
	) {
		return NewBuilder_php(bs)
	}
}

type Builder_php struct {
	*Builder_std
}

func NewBuilder_php(bs basictypes.BuildingSiteCtlI) (*Builder_php, error) {

	self := new(Builder_php)

	self.Builder_std = NewBuilder_std(bs)

	self.EditConfigureArgsCB = self.EditConfigureArgs

	return self, nil
}

func (self *Builder_php) EditConfigureArgs(log *logger.Logger, ret []string) ([]string, error) {

	ret = append(
		ret,
		[]string{
			"--enable-ftp",
			"--with-openssl",
			"--enable-mbstring",
			"--with-sqlite",
			"--enable-sqlite-utf8",
			"--with-pdo-sqlite",
			"--with-gd",
			"--with-jpeg-dir",
			"--with-png-dir",
			"--with-zlib-dir",
			"--with-ttf",
			"--with-freetype-dir",
			"--with-pdo-pgsql",
			"--with-pgsql",
			"--with-mysql",
			"--with-ncurses",
			"--with-pdo-mysql",
			"--with-mysqli",
			"--with-readline",
			// "--enable-embed",
			"--enable-fpm",
			"--enable-fastcgi",
			//            "--with-apxs2={}".format(
			//                wayround_i2p.utils.file.which(
			//                    "apxs",
			//                    self.get_host_dir()
			//                    )
			//                ),
		}...,
	)

	return ret, nil
}
