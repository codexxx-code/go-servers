package service

import (
	"context"

	"exchange/internal/services/analyticWriter/model"
)

func (s *AnalyticWriterService) CreateSSPBillingEventReq(ctx context.Context, req model.CreateSSPBillingEventReq) error {
	return s.analyticWriterRepository.CreateSSPBillingEventReq(ctx, req)
}
