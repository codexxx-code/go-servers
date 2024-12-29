package repository

import (
	"context"

	"exchange/internal/config"
	"exchange/internal/services/analyticWriter/model"
)

func (r *AnalyticWriterRepository) CreateExchangeToDSPRequestEvent(ctx context.Context, req model.CreateExchangeToDSPRequestEventReq) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Analytic.ExchangeToDSPRequests)
}
