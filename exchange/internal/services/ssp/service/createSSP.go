package service

import (
	"context"

	"exchange/internal/services/ssp/model"
)

func (s *SSPService) CreateSSP(ctx context.Context, req model.CreateSSPReq) error {
	return s.sspRepository.CreateSSP(ctx, req)
}
