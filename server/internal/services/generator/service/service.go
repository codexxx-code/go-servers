package service

import (
	"context"

	forecastModel "server/internal/services/forecast/model"
	generatorModel "server/internal/services/generator/model"
	promptModel "server/internal/services/prompt/model"
	tgBotModel "server/internal/services/tgBot/model"
)

type GeneratorService struct {
	promptService   PromptService
	neuralNetwork   NeuralNetwork
	forecastService ForecastService
	tgBotService    TgBotService
}

type TgBotService interface {
	SendMessage(ctx context.Context, req tgBotModel.SendMessageReq) error
}

type ForecastService interface {
	CreateForecast(ctx context.Context, req forecastModel.CreateForecastReq) (uint32, error)
}

type NeuralNetwork interface {
	Generate(ctx context.Context, req generatorModel.GenerateReq) (generatorModel.GenerateRes, error)
}

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
