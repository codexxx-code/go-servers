package repository

import (
	"time"

	dspModel "exchange/internal/services/dsp/model"
	"pkg/cache"
	"pkg/sql"
)

type DSPRepository struct {
	pgsql sql.SQL
	cache *cache.ListCache[dspModel.DSP]
}

func NewDSPRepository(
	pgsql sql.SQL,
) *DSPRepository {
	return &DSPRepository{
		pgsql: pgsql,
		cache: cache.NewListCache[dspModel.DSP](5 * time.Second),
	}
}
