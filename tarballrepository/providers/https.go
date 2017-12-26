package providers

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/htmlwalk"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/version"
)

type ProviderHttps struct {
	repo                types.RepositoryI
	pkg_name            string
	pkg_info            *basictypes.PackageInfo
	sys                 basictypes.SystemI
	tarballs_output_dir string
	cache               *cache01.CacheDir

	args []string

	htw *htmlwalk.HTMLWalk

	scheme string
	host   string
	path   string

	log *logger.Logger
}

func NewProviderHttps(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	cache *cache01.CacheDir,
	log *logger.Logger,
) (*ProviderHttps, error) {

	ret := new(ProviderHttps)
	ret.repo = repo
	ret.pkg_name = pkg_name
	ret.pkg_info = pkg_info
	ret.sys = sys
	ret.tarballs_output_dir = tarballs_output_dir
	ret.cache = cache

	ret.log = log

	ret.args = pkg_info.TarballProviderArguments

	switch len(ret.args) {
	case 0:
	case 1:
		u, err := url.Parse(ret.args[0])
		if err != nil {
			return nil, err
		}
		ret.scheme = u.Scheme
		ret.host = u.Host
		ret.path = u.Path
	default:
		return nil, errors.New("invalid arguments number")
	}

	return ret, nil
}

func (self *ProviderHttps) ProviderDescription() string {
	return ""
}

func (self *ProviderHttps) ArgCount() int {
	return 1
}

func (self *ProviderHttps) CanListArg(i int) bool {
	return false
}

func (self *ProviderHttps) ListArg(i int) ([]string, error) {
	return []string{}, errors.New("not supported")
}

func (self *ProviderHttps) TarballNames() ([]string, error) {
	return make([]string, 0), nil
}

func (self *ProviderHttps) Tarballs() ([]string, error) {
	return make([]string, 0), nil
}

func (self *ProviderHttps) _GetHTW() (*htmlwalk.HTMLWalk, error) {
	if self.htw == nil {

		h, err := htmlwalk.NewHTMLWalk(
			self.scheme,
			self.host,
			self.cache,
			self.log,
		)
		if err != nil {
			return nil, err
		}
		self.htw = h
	}
	return self.htw, nil
}

func (self *ProviderHttps) PerformUpdate() error {
	htw, err := self._GetHTW()
	if err != nil {
		return err
	}

	tree, err := htw.Tree(self.path)
	if err != nil {
		return err
	}

	tree_keys := make([]string, 0)
	for k, _ := range tree {
		tree_keys = append(tree_keys, k)
	}

	sort.Strings(tree_keys)

	filtered_keys := make([]string, 0)

	parser, err := tarballnameparsers.Get(self.pkg_info.TarballFileNameParser)
	if err != nil {
		return err
	}

	for _, i := range tree_keys {
		if strings.HasSuffix(i, "/") {
			continue
		}

		err := tarballname.IsPossibleTarballNameErr(i)
		if err != nil {
			continue
		}

		parse_res, err := parser.ParseName(i)
		if err != nil {
			continue
		}

		if parse_res.Name != self.pkg_info.TarballName {
			continue
		}

		fres, err := pkginfodb.ApplyInfoFilter(self.pkg_info, []string{i})
		if err != nil {
			return err
		}

		if len(fres) != 1 {
			continue
		}

		filtered_keys = append(filtered_keys, i)
	}

	self.log.Info("tarball list gotten from site")
	for _, i := range filtered_keys {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	version_tree, err := version.NewVersionTree(
		self.pkg_info.TarballName,
		self.pkg_info.TarballFileNameParser,
	)
	if err != nil {
		return err
	}

	for _, i := range filtered_keys {
		version_tree.Add(path.Base(i))
	}

	depth := self.pkg_info.TarballProviderVersionSyncDepth
	if depth == 0 {
		depth = 3
	}

	version_tree.TruncateByVersionDepth(nil, depth)

	self.log.Info("-----------------")
	self.log.Info("tarball versioned truncation result")

	res := version_tree.Basenames(tarballname.ACCEPTABLE_TARBALL_EXTENSIONS)
	for _, i := range res {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	err = version.SortByVersion(res, parser)
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
	self.log.Info("to download")

	for _, i := range res {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	downloading_errors := 0
	for _, i := range res {
		uri, err := self.GetDownloadingURIForTarball(i)
		if err != nil {
			return err
		}

		res_err := self.repo.PerformDownload(self.pkg_name, i, uri)
		if res_err != nil {
			downloading_errors++
		}
	}

	if downloading_errors != 0 {
		return errors.New("some files hasn't been downloaded successfully")
	}

	// WARNING: do not move existing tarballs deletion before download!
	//          deletions should be done only if all downloads done successfully!
	lst, err := self.repo.ListLocalTarballs(self.pkg_name)
	if err != nil {
		return err
	}

	to_delete := make([]string, 0)

	for _, i := range lst {
		found := false
		for _, j := range res {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			to_delete = append(to_delete, i)
		}
	}

	self.log.Info("-----------------")
	self.log.Info("to delete")

	for _, i := range to_delete {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	for _, i := range to_delete {
		err = self.repo.DeleteFile(self.pkg_name, i)
		if err != nil {
			return err
		}
	}

	// needed_tarball_name := self.pkg_info.TarballName
	// var tarball_name_parser tarballnameparsers.TarballNameParserI
	// if tarball_name_parser_c, ok :=
	// 	tarballnameparsers.Index[self.pkg_info.TarballFileNameParser]; !ok {
	// 	return errors.New("tarball name parser not found")
	// } else {
	// 	tarball_name_parser = tarball_name_parser_c()
	// }
	//
	// for k, v := range tree {
	//
	// }
	//
	// tarballs := make([]string, 0)

	return nil
}

func (self *ProviderHttps) GetDownloadingURIForTarball(name string) (string, error) {
	name = path.Base(name)

	htw, err := self._GetHTW()
	if err != nil {
		return "", err
	}

	tree, err := htw.Tree(self.path)
	if err != nil {
		return "", err
	}

	for k, _ := range tree {
		if path.Base(k) == name {
			u := &url.URL{
				Scheme: self.scheme,
				Host:   self.host,
				Path:   k,
			}
			return u.String(), nil
		}
	}

	return "", errors.New("not found")
}
