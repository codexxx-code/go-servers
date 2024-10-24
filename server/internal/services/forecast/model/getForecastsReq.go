package model

import (
	"pkg/datetime"
	"server/internal/services/forecast/model/zodiac"
)

type GetForecastsReq struct {
	Dates   []datetime.Date
	Zodiacs []zodiac.Zodiac
}
