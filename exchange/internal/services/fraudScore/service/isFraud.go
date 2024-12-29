package service

import (
	"context"

	"exchange/internal/services/fraudScore/model"
	"pkg/errors"
)

// IsFraud возвращает false, если по пользователю не обнаружено нарушений.
func (fs *FraudScoreService) IsFraud(ctx context.Context, req model.IsFraudReq) (bool, error) {

	// Singleflight для предотвращения дублирующихся запросов.
	v, err, _ := fs.singleflight.Do(req.GetHash(), func() (any, error) {

		// Вызов слоя сети.
		return fs.client.CheckFraudScore(ctx, req)
	})
	if err != nil {
		return true, err
	}
	isFraud, ok := v.(bool)
	if !ok {
		return false, errors.InternalServer.New("Variable is not bool")
	}
	return isFraud, nil
}
