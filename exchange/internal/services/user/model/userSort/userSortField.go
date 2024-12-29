package userSort

import (
	"exchange/internal/services/user/repository/userDDL"
	"pkg/errors"
)

type UserSortField int

const (
	UUID = iota + 1
	Email
	FirstName
	Permissions
	DeletedAt
	AuthorUUID
	CreatedAt
	LastLoginAt
	LastName
)

var mappingModelToDDL = map[UserSortField]string{
	UUID:        userDDL.ColumnID,
	Email:       userDDL.ColumnEmail,
	FirstName:   userDDL.ColumnFirstName,
	Permissions: userDDL.ColumnPermissions,
	DeletedAt:   userDDL.ColumnIsDeleted,
	AuthorUUID:  userDDL.ColumnAuthorID,
	CreatedAt:   "",
	LastLoginAt: userDDL.ColumnLastLoginAt,
	LastName:    userDDL.ColumnLastName,
}

// Validate проверяет, что UserSortField имеет допустимое значение
func (s UserSortField) Validate() error {
	switch s {
	case UUID, Email,
		FirstName, Permissions, DeletedAt,
		AuthorUUID, CreatedAt, LastLoginAt,
		LastName:
		return nil
	default:
		return errors.BadRequest.New("UserSortField undefined")
	}
}

func (s UserSortField) ConvertToDDL() string {
	return mappingModelToDDL[s]
}
