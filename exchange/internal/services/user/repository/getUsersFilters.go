package repository

import (
	"github.com/Masterminds/squirrel"

	"exchange/internal/services/user/model/userFilters"
	"exchange/internal/services/user/repository/userDDL"
)

// getUsersFilters применяет фильтры для запроса пользователей
func getUsersFilters(filters userFilters.UserFilters, q squirrel.SelectBuilder) squirrel.SelectBuilder {
	// Фильтрация по ID
	if len(filters.IDs) > 0 {
		q = q.Where(squirrel.Eq{userDDL.ColumnID: filters.IDs})
	}

	// Фильтрация по LastNames
	if len(filters.LastNames) > 0 {
		q = q.Where(squirrel.Eq{userDDL.ColumnLastName: filters.LastNames})
	}

	// Фильтрация по FirstNames
	if len(filters.FirstNames) > 0 {
		q = q.Where(squirrel.Eq{userDDL.ColumnFirstName: filters.FirstNames})
	}

	if filters.IsDeleted != nil {
		q = q.Where(squirrel.Eq{userDDL.ColumnIsDeleted: filters.IsDeleted})
	}

	// Фильтрация по Email
	if len(filters.Emails) > 0 {
		q = q.Where(squirrel.Eq{userDDL.ColumnEmail: filters.Emails})
	}

	// Фильтрация по AuthorID
	if len(filters.AuthorUuids) > 0 {
		q = q.Where(squirrel.Eq{userDDL.ColumnAuthorID: filters.AuthorUuids})
	}

	// Фильтрация по LastLoginAt (диапазон)
	if filters.LastLoginAtFrom != nil {
		q = q.Where(squirrel.GtOrEq{userDDL.ColumnLastLoginAt: filters.LastLoginAtFrom.AsTime()})
	}
	if filters.LastLoginAtTo != nil {
		q = q.Where(squirrel.LtOrEq{userDDL.ColumnLastLoginAt: filters.LastLoginAtTo.AsTime()})
	}

	return q
}
