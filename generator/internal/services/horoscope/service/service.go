package service

import (
	"context"

	generatorModel "generator/internal/services/generator/model"
	generatorService "generator/internal/services/generator/service"
	"generator/internal/services/horoscope/model"
	horoscopeRepository "generator/internal/services/horoscope/repository"
)

var _ HoroscopeRepository = new(horoscopeRepository.HoroscopeRepository)

type HoroscopeRepository interface {
	GetHoroscopes(context.Context, model.GetHoroscopesReq) ([]model.Horoscope, error)
	CreateHoroscope(context.Context, model.CreateHoroscopeReq) (uint32, error)
}

var _ GeneratorService = new(generatorService.GeneratorService)

type GeneratorService interface {
	GenerateDailyHoroscope(ctx context.Context, req generatorModel.GenerateDailyHoroscopeReq) (generatorModel.GenerateDailyHoroscopeRes, error)
}

type HoroscopeService struct {
	horoscopeRepository HoroscopeRepository
	generatorService    GeneratorService
}

func NewHoroscopeService(
	horoscopeRepository HoroscopeRepository,
	generatorService GeneratorService,
) *HoroscopeService {
	return &HoroscopeService{
		horoscopeRepository: horoscopeRepository,
		generatorService:    generatorService,
	}
}
