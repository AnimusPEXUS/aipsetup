package providers

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/aipsetup/repository/types"
	"github.com/AnimusPEXUS/utils/logger"
)

var _ types.ProviderI = &ProviderGitHub{}

func init() {
	// TODO: this provider needs completion
	//	Index["srs"] = NewProviderGitHub
}

type ProviderGitHub struct {
	repo                types.RepositoryI
	pkg_name            string
	pkg_info            *basictypes.PackageInfo
	sys                 basictypes.SystemI
	tarballs_output_dir string
	log                 *logger.Logger

	githubv4_token string
}

func NewProviderGitHub(
	repo types.RepositoryI,
	pkg_name string,
	pkg_info *basictypes.PackageInfo,
	sys basictypes.SystemI,
	tarballs_output_dir string,
	log *logger.Logger,
) (types.ProviderI, error) {
	self := &ProviderGitHub{
		repo:                repo,
		pkg_name:            pkg_name,
		pkg_info:            pkg_info,
		sys:                 sys,
		tarballs_output_dir: tarballs_output_dir,
		log:                 log,
	}

	d, err := ioutil.ReadFile("~/.aipsetup/github_token")
	if err == nil {
		ds := string(d)
		self.githubv4_token = strings.Trim(ds, "\n \t\r")
	}

	return self, nil
}

func (self *ProviderGitHub) ProviderDescription() string {
	return "githubv4"
}

func (self *ProviderGitHub) ArgCount() int {
	return 1
}

func (self *ProviderGitHub) CanListArg(i int) bool {
	return false
}

func (self *ProviderGitHub) ListArg(i int) ([]string, error) {
	return nil, errors.New("not supported")
}

func (self *ProviderGitHub) Tarballs() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderGitHub) TarballNames() ([]string, error) {
	return []string{}, nil
}

func (self *ProviderGitHub) PerformUpdate() error {
	//	c := &http.Client{}
	//	gh4 := githubv4.NewClient(c)
	return nil
}
