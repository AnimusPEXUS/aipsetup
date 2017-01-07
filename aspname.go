package aipsetup

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

var ASP_NAME_REGEXPS_AIPSETUP3 string = `` +
	`^\((?P<name>.+?)\)-\((?P<version>(\d+\.??)+)\)-\((?P<status>.*?)\)` +
	`-\((?P<timestamp>\d{8}\.\d{6}\.\d{7})\)-\((?P<host>.*)\)` +
	`-\((?P<arch>.*)\)$`

var ASP_NAME_REGEXPS_AIPSETUP3_COMPILED *regexp.Regexp = regexp.MustCompile(
	ASP_NAME_REGEXPS_AIPSETUP3,
)

type ASPNameParsed struct {
	Name      string
	Version   string
	Status    string
	TimeStamp string
	Host      string
	Arch      string
}

func (self *ASPNameParsed) String() string {
	return fmt.Sprintf(
		"(%s)-(%s)-(%s)-(%s)-(%s)-(%s)",
		self.Name,
		self.Version,
		self.Status,
		self.TimeStamp,
		self.Host,
		self.Arch,
	)
}

func NormalizeASPName(aspname string) string {

	aspname = path.Base(aspname)

	for _, i := range []string{".tar.xz", ".asp", ".xz"} {
		if strings.HasSuffix(aspname, i) {
			aspname = aspname[:len(aspname)-len(i)]
			break
		}
	}

	return aspname
}

func NewASPNameParsedFromString(str string) *ASPNameParsed {

	var ret *ASPNameParsed = nil

	str = path.Base(str)

	str = NormalizeASPName(str)

	parsed_strs :=
		ASP_NAME_REGEXPS_AIPSETUP3_COMPILED.FindStringSubmatch(str)

	if parsed_strs != nil {

		ret = new(ASPNameParsed)

		for ii, i := range ASP_NAME_REGEXPS_AIPSETUP3_COMPILED.SubexpNames() {
			switch i {
			case "name":
				ret.Name = parsed_strs[ii]
			case "version":
				ret.Version = parsed_strs[ii]
			case "status":
				ret.Status = parsed_strs[ii]
			case "timestamp":
				ret.TimeStamp = parsed_strs[ii]
			case "host":
				ret.Host = parsed_strs[ii]
			case "arch":
				ret.Arch = parsed_strs[ii]
			}
		}

	}

	return ret
}

type ASPNameSorter []string

func (self ASPNameSorter) Len() int {
	return len(self)
}

func (self ASPNameSorter) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self ASPNameSorter) Less(i, j int) bool {
	ni := NewASPNameParsedFromString(self[i])
	nj := NewASPNameParsedFromString(self[j])

	if ni.Host != nj.Host || ni.Arch != nj.Arch {
		panic("programming error: Hosts or Archs missmatch")
	}

	return ni.TimeStamp < nj.TimeStamp
}
