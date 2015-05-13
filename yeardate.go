package gofindadate

import (
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/weberr13/errorable"
	"regexp"
	"strconv"
)

type YearDate struct {
	year  int
	month int
	day   int
}
type StringWithDate string

var NoInput = errors.New("No data for date")
var InvalidDateLength = errors.New("Invalid date length")
var InvalidDateString = errors.New("Invalid date string")
var dateReg *regexp.Regexp
var initErr error

// Static initialization
func init() {

	dateReg, initErr = regexp.Compile(`.*(\d{4})[-_]+(\d{2})[-_]+(\d{2}).*`)
	if initErr != nil {
		log.Error(initErr)
	}
}

// Use the yeardate format to parse a string slice for the oldest, or the first malformed
func GetOldestOfSlice(slice []string) (oldest string) {

	for _, v := range slice {
		if oldest == "" {
			oldest = v
			continue
		} else {
			date, err := NewYearDate(GetDateWithYYYYMMDD, v)
			if nil != err {
				log.Error("Invalid index " + v + " " + err.Error())
			}
			oldDate, _ := NewYearDate(GetDateWithYYYYMMDD, oldest)
			if less, ok := date.InvalidIsLess(oldDate); !ok {
				if less {
					oldest = v
				}
			} else if date.Less(oldDate) {
				oldest = v
			}
		}
	}
	return oldest
}

// Construct a yeardate from either a string or a sequcence of YYYY,MM,DD
func NewYearDate(f func(string) (*YearDate, error), s ...string) (date *YearDate, err error) {
	date = &YearDate{year: 0, month: 0, day: 0}
	if len(s) == 0 {
		log.Error(NoInput)
		return nil, NoInput
	}
	if len(s) == 1 {
		return f(s[0])
	}
	ints := errorable.NewIntErrs(strconv.Atoi, s[1:]...)
	if err = ints.GetFirstErr(); nil != err {
		log.Error(err)
		return nil, err
	}

	if ints.Len() < 3 {
		log.Error(fmt.Sprint("invalid length ", s))
		return nil, InvalidDateLength
	}
	date.year = ints.Get(0)
	date.month = ints.Get(1)
	date.day = ints.Get(2)
	return date, nil
}

// The invalid dates float to Less
func (this *YearDate) InvalidIsLess(that *YearDate) (less bool, ok bool) {
	if nil == that {
		return false, false
	}
	if nil == this {
		return true, false
	}
	return false, true
}

// For sorting
func (this *YearDate) Less(that *YearDate) (less bool) {
	if this.year == that.year {
		if this.month == that.month {
			return this.day < that.day
		}
		return this.month < that.month
	}
	return this.year < that.year
}

// Extract a slice of YYYY,MM,DD from the complied in regex pattern
func GetDateWithYYYYMMDD(s string) (date *YearDate, err error) {
	ms := dateReg.FindStringSubmatch(string(s))
	if ms == nil {
		log.Error("No date found")
		return nil, InvalidDateString
	}
	return NewYearDate(GetDateWithYYYYMMDD, ms...)
}
