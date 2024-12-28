package timeframe

import "pkg/errors"

type Timeframe string

// enums:"day,week,month,year"
const (
	Daily   Timeframe = "day"
	Weekly  Timeframe = "week"
	Monthly Timeframe = "month"
	Yearly  Timeframe = "year"
)

func (h Timeframe) Validate() error {
	switch h {
	case Daily, Weekly, Monthly, Yearly:
		return nil
	default:
		return errors.BadRequest.New("Unknown Timeframe", errors.ParamsOption("Timeframe", h))
	}
}
