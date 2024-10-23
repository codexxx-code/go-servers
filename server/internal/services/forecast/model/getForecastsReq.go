package model

import (
	"pkg/datetime"
	"server/internal/services/forecast/model/zodiac"
)

type GetForecastsReq struct {
	Dates   []datetime.Date `query:"dates"`
	Zodiacs []zodiac.Zodiac `query:"zodiacs"`
}
