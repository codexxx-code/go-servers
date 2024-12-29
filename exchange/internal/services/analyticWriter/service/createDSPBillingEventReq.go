package service

import (
	"context"

	"exchange/internal/services/analyticWriter/model"
)

func (s *AnalyticWriterService) CreateDSPBillingEventReq(ctx context.Context, req model.CreateDSPBillingEventReq) error {
	return s.analyticWriterRepository.CreateDSPBillingEventReq(ctx, req)
}
