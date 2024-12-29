package userFilters

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"exchange/internal/enum/permission"
	"pkg/errors"
)

type UserFilters struct {
	IDs             []string                // Список уникальных идентификаторов пользователей
	Emails          []string                // Список email-адресов пользователей
	FirstNames      []string                // Список имен пользователей
	LastNames       []string                // Список фамилий пользователей
	IsDeleted       *bool                   // Флаг удаленности пользователя
	DeletedAtFrom   *timestamppb.Timestamp  // Верхняя граница фильтра по дате удаления
	DeletedAtTo     *timestamppb.Timestamp  // Нижняя граница фильтра по дате удаления
	AuthorUuids     []string                // Список создателей пользователей
	CreatedAtFrom   *timestamppb.Timestamp  // Верхняя граница фильтра по дате создания
	CreatedAtTo     *timestamppb.Timestamp  // Нижняя граница фильтра по дате создания
	LastLoginAtFrom *timestamppb.Timestamp  // Верхняя граница фильтра по дате последнего входа
	LastLoginAtTo   *timestamppb.Timestamp  // Нижняя граница фильтра по дате последнего входа
	Permissions     []permission.Permission // Список прав доступа
}

func (u UserFilters) Validate() error {
	for _, permission := range u.Permissions {
		if err := permission.Validate(); err != nil {
			return err
		}
	}

	if u.DeletedAtTo != nil && !u.DeletedAtTo.IsValid() {
		return errors.BadRequest.New("DeletedAtTo is invalid")
	}

	if u.DeletedAtFrom != nil && !u.DeletedAtFrom.IsValid() {
		return errors.BadRequest.New("DeletedAtFrom is invalid")
	}

	if u.CreatedAtTo != nil && !u.CreatedAtTo.IsValid() {
		return errors.BadRequest.New("CreatedAtTo is invalid")
	}

	if u.CreatedAtFrom != nil && !u.CreatedAtFrom.IsValid() {
		return errors.BadRequest.New("CreatedAtFrom is invalid")
	}

	if u.LastLoginAtTo != nil && !u.LastLoginAtTo.IsValid() {
		return errors.BadRequest.New("LastLoginAtTo is invalid")
	}

	if u.LastLoginAtFrom != nil && !u.LastLoginAtFrom.IsValid() {
		return errors.BadRequest.New("LastLoginAtFrom is invalid")
	}

	return nil
}
