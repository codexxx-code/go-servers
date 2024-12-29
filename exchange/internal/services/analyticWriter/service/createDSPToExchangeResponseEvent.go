package service

import (
	"context"

	"exchange/internal/services/analyticWriter/model"
)

func (s *AnalyticWriterService) CreateDSPToExchangeResponseEvent(ctx context.Context, req model.CreateDSPToExchangeResponseEventReq) error {
	return s.analyticWriterRepository.CreateDSPToExchangeResponseEvent(ctx, req)
}
