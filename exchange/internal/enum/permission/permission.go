package permission

import (
	"pkg/errors"
)

type Permission string

const (
	Root                    = "root"
	Admin                   = "admin"
	UserCreate              = "user_create"
	UserUpdate              = "user_update"
	UserPermissionsManaging = "user_permissions_managing"
	UserDelete              = "user_delete"
	UserRemove              = "user_remove"
)

// Validate проверяет, что Permission имеет допустимое значение
func (p Permission) Validate() error {
	switch p {
	case Root, Admin, UserCreate,
		UserUpdate, UserPermissionsManaging, UserDelete, UserRemove:
		return nil
	default:
		return errors.BadRequest.New("Permission undefined")
	}
}

// IsRoot проверяет, имеет ли пользователь корневые права доступа
func IsRoot(permissions []Permission) bool {
	for _, permission := range permissions {
		if permission == Root {
			return true
		}
	}
	return false
}

// ContainsPermissions проверяет, содержит ли пользователь необходимые права доступа
func ContainsPermissions(userPermissions []Permission, neededPermissions []Permission) bool {
	if len(userPermissions) == 0 {
		return len(neededPermissions) == 0
	}

	userPermissionsMap := make(map[Permission]struct{}, len(userPermissions))
	for _, userPermission := range userPermissions {
		userPermissionsMap[userPermission] = struct{}{}
	}

	for _, neededPermission := range neededPermissions {
		if _, ok := userPermissionsMap[neededPermission]; ok {
			return true
		}
	}

	return false
}
