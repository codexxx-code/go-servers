package repository

import (
	"pkg/sql"
)

type ForecastRepository struct {
	db sql.SQL
}

func NewForecastRepository(db sql.SQL, ) *ForecastRepository {
	return &ForecastRepository{
		db: db,
	}
}
