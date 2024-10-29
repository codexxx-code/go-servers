package config

import (
	"github.com/caarlos0/env/v11"

	"pkg/errors"
)

// Config - общая структура конфига
type Config struct {

	// Адрес для http-сервера
	HTTP string `env:"LISTEN_HTTP" envDefault:":8080"`

	Ebay EbayConfig

	ServiceName string `env:"SERVICE_NAME" envDefault:"zodiac"`
}

type EbayConfig struct {
	ClientID     string `env:"EBAY_CLIENT"`
	ClientSecret string `env:"EBAY_CLIENT_SECRET"`
	IsSandbox    bool   `env:"EBAY_IsSandbox"`
}

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, errors.InternalServer.Wrap(err)
	}
	return config, nil
}
