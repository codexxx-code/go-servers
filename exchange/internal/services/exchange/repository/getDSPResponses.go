package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/repository/dspResponseDDL"
	"pkg/errors"
)

func (r *ExchangeRepository) GetDSPResponses(ctx context.Context, req model.GetDSPResponsesReq) (bidResponses []model.DSPResponse, err error) {

	var filters bson.M

	// Выставляем фильтры
	if len(req.BidIDs) != 0 {
		filters = bson.M{
			dspResponseDDL.FieldBidID: bson.M{
				"$in": req.BidIDs,
			},
		}
	}

	if len(req.RequestIDs) != 0 {
		filters = bson.M{
			dspResponseDDL.FieldRequestID: bson.M{
				"$in": req.RequestIDs,
			},
		}
	}

	// Получаем ответы DSP
	cur, err := r.mongo.Collection(dspResponseDDL.Collection).Find(ctx, filters, &options.FindOptions{ //nolint:exhaustruct
		Limit: req.Limit,
	})
	if err != nil {
		switch {
		case errors.IsContextError(err):
			return bidResponses, errors.Timeout.Wrap(err)
		default:
			return bidResponses, errors.InternalServer.Wrap(err)
		}
	}
	for cur.Next(ctx) {

		// Декодируем запись в структуру
		var response model.DSPResponse
		if err = cur.Decode(&response); err != nil {
			return bidResponses, errors.InternalServer.Wrap(err)
		}
		bidResponses = append(bidResponses, response)
	}

	return bidResponses, nil
}
