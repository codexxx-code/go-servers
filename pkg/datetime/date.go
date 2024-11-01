package datetime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const JSONNull = "null"
const DateFormat = "2006-01-02"

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Format(DateFormat)
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func Now() Date {
	date := time.Now()
	return Date{time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)}
}

func ParseDate(s string) (Date, error) {
	t, err := time.Parse(DateFormat, s)
	if err != nil {
		return Date{}, err
	}
	return Date{t}, nil
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == JSONNull || s == "" {
		return nil
	}

	d.Time, err = time.Parse(DateFormat, s)
	if err != nil {
		return err
	}

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {

	if d.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%v"`, d.Format(DateFormat))), nil
}

func (d *Date) UnmarshalText(b []byte) (err error) {

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == JSONNull || s == "" {
		return nil
	}

	d.Time, err = time.Parse(DateFormat, s)
	if err != nil {
		return err
	}

	return nil
}

func (d *Date) Scan(src any) error {
	if t, ok := src.(time.Time); ok {
		d.Time = t
	}
	return nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}
