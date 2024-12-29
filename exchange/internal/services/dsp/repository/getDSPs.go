package repository

import (
	"context"

	dspModel "exchange/internal/services/dsp/model"
	"pkg/slices"
)

// GetDSPs Возвращает DSP по типизированным фильтрам из кэша
func (r *DSPRepository) GetDSPs(ctx context.Context, req dspModel.GetDSPsReq) (dsps []dspModel.DSP, err error) {

	// Получаем данные из кэша
	dsps, ok := r.cache.Get()
	if !ok { // Если данные протухли

		// Получаем данные из БД
		if dsps, err = r.FindDSPs(ctx, dspModel.FindDSPsReq{}); err != nil { //nolint:exhaustruct
			return dsps, err
		}

		// Обновляем кэш
		r.cache.Set(dsps)
	}

	// Динамически фильтруем объекты
	return slices.Filter(dsps, func(dsp dspModel.DSP) bool {

		// Включенность DSP
		if req.IsEnable != nil && dsp.IsEnable != *req.IsEnable {
			return false
		}

		// Фильтр по слагам
		if len(req.Slugs) != 0 && !slices.Contains(req.Slugs, dsp.Slug) {
			return false
		}

		// Фильтр по типу интеграции
		if len(req.SourceTrafficType) != 0 && !slices.ContainsSlice(dsp.SourceTrafficTypes, req.SourceTrafficType...) {
			return false
		}

		return true
	}), nil
}
