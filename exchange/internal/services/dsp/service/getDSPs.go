package service

import (
	"context"

	"exchange/internal/services/dsp/model"
)

func (s *DSPService) GetDSPs(ctx context.Context, req model.GetDSPsReq) (res []model.DSP, err error) {
	return s.dspRepository.GetDSPs(ctx, req)
}
