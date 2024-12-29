package repository

import (
	"context"

	"exchange/internal/config"
	"exchange/internal/services/analyticWriter/model"
)

func (r *AnalyticWriterRepository) CreateDSPToExchangeResponseEvent(ctx context.Context, req model.CreateDSPToExchangeResponseEventReq) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Analytic.DSPToExchangeResponses)
}
