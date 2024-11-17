package service

import "pkg/sql"

type TemplaterRepository struct {
	sql sql.SQL
}

func NewTemplaterRepository(
	sql sql.SQL,
) *TemplaterRepository {
	return &TemplaterRepository{
		sql: sql,
	}
}
