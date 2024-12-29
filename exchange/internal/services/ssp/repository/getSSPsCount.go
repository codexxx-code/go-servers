package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	sspModel "exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/repository/sspDDL"
	"pkg/ddlHelper"
)

func (r *SSPRepository) GetSSPsCount(ctx context.Context, req sspModel.FindSSPsReq) (count int, err error) {

	// Формируем первичный запрос
	q := sq.
		Select(ddlHelper.Count(ddlHelper.SelectAll)).
		From(sspDDL.Table)

	// Добавляем фильтры
	q = GetSSPsFilters(req.Filters, q)

	// Выполняем запрос
	row, err := r.pgsql.QueryRow(ctx, q)
	if err != nil {
		return 0, err
	}

	// Получаем количество пользователей
	return count, row.Scan(&count)
}
