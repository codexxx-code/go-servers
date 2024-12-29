package repository

import (
	"context"
	"time"

	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/repository/dspResponseDDL"

	"pkg/errors"
)

func (r *ExchangeRepository) CreateDSPResponse(ctx context.Context, _ int, res model.DSPResponse) error {

	res.RecordDateTime = time.Now()
	_, err := r.mongo.Collection(dspResponseDDL.Collection).InsertOne(ctx, res)
	if err != nil {
		switch {
		case errors.IsContextError(err):
			return errors.Timeout.Wrap(err)
		default:
			return errors.InternalServer.Wrap(err)
		}
	}
	return nil
}
