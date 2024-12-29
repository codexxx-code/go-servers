package service

import (
	"context"

	"exchange/internal/services/dsp/model"
)

func (s *DSPService) CreateDSP(ctx context.Context, req model.CreateDSPReq) error {

	// Создаем DSP в репозитории
	return s.dspRepository.CreateDSP(ctx, req)
}
