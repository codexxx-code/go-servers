package model

import (
	"generator/internal/services/horoscope/model/zodiac"

	"pkg/datetime"
)

type GetHoroscopeWithGenerationReq struct {
	Date   datetime.Date `schema:"date" json:"date" swaggertype:"string" format:"date"`
	Zodiac zodiac.Zodiac `schema:"zodiacs" json:"zodiacs" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`
}
