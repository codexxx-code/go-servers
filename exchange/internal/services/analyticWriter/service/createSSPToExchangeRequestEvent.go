package service

import (
	"context"

	"exchange/internal/services/analyticWriter/model"
)

func (s *AnalyticWriterService) CreateSSPToExchangeRequestEvent(ctx context.Context, req model.CreateSSPToExchangeRequestEventReq) error {
	return s.analyticWriterRepository.CreateSSPToExchangeRequestEvent(ctx, req)
}
