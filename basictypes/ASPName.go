package basictypes

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
		`^\((?P<name>.+?)\)` +
		`-\((?P<version>\d+(\.\d+)*)\)` +
		`-\((?P<status>.*?)\)` +
		`-\((?P<timestamp>\d{8}\.\d{6}\.\d{7})\)` +
		`-\((?P<host>.*)\)` +
		`(-\((?P<hostarch>.*)\))?` +
		`(-\((?P<crossbuilder_target>crossbuilder\-target\:.*)\))?` +
		`((\.tar.xz)|(\.asp)|(\.xz))?$`

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
	Name               string
	Version            string
	Status             string
	TimeStamp          string
	Host               string
	HostArch           string
	CrossbuilderTarget string
}

func (self *ASPName) IsEqual(other *ASPName) bool {
	return self.Name == other.Name &&
		self.Version == other.Version &&
		self.Status == other.Status &&
		self.TimeStamp == other.TimeStamp &&
		self.Host == other.Host &&
		self.HostArch == other.HostArch &&
		self.CrossbuilderTarget == other.CrossbuilderTarget
}

func (self *ASPName) String() string {

	has_arch_part := self.HostArch != self.Host
	has_target_part := self.CrossbuilderTarget != ""

	arch_part := ""
	if has_arch_part {
		arch_part = fmt.Sprintf("-(%s)", self.HostArch)
	}

	target_part := ""
	if has_target_part {
		target_part = fmt.Sprintf("-(crossbuilder-target:%s)", self.CrossbuilderTarget)
	}

	ret := fmt.Sprintf(
		"(%s)-(%s)-(%s)-(%s)-(%s)%s%s%s",
		self.Name,
		self.Version,
		self.Status,
		self.TimeStamp,
		self.Host,
		arch_part,
		target_part,
	)

	return ret
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
			ret.HostArch = parsed_strs[ii]
		}
	}

	if ret.HostArch == "" {
		ret.HostArch = ret.Host
	}

	return ret, nil
}

func (self *ASPName) StringD() string {
	ret := ""
	ret += "Name:      " + self.Name + "\n"
	ret += "Version:   " + self.Version + "\n"
	ret += "Status:    " + self.Status + "\n"
	ret += "TimeStamp: " + self.TimeStamp + "\n"
	ret += "Host:      " + self.Host + "\n"
	ret += "HostArch:  " + self.HostArch + "\n"
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
	// sorter construction. only bool is valid return type

	ni, err := NewASPNameFromString(self[i])
	if err != nil {
		panic(err)
	}
	nj, err := NewASPNameFromString(self[j])
	if err != nil {
		panic(err)
	}

	if ni.Host != nj.Host || ni.HostArch != nj.HostArch {
		panic("Hosts or HostArchs missmatch")
	}

	return ni.TimeStamp < nj.TimeStamp
}
