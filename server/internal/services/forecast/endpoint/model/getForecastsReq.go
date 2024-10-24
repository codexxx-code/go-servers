package model

import (
	"pkg/datetime"
	"server/internal/services/forecast/model"
	"server/internal/services/forecast/model/zodiac"
)

type GetForecastsReq struct {
	Dates   []string        `schema:"dates" json:"dates" swaggertype:"array,string" example:"2024-01-01,2024-01-02"`
	Zodiacs []zodiac.Zodiac `schema:"zodiacs" json:"zodiacs" example:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`
}

func (r GetForecastsReq) Validate() error {

	for _, zodiac := range r.Zodiacs {
		if err := zodiac.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r GetForecastsReq) ConvertToBusinessModel() (res model.GetForecastsReq, err error) {

	dates := make([]datetime.Date, 0, len(r.Dates))

	for _, dateString := range r.Dates {
		date, err := datetime.ParseDate(dateString)
		if err != nil {
			return model.GetForecastsReq{}, err
		}
		dates = append(dates, date)
	}

	return model.GetForecastsReq{
		Dates:   dates,
		Zodiacs: r.Zodiacs,
	}, nil
}
