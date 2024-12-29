package service

import (
	"context"

	"exchange/internal/services/dsp/model"
)

func (s *DSPService) FindDSPs(ctx context.Context, req model.FindDSPsReq) (res model.FindDSPsRes, err error) {

	// Получаем DSP
	dsps, err := s.dspRepository.FindDSPs(ctx, req)
	if err != nil {
		return res, err
	}

	// Получаем количество DSP
	dspsCount, err := s.dspRepository.GetDSPsCount(ctx, req)
	if err != nil {
		return res, err
	}

	return model.FindDSPsRes{
		DSPs:      dsps,
		DSPsCount: dspsCount,
	}, nil
}
