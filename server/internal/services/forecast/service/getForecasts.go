package service

import (
	"context"

	"server/internal/services/forecast/model"
)

func (s *ForecastService) GetForecasts(ctx context.Context, req model.GetForecastsReq) ([]model.Forecast, error) {
	return s.forecastRepository.GetForecasts(ctx, req)
}
