package repository

import (
	"pkg/sql"
)

type UserRepository struct {
	pgsql sql.SQL
}

func NewUserRepository(pgsql sql.SQL) *UserRepository {
	return &UserRepository{
		pgsql: pgsql,
	}
}
