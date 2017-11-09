package tarballnameparsers

import (
	"strconv"

	"github.com/AnimusPEXUS/tarball"
)

type Status uint

const (
	PreAlpha Status = iota
	Alpha
	Beta
	RC
	RTM
	GA
	Gold
)

type ParseResult struct {
	Name        string
	HaveVersion bool
	Version     []uint
	HaveStatus  bool
	Status      Status
	HaveBuildId bool
	BuildId     string
}

type TarballNameParserI interface {
	ParseName(value string) (*ParseResult, error)
}

var Index = map[string](func() TarballNameParserI){
	"std": func() TarballNameParserI {
		return new(TarballNameParser_Std)
	},
}

type TarballNameParser_Std struct{}

func (self *TarballNameParser_Std) ParseName(value string) (
	*ParseResult,
	error,
) {

	result, err := tarball.Parse(value)
	if err != nil {
		return nil, err
	}

	ret := new(ParseResult)

	ret.HaveVersion = true

	ret.Version = make([]uint, 0)

	for _, i := range result.Version.ParsedVersionOrStatus.Arr {
		ii, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		ret.Version = append(ret.Version, uint(ii))
	}

	ret.HaveStatus = false
	ret.HaveBuildId = false

	ret.Name = result.Name

	return ret, nil
}
