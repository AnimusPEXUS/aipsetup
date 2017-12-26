package providers

/*
SRS - stends for Source Revision System

this provider aims to keep local copy of git/svn/etc. and building tarballs
from them by request.

PackageInfo's provider attributes
=================================

  (! - necissery, ? - optional)

 !engine: for instance 'git', 'svn', etc..
 !uri: repository uri, for example "https://github.com/SELinuxProject/selinux.git"

 ?shared_repo:aipsetup_package_name: for example 'shared_repo:libselinux'
          will use srs directory under libselinux package's directory.
          If omited, equals to current package name

  needed_tag_re_prefix_is - string, regular expression, which allows to select
          prefix part of tag string

  needed_tag_re_suffix_is - same as needed_tag_re_prefix_is, but for suffix

  needed_tag_re - same as needed_tag_re_prefix_is, but for entire tag string

  tarball_format - extension for output tarballs string.
          ...only ".tar.xz" is supported



*/

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/infoeditor"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tags"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
)

var _ types.ProviderI = &ProviderSRS{}

type ProviderSRS struct {
	repo                types.RepositoryI
	pkg_name            string
	pkg_info            *basictypes.PackageInfo
	sys                 basictypes.SystemI
	tarballs_output_dir string
	cache               *cache01.CacheDir
	log                 *logger.Logger
}

func NewProviderSRS(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	cache *cache01.CacheDir,
	log *logger.Logger,
) (*ProviderSRS, error) {
	self := &ProviderSRS{
		repo:                repo,
		pkg_name:            pkg_name,
		pkg_info:            pkg_info,
		sys:                 sys,
		tarballs_output_dir: tarballs_output_dir,
		cache:               cache,
		log:                 log,
	}
	return self, nil
}

func (self *ProviderSRS) ProviderDescription() string {
	return "git/svn/etc..."
}

func (self *ProviderSRS) ArgCount() int {
	return 2
}

func (self *ProviderSRS) CanListArg(i int) bool {
	switch i {
	default:
		return false
	case 0:
		return true
	}
}

func (self *ProviderSRS) ListArg(i int) ([]string, error) {
	switch i {
	default:
		return []string{}, errors.New("not supported")
	case 0:
		return []string{"git"}, nil
	}
}

func (self *ProviderSRS) Tarballs() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderSRS) TarballNames() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderSRS) PerformUpdate() error {
	switch self.pkg_info.TarballProviderArguments[0] {
	default:
		return errors.New("not supported SRS system")
	case "git":
		p := self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2])

		self.log.Info(fmt.Sprintf("working inside %s", p))

		err := self.GetAndUpdateGit(
			p,
			self.pkg_info.TarballProviderArguments[1],
		)
		if err != nil {
			return err
		}

		err = self.MakeTarballsGit(
			self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2]),
			self.repo.GetPackageTarballsPath(self.pkg_name),
			self.pkg_info,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *ProviderSRS) GetAndUpdateGit(
	git_dir string,
	git_source_url string,
	// info *basictypes.PackageInfo,
) error {

	// git_source_url := info.TarballProviderArguments[2]
	git_subdir := path.Join(git_dir, ".git")

	new_download_mode := false

	if _, err := os.Stat(git_subdir); err == nil {
		new_download_mode = false
	} else {
		new_download_mode = true
	}

	var err error

	if new_download_mode {
		err = os.RemoveAll(git_dir)
		if err != nil {
			return err
		}

		err = os.MkdirAll(git_dir, 0700)
		if err != nil {
			return err
		}

		self.log.Info(fmt.Sprintf("getting %s", git_source_url))

		c := exec.Command("git", "clone", git_source_url, ".")
		c.Dir = git_dir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		return c.Run()

	} else {

		self.log.Info(fmt.Sprintf("updating %s", git_source_url))

		c := exec.Command("git", "pull")
		c.Dir = git_dir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		return c.Run()

	}

	return nil
}

func (self *ProviderSRS) MakeTarballsGit(
	git_dir string,
	output_dir string,
	info *basictypes.PackageInfo,
) error {

	basename := info.TarballName
	if basename == "" {
		basename = infoeditor.GithubDefaultTarballName
	}

	t := tags.New(info.TarballProviderArguments)

	needed_tag_re_prefix_is, _ := t.GetSingle("TagPrefixRegExp", true)
	if needed_tag_re_prefix_is == "" {
		needed_tag_re_prefix_is = infoeditor.GithubDefaultTagPrefixRegExp
	}

	needed_tag_re_suffix_is, _ := t.GetSingle("TagSuffixRegExp", true)
	if needed_tag_re_suffix_is == "" {
		needed_tag_re_suffix_is = infoeditor.GithubDefaultTagSuffixRegExp
	}

	needed_tag_re, _ := t.GetSingle("WholeTagRegExp", true)
	if needed_tag_re == "" {
		needed_tag_re = infoeditor.GithubDefaultWholeTagRegExp
	}

	tarball_format, _ := t.GetSingle("TarballFormat", true)
	if tarball_format == "" {
		tarball_format = infoeditor.GithubDefaultTarballFormat
	}

	truncate_versions := info.TarballProviderVersionSyncDepth
	if truncate_versions == 0 {
		truncate_versions = 3
	}

	acceptable_tags := make([]string, 0)

	{
		b := &bytes.Buffer{}
		c := exec.Command("git", "tag")
		c.Dir = git_dir
		c.Stdout = b
		err := c.Run()
		if err != nil {
			return err
		}

		tags := strings.Split(b.String(), "\n")

		// for ii, i := range tags {
		// 	fmt.Println(ii, i)
		// }

		parser, err := tarballnameparsers.Get(info.TarballFileNameParser)
		if err != nil {
			return err
		}

		for _, i := range tags {
			parse_res, err := parser.ParseName(i)
			if err != nil {
				// fmt.Println("tag parsing error:", err.Error())
				continue
			}

			if parse_res.Name != info.TarballName {
				continue
			}

			// if parse_res.Name != basename {
			// 	fmt.Println(parse_res.Name, "!=", basename)
			// 	continue
			// }

			fres, err := pkginfodb.ApplyInfoFilter(info, []string{i})
			if err != nil {
				return err
			}

			if len(fres) != 1 {
				continue
			}

			acceptable_tags = append(acceptable_tags, i)
		}
	}

	for ii, i := range acceptable_tags {
		fmt.Println(ii, i)
	}

	return nil
}
