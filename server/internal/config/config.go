package config

import (
	env "github.com/caarlos0/env/v11"

	"pkg/database/postgresql"
	"pkg/errors"
)

// Config - общая структура конфига
type Config struct {

	// Адрес для http-сервера
	HTTP string `env:"LISTEN_HTTP" envDefault:":8080"`

	// Данные базы данных
	Repository postgresql.PostgreSQLConfig
	DBName     string `env:"DB_NAME" envDefault:"coin"`

	ServiceName string `env:"SERVICE_NAME" envDefault:"coin"`
}

// GetConfig возвращает конфигурацию из .env файла
func GetConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, errors.InternalServer.Wrap(err)
	}
	return config, nil
}
