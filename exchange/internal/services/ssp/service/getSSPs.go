package service

import (
	"context"

	"exchange/internal/services/ssp/model"
)

func (s *SSPService) GetSSPs(ctx context.Context, req model.GetSSPsReq) (res []model.SSP, err error) {
	return s.sspRepository.GetSSPs(ctx, req)
}
