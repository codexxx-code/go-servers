package repository

import (
	"pkg/sql"
)

type ZodiacRepository struct {
	db sql.SQL
}

func NewZodiacRepository(db sql.SQL, ) *ZodiacRepository {
	return &ZodiacRepository{
		db: db,
	}
}
