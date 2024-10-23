package service

import (
	"context"

	"server/internal/services/forecast/model"
	forecastRepository "server/internal/services/forecast/repository"
)

var _ ForecastRepository = new(forecastRepository.ForecastRepository)

type ForecastRepository interface {
	GetForecasts(context.Context, model.GetForecastsReq) ([]model.Forecast, error)
	CreateForecast(context.Context, model.CreateForecastReq) (uint32, error)
}

type ForecastService struct {
	forecastRepository ForecastRepository
}

func NewForecastService(
	forecastRepository ForecastRepository,
) *ForecastService {
	return &ForecastService{
		forecastRepository: forecastRepository,
	}
}
