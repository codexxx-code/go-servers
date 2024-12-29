package config

import (
	"time"

	"exchange/internal/metrics"
	"pkg/decimal"
	"pkg/errors"
	"pkg/jwtManager"
	"pkg/stackTrace"
)

func InitSingletones(conf *Config) error {

	// Назначаем имя сервиса, чтобы стектрейсы получались только с вызовами нашего приложения
	stackTrace.Init(conf.ServiceName)

	// Параметр, отвечающий за то, что decimal при (де-)сериализации будет форматироваться как float, иначе будет как string "3.14"
	decimal.OffQuotesInJSON()

	// Инициализируем метрики
	if err := metrics.Init(conf.ServiceName); err != nil {
		return err
	}

	// Инициализируем jwt manager
	accessTokenTTL, err := time.ParseDuration(conf.Auth.AccessTokenTTL)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	refreshTokenTTL, err := time.ParseDuration(conf.Auth.RefreshTokenTTL)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}
	jwtManager.Init(
		[]byte(conf.Auth.SigningKey),
		accessTokenTTL,
		refreshTokenTTL,
	)

	return nil
}
