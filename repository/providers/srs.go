package providers

/*

SRS - stends for Source Revision System

this provider aims to keep local copy of git/svn/etc. and building tarballs
from them by request.

PackageInfo provider's attributes
=================================

  #0 engine name, for instance 'git', 'svn', etc.. (only git is currently supported)

  #1 repository url (required)

  #2 PackageInfo name, which use to place srs repository

  if engine name equals to git or hg:

	TagParser default is equal to TarballFileNameParser

	TagName regexp default is "^v"

	TagStatus regexp default is "^$"

	TagComparator default equals to TarballVersionComparator

	TagFilters if defined, shold be other package info name - it's filters will
		be used

	TagTarballRenderer name of tarball name parser, which shoild be used to render
		tarball name for tag. default is equal to TagParser

	EnableSubmodules (only git) get submodules and use git-archive-all python script
		to archive tarballs

*/

import (
	"errors"
	"fmt"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/repository/types"
	"github.com/AnimusPEXUS/utils/logger"
)

// const GithubDefaultTagParser = "std"
// const GithubDefaultTagName = "v"
// const GithubDefaultTagStatus = `^$`

var _ types.ProviderI = &ProviderSRS{}

func init() {
	Index["srs"] = NewProviderSRS
}

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
) (types.ProviderI, error) {
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
		//		return []string{"git", "svn"}, nil
		return []string{"git", "hg"}, nil
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
	case "hg":

		subtool := NewSRSHg(self)

		p := self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2])

		self.log.Info(fmt.Sprintf("working srs dir is %s", p))

		err := subtool.GetAndUpdate(
			p,
			self.pkg_info.TarballProviderArguments[1],
		)
		if err != nil {
			return err
		}

		err = subtool.MakeTarballs(
			self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2]),
			self.repo.GetPackageTarballsPath(self.pkg_name),
		)
		if err != nil {
			return err
		}
	case "git":

		subtool := NewSRSGit(self)

		p := self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2])

		self.log.Info(fmt.Sprintf("working srs dir is %s", p))

		err := subtool.GetAndUpdate(
			p,
			self.pkg_info.TarballProviderArguments[1],
		)
		if err != nil {
			return err
		}

		err = subtool.MakeTarballs(
			self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2]),
			self.repo.GetPackageTarballsPath(self.pkg_name),
		)
		if err != nil {
			return err
		}
		//	case "svn":
		//		p := self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2])

		//		self.log.Info(fmt.Sprintf("working inside %s", p))

		//		err := self.GetAndUpdateSvn(
		//			p,
		//			self.pkg_info.TarballProviderArguments[1],
		//		)
		//		if err != nil {
		//			return err
		//		}

		//		err = self.MakeTarballsGit(
		//			self.repo.GetPackageSRSPath(self.pkg_info.TarballProviderArguments[2]),
		//			self.repo.GetPackageTarballsPath(self.pkg_name),
		//			self.pkg_info,
		//		)
		//		if err != nil {
		//			return err
		//		}
	}
	return nil
}
