package service

import (
	"context"

	"exchange/internal/services/setting/model"
)

func (s *SettingService) GetSettings(ctx context.Context) (settings model.Settings, err error) {
	return s.settingsRepository.GetSettings(ctx)
}
