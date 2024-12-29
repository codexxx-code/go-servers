package repository

import (
	"context"

	"pkg/errors"
)

func (r *ExchangeRepository) GetADM(ctx context.Context, id string) (string, error) {

	// Получаем из Redis
	adm, err := r.redisADMs.Get(ctx, id).Result()
	if err != nil {
		switch {
		case errors.IsContextError(err):
			return "", errors.Timeout.Wrap(err)
		default:
			return "", errors.InternalServer.Wrap(err)
		}
	}

	return adm, nil
}
