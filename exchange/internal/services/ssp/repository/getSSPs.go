package repository

import (
	"context"

	sspModel "exchange/internal/services/ssp/model"
	"pkg/slices"
)

// GetSSPs Возвращает SSP по типизированным фильтрам из кэша
func (r *SSPRepository) GetSSPs(ctx context.Context, req sspModel.GetSSPsReq) (ssps []sspModel.SSP, err error) {

	// Получаем данные из кэша
	ssps, ok := r.cache.Get()
	if !ok { // Если данные протухли

		// Получаем данные из БД
		if ssps, err = r.FindSSPs(ctx, sspModel.FindSSPsReq{}); err != nil { //nolint:exhaustruct
			return ssps, err
		}

		// Обновляем кэш
		r.cache.Set(ssps)
	}

	// Динамически фильтруем объекты
	return slices.Filter(ssps, func(dsp sspModel.SSP) bool {

		// Включенность DSP
		if req.IsEnable != nil && dsp.IsEnable != *req.IsEnable {
			return false
		}

		// Фильтр по слагам
		if len(req.Slugs) != 0 && !slices.Contains(req.Slugs, dsp.Slug) {
			return false
		}

		return true

	}), nil
}
