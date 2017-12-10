package providers

import (
	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/tarballdownloader/types"
)

var Index = map[string](func(
	pkg_info *basictypes.PackageInfo,
	sys *aipsetup.System,
	args []string,
) (types.ProviderI, error)){
	"https": func(
		pkg_info *basictypes.PackageInfo,
		sys *aipsetup.System,
		args []string,
	) (types.ProviderI, error) {
		return NewProviderHttps(pkg_info, sys, args)
	},
}
