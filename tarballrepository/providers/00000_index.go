package providers

import (
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/logger"
)

func Get(
	name string,
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (types.ProviderI, error) {
	if t, ok := Index[name]; ok {
		return t(
			repo,
			pkg_name,
			pkg_info,
			sys,
			tarballs_output_dir,
			log,
		)
	} else {
		return nil, errors.New("provider not found")
	}
}

var Index = map[string](func(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (types.ProviderI, error)){
	"https": func(
		repo types.RepositoryI,
		pkg_name string,
		pkg_info *basictypes.PackageInfo,
		sys basictypes.SystemI,
		tarballs_output_dir string,
		log *logger.Logger,
	) (types.ProviderI, error) {
		return NewProviderHttps(
			repo,
			pkg_name,
			pkg_info,
			sys,
			tarballs_output_dir,
			log,
		)
	},
	"sf.net": func(
		repo types.RepositoryI,
		pkg_name string,
		pkg_info *basictypes.PackageInfo,
		sys basictypes.SystemI,
		tarballs_output_dir string,
		log *logger.Logger,
	) (types.ProviderI, error) {
		return NewProviderSFNet(
			repo,
			pkg_name,
			pkg_info,
			sys,
			tarballs_output_dir,
			log,
		)
	},
	"launchpad.net": func(
		repo types.RepositoryI,
		pkg_name string,
		pkg_info *basictypes.PackageInfo,
		sys basictypes.SystemI,
		tarballs_output_dir string,
		log *logger.Logger,
	) (types.ProviderI, error) {
		return NewProviderLaunchpadNet(
			repo,
			pkg_name,
			pkg_info,
			sys,
			tarballs_output_dir,
			log,
		)
	},
	"srs": func(
		repo types.RepositoryI,
		pkg_name string,
		pkg_info *basictypes.PackageInfo,
		sys basictypes.SystemI,
		tarballs_output_dir string,
		log *logger.Logger,
	) (types.ProviderI, error) {
		return NewProviderSRS(
			repo,
			pkg_name,
			pkg_info,
			sys,
			tarballs_output_dir,
			log,
		)
	},
}
