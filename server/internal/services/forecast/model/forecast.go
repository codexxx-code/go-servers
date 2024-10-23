package model

import (
	"pkg/datetime"
	"server/internal/services/forecast/model/zodiac"
)

type Forecast struct {
	ID     uint32        `db:"id" json:"id"`
	Date   datetime.Date `db:"date" json:"date"`
	Zodiac zodiac.Zodiac `db:"zodiac" json:"zodiac"`
	Text   string        `db:"text" json:"text"`
}
