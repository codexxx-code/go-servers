package model

import (
	"exchange/internal/enum/permission"
)

type UpdateUserPermissionsReq struct {
	ID          string
	Permissions []permission.Permission
}

func (u UpdateUserPermissionsReq) Validate() error {
	for _, permission := range u.Permissions {
		if err := permission.Validate(); err != nil {
			return err
		}
	}
	return nil
}
