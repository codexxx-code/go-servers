package model

import (
	"generator/internal/services/horoscope/model/zodiac"

	"pkg/datetime"
)

type GetHoroscopesReq struct {
	Dates   []datetime.Date
	Zodiacs []zodiac.Zodiac
}
