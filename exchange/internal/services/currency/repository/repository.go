package repository

import (
	"time"

	"exchange/internal/services/currency/model"
	"pkg/cache"
	"pkg/sql"
)

type CurrencyRepository struct {
	pgsql sql.SQL
	cache *cache.ListCache[model.Currency]
}

func NewCurrencyRepository(pgsql sql.SQL) *CurrencyRepository {
	return &CurrencyRepository{
		pgsql: pgsql,
		cache: cache.NewListCache[model.Currency](5 * time.Second),
	}
}
