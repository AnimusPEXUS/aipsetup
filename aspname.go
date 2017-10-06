package aipsetup

import (
	"errors"
	"fmt"
	"path"
	"regexp"
	"strconv"
	"time"
)

var (
	ErrCantParseTimestamp = errors.New(
		"Can't parse given str as ASP name timestamp",
	)
)

const (
	ASP_NAME_REGEXPS_AIPSETUP3 string = `` +
		`^\((?P<name>.+?)\)-\((?P<version>(\d+\.??)+)\)-\((?P<status>.*?)\)` +
		`-\((?P<timestamp>\d{8}\.\d{6}\.\d{7})\)-\((?P<host>.*)\)` +
		`-\((?P<arch>.*)\)((\.tar.xz)|(\.asp)|(\.xz))?$`

	ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP = `` +
		`(?P<year>\d{4})(?P<month>\d{2})(?P<day>\d{2})\.` +
		`(?P<hour>\d{2})(?P<min>\d{2})(?P<sec>\d{2})\.` +
		`(?P<nsec>\d+)`
)

var (
	ASP_NAME_REGEXPS_AIPSETUP3_COMPILED *regexp.Regexp = regexp.MustCompile(
		ASP_NAME_REGEXPS_AIPSETUP3,
	)

	ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED *regexp.Regexp = regexp.MustCompile(
		ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP,
	)
)

type ASPName struct {
	Name      string
	Version   string
	Status    string
	TimeStamp string
	Host      string
	Arch      string
}

func (self *ASPName) String() string {
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

func (self *ASPName) TimeStampTime() (*time.Time, error) {

	var tmp struct {
		year                      int
		month                     time.Month
		day, hour, min, sec, nsec int
	}

	if !ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED.MatchString(self.TimeStamp) {
		return nil, errors.New("not matching ASP name timestamp regexp")
	}

	parsed_strs :=
		ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED.FindStringSubmatch(self.TimeStamp)

	if parsed_strs == nil {
		return nil, ErrCantParseTimestamp
	}

	for ii, i := range ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED.SubexpNames() {
		switch i {

		case "year":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return nil, ErrCantParseTimestamp
			}
			tmp.year = t

		case "month":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return nil, ErrCantParseTimestamp
			}
			tmp.month = time.Month(t)

		case "day":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return nil, ErrCantParseTimestamp
			}
			tmp.day = t

		case "hour":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return nil, ErrCantParseTimestamp
			}
			tmp.hour = t

		case "min":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return nil, ErrCantParseTimestamp
			}
			tmp.min = t

		case "sec":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return nil, ErrCantParseTimestamp
			}
			tmp.sec = t

		case "nsec":
			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return nil, ErrCantParseTimestamp
			}
			tmp.nsec = int(time.Duration(t) * time.Microsecond)
		}
	}

	ret := time.Date(
		tmp.year,
		tmp.month,
		tmp.day, tmp.hour, tmp.min, tmp.sec, tmp.nsec,
		time.UTC,
	)

	return &ret, nil
}

func NormalizeASPName(aspname string) (string, error) {

	res, err := NewASPNameFromString(aspname)
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func NewASPNameFromString(str string) (*ASPName, error) {

	var ret *ASPName = nil

	str = path.Base(str)

	if !ASP_NAME_REGEXPS_AIPSETUP3_COMPILED.MatchString(str) {
		return nil, errors.New("not matching ASP name regexp")
	}

	parsed_strs :=
		ASP_NAME_REGEXPS_AIPSETUP3_COMPILED.FindStringSubmatch(str)

	if parsed_strs == nil {
		return nil, errors.New("Can't parse given str as ASP name")
	}

	ret = new(ASPName)

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

	return ret, nil
}

type ASPNameSorter []string

func (self ASPNameSorter) Len() int {
	return len(self)
}

func (self ASPNameSorter) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self ASPNameSorter) Less(i, j int) bool {
	// sorter construction. only bool is valid return type

	ni, err := NewASPNameFromString(self[i])
	if err != nil {
		panic(err)
	}
	nj, err := NewASPNameFromString(self[j])
	if err != nil {
		panic(err)
	}

	if ni.Host != nj.Host || ni.Arch != nj.Arch {
		panic("Hosts or Archs missmatch")
	}

	return ni.TimeStamp < nj.TimeStamp
}
