package basictypes

import (
	"bytes"
	"encoding/json"
)

type PackageInfo struct {
	Description string
	HomePage    string

	BuilderName string

	Removable                 bool
	Reducible                 bool
	AutoReduce                bool
	DontPreserveSharedObjects bool
	NonBuildable              bool
	NonInstallable            bool
	Deprecated                bool
	PrimaryInstallOnly        bool

	BuildPkgDeps []string

	BuildDeps   []string
	SODeps      []string
	RunTimeDeps []string

	Tags     []string
	Category string
	Groups   []string

	TarballFileNameParser string

	TarballName string

	TarballFilters []string

	TarballProvider                 string
	TarballProviderArguments        []string
	TarballProviderVersionSyncDepth int

	TarballStabilityClassifier string
	TarballVersionComparator   string

	DownloadPatches              bool
	PatchesDownloadingScriptText string
}

func (self *PackageInfo) RenderJSON() (string, error) {
	data, err := json.Marshal(self)
	if err != nil {
		return "", err
	}

	b := &bytes.Buffer{}

	err = json.Indent(b, data, "  ", "  ")
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
