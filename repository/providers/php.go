package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/pkginfodb"
	"github.com/AnimusPEXUS/aipsetup/repository/types"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/logger"
	"github.com/AnimusPEXUS/utils/set"
	"github.com/AnimusPEXUS/utils/tarballname"
	"github.com/AnimusPEXUS/utils/tarballname/tarballnameparsers"
	"github.com/AnimusPEXUS/utils/tarballstabilityclassification"
	"github.com/AnimusPEXUS/utils/tarballversion"
	"github.com/AnimusPEXUS/utils/tarballversion/versioncomparators"
	"github.com/antchfx/htmlquery"
)

// "/get/php-7.2.10.tar.bz2/from/a/mirror"
var PHP_MIRROR_DOWNLOADING_URI = regexp.MustCompile(`^/get/(php-((?:\d+.?)+)\.tar(\.?[a-zA-Z0-9]+))/from/a/mirror$`)

var _ types.ProviderI = &ProviderPHPnet{}

func init() {
	Index["php.net"] = NewProviderPHPnet
}

type ProviderPHPnet struct {
	repo                types.RepositoryI
	pkg_name            string
	pkg_info            *basictypes.PackageInfo
	sys                 basictypes.SystemI
	tarballs_output_dir string
	log                 *logger.Logger

	cache *cache01.CacheDir
}

func NewProviderPHPnet(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (types.ProviderI, error) {

	self := &ProviderPHPnet{
		repo:                repo,
		pkg_name:            pkg_name,
		pkg_info:            pkg_info,
		sys:                 sys,
		tarballs_output_dir: tarballs_output_dir,
		log:                 log,
	}

	if t, err := cache01.NewCacheDir(
		path.Join(
			self.repo.GetCachesDir(),
			"php.net",
		),
		nil,
	); err != nil {
		return nil, err
	} else {
		self.cache = t
	}

	return self, nil
}

func (self *ProviderPHPnet) ProviderDescription() string {
	return "download tarballs from php.net"
}

func (self *ProviderPHPnet) ArgCount() int {
	return 0
}

func (self *ProviderPHPnet) CanListArg(i int) bool {
	return false
}

func (self *ProviderPHPnet) ListArg(i int) ([]string, error) {
	return []string{}, errors.New("not supported")
}

func (self *ProviderPHPnet) Tarballs() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderPHPnet) TarballNames() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderPHPnet) LogI(txt string) {
	if self.log != nil {
		self.log.Info(txt)
	}
}

func (self *ProviderPHPnet) LogE(txt string) {
	if self.log != nil {
		self.log.Error(txt)
	}
}

func (self *ProviderPHPnet) readPageNC() ([]string, error) {

	self.LogI("updating php.net downloads cache")

	u := &url.URL{
		Scheme: "https",
		Host:   "secure.php.net",
		Path:   "/downloads.php",
	}

	http_res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)
	_, err = io.Copy(b, http_res.Body)
	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(b)

	ret := set.NewSetString()

	file_list_table_res := htmlquery.Find(doc, `.//a`)

	for _, i := range file_list_table_res {

		href := ""

		for _, j := range i.Attr {
			if j.Key == "href" {
				href = j.Val
				break
			}
		}

		if href == "" {
			continue
		}

		if PHP_MIRROR_DOWNLOADING_URI.MatchString(href) {
			href_basename :=
				path.Base(PHP_MIRROR_DOWNLOADING_URI.FindStringSubmatch(href)[1])
			if tarballname.IsPossibleTarballName(href_basename) {
				ret.AddStrings(href_basename)
			}
		}

	}

	t := ret.ListStrings()
	ret = nil

	sort.Strings(t)

	return t, nil
}

func (self *ProviderPHPnet) readPage() ([]string, error) {
	c, err := self.cache.Cache(
		"downloads_page",
		func() ([]byte, error) {
			res, err := self.readPageNC()
			if err != nil {
				return nil, err
			}

			ret, err := json.Marshal(res)
			if err != nil {
				return nil, err
			}

			return ret, nil
		},
	)

	ret := make([]string, 0)

	res, err := c.GetValue()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (self *ProviderPHPnet) PerformUpdate() error {

	// TODO: PerformUpdate() functions of providers have many common code. some
	//       simple unification required.

	page_parse_result, err := self.readPage()
	if err != nil {
		return err
	}

	filtered_page_parse_result := make([]string, 0)

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

	for _, i := range page_parse_result {
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

		{
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
		}

		if ok, err := stability_classifier.IsStable(parse_res); err != nil {
			return err
		} else {
			if !ok {
				continue
			}
		}

		filtered_page_parse_result = append(filtered_page_parse_result, i)
	}

	self.log.Info("tarball list gotten from site")
	for _, i := range filtered_page_parse_result {
		self.log.Info(fmt.Sprintf("  %s", i))
	}

	version_tree, err := tarballversion.NewVersionTree(
		self.pkg_info.TarballName,
		false,
		parser,
		comparator,
	)
	if err != nil {
		return err
	}

	for _, i := range filtered_page_parse_result {
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

func (self *ProviderPHPnet) GetDownloadingURIForFile(name string) (string, error) {

	name = path.Base(name)

	//	_, info, err := pkginfodb.DetermineTarballPackageInfoSingle(name)
	//	if err != nil {
	//		return "", err
	//	}

	//	parser, err := tarballnameparsers.Get(info.TarballFileNameParser)
	//	if err != nil {
	//		return "", err
	//	}

	//	parsed, err := parser.Parse(name)
	//	if err != nil {
	//		return "", err
	//	}

	// https://php.net/get/php-7.2.10.tar.xz/from/this/mirror

	l := fmt.Sprintf(
		"https://php.net/get/%s/from/this/mirror",
		name,
		//		parsed.Version.StrSliceString("."),
	)

	return l, nil
}
