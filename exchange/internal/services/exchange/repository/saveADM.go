package repository

import (
	"context"
	"time"

	"pkg/errors"
)

func (r *ExchangeRepository) SaveADM(ctx context.Context, id, adm string) error {

	// Сохраняем в Redis
	_, err := r.redisADMs.Set(ctx, id, adm, 1*time.Hour).Result()
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
