package service

import (
	"context"

	"exchange/internal/services/ssp/model"
)

func (s *SSPService) UpdateSSP(ctx context.Context, req model.UpdateSSPReq) error {
	return s.sspRepository.UpdateSSP(ctx, req)
}
