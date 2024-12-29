package service

import (
	"context"

	"exchange/internal/services/analyticWriter/model"
)

func (s *AnalyticWriterService) CreateExchangeToSSPResponseEvent(ctx context.Context, req model.CreateExchangeToSSPResponseEventReq) error {
	return s.analyticWriterRepository.CreateExchangeToSSPResponseEvent(ctx, req)
}
