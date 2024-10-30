package service

import (
	"context"

	"generator/internal/services/horoscope/model"
)

func (s *HoroscopeService) CreateHoroscope(ctx context.Context, req model.CreateHoroscopeReq) (uint32, error) {
	return s.horoscopeRepository.CreateHoroscope(ctx, req)
}
