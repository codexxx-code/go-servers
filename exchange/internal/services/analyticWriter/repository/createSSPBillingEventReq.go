package repository

import (
	"context"

	"exchange/internal/config"
	"exchange/internal/services/analyticWriter/model"
)

func (r *AnalyticWriterRepository) CreateSSPBillingEventReq(ctx context.Context, req model.CreateSSPBillingEventReq) error {
	return r.writeToTopic(ctx, req, config.Load().Queue.Topic.Analytic.SSPBillings)
}
