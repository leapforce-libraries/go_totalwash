package timecamp

import (
	"strconv"
	"strings"
	"time"
)

const (
	DateTimeFormat string = "2006-01-02T15:04:05.999-07:00"
)

type DateTimeString time.Time

func (d *DateTimeString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	unquoted = strings.TrimSpace(unquoted)

	if unquoted == "" || unquoted == "0001-01-01T00:00:00" {
		d = nil
		return nil
	}

	_t, err := time.Parse(DateTimeFormat, unquoted)
	if err != nil {
		return err
	}

	*d = DateTimeString(_t)
	return nil
}

func (d *DateTimeString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}

func (d DateTimeString) Value() time.Time {
	return time.Time(d)
}
