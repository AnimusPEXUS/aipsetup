package aipsetup

import (
	"errors"
	"fmt"
	"regexp"
)

const SYSTEM_TRIPLET_REGEXP = "" +
	`(?P<cpu>.*?)-` +
	`(?P<company>.*?)-` +
	`(?P<system>((?P<kernel>.*?)-)?(?P<os>.*))`

var SYSTEM_TRIPLET_REGEXP_COMPILED = regexp.MustCompile(
	SYSTEM_TRIPLET_REGEXP,
)

var ErrCantParseSystemTriplet = errors.New("Can't parse")

type SystemTriplet struct {
	CPU     string
	Company string
	Kernel  string
	OS      string
}

func NewSystemTripletFromString(value string) (*SystemTriplet, error) {
	if !SYSTEM_TRIPLET_REGEXP_COMPILED.MatchString(value) {
		return nil, errors.New("not matching system triplet regexp")
	}

	parsed_strs := SYSTEM_TRIPLET_REGEXP_COMPILED.FindStringSubmatch(value)

	if parsed_strs == nil {
		return nil, ErrCantParseSystemTriplet
	}

	ret := new(SystemTriplet)
	ret.Kernel = ""

	for ii, i := range SYSTEM_TRIPLET_REGEXP_COMPILED.SubexpNames() {
		switch i {
		case "cpu":
			ret.CPU = parsed_strs[ii]
		case "company":
			ret.Company = parsed_strs[ii]
		case "kernel":
			ret.Kernel = parsed_strs[ii]
		case "os":
			ret.OS = parsed_strs[ii]
		}
	}

	return ret, nil
}

func (self *SystemTriplet) String() string {
	kernel_str := ""
	if self.Kernel != "" {
		kernel_str = fmt.Sprintf("%s-", self.Kernel)
	}
	return fmt.Sprintf(
		"%s-%s-%s%s",
		self.CPU,
		self.Company,
		kernel_str,
		self.OS,
	)
}
