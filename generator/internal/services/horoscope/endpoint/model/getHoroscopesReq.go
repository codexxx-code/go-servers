package model

import (
	"generator/internal/services/horoscope/model"
	"generator/internal/services/horoscope/model/zodiac"

	"pkg/datetime"
)

type GetHoroscopesReq struct {
	Dates   []string        `schema:"dates" json:"dates" swaggertype:"array,string" example:"2024-01-01,2024-01-02"`
	Zodiacs []zodiac.Zodiac `schema:"zodiacs" json:"zodiacs" example:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`
}

func (r GetHoroscopesReq) Validate() error {

	for _, zodiac := range r.Zodiacs {
		if err := zodiac.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r GetHoroscopesReq) ConvertToBusinessModel() (res model.GetHoroscopesReq, err error) {

	dates := make([]datetime.Date, 0, len(r.Dates))

	for _, dateString := range r.Dates {
		date, err := datetime.ParseDate(dateString)
		if err != nil {
			return model.GetHoroscopesReq{}, err
		}
		dates = append(dates, date)
	}

	return model.GetHoroscopesReq{
		Dates:   dates,
		Zodiacs: r.Zodiacs,
	}, nil
}
