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

	!tag_parser:string - name of parser to retrive name
	!tag_name:string - value thich should match 'name' result of
					tag_parser:string work.

 	?shared_repo:aipsetup_package_name: for example 'shared_repo:libselinux'
          will use srs directory under libselinux package's directory.
          If omited, equals to current package name

  ?tarball_format - extension for output tarballs string.
          ...only ".tar.xz" is supported



*/

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tags"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/version"
	"github.com/AnimusPEXUS/utils/version/versioncomparators"
)

// const GithubDefaultTagParser = "std"
// const GithubDefaultTagName = "v"
// const GithubDefaultTagStatus = `^$`

var _ types.ProviderI = &ProviderSRS{}

type ProviderSRS struct {
	repo                types.RepositoryI
	pkg_name            string
	pkg_info            *basictypes.PackageInfo
	sys                 basictypes.SystemI
	tarballs_output_dir string
	log                 *logger.Logger
}

func NewProviderSRS(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (*ProviderSRS, error) {
	self := &ProviderSRS{
		repo:                repo,
		pkg_name:            pkg_name,
		pkg_info:            pkg_info,
		sys:                 sys,
		tarballs_output_dir: tarballs_output_dir,
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

	t := tags.New(info.TarballProviderArguments)

	{
		tp := self.repo.GetPackageTarballsPath(self.pkg_name)

		err := os.MkdirAll(tp, 0700)
		if err != nil {
			return err
		}
	}

	TagParser, _ := t.GetSingle("TagParser", true)
	if TagParser == "" {
		TagParser = self.pkg_info.TarballFileNameParser
	}

	TagName, _ := t.GetSingle("TagName", true)
	if TagName == "" {
		TagName = "v"
	}

	TagStatus, _ := t.GetSingle("TagStatus", true)
	if TagStatus == "" {
		TagStatus = "^$"
	}

	TagComparator, _ := t.GetSingle("TagComparator", true)
	if TagComparator == "" {
		TagComparator = self.pkg_info.TarballVersionComparator
	}

	TagFilters, TagFiltersUse := t.GetSingle("TagFilters", true)

	// TODO: do srs SyncDepth should also be moved from InfoDB to srs args?
	truncate_versions := info.TarballProviderVersionSyncDepth
	if truncate_versions == 0 {
		truncate_versions = 3
	}

	parser, err := tarballnameparsers.Get(TagParser)
	if err != nil {
		return err
	}

	comparator, err := versioncomparators.Get(TagComparator)
	if err != nil {
		return err
	}

	acceptable_tags := make([]string, 0)

	{
		var tags []string

		{
			b := &bytes.Buffer{}
			c := exec.Command("git", "tag")
			c.Dir = git_dir
			c.Stdout = b
			err := c.Run()
			if err != nil {
				return err
			}
			tags = strings.Split(b.String(), "\n")
		}

		for _, i := range tags {
			parse_res, err := parser.Parse(i)
			if err != nil {
				continue
			}

			matched, err := regexp.MatchString(TagName, parse_res.Name)
			if err != nil {
				return err
			}

			if !matched {
				continue
			}

			matched, err = regexp.MatchString(TagStatus, parse_res.Status.DirtyStr)
			if err != nil {
				return err
			}

			if !matched {
				continue
			}

			if TagFiltersUse {
				info := self.pkg_info
				switch TagFilters {
				case "+":
					info = self.pkg_info
				default:
					var err error
					info, err = pkginfodb.Get(TagFilters)
					if err != nil {
						return errors.New("can't get named info filters for srs")
					}
				}
				fres, err := pkginfodb.ApplyInfoFilter(
					info,
					[]string{i},
				)
				if err != nil {
					return err
				}

				if len(fres) != 1 {
					continue
				}
			}

			acceptable_tags = append(acceptable_tags, i)
		}
	}

	{

		version_tree, err := version.NewVersionTree(
			TagName,
			parser,
			comparator,
		)
		if err != nil {
			return err
		}

		for _, i := range acceptable_tags {
			b := path.Base(i)

			err = version_tree.Add(b)
			if err != nil {
				return err
			}
		}

		depth := self.pkg_info.TarballProviderVersionSyncDepth
		if depth == 0 {
			depth = 3
		}

		self.log.Info("-----------------")
		self.log.Info("tags before versioned truncation")

		res, err := version_tree.Basenames([]string{""})
		if err != nil {
			return err
		}
		for _, i := range res {
			self.log.Info(fmt.Sprintf("  %s", i))
		}

		version_tree.TruncateByVersionDepth(nil, depth)

		self.log.Info("-----------------")
		self.log.Info("tag versioned truncation result")

		res, err = version_tree.Basenames([]string{""})
		if err != nil {
			return err
		}
		for _, i := range res {
			self.log.Info(fmt.Sprintf("  %s", i))
		}

		err = comparator.SortStrings(res, parser)
		if err != nil {
			return err
		}

		self.log.Info("-----------------")
		self.log.Info("sorted by version")

		for _, i := range res {
			self.log.Info(fmt.Sprintf("  %s", i))
		}

		{
			len_res := len(res)
			t := make([]string, len_res)
			for i := range res {
				t[i] = res[len_res-i-1]
			}
			res = t
		}

		self.log.Info("-----------------")
		self.log.Info("to archive")

		for _, i := range res {
			self.log.Info(fmt.Sprintf("  %s", i))
		}

		acceptable_tags = res

	}

	was_errors := false
	downloaded_files := make([]string, 0)

	{
		for _, i := range acceptable_tags {

			i_parsed, err := parser.Parse(i)
			if err != nil {
				return err
			}

			tag_filename_noext, err := parser.Render(&tarballname.ParsedTarballName{
				Name:    info.TarballName,
				Version: i_parsed.Version,
				Status:  i_parsed.Status,
			})
			if err != nil {
				return err
			}

			tag_filename := tag_filename_noext + ".tar.xz"

			tag_filename_done := self.repo.GetTarballDoneFilePath(
				self.pkg_name,
				tag_filename,
			)

			if _, err := os.Stat(tag_filename_done); err == nil {
				downloaded_files = append(downloaded_files, tag_filename)
				continue
			}

			target_file := self.repo.GetTarballFilePath(self.pkg_name, tag_filename)

			self.log.Info(fmt.Sprintf("  archiving %s (%s)", i, tag_filename))

			c := exec.Command(
				"git",
				"archive",
				fmt.Sprintf("--prefix=%s/", tag_filename_noext),
				"-o", target_file,
				i,
			)
			c.Dir = git_dir
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			if c.Run() != nil {
				was_errors = true
			} else {
				downloaded_files = append(downloaded_files, tag_filename)
				if f, err := os.Create(tag_filename_done); err != nil {
					return err
				} else {
					f.Close()
				}
			}

		}
	}

	if was_errors {
		return errors.New("there was errors making tarballs")
	}

	// TODO: maybe rest of this func shold be moved to PerformUpdate

	lst, err := self.repo.PrepareTarballCleanupListing(self.pkg_name, downloaded_files)
	if err != nil {
		return err
	}

	self.log.Info("-----------------")
	self.log.Info("to delete")

	for _, i := range lst {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	err = self.repo.DeleteFiles(self.pkg_name, lst)
	if err != nil {
		return err
	}

	return nil
}
