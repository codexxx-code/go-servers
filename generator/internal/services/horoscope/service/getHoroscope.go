package service

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	"generator/internal/enum/promptCase"
	"generator/internal/enum/timeframe"
	generatorModel "generator/internal/services/generator/model"
	"generator/internal/services/horoscope/model"
	promptModel "generator/internal/services/promptTemplate/model"
	"pkg/datetime"
	"pkg/errors"
	"pkg/log"
	"pkg/pointer"
	"pkg/slices"
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

	log.Info(ctx, fmt.Sprintf("Не нашли гороскоп со следующими параметрами: %+v", req))

	// Получаем промпт для генерации гороскопа из базы данных
	promptRes, err := slices.FirstWithError(
		s.promptTemplateService.GetPromptTemplates(ctx, promptModel.GetPromptTemplatesReq{ //nolint:exhaustruct
			Cases: []promptCase.PromptCase{promptCase.CreateHoroscope},
		}),
	)
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	// Считаем начало и конец периода гороскопа
	var dateFrom, dateTo datetime.Date
	switch req.Timeframe {
	case timeframe.Daily:
		dateFrom = req.DateFrom
		dateTo = req.DateFrom
	case timeframe.Weekly:
		dateFrom = req.DateFrom.GetStartOfWeek()
		dateTo = req.DateFrom.GetEndOfWeek()
	case timeframe.Monthly:
		dateFrom = req.DateFrom.GetStartOfMonth()
		dateTo = req.DateFrom.GetEndOfMonth()
	case timeframe.Yearly:
		dateFrom = req.DateFrom.GetStartOfYear()
		dateTo = req.DateFrom.GetEndOfYear()
	}

	// Компилируем го темплейт промпта для генерации гороскопа
	t, err := template.New(string(promptCase.CreateHoroscope)).Parse(promptRes.Template)
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	var secondaryZodiacString *string
	if req.SecondaryZodiac != nil {
		secondaryZodiacString = pointer.Pointer(string(*req.SecondaryZodiac))
	}

	// Заполняем темплейт данными
	var buf bytes.Buffer
	if err = t.Execute(&buf, struct {
		DateFrom        time.Time
		DateTo          time.Time
		PrimaryZodiac   string
		SecondaryZodiac *string
		Language        string
		Timeframe       string
		Type            string
	}{
		DateFrom:        dateFrom.Time,
		DateTo:          dateTo.Time,
		PrimaryZodiac:   string(req.PrimaryZodiac),
		SecondaryZodiac: secondaryZodiacString,
		Language:        string(req.Language),
		Timeframe:       string(req.Timeframe),
		Type:            string(req.HoroscopeType),
	}); err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	prompt := buf.String()

	log.Debug(ctx, prompt)

	// Генерируем прогноз
	generateRes, err := s.generatorService.Generate(ctx, generatorModel.GenerateReq{
		System: "You are generator of forecasts for zodiac signs",
		Prompt: prompt,
	})
	if err != nil {
		return res, err
	}

	// Добавляем гороскоп в базу данных
	id, err := s.horoscopeRepository.CreateHoroscope(ctx, model.CreateHoroscopeReq{
		DateFrom:        dateFrom,
		DateTo:          dateTo,
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
		DateFrom:        dateFrom,
		DateTo:          dateTo,
		PrimaryZodiac:   req.PrimaryZodiac,
		SecondaryZodiac: req.SecondaryZodiac,
		Language:        req.Language,
		Timeframe:       req.Timeframe,
		HoroscopeType:   req.HoroscopeType,
		Text:            generateRes.Text,
	}, nil
}
