package repository

import (
	"context"

	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/repository/analyticDTODDL"
	"pkg/errors"
)

func (r *ExchangeRepository) SaveAnalyticDTOs(ctx context.Context, analyticDTOs []model.AnalyticDTO) error {

	analyticDTOsAny := make([]any, len(analyticDTOs))
	for i, analyticDTO := range analyticDTOs {
		analyticDTOsAny[i] = analyticDTO
	}

	if _, err := r.mongo.Collection(analyticDTODDL.Collection).InsertMany(ctx, analyticDTOsAny); err != nil {
		switch {
		case errors.IsContextError(err):
			return errors.Timeout.Wrap(err)
		default:
			return errors.InternalServer.Wrap(err)
		}
	}

	return nil
}
