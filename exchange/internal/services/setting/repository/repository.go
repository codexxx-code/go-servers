package repository

import (
	"time"

	settingsModel "exchange/internal/services/setting/model"
	"pkg/cache"
	"pkg/sql"
)

type SettingsRepository struct {
	pgsql sql.SQL
	cache *cache.ListCache[settingsModel.Settings]
}

func NewSettingsRepository(
	pgsql sql.SQL,
) *SettingsRepository {
	return &SettingsRepository{
		pgsql: pgsql,
		cache: cache.NewListCache[settingsModel.Settings](5 * time.Second),
	}
}
