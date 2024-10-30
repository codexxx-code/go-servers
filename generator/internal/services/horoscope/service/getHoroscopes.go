package service

import (
	"context"

	"generator/internal/services/horoscope/model"
)

func (s *HoroscopeService) GetHoroscopes(ctx context.Context, req model.GetHoroscopesReq) ([]model.Horoscope, error) {
	return s.horoscopeRepository.GetHoroscopes(ctx, req)
}
