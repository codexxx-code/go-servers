package model

import (
	"generator/internal/services/horoscope/model/zodiac"

	"pkg/datetime"
)

type GenerateDailyHoroscopeReq struct {
	Date   datetime.Date
	Zodiac zodiac.Zodiac
}
