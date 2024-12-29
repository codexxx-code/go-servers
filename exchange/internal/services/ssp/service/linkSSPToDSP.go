package service

import (
	"context"

	"exchange/internal/services/ssp/model"
)

func (s *SSPService) LinkSSPToDSP(ctx context.Context, req model.LinkSSPToDSPsReq) error {
	return s.sspRepository.LinkSSPToDSPs(ctx, req)
}
