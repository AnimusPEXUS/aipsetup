package providers

import (
	"errors"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/repository/types"
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

func init() {
	Index = make(map[string]types.NewProviderFunc)
}

var Index map[string]types.NewProviderFunc
