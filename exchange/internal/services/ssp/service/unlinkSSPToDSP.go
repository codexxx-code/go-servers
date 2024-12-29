package service

import (
	"context"

	"exchange/internal/services/ssp/model"
)

func (s *SSPService) UnlinkSSPToDSP(ctx context.Context, req model.UnlinkSSPToDSPsReq) error {
	return s.sspRepository.UnlinkSSPToDSPs(ctx, req)
}
