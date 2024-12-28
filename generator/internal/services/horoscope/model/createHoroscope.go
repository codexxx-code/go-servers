package model

import (
	"generator/internal/enum/horoscopeType"
	"generator/internal/enum/language"
	"generator/internal/enum/timeframe"
	"generator/internal/enum/zodiac"
	"pkg/datetime"
)

type CreateHoroscopeReq struct {
	DateFrom        datetime.Date               // Дата гороскопа (относительно которой считается период)
	DateTo          datetime.Date               // Дата окончания периода гороскопа (верхняя граница)
	PrimaryZodiac   zodiac.Zodiac               // Знак зодиака, для которого генерируется гороскоп
	SecondaryZodiac *zodiac.Zodiac              // Знак зодиака партнера (для гороскопа на пару)
	Language        language.Language           // Язык текста гороскопа
	Timeframe       timeframe.Timeframe         // Период гороскопа
	HoroscopeType   horoscopeType.HoroscopeType // Тип гороскопа. Парный или для одного знака
	Text            string                      // Текст гороскопа
}
