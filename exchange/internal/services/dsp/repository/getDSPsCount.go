package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/dsp/repository/dspDDL"
	"pkg/ddlHelper"
)

func (r *DSPRepository) GetDSPsCount(ctx context.Context, req dspModel.FindDSPsReq) (count int, err error) {

	// Формируем первичный запрос
	q := sq.
		Select(ddlHelper.Count(ddlHelper.SelectAll)).
		From(dspDDL.Table)

	// Применяем фильтрацию
	q = GetDSPsFilters(req.Filters, q)

	// Выполняем запрос
	row, err := r.pgsql.QueryRow(ctx, q)
	if err != nil {
		return 0, err
	}

	// Получаем количество пользователей
	return count, row.Scan(&count)
}
