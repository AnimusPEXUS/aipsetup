package types

type ProviderI interface {
	// can be accessed by user?
	// Enabled() bool
	// commented out. to disable - comment in providers/00000_index.go

	// SetArgs([]string) error

	ProviderDescription() string

	ArgCount() int

	CanListArg(i int) bool
	ListArg(i int) ([]string, error)

	Tarballs() ([]string, error)
	TarballNames() ([]string, error)

	PerformUpdate() error
}
