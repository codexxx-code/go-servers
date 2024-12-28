package model

import (
	"generator/internal/enum/horoscopeType"
	"generator/internal/enum/language"
	"generator/internal/enum/timeframe"
	"generator/internal/enum/zodiac"
	"pkg/datetime"
)

type Horoscope struct {
	ID              uint32                      `db:"id" json:"id"`                                                                                                                             // Идентификатор гороскопа
	DateFrom        datetime.Date               `db:"date_from" json:"dateFrom"`                                                                                                                // Дата гороскопа (относительно которой считается период)
	DateTo          datetime.Date               `db:"date_to" json:"dateTo"`                                                                                                                    // Дата окончания периода гороскопа (верхняя граница)
	PrimaryZodiac   zodiac.Zodiac               `db:"primary_zodiac" json:"primaryZodiac" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"`     // Знак зодиака, для которого генерируется гороскоп
	SecondaryZodiac *zodiac.Zodiac              `db:"secondary_zodiac" json:"secondaryZodiac" enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"` // Знак зодиака партнера (для гороскопа на пару)
	Language        language.Language           `db:"language" json:"language" enums:"russian,english"`                                                                                         // Язык текста гороскопа
	Timeframe       timeframe.Timeframe         `db:"timeframe" json:"timeframe" enums:"day,week,month,year"`                                                                                   // Период гороскопа
	HoroscopeType   horoscopeType.HoroscopeType `db:"horoscope_type" json:"horoscopeType" enums:"single,couple"`                                                                                // Тип гороскопа. Парный или для одного знака
	Text            string                      `db:"text" json:"text"`                                                                                                                         // Текст гороскопа
}
