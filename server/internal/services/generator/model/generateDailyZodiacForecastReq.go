package model

import (
	"pkg/datetime"
	"server/internal/services/forecast/model/zodiac"
)

type GenerateDailyZodiacForecastReq struct {
	Date   datetime.Date `schema:"date" validate:"required" swaggertype:"string" format:"date" example:"2024-01-01"`
	Zodiac zodiac.Zodiac `schema:"zodiac" validate:"required" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`
}

func (r GenerateDailyZodiacForecastReq) Validate() error {
	return r.Zodiac.Validate()
}
