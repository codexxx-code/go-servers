package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/setting/model"
	"exchange/internal/services/setting/repository/settingDDL"
	"pkg/ddlHelper"
	"pkg/errors"
)

func (r *SettingsRepository) GetSettings(ctx context.Context) (setting model.Settings, err error) {

	// Получаем данные из кэша
	settings, ok := r.cache.Get()
	if !ok { // Если данные протухли

		// Получаем данные из БД
		if err = r.pgsql.Get(ctx, &setting, sq.
			Select(ddlHelper.SelectAll).
			From(settingDDL.Table),
		); err != nil {
			return setting, err
		}

		// Обновляем кэш
		r.cache.Set([]model.Settings{setting})

		return setting, nil
	}

	if len(settings) != 0 {
		return settings[0], nil
	} else {
		return setting, errors.InternalServer.New("Settings not found")
	}
}
