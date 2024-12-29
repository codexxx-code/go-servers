package repository

import (
	"context"

	"exchange/internal/config"
	"exchange/internal/services/analyticWriter/model"
)

func (r *AnalyticWriterRepository) CreateSSPToExchangeRequestEvent(ctx context.Context, req model.CreateSSPToExchangeRequestEventReq) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Analytic.SSPToExchangeRequests)
}
