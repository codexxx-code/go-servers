package repository

import (
	"context"

	"exchange/internal/config"
	"pkg/openrtb"
)

func (r *EventRepository) CreateExchangeBidResponseToSSPEvent(ctx context.Context, req openrtb.BidResponse) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Event.ExchangeBidResponsesToSSP)
}
