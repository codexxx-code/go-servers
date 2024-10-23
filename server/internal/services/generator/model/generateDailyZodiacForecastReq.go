package model

import (
	"pkg/datetime"
	"server/internal/services/forecast/model/zodiac"
)

type GenerateDailyZodiacForecastReq struct {
	Date   datetime.Date `schema:"date" validate:"required"`
	Zodiac zodiac.Zodiac `schema:"zodiac" validate:"required"`
}
