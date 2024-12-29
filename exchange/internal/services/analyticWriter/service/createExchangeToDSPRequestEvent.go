package service

import (
	"context"

	"exchange/internal/services/analyticWriter/model"
)

func (s *AnalyticWriterService) CreateExchangeToDSPRequestEvent(ctx context.Context, req model.CreateExchangeToDSPRequestEventReq) error {
	return s.analyticWriterRepository.CreateExchangeToDSPRequestEvent(ctx, req)
}
