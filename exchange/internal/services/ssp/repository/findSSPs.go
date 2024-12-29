package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/enum/table/direction"
	sspModel "exchange/internal/services/ssp/model"
	sspRepoModel "exchange/internal/services/ssp/repository/model"
	"exchange/internal/services/ssp/repository/sspDDL"
	"pkg/ddlHelper"
	"pkg/slices"
)

// FindSSPs Возвращает SSP по фильтрам, сортировкам и пагинации
func (r *SSPRepository) FindSSPs(ctx context.Context, req sspModel.FindSSPsReq) (ssps []sspModel.SSP, err error) {

	// Формируем начальный запрос
	q := sq.
		Select(ddlHelper.SelectAll).
		From(sspDDL.Table)

	// Применяем сортировки
	if len(req.Sorts) != 0 {
		for _, sort := range req.Sorts {
			fieldName := sort.SSPField.ConvertToDDL()
			switch {
			case sort.Direction == direction.Asc:
				q = q.OrderBy(ddlHelper.Asc(fieldName))
			case sort.Direction == direction.Desc:
				q = q.OrderBy(ddlHelper.Desc(fieldName))
			}
		}
	} else {
		// Сортировка по умолчанию по slug
		q = q.OrderBy(ddlHelper.Asc(sspDDL.ColumnSlug))
	}

	// Применяем фильтрацию
	q = GetSSPsFilters(req.Filters, q)

	// Пагинация
	if req.Pagination.Size != 0 && req.Pagination.Page != 0 {
		q = q.Offset(uint64((req.Pagination.Page - 1) * req.Pagination.Size)).
			Limit(uint64(req.Pagination.Size))
	}

	// Выполняем SQL-запрос
	var repoSSPs []sspRepoModel.SSP
	if err = r.pgsql.Select(ctx, &repoSSPs, q); err != nil {
		return ssps, err
	}

	ssps = make([]sspModel.SSP, 0, len(repoSSPs))
	for _, ssp := range repoSSPs {
		ssps = append(ssps, ssp.ConvertToModel())
	}

	// Получаем связки SSP и DSP
	sspSlugs := slices.GetFields(repoSSPs, func(ssp sspRepoModel.SSP) string { return ssp.Slug })
	dspsToSSPs, err := r.getSSPsToDSPs(ctx, sspSlugs)
	if err != nil {
		return nil, err
	}

	// Связываем SSP и DSP
	mapDSPsToSSPs := make(map[string][]string)
	for _, dspToSSP := range dspsToSSPs {
		mapDSPsToSSPs[dspToSSP.SSPSlug] = append(mapDSPsToSSPs[dspToSSP.SSPSlug], dspToSSP.DSPSlug)
	}
	for i, ssp := range ssps {
		ssps[i].DSPs = mapDSPsToSSPs[ssp.Slug]
	}

	// Выполнение SQL-запроса для получения списка SSP
	return ssps, nil
}
