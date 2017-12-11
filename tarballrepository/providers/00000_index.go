package providers

import (
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/cache01"
)

func Get(
	name string,
	repo types.RepositoryI,
	pkg_name string,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	cache *cache01.CacheDir,
	args []string,
) (types.ProviderI, error) {
	if t, ok := Index[name]; ok {
		return t(
			repo,
			pkg_name,
			sys,
			tarballs_output_dir,
			cache,
			args,
		)
	} else {
		return nil, errors.New("provider not found")
	}
}

var Index = map[string](func(
	repo types.RepositoryI,
	pkg_name string,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	cache *cache01.CacheDir,
	args []string,
) (types.ProviderI, error)){
	"https": func(
		repo types.RepositoryI,
		pkg_name string,
		sys basictypes.SystemI,
		tarballs_output_dir string,
		cache *cache01.CacheDir,
		args []string,
	) (types.ProviderI, error) {
		return NewProviderHttps(
			repo,
			pkg_name,
			sys,
			tarballs_output_dir,
			cache,
			args,
		)
	},
}
