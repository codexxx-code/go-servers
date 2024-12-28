package service

import (
	"context"

	generatorModel "generator/internal/services/generator/model"
	"generator/internal/services/generator/neuralNetworks/chatGPT"
	"generator/internal/services/horoscope/model"
	horoscopeRepository "generator/internal/services/horoscope/repository"
	promptModel "generator/internal/services/promptTemplate/model"
	promptService "generator/internal/services/promptTemplate/service"
)

var _ HoroscopeRepository = new(horoscopeRepository.HoroscopeRepository)

type HoroscopeRepository interface {
	GetHoroscope(context.Context, model.GetHoroscopeReq) (model.Horoscope, error)
	CreateHoroscope(context.Context, model.CreateHoroscopeReq) (uint32, error)
}

var _ GeneratorService = new(chatGPT.ChatGPTService)

type GeneratorService interface {
	Generate(ctx context.Context, req generatorModel.GenerateReq) (generatorModel.GenerateRes, error)
}

var _ PromptTemplateService = new(promptService.PromptTemplateService)

type PromptTemplateService interface {
	GetPromptTemplates(ctx context.Context, req promptModel.GetPromptTemplatesReq) ([]promptModel.PromptTemplate, error)
}

type HoroscopeService struct {
	horoscopeRepository   HoroscopeRepository
	generatorService      GeneratorService
	promptTemplateService PromptTemplateService
}

func NewHoroscopeService(
	horoscopeRepository HoroscopeRepository,
	generatorService GeneratorService,
	promptTemplateService PromptTemplateService,
) *HoroscopeService {
	return &HoroscopeService{
		horoscopeRepository:   horoscopeRepository,
		generatorService:      generatorService,
		promptTemplateService: promptTemplateService,
	}
}
