package model

import "exchange/internal/enum/permission"

type Claims struct {
	Permissions []permission.Permission `json:"permissions"`
	ID          string                  `json:"id"`
}
