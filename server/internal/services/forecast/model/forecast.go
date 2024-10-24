package model

import (
	"pkg/datetime"
	"server/internal/services/forecast/model/zodiac"
)

type Forecast struct {
	ID     uint32        `db:"id" json:"id"`
	Date   datetime.Date `db:"date" json:"date" swaggertype:"string" format:"date"`
	Zodiac zodiac.Zodiac `db:"zodiac" json:"zodiac" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`
	Text   string        `db:"text" json:"text"`
}
