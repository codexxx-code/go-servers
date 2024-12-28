package model

import (
	"generator/internal/enum/horoscopeType"
	"generator/internal/enum/language"
	"generator/internal/enum/timeframe"
	"generator/internal/enum/zodiac"
	"pkg/datetime"
	"pkg/errors"
)

type GetHoroscopeReq struct {
	DateFrom        datetime.Date               `schema:"date_from" json:"date_from" swaggertype:"string" format:"date" validate:"required"`                                                                         // Дата гороскопа (относительно которой считается период)
	Timeframe       timeframe.Timeframe         `schema:"timeframe" json:"timeframe" enums:"day,week,month,year" validate:"required"`                                                                                // Период гороскопа
	PrimaryZodiac   zodiac.Zodiac               `schema:"primary_zodiac" json:"primary_zodiac" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces" validate:"required"` // Знак зодиака, для которого генерируется гороскоп
	SecondaryZodiac *zodiac.Zodiac              `schema:"secondary_zodiac" json:"secondary_zodiac" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`                 // Знак зодиака партнера (для гороскопа на пару)
	Language        language.Language           `schema:"language" json:"language" enums:"russian,english" validate:"required"`                                                                                      // Язык текста гороскопа
	HoroscopeType   horoscopeType.HoroscopeType `schema:"horoscope_type" json:"horoscope_type" enums:"single,couple" validate:"required"`                                                                            // Тип гороскопа. Парный или для одного знака
}

func (r GetHoroscopeReq) Validate() error {

	// Проверяем дату
	if r.DateFrom.IsZero() {
		return errors.BadRequest.New("date_from is zero value")
	}

	// Проверяем таймфрейм
	if err := r.Timeframe.Validate(); err != nil {
		return err
	}

	// Проверяем первый знак зодиака
	if err := r.PrimaryZodiac.Validate(); err != nil {
		return err
	}

	// Проверяем второй знак зодиака, если передан
	if r.SecondaryZodiac != nil {
		if err := r.SecondaryZodiac.Validate(); err != nil {
			return err
		}
	}

	// Проверяем язык
	if err := r.Language.Validate(); err != nil {
		return err
	}

	// Проверяем тип гороскопа
	if err := r.HoroscopeType.Validate(); err != nil {
		return err
	}

	// Если передан парный гороскоп, то второй знак зодиака должен быть передан
	if r.HoroscopeType == horoscopeType.Couple && r.SecondaryZodiac == nil {
		return errors.BadRequest.New("secondary_zodiac is required for pair horoscope")
	}

	return nil
}
