package providers

/*
SRS - stends for Source Revision System

this provider aims to keep local copy of git/svn/etc. and building tarballs
from them by request.

PackageInfo's provider attributes
=================================

  (! - necissery, ? - optional)

 !engine: for instance 'git', 'svn', etc..
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
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tags"
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
		err := self.GetAndUpdateGit(
			self.repo.GetPackageSRSPath(self.pkg_name),
			self.pkg_info.TarballProviderArguments[1],
		)
		if err != nil {
			return err
		}

		t := tags.New(self.pkg_info.TarballProviderArguments)

		needed_tag_re_prefix_is, _ := t.GetSingle("needed_tag_re_prefix_is", true)
		needed_tag_re_suffix_is, _ := t.GetSingle("needed_tag_re_suffix_is", true)
		needed_tag_re, _ := t.GetSingle("needed_tag_re", true)
		tarball_format, _ := t.GetSingle("tarball_format", true)

		err = self.MakeTarballsGit(
			self.repo.GetPackageSRSPath(self.pkg_name),
			self.repo.GetPackageTarballsPath(self.pkg_name),
			self.pkg_info.TarballName,
			needed_tag_re_prefix_is,
			needed_tag_re_suffix_is,
			needed_tag_re,
			tarball_format,
			self.pkg_info.TarballProviderVersionSyncDepth,
		)
	}
	return nil
}

func (self *ProviderSRS) GetAndUpdateGit(git_dir string, git_source_url string) error {
	return nil
}

func (self *ProviderSRS) MakeTarballsGit(
	git_dir string,
	output_dir string,
	basename string,
	needed_tag_re_prefix_is string,
	needed_tag_re_suffix_is string,
	needed_tag_re string,
	tarball_format string,
	truncate_versions int,
) error {
	if basename == "" {
		basename = "v"
	}

	if needed_tag_re_prefix_is == "" {
		needed_tag_re_prefix_is = "v"
	}

	if needed_tag_re_suffix_is == "" {
		needed_tag_re_suffix_is = `^$`
	}

	if needed_tag_re == "" {
		needed_tag_re =
			`^((?P<prefix>.*?)[\-\_]?)?(?P<version>\d+(?P<delim>[\_\-\.])?` +
				`(\d+(?P=delim)?)*)([\-\_]??(?P<suffix>.*?)??)??$`
	}

	if tarball_format == "" {
		tarball_format = "tar.xz"
	}

	if truncate_versions == 0 {
		truncate_versions = 3
	}

	return nil
}
