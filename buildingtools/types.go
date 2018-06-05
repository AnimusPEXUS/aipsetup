package buildingtools

type (
	EnvironmentOperationMode uint
)

const (
	Copy EnvironmentOperationMode = iota
	Clean
)
