package providers

import (
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/launchpadnetwalk"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/version"
	"github.com/AnimusPEXUS/utils/version/versioncomparators"
)

var _ types.ProviderI = &ProviderLaunchpadNet{}

type ProviderLaunchpadNet struct {
	repo                types.RepositoryI
	pkg_name            string
	pkg_info            *basictypes.PackageInfo
	sys                 basictypes.SystemI
	tarballs_output_dir string
	log                 *logger.Logger

	cache *cache01.CacheDir

	lpw *launchpadnetwalk.LaunchpadNetWalk

	project string
}

func NewProviderLaunchpadNet(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (*ProviderLaunchpadNet, error) {

	self := &ProviderLaunchpadNet{
		repo:                repo,
		pkg_name:            pkg_name,
		pkg_info:            pkg_info,
		sys:                 sys,
		tarballs_output_dir: tarballs_output_dir,
		log:                 log,
	}

	switch len(pkg_info.TarballProviderArguments) {
	case 0:
	case 1:
		self.project = pkg_info.TarballProviderArguments[0]
	default:
		return nil, errors.New("invalid arguments count")
	}

	if t, err := cache01.NewCacheDir(
		path.Join(
			self.repo.GetCachesDir(),
			"launchpad.net",
			self.project,
		),
		nil,
	); err != nil {
		return nil, err
	} else {
		self.cache = t
	}

	return self, nil
}

func (self *ProviderLaunchpadNet) ProviderDescription() string {
	return "launchpad.net"
}

func (self *ProviderLaunchpadNet) ArgCount() int {
	return 1
}

func (self *ProviderLaunchpadNet) CanListArg(i int) bool {
	return false
}

func (self *ProviderLaunchpadNet) ListArg(i int) ([]string, error) {
	return []string{}, errors.New("not supported")
}

func (self *ProviderLaunchpadNet) Tarballs() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderLaunchpadNet) TarballNames() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderLaunchpadNet) _GetLPW() (*launchpadnetwalk.LaunchpadNetWalk, error) {
	if self.lpw == nil {

		h, err := launchpadnetwalk.NewLaunchpadNetWalk(
			self.project,
			self.cache,
			self.log,
		)
		if err != nil {
			return nil, err
		}
		self.lpw = h
	}
	return self.lpw, nil
}

func (self *ProviderLaunchpadNet) PerformUpdate() error {
	htw, err := self._GetLPW()
	if err != nil {
		return err
	}

	tree, err := htw.Tree("/")
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
		parser,
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

	err = self.repo.DeleteFiles(self.pkg_name, lst)
	if err != nil {
		return err
	}
	return nil
}

func (self *ProviderLaunchpadNet) GetDownloadingURIForFile(name string) (string, error) {
	name = path.Base(name)

	htw, err := self._GetLPW()
	if err != nil {
		return "", err
	}

	return htw.GetDownloadingURIForFile(name)
}
