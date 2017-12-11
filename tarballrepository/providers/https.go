package providers

import (
	"errors"
	"net/url"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/cache01"
	"github.com/AnimusPEXUS/utils/htmlwalk"
)

type ProviderHttps struct {
	repo                types.RepositoryI
	pkg_name            string
	sys                 basictypes.SystemI
	tarballs_output_dir string
	cache               *cache01.CacheDir
	args                []string

	pkg_info *basictypes.PackageInfo

	htw *htmlwalk.HTMLWalk

	scheme string
	host   string
	path   string
}

func NewProviderHttps(
	repo types.RepositoryI,
	pkg_name string,
	// pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	cache *cache01.CacheDir,
	args []string,
) (*ProviderHttps, error) {

	ret := new(ProviderHttps)
	ret.repo = repo
	ret.sys = sys
	ret.tarballs_output_dir = tarballs_output_dir
	ret.cache = cache
	ret.args = args

	pkg_info, err := distropkginfodb.Get(pkg_name)
	if err != nil {
		return nil, err
	}

	ret.pkg_info = pkg_info

	switch len(args) {
	case 0:
	case 1:
		u, err := url.Parse(args[0])
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
		)
		if err != nil {
			return nil, err
		}
		self.htw = h
	}
	return self.htw, nil
}

func (self *ProviderHttps) PerformUpdate() error {
	// htw, err := self._GetHTW()
	// if err != nil {
	// 	return err
	// }
	//
	// // tree, err := htw.Tree(self.path)
	// // if err != nil {
	// // 	return err
	// // }
	//
	// // needed_tarball_name := self.pkg_info.TarballName
	// // var tarball_name_parser tarballnameparsers.TarballNameParserI
	// if tarball_name_parser_c, ok :=
	// 	tarballnameparsers.Index[self.pkg_info.TarballFileNameParser]; !ok {
	// 	return errors.New("tarball name parser not found")
	// } else {
	// 	// tarball_name_parser = tarball_name_parser_c()
	// }
	//
	// // for k, v := range tree {
	// //
	// // }
	//
	// tarballs := make([]string, 0)

	return nil
}
