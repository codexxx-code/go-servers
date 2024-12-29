package repository

import (
	"context"

	userModel "exchange/internal/services/user/model"
	userRepoModel "exchange/internal/services/user/repository/model"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/enum/table/direction"
	"exchange/internal/services/user/repository/userDDL"
	"pkg/ddlHelper"

	"exchange/internal/services/user/model"
)

// FindUsers Возвращает пользователей по фильтрам
func (r *UserRepository) FindUsers(ctx context.Context, req model.FindUsersReq) (users []model.User, err error) {

	// Формируем первичный запрос
	q := sq.
		Select(ddlHelper.SelectAll).
		From(userDDL.Table)

	q = getUsersFilters(req.Filters, q)

	// Добавление сортировок
	if len(req.Sorts) != 0 {
		for _, sort := range req.Sorts {
			fieldName := sort.UserField.ConvertToDDL()
			// Добавляем сортировку
			switch {
			case sort.Direction == direction.Asc:
				q = q.OrderBy(ddlHelper.Asc(fieldName))
			case sort.Direction == direction.Desc:
				q = q.OrderBy(ddlHelper.Desc(fieldName))
			}
		}
	} else {

		// Сортировка по умолчанию по id
		q = q.OrderBy(ddlHelper.Asc(userDDL.ColumnID))
	}

	// Пагинация
	if req.Pagination.Size != 0 && req.Pagination.Page != 0 {
		q = q.Offset(uint64((req.Pagination.Page - 1) * req.Pagination.Size))
		q = q.Limit(uint64(req.Pagination.Size))
	}

	// Выполняем SQL-запрос
	var repoUsers []userRepoModel.User
	if err = r.pgsql.Select(ctx, &repoUsers, q); err != nil {
		return users, err
	}

	users = make([]userModel.User, 0, len(repoUsers))
	for _, user := range repoUsers {
		users = append(users, user.ConvertToModel())
	}

	// Выполнение SQL-запроса для получения списка SSP
	return users, nil
}
