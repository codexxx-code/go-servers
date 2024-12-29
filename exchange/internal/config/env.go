package config

import (
	"context"
	"sync/atomic"

	"github.com/caarlos0/env/v11"

	"pkg/database/clickhouse"
	"pkg/database/mongo"
	"pkg/database/pgsql"
	"pkg/database/redis"
	"pkg/errors"
	"pkg/http/fiber"
	"pkg/kafka"
	"pkg/log"
)

// A Config is an implementation for environment variables.
type Config struct {
	FiberServer fiber.ServerSettingsEnv

	Kafka kafka.KafkaSettingsEnv

	Pgsql pgsql.PostgreSQLConfig

	Clickhouse clickhouse.ClickhouseConfig

	Mongo mongo.SettingsMongoConfig

	Redis redis.RedisConfigEnv

	GeoLiteFilePath string `env:"GEO_LITE_FILE_PATH,required"`

	Queue struct {
		Topic struct {
			Event struct {
				SSPBidRequestsToExchange  string `env:"QUEUE_TOPIC_EVENT_SSP_BID_REQUESTS_TOPIC_NAME,required"`       // Запросы от SSP в AdEx
				ExchangeBidResponsesToSSP string `env:"QUEUE_TOPIC_EVENT_EXCHANGE_BID_RESPONSES_TOPIC_NAME,required"` // Ответы от AdEx в SSP
				ExchangeBidRequestsToDSP  string `env:"QUEUE_TOPIC_EVENT_EXCHANGE_BID_REQUESTS_TOPIC_NAME,required"`  // Запросы от AdEx в DSP
			}

			OldAnalytic struct {
				SSPWins string `env:"QUEUE_TOPIC_OLD_ANALYTIC_SSP_WINS_TOPIC_NAME,required"`
				DSPWins string `env:"QUEUE_TOPIC_OLD_ANALYTIC_DSP_WINS_TOPIC_NAME,required"`
			}

			Analytic struct {
				SSPToExchangeRequests  string `env:"QUEUE_TOPIC_ANALYTIC_SSP_TO_EXCHANGE_REQUESTS_TOPIC_NAME,required"`
				ExchangeToDSPRequests  string `env:"QUEUE_TOPIC_ANALYTIC_EXCHANGE_TO_DSP_REQUESTS_TOPIC_NAME,required"`
				DSPToExchangeResponses string `env:"QUEUE_TOPIC_ANALYTIC_DSP_TO_EXCHANGE_RESPONSES_TOPIC_NAME,required"`
				ExchangeToSSPResponses string `env:"QUEUE_TOPIC_ANALYTIC_EXCHANGE_TO_SSP_RESPONSES_TOPIC_NAME,required"`
				SSPBillings            string `env:"QUEUE_TOPIC_ANALYTIC_SSP_BILLINGS_TOPIC_NAME,required"`
				DSPBillings            string `env:"QUEUE_TOPIC_ANALYTIC_DSP_BILLINGS_TOPIC_NAME,required"`
			}
		}
	}

	Auth struct {
		GeneralSalt       string `env:"AUTH_GENERAL_SALT,required"`
		SigningKey        string `env:"AUTH_SIGNING_KEY,required"`
		AccessTokenTTL    string `env:"AUTH_ACCESS_TOKEN_TTL,required"`
		RefreshTokenTTL   string `env:"AUTH_REFRESH_TOKEN_TTL,required"`
		RootAdminLogin    string `env:"AUTH_ROOT_ADMIN_LOGIN,required"`
		RootAdminPassword string `env:"AUTH_ROOT_ADMIN_PASSWORD,required"`
	}

	FraudScoreKey string `env:"FRAUD_SCORE_KEY,required"`
	ServiceName   string `env:"SERVICE_NAME,required"`
	Environment   string `env:"ENV,required"`

	Host string `env:"HOST,required"`

	// Переменная, которая отключает отправку нашего типа ошибок клиенту, а вместо этого шлет просто 204 в случае ошибки
	// false - отправляем ошибки на запросы, нужно для разработки
	// true - обычное состояние, используемое на проде, чтобы не выдавать свои проблемы клиенту
	IsSendNoContentInsteadErrorsInResponse bool `env:"IS_SEND_NO_CONTENT_INSTEAD_ERRORS_IN_RESPONSE,required"`

	// Переменная, которая отключает отправку успешных ответов клиенту, а вместо этого шлет просто 204, если аукцион прошел успешно и мы готовы отправить рекламу клиенту
	// false - отправляем настоящие ответы ortb.BidResponse клиенту, это нормальное состояние прода и разработки
	// true - отправляем на успешный аукцион все равно 204, используем, если сервис находится на этапе разработки и не готовы торговать по настоящему, но готовы получать трафик
	IsSendNoContentForSuccessResponse bool `env:"IS_SEND_NO_CONTENT_FOR_SUCCESS_RESPONSE,required"`
}

var instance *Config
var alreadyCalled atomic.Bool

// Load returns a new Config.
func Load() *Config {

	// Если уже вызывали, то возвращаем тот же инстанс
	if alreadyCalled.Load() {
		return instance
	}

	// Обозначаем, что уже вызывали эту функцию
	alreadyCalled.Store(true)

	// Создаем новый инстанс
	instance = new(Config)

	// Парсим env'ы в инстанс
	if err := env.Parse(instance); err != nil {
		log.Fatal(context.Background(), errors.InternalServer.Wrap(err))
	}

	return instance
}

func Reset() {

	// Сбрасываем положение индикатора вызова функции
	alreadyCalled.Store(false)
}
