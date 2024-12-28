package config

import (
	"github.com/caarlos0/env/v11"

	"pkg/errors"
)

// Config - общая структура конфига
type Config struct {

	// Адрес для http-сервера
	HTTP string `env:"LISTEN_HTTP,required"`

	Ebay EbayConfig

	ServiceName string `env:"SERVICE_NAME,required"`
}

type EbayConfig struct {
	ClientID     string `env:"EBAY_CLIENT_ID,required"`
	ClientSecret string `env:"EBAY_CLIENT_SECRET,required"`
	IsSandbox    bool   `env:"EBAY_IS_SANDBOX,required"`
	CampaignID   string `env:"EBAY_CAMPAIGN_ID,required"`
}

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, errors.InternalServer.Wrap(err)
	}
	return config, nil
}
