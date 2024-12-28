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

func (d Date) GetStartOfWeek() Date {
	weekday := int(d.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return d.AddDate(0, 0, -weekday+1)
}

func (d Date) GetEndOfWeek() Date {
	weekday := int(d.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return d.AddDate(0, 0, 7-weekday)
}

func (d Date) GetStartOfMonth() Date {
	return NewDate(d.Year(), d.Month(), 1)
}

func (d Date) GetEndOfMonth() Date {
	return d.GetStartOfMonth().AddDate(0, 1, -1)
}

func (d Date) GetStartOfYear() Date {
	return NewDate(d.Year(), time.January, 1)
}

func (d Date) GetEndOfYear() Date {
	return NewDate(d.Year(), time.December, 31)
}

func (d Date) AddDate(years int, months int, days int) Date {
	return Date{d.Time.AddDate(years, months, days)}
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
