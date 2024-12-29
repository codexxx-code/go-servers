package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/enum/table/direction"
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/dsp/repository/dspDDL"
	dspRepoModel "exchange/internal/services/dsp/repository/model"
	"pkg/ddlHelper"
)

// FindDSPs Возвращает DSP по фильтрам, сортировкам и пагинации
func (r *DSPRepository) FindDSPs(ctx context.Context, req dspModel.FindDSPsReq) (dsps []dspModel.DSP, err error) {

	// Формируем начальный запрос
	q := sq.
		Select(ddlHelper.SelectAll).
		From(dspDDL.Table)

	// Применяем сортировки
	if len(req.Sorts) != 0 {
		for _, sort := range req.Sorts {

			sortField := sort.DSPField.ConvertToDDL()
			if sortField == "" {
				continue
			}

			switch {
			case sort.Direction == direction.Asc:
				q = q.OrderBy(ddlHelper.Asc(sortField))
			case sort.Direction == direction.Desc:
				q = q.OrderBy(ddlHelper.Desc(sortField))
			}
		}
	} else {
		// Сортировка по умолчанию по id
		q = q.OrderBy(ddlHelper.Asc(dspDDL.ColumnSlug))
	}

	// Пагинация
	if req.Pagination.Size != 0 && req.Pagination.Page != 0 {
		q = q.Offset(uint64((req.Pagination.Page - 1) * req.Pagination.Size)).
			Limit(uint64(req.Pagination.Size))
	}

	// Применяем фильтрацию
	q = GetDSPsFilters(req.Filters, q)

	// Выполняем SQL-запрос
	var repoDSPs []dspRepoModel.DSP
	if err = r.pgsql.Select(ctx, &repoDSPs, q); err != nil {
		return dsps, err
	}

	// Маппим модель репозитория в модель бизнеса
	dsps = make([]dspModel.DSP, 0, len(repoDSPs))
	for _, dsp := range repoDSPs {
		dsps = append(dsps, dsp.ConvertToModel())
	}

	return dsps, nil
}
