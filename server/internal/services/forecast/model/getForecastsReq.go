package model

import "pkg/datetime"

type GetForecastsReq struct {
	Dates   []datetime.Date `query:"dates"`
	Zodiacs []string        `query:"zodiacs"`
}
