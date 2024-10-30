package model

import (
	"generator/internal/services/horoscope/model/zodiac"

	"pkg/datetime"
)

type Horoscope struct {
	ID     uint32        `db:"id" json:"id"`
	Date   datetime.Date `db:"date" json:"date" swaggertype:"string" format:"date"`
	Zodiac zodiac.Zodiac `db:"zodiac" json:"zodiac" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`
	Text   string        `db:"text" json:"text"`
}
