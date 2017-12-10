package tarballnameparsers

import (
	"strconv"

	"github.com/AnimusPEXUS/utils/tarballname"
)

type TarballNameParser_Std struct{}

func (self *TarballNameParser_Std) ParseName(value string) (
	*ParseResult,
	error,
) {

	result, err := tarballname.Parse(value)
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
