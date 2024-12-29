package service

import (
	"context"

	"pkg/openrtb"
)

func (s *EventService) CreateExchangeBidResponseToSSPEvent(ctx context.Context, req openrtb.BidResponse) error {
	return s.eventRepository.CreateExchangeBidResponseToSSPEvent(ctx, req)
}
