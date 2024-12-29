package service

import (
	"context"

	"exchange/internal/services/ssp/model"
)

func (s *SSPService) DeleteSSP(ctx context.Context, req model.DeleteSSPReq) error {
	return s.sspRepository.DeleteSSP(ctx, req)
}
