package model

import (
	"time"

	"github.com/lib/pq"

	"exchange/internal/enum/permission"
	"exchange/internal/services/user/model"
)

type User struct {
	ID           string         `db:"id"`
	Email        string         `db:"email"`
	FirstName    string         `db:"first_name"`
	LastName     string         `db:"last_name"`
	Permissions  pq.StringArray `db:"permissions"`
	AuthorID     *string        `db:"author_id"`
	PasswordHash []byte         `db:"password_hash"`
	PasswordSalt []byte         `db:"password_salt"`
	CreatedAt    time.Time      `db:"-"`
	DeletedAt    time.Time      `db:"-"`
	LastLoginAt  *time.Time     `db:"last_login_at"`
}

func (d *User) ConvertToModel() model.User {

	permissions := make([]permission.Permission, 0, len(d.Permissions))
	for _, _permission := range d.Permissions {
		permissions = append(permissions, permission.Permission(_permission))
	}

	return model.User{
		ID:           d.ID,
		Email:        d.Email,
		FirstName:    d.FirstName,
		LastName:     d.LastName,
		Permissions:  permissions,
		AuthorID:     d.AuthorID,
		PasswordHash: d.PasswordHash,
		PasswordSalt: d.PasswordSalt,
		CreatedAt:    d.CreatedAt,
		DeletedAt:    d.DeletedAt,
		LastLoginAt:  d.LastLoginAt,
	}
}
