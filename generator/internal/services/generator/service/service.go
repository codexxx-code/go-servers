package service

import (
	"context"

	generatorModel "generator/internal/services/generator/model"
	promptModel "generator/internal/services/prompt/model"
	promptService "generator/internal/services/prompt/service"
)

type GeneratorService struct {
	promptService PromptService
	neuralNetwork NeuralNetwork
}

var _ PromptService = new(promptService.PromptService)

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
) *GeneratorService {
	return &GeneratorService{
		neuralNetwork: neuralNetwork,
		promptService: promptService,
	}
}
