package service

import (
	"context"

	forecastModel "generator/internal/services/forecast/model"
	forecastService "generator/internal/services/forecast/service"
	generatorModel "generator/internal/services/generator/model"
	promptModel "generator/internal/services/prompt/model"
	promptService "generator/internal/services/prompt/service"
	tgBotModel "generator/internal/services/tgBot/model"
	tgBotService "generator/internal/services/tgBot/service"
)

type GeneratorService struct {
	promptService   PromptService
	neuralNetwork   NeuralNetwork
	forecastService ForecastService
	tgBotService    TgBotService
}

var _ TgBotService = new(tgBotService.TgBotService)

type TgBotService interface {
	SendMessage(ctx context.Context, req tgBotModel.SendMessageReq) error
}

var _ PromptService = new(promptService.PromptService)

type ForecastService interface {
	CreateForecast(ctx context.Context, req forecastModel.CreateForecastReq) (uint32, error)
}

var _ ForecastService = new(forecastService.ForecastService)

type NeuralNetwork interface {
	Generate(ctx context.Context, req generatorModel.GenerateReq) (generatorModel.GenerateRes, error)
}

var _ PromptService = new(promptService.PromptService)

type PromptService interface {
	GetPrompts(ctx context.Context, req promptModel.GetPromptsReq) ([]promptModel.Prompt, error)
}

func NewGeneratorService(
	neuralNetwork NeuralNetwork,
	promptService PromptService,
	forecastService ForecastService,
	tgBotService TgBotService,
) *GeneratorService {
	return &GeneratorService{
		neuralNetwork:   neuralNetwork,
		promptService:   promptService,
		forecastService: forecastService,
		tgBotService:    tgBotService,
	}
}
