package repository

import (
	"context"

	"exchange/internal/config"
	"exchange/internal/services/analyticWriter/model"
)

func (r *AnalyticWriterRepository) CreateDSPBillingEventReq(ctx context.Context, req model.CreateDSPBillingEventReq) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Analytic.DSPBillings)
}
