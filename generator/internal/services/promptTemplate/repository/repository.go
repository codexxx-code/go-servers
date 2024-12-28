package repository

import (
	"pkg/sql"
)

type PromptTemplateRepository struct {
	db sql.SQL
}

func NewPromptTemplateRepository(db sql.SQL) *PromptTemplateRepository {
	return &PromptTemplateRepository{
		db: db,
	}
}
