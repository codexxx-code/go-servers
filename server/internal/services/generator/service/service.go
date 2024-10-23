package service

import (
	"context"

	forecastModel "server/internal/services/forecast/model"
	generatorModel "server/internal/services/generator/model"
	tgBotModel "server/internal/services/tgBot/model"
)

type GeneratorService struct {
	generatorRepository GeneratorRepository
	neuralNetwork       NeuralNetwork
	forecastService     ForecastService
	tgBotService        TgBotService
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

type GeneratorRepository interface {
	GetPrompts(ctx context.Context, req generatorModel.GetPromptsReq) ([]generatorModel.Prompt, error)
}

func NewGeneratorService(
	neuralNetwork NeuralNetwork,
	generatorRepository GeneratorRepository,
	forecastService ForecastService,
	tgBotService TgBotService,
) *GeneratorService {
	return &GeneratorService{
		neuralNetwork:       neuralNetwork,
		generatorRepository: generatorRepository,
		forecastService:     forecastService,
		tgBotService:        tgBotService,
	}
}
