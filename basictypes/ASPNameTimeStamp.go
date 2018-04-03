package basictypes

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP = `` +
		`(?P<year>\d{4})(?P<month>\d{2})(?P<day>\d{2})\.` +
		`(?P<hour>\d{2})(?P<min>\d{2})(?P<sec>\d{2})\.` +
		`(?P<nsec>\d+)`
)

var (
	ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED = regexp.MustCompile(
		ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP,
	)
)

type ASPNameTimeStamp struct {
	time.Time
}

func NewASPTimeStampFromCurrentTime() ASPNameTimeStamp {
	return NewASPTimeStampFromTime(time.Now().UTC())
}

func NewASPTimeStampFromTime(t time.Time) ASPNameTimeStamp {
	return ASPNameTimeStamp{Time: t.UTC()}
}

func NewASPTimeStampFromString(text string) (ASPNameTimeStamp, error) {
	var tmp struct {
		year                      int
		month                     time.Month
		day, hour, min, sec, nsec int
	}

	if !ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED.MatchString(text) {
		return ASPNameTimeStamp{}, errors.New("not matching ASP name timestamp regexp")
	}

	parsed_strs :=
		ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED.FindStringSubmatch(text)

	if parsed_strs == nil {
		return ASPNameTimeStamp{}, ErrCantParseTimestamp
	}

	for ii, i := range ASP_NAME_REGEXPS_AIPSETUP3_TIMESTAMP_COMPILED.SubexpNames() {
		switch i {

		case "year":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return ASPNameTimeStamp{}, ErrCantParseTimestamp
			}
			tmp.year = t

		case "month":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return ASPNameTimeStamp{}, ErrCantParseTimestamp
			}
			tmp.month = time.Month(t)

		case "day":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return ASPNameTimeStamp{}, ErrCantParseTimestamp
			}
			tmp.day = t

		case "hour":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return ASPNameTimeStamp{}, ErrCantParseTimestamp
			}
			tmp.hour = t

		case "min":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return ASPNameTimeStamp{}, ErrCantParseTimestamp
			}
			tmp.min = t

		case "sec":

			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return ASPNameTimeStamp{}, ErrCantParseTimestamp
			}
			tmp.sec = t

		case "nsec":
			t, err := strconv.Atoi(parsed_strs[ii])
			if err != nil {
				return ASPNameTimeStamp{}, ErrCantParseTimestamp
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

	return ASPNameTimeStamp{Time: ret}, nil
}

func (self ASPNameTimeStamp) GetTime() time.Time {
	return self.Time
}

func (self ASPNameTimeStamp) String() string {
	return fmt.Sprintf(
		"%04d%02d%02d.%02d%02d%02d.%d",
		self.Year(),
		self.Month(),
		self.Day(),
		self.Hour(),
		self.Minute(),
		self.Second(),
		self.Nanosecond(),
	)
}
