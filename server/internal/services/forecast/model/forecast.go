package model

import "pkg/datetime"

type Forecast struct {
	ID     int           `db:"id" json:"id"`
	Date   datetime.Date `db:"date" json:"date"`
	Zodiac string        `db:"zodiac" json:"zodiac"`
	Text   string        `db:"text" json:"text"`
}
