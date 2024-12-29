package repository

import (
	"context"

	"exchange/internal/config"
	"exchange/internal/services/analyticWriter/model"
)

func (r *AnalyticWriterRepository) CreateExchangeToSSPResponseEvent(ctx context.Context, req model.CreateExchangeToSSPResponseEventReq) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Analytic.ExchangeToSSPResponses)
}
