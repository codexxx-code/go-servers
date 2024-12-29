package service

import (
	"context"

	"exchange/internal/services/analyticWriter/model"
	analyticWriterRepository "exchange/internal/services/analyticWriter/repository"
)

type AnalyticWriterService struct {
	analyticWriterRepository AnalyticWriterRepository
}

var _ AnalyticWriterRepository = new(analyticWriterRepository.AnalyticWriterRepository)

type AnalyticWriterRepository interface {
	CreateSSPToExchangeRequestEvent(context.Context, model.CreateSSPToExchangeRequestEventReq) error
	CreateExchangeToDSPRequestEvent(context.Context, model.CreateExchangeToDSPRequestEventReq) error
	CreateDSPToExchangeResponseEvent(context.Context, model.CreateDSPToExchangeResponseEventReq) error
	CreateExchangeToSSPResponseEvent(context.Context, model.CreateExchangeToSSPResponseEventReq) error
	CreateSSPBillingEventReq(context.Context, model.CreateSSPBillingEventReq) error
	CreateDSPBillingEventReq(context.Context, model.CreateDSPBillingEventReq) error
}

func NewAnalyticWriterService(
	analyticWriterRepository AnalyticWriterRepository,
) *AnalyticWriterService {
	return &AnalyticWriterService{
		analyticWriterRepository: analyticWriterRepository,
	}
}
