package service

import (
	"context"
	"strings"

	generatorModel "generator/internal/services/generator/model"
	promptModel "generator/internal/services/prompt/model"

	"pkg/slices"
)

const (
	zodiacMacros = "${ZODIAC}"
	dateMacros   = "${DATE}"
)

const zodiacHoroscope = "zodiac_horoscope"

func (s *GeneratorService) GenerateDailyHoroscope(ctx context.Context, req generatorModel.GenerateDailyHoroscopeReq) (res generatorModel.GenerateDailyHoroscopeRes, err error) {

	// Получаем промпт
	promptRes, err := slices.FirstWithError(
		s.promptService.GetPrompts(ctx, promptModel.GetPromptsReq{ //nolint:exhaustruct
			Cases: []string{zodiacHoroscope},
		}),
	)
	if err != nil {
		return res, err
	}

	// Раскрываем макросы
	prompt := strings.NewReplacer(
		zodiacMacros, req.Zodiac.ToRussian(),
		dateMacros, req.Date.String(),
	).Replace(promptRes.Text)

	// Генерируем прогноз
	generateRes, err := s.neuralNetwork.Generate(ctx, generatorModel.GenerateReq{
		System: "You are generator of forecasts for zodiac signs",
		Prompt: prompt,
	})
	if err != nil {
		return res, err
	}

	return generatorModel.GenerateDailyHoroscopeRes{
		Text: generateRes.Text,
	}, nil
}
