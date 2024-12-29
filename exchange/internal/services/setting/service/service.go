package service

import (
	"context"

	"exchange/internal/services/setting/model"
	settingsRepository "exchange/internal/services/setting/repository"
)

type SettingService struct {
	settingsRepository SettingsRepository
}

func NewSettingsService(
	settingsRepository SettingsRepository,
) *SettingService {
	return &SettingService{
		settingsRepository: settingsRepository,
	}
}

var _ SettingsRepository = &settingsRepository.SettingsRepository{}

type SettingsRepository interface {
	GetSettings(context.Context) (model.Settings, error)
}
