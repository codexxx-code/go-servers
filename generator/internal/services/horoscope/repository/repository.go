package repository

import (
	"pkg/sql"
)

type HoroscopeRepository struct {
	db sql.SQL
}

func NewHoroscopeRepository(db sql.SQL, ) *HoroscopeRepository {
	return &HoroscopeRepository{
		db: db,
	}
}
