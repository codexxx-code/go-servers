package repository

import (
	"time"

	"exchange/internal/services/ssp/model"
	"pkg/cache"
	"pkg/sql"
)

type SSPRepository struct {
	pgsql sql.SQL
	cache *cache.ListCache[model.SSP]
}

func NewSSPRepository(
	pgsql sql.SQL,
) *SSPRepository {
	return &SSPRepository{
		pgsql: pgsql,
		cache: cache.NewListCache[model.SSP](5 * time.Second),
	}
}
