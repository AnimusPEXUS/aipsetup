package providers

import (
	"errors"
	"os"

	"github.com/AnimusPEXUS/aipsetup"
	"github.com/AnimusPEXUS/aipsetup/basictypes"
)

type ProviderHttps struct {
}

func NewProviderHttps(
	pkg_info *basictypes.PackageInfo,
	sys *aipsetup.System,
	args []string,
) (*ProviderHttps, error) {

	ret := new(ProviderHttps)

	return ret, nil
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

func (self *ProviderHttps) ListDir(pth string) ([]os.FileInfo, []os.FileInfo, error) {

}

func (self *ProviderHttps) Walk(
	pth string,
	target func(
		dir string,
		dirs []os.FileInfo,
		files []os.FileInfo,
	) error,
) error {
}

func (self *ProviderHttps) Tree() []string {

}
