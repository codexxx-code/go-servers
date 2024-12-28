package service

import (
	"context"

	generatorModel "generator/internal/services/generator/model"
	"generator/internal/services/horoscope/model"
	"generator/internal/services/horoscope/service/utils"
	"pkg/errors"
	"pkg/sql"
)

func (s *HoroscopeService) GetHoroscope(ctx context.Context, req model.GetHoroscopeReq) (res model.Horoscope, err error) {

	// Получаем гороскоп из базы данных
	horoscope, err := s.horoscopeRepository.GetHoroscope(ctx, req)

	// Если ошибок нет = гороскоп найден, возвращаем его
	if err == nil {
		return horoscope, nil
	}

	// Если ошибка не связана с пустым ответом
	if !errors.Is(err, sql.ErrNoRows) {

		// Возвращаем ошибку
		return res, err

	}

	// Считаем начало и конец периода гороскопа
	req.DateFrom, req.DateTo = utils.GetDateRangeForTimeframe(req)

	// Получаем промпт для генерации гороскопа
	getPromptRes, err := s.getHoroscopePrompt(ctx, req)
	if err != nil {
		return res, err
	}

	// Генерируем прогноз
	generateRes, err := s.generatorService.Generate(ctx, generatorModel.GenerateReq{
		System: "You are generator of forecasts for zodiac signs",
		Prompt: getPromptRes.Prompt,
	})
	if err != nil {
		return res, err
	}

	// Добавляем гороскоп в базу данных
	id, err := s.horoscopeRepository.CreateHoroscope(ctx, model.CreateHoroscopeReq{
		DateFrom:        req.DateFrom,
		DateTo:          req.DateTo,
		PrimaryZodiac:   req.PrimaryZodiac,
		SecondaryZodiac: req.SecondaryZodiac,
		Language:        req.Language,
		Timeframe:       req.Timeframe,
		HoroscopeType:   req.HoroscopeType,
		Text:            generateRes.Text,
	})
	if err != nil {
		return res, err
	}

	// Возвращаем гороскоп
	return model.Horoscope{
		ID:              id,
		DateFrom:        req.DateFrom,
		DateTo:          req.DateTo,
		PrimaryZodiac:   req.PrimaryZodiac,
		SecondaryZodiac: req.SecondaryZodiac,
		Language:        req.Language,
		Timeframe:       req.Timeframe,
		HoroscopeType:   req.HoroscopeType,
		Text:            generateRes.Text,
	}, nil
}
