package repository

import (
	"context"
	"encoding/json"

	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/repository/unusedBidsDDL"
	"pkg/errors"
)

func (r *ExchangeRepository) SaveUnusedBids(ctx context.Context, bids []model.BidPointer) error {

	// Преобразуем биды в JSON
	bidJSONs := make([]any, 0, len(bids))
	for _, bid := range bids {
		bidJSON, err := json.Marshal(bid)
		if err != nil {
			return errors.InternalServer.Wrap(err)
		}
		bidJSONs = append(bidJSONs, string(bidJSON))
	}

	// Добавляем запись с конца
	if err := r.redisUnusedBids.RPush(ctx, unusedBidsDDL.UnusedBidsListName, bidJSONs...).Err(); err != nil {
		switch {
		case errors.IsContextError(err):
			return errors.Timeout.Wrap(err)
		default:
			return errors.InternalServer.Wrap(err)
		}
	}

	return nil
}
