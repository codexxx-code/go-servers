package utils

import (
	"generator/internal/enum/timeframe"
	"generator/internal/services/horoscope/model"
	"pkg/datetime"
)

func GetDateRangeForTimeframe(req model.GetHoroscopeReq) (datetime.Date, datetime.Date) {
	switch req.Timeframe {
	case timeframe.Daily:
		return req.Date, req.Date
	case timeframe.Weekly:
		return req.Date.GetStartOfWeek(), req.Date.GetEndOfWeek()
	case timeframe.Monthly:
		return req.Date.GetStartOfMonth(), req.Date.GetEndOfMonth()
	case timeframe.Yearly:
		return req.Date.GetStartOfYear(), req.Date.GetEndOfYear()
	}

	return req.Date, req.Date
}
