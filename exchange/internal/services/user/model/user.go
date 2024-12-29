package model

import (
	"time"

	"exchange/internal/enum/permission"
)

type User struct {
	ID           string
	Email        string
	FirstName    string
	LastName     string
	Permissions  []permission.Permission
	AuthorID     *string
	PasswordHash []byte
	PasswordSalt []byte
	CreatedAt    time.Time
	DeletedAt    time.Time
	LastLoginAt  *time.Time
}
