package repository

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"

	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/repository/unusedBidsDDL"
	"pkg/errors"
)

func (r *ExchangeRepository) PopUnusedBids(ctx context.Context, needCount int) (bids []model.BidPointer, err error) {

	// Забираем записи из начала списка с удалением
	_, unusedResponses, err := r.redisUnusedBids.LMPop(ctx, "left", int64(needCount), unusedBidsDDL.UnusedBidsListName).Result()
	if err != nil {
		switch {
		case errors.IsContextError(err):
			return nil, errors.Timeout.Wrap(err)
		case errors.Is(err, redis.Nil):
			return nil, nil
		default:
			return nil, errors.InternalServer.Wrap(err)
		}
	}

	// Если не нашлось нужное количество записей
	switch {
	case len(unusedResponses) == 0:
		return nil, nil
	case len(unusedResponses) < needCount:

		interfaces := make([]interface{}, 0, len(unusedResponses))
		for _, value := range unusedResponses {
			interfaces = append(interfaces, value)
		}

		// Снова добавляем записи
		if err = r.redisUnusedBids.RPush(context.TODO(), unusedBidsDDL.UnusedBidsListName, interfaces...).Err(); err != nil {
			switch {
			case errors.IsContextError(err):
				return nil, errors.Timeout.Wrap(err)
			default:
				return nil, errors.InternalServer.Wrap(err)
			}
		}

		return nil, nil
	}

	// Преобразуем JSON в биды
	bids = make([]model.BidPointer, 0, len(unusedResponses))
	for _, value := range unusedResponses {
		var bid model.BidPointer
		if err = json.Unmarshal([]byte(value), &bid); err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}

	return bids, nil
}
