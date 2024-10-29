package service

import (
	"context"
	"fmt"
	"strings"

	forecastModel "generator/internal/services/forecast/model"
	generatorModel "generator/internal/services/generator/model"
	promptModel "generator/internal/services/prompt/model"
	tgBotModel "generator/internal/services/tgBot/model"
	tgBotService "generator/internal/services/tgBot/service"

	"pkg/log"
	"pkg/slices"
)

const (
	zodiacMacros = "${ZODIAC}"
	dateMacros   = "${DATE}"
)

const zodiacForecast = "zodiac_forecast"

const generatedForecastMessageTemplate = "*Прогноз для знака зодиака успешно сгенерирован*\n\n%s: %s\n\n%s"

func (s *GeneratorService) GenerateDailyZodiacForecast(ctx context.Context, req generatorModel.GenerateDailyZodiacForecastReq) error {

	var tgMessage tgBotModel.SendMessageReq

	defer func() {
		err := s.tgBotService.SendMessage(ctx, tgMessage)
		if err != nil {
			log.Error(ctx, err)
		}
	}()

	// Получаем промпт
	promptRes, err := slices.FirstWithError(
		s.promptService.GetPrompts(ctx, promptModel.GetPromptsReq{ //nolint:exhaustruct
			Cases: []string{zodiacForecast},
		}),
	)
	if err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли получить промпт из базы данных\n\n%v", tgBotService.Replacer.Replace(err.Error()))
		return err
	}

	// Раскрываем макросы
	prompt := strings.NewReplacer(
		zodiacMacros, req.Zodiac.ToRussian(),
		dateMacros, req.Date.String(),
	).Replace(promptRes.Text)

	// Генерируем прогноз
	res, err := s.neuralNetwork.Generate(ctx, generatorModel.GenerateReq{
		System: "You are generator of forecasts for zodiac signs",
		Prompt: prompt,
	})
	if err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли сгенерировать прогноз\n\n%v", tgBotService.Replacer.Replace(err.Error()))
		return err
	}

	// Сохраняем прогноз в базу данных
	if _, err := s.forecastService.CreateForecast(ctx, forecastModel.CreateForecastReq{
		Date:   req.Date,
		Zodiac: req.Zodiac,
		Text:   res.Text,
	}); err != nil {
		tgMessage.Message += fmt.Sprintf("Не смогли сохранить прогноз в базу данных\n\n%v", tgBotService.Replacer.Replace(err.Error()))
		return err
	}

	tgMessage.Message = fmt.Sprintf(generatedForecastMessageTemplate, req.Zodiac.ToRussian(), tgBotService.Replacer.Replace(req.Date.String()), tgBotService.Replacer.Replace(res.Text))

	return nil
}
