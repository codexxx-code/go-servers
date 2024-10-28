package model

import (
	"generator/internal/services/forecast/model/zodiac"

	"pkg/datetime"
)

type GetForecastsReq struct {
	Dates   []datetime.Date
	Zodiacs []zodiac.Zodiac
}
