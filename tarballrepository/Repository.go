package tarballrepository

import (
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/distropkginfodb"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/cachepresets"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/providers"
	"github.com/AnimusPEXUS/aipsetup/tarballrepository/types"
	"github.com/AnimusPEXUS/utils/cache01"
)

type Repository struct {
	sys basictypes.SystemI
}

func NewRepository(sys basictypes.SystemI) (*Repository, error) {
	ret := new(Repository)
	ret.sys = sys
	return ret, nil
}

func (self *Repository) GetRepositoryPath() string {
	return self.sys.GetTarballRepoRootDir()
}

func (self *Repository) GetCachesDir() string {
	return path.Join(self.GetRepositoryPath(), "cache")
}

func (self *Repository) GetPackagePath(name string) string {
	return path.Join(self.GetRepositoryPath(), "packages", name)
}

func (self *Repository) GetPackageTarballsPath(name string) string {
	return path.Join(self.GetPackagePath(name), "tarballs")
}

func (self *Repository) GetPackageCachePath(name string) string {
	return path.Join(self.GetCachesDir(), name)
}

func (self *Repository) CreateCacheObjectForPackage(name string) (
	*cache01.CacheDir,
	error,
) {
	info, err := distropkginfodb.Get(name)
	if err != nil {
		return nil, err
	}

	var preset *cache01.Settings

	if i, err :=
		cachepresets.Get(info.TarballProviderCachePresetName); err != nil {
		return nil, err
	} else {
		preset = i
	}

	return cache01.NewCacheDir(self.GetPackageCachePath(name), preset)
}

func (self *Repository) PerformPackageTarballsUpdate(name string) error {

	var info *basictypes.PackageInfo
	var prov types.ProviderI

	if i, err := distropkginfodb.Get(name); err != nil {
		return err
	} else {
		info = i
	}

	cache, err := self.CreateCacheObjectForPackage(name)
	if err != nil {
		return err
	}

	if i, err := providers.Get(
		info.TarballProvider,
		self,
		name,
		self.sys,
		self.GetPackageTarballsPath(name),
		cache,
		info.TarballProviderArguments,
	); err != nil {
		return err
	} else {
		prov = i
	}

	err = prov.PerformUpdate()
	if err != nil {
		return err
	}

	return nil
}
