package repository

import "pkg/sql"

type GeneratorRepository struct {
	db sql.SQL
}

func NewGeneratorRepository(db sql.SQL) *GeneratorRepository {
	return &GeneratorRepository{
		db: db,
	}
}
