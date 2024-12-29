package service

import (
	"context"

	eventRepository "exchange/internal/services/event/repository"
	exchangeModel "exchange/internal/services/exchange/model"
	exchangeRepository "exchange/internal/services/exchange/repository"
	"pkg/openrtb"
)

type EventService struct {
	eventRepository    EventRepository
	exchangeRepository ExchangeRepository
}

var _ EventRepository = new(eventRepository.EventRepository)

type EventRepository interface {
	CreateSSPBidRequestToExchangeEvent(context.Context, openrtb.BidRequest) error
	CreateExchangeBidResponseToSSPEvent(context.Context, openrtb.BidResponse) error
	CreateExchangeBidRequestToDSPEvent(context.Context, openrtb.BidRequest) error

	IncrementLoadsForPublisher(context.Context, string) error
	IncrementViewsForPublisher(context.Context, string) error
}

var _ ExchangeRepository = new(exchangeRepository.ExchangeRepository)

type ExchangeRepository interface {
	GetDSPResponses(ctx context.Context, req exchangeModel.GetDSPResponsesReq) (bidResponses []exchangeModel.DSPResponse, err error)
}

func NewEventService(
	eventRepository EventRepository,
	exchangeRepository ExchangeRepository,
) *EventService {
	return &EventService{
		eventRepository:    eventRepository,
		exchangeRepository: exchangeRepository,
	}
}
