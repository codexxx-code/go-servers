package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/user/model"
	"exchange/internal/services/user/repository/userDDL"
	"pkg/ddlHelper"
)

// GetUsersCount Возвращает количество пользователей
func (r *UserRepository) GetUsersCount(ctx context.Context, req model.FindUsersReq) (count int, err error) {

	// Формируем первичный запрос
	q := sq.
		Select(ddlHelper.Count(ddlHelper.SelectAll)).
		From(userDDL.Table)

	// Добавляем фильтры
	q = getUsersFilters(req.Filters, q)

	// Выполняем запрос
	row, err := r.pgsql.QueryRow(ctx, q)
	if err != nil {
		return 0, err
	}

	// Получаем количество пользователей
	return count, row.Scan(&count)
}
