package service

import (
	"context"

	"exchange/internal/services/dsp/model"
)

func (s *DSPService) DeleteDSP(ctx context.Context, req model.DeleteDSPReq) error {
	return s.dspRepository.DeleteDSP(ctx, req)
}
