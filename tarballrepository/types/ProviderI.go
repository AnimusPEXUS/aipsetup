package types

import (
	"time"

	"github.com/AnimusPEXUS/utils/filetools"
)

type ProjectI interface {
	ProjectName() string
}

type ProviderI interface {
	// can be accessed by user?
	// Enabled() bool
	// commented. to disable - comment in providers/00000_index.go

	ProviderName() string

	MainURI() string

	ArgCount() int

	CanListArg(i int) bool
	ListArg(i int) ([]string, error)

	ListDirTimeout() time.Duration

	filetools.WalkerI

	Tarballs(project string) []string
	TarballNames(project string) []string
}
