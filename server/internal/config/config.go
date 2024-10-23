package config

import (
	"github.com/caarlos0/env/v11"

	"pkg/database/postgresql"
	"pkg/errors"
)

// Config - общая структура конфига
type Config struct {

	// Адрес для http-сервера
	HTTP string `env:"LISTEN_HTTP" envDefault:":8080"`

	// Данные базы данных
	Repository postgresql.PostgreSQLConfig
	DBName     string `env:"DB_NAME"`

	// Доступы к телеграм-боту
	Telegram struct {
		Enabled bool   `env:"TG_BOT_ENABLED"`
		Token   string `env:"TG_BOT_TOKEN"`
		ChatID  int64  `env:"TG_CHAT_ID"`
	}

	ChatGPTApiKey string `env:"CHATGPT_APIKEY"`

	ServiceName string `env:"SERVICE_NAME" envDefault:"zodiac"`
}

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, errors.InternalServer.Wrap(err)
	}
	return config, nil
}
