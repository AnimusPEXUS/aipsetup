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
	"github.com/AnimusPEXUS/aipsetup/repository/types"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/gnupghtmlftpwalk"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
	"github.com/AnimusPEXUS/utils/version"
	"github.com/AnimusPEXUS/utils/version/versioncomparators"
)

var _ types.ProviderI = &ProviderGNUPGOrg{}

func init() {
	Index["gnupg.org"] = NewProviderGNUPGOrg
}

type ProviderGNUPGOrg struct {
	repo                types.RepositoryI
	pkg_name            string
	pkg_info            *basictypes.PackageInfo
	sys                 basictypes.SystemI
	tarballs_output_dir string
	log                 *logger.Logger

	cache *cache01.CacheDir

	gpgw *gnupghtmlftpwalk.Walk

	scheme string
	host   string
	path   string
}

func NewProviderGNUPGOrg(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (types.ProviderI, error) {

	self := &ProviderGNUPGOrg{
		repo:                repo,
		pkg_name:            pkg_name,
		pkg_info:            pkg_info,
		sys:                 sys,
		tarballs_output_dir: tarballs_output_dir,
		log:                 log,
	}

	u, err := url.Parse(pkg_info.TarballProviderArguments[0])
	if err != nil {
		return nil, err
	}
	self.scheme = u.Scheme
	self.host = u.Host
	self.path = u.Path

	cache_uri := (&url.URL{
		Scheme: self.scheme,
		Host:   self.host,
	}).String()

	if t, err := cache01.NewCacheDir(
		path.Join(
			self.repo.GetCachesDir(),
			"gnupg.org",
			url.PathEscape(cache_uri),
		),
		nil,
	); err != nil {
		return nil, err
	} else {
		self.cache = t
	}

	return self, nil
}

func (self *ProviderGNUPGOrg) ProviderDescription() string {
	return ""
}

func (self *ProviderGNUPGOrg) ArgCount() int {
	return 1
}

func (self *ProviderGNUPGOrg) CanListArg(i int) bool {
	return false
}

func (self *ProviderGNUPGOrg) ListArg(i int) ([]string, error) {
	return []string{}, errors.New("not supported")
}

func (self *ProviderGNUPGOrg) Tarballs() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderGNUPGOrg) TarballNames() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderGNUPGOrg) _GetGPGW() (*gnupghtmlftpwalk.Walk, error) {
	if self.gpgw == nil {

		h, err := gnupghtmlftpwalk.NewWalk(
			self.scheme,
			self.host,
			self.cache,
			self.log,
		)
		if err != nil {
			return nil, err
		}
		self.gpgw = h
	}
	return self.gpgw, nil
}

func (self *ProviderGNUPGOrg) PerformUpdate() error {
	htw, err := self._GetGPGW()
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

	comparator, err := versioncomparators.Get(self.pkg_info.TarballVersionComparator)
	if err != nil {
		return err
	}

	stability_classifier, err := tarballstabilityclassification.Get(
		self.pkg_info.TarballStabilityClassifier,
	)
	if err != nil {
		return err
	}

	for _, i := range tree_keys {
		if strings.HasSuffix(i, "/") {
			continue
		}

		if !tarballname.IsPossibleTarballName(i) {
			continue
		}

		parse_res, err := parser.Parse(i)
		if err != nil {
			continue
		}

		if parse_res.Name != self.pkg_info.TarballName {
			continue
		}

		fres, err := pkginfodb.ApplyInfoFilter(
			self.pkg_info,
			[]string{i},
		)
		if err != nil {
			return err
		}

		if len(fres) != 1 {
			continue
		}

		if ok, err := stability_classifier.IsStable(parse_res); err != nil {
			return err
		} else {
			if !ok {
				continue
			}
		}

		filtered_keys = append(filtered_keys, i)
	}

	self.log.Info("tarball list gotten from site")
	for _, i := range filtered_keys {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	version_tree, err := version.NewVersionTree(
		self.pkg_info.TarballName,
		parser,
		comparator,
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

	res, err := version_tree.Basenames(tarballname.ACCEPTABLE_TARBALL_EXTENSIONS)
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
	self.log.Info("to download")

	for _, i := range res {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	downloading_errors := false
	for _, i := range res {
		uri, err := self.GetDownloadingURIForFile(i)
		if err != nil {
			return err
		}

		res_err := self.repo.PerformDownload(self.pkg_name, i, uri)
		if res_err != nil {
			downloading_errors = true
		}
	}

	if downloading_errors {
		return errors.New("some files hasn't been downloaded successfully")
	}

	// WARNING: do not move existing tarballs deletion before download!
	//          deletions should be done only if all downloads done successfully!
	lst, err := self.repo.PrepareTarballCleanupListing(self.pkg_name, res)
	if err != nil {
		return err
	}

	self.log.Info("-----------------")
	self.log.Info("to delete")

	for _, i := range lst {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	err = self.repo.DeleteTarballFiles(self.pkg_name, lst)
	if err != nil {
		return err
	}
	return nil
}

func (self *ProviderGNUPGOrg) GetDownloadingURIForFile(name string) (string, error) {
	name = path.Base(name)

	htw, err := self._GetGPGW()
	if err != nil {
		return "", err
	}

	return htw.GetDownloadingURIForFile(name, self.path)
}
