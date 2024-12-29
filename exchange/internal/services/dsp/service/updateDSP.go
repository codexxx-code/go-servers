package service

import (
	"context"

	"exchange/internal/services/dsp/model"
)

func (s *DSPService) UpdateDSP(ctx context.Context, req model.UpdateDSPReq) error {

	// Обновляем DSP в репозитории
	return s.dspRepository.UpdateDSP(ctx, req)
}
