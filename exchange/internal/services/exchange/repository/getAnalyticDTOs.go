package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/repository/analyticDTODDL"
	"pkg/errors"
)

func (r *ExchangeRepository) GetAnalyticDTOs(ctx context.Context, req model.GetAnalyticDTOsReq) (analyticDTOs []model.AnalyticDTO, err error) {

	var filters bson.M

	// Выставляем фильтры
	if len(req.BidIDs) != 0 {
		filters = bson.M{
			analyticDTODDL.FieldExchangeBidID: bson.M{
				"$in": req.BidIDs,
			},
		}
	}

	if len(req.RequestIDs) != 0 {
		filters = bson.M{
			analyticDTODDL.FieldExchangeID: bson.M{
				"$in": req.RequestIDs,
			},
		}
	}

	// Получаем ответы DSP
	cur, err := r.mongo.Collection(analyticDTODDL.Collection).Find(ctx, filters)
	if err != nil {
		switch {
		case errors.IsContextError(err):
			return analyticDTOs, errors.Timeout.Wrap(err)
		default:
			return analyticDTOs, errors.InternalServer.Wrap(err)
		}
	}
	for cur.Next(ctx) {

		// Декодируем запись в структуру
		var analyticDTO model.AnalyticDTO
		if err = cur.Decode(&analyticDTO); err != nil {
			return analyticDTOs, errors.InternalServer.Wrap(err)
		}
		analyticDTOs = append(analyticDTOs, analyticDTO)
	}

	return analyticDTOs, nil
}
