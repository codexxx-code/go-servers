package repository

import "pkg/sql"

type PromptRepository struct {
	db sql.SQL
}

func NewPromptRepository(db sql.SQL) *PromptRepository {
	return &PromptRepository{
		db: db,
	}
}
