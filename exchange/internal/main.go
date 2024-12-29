package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
	swaggerHandler "github.com/gofiber/swagger"
	"github.com/valyala/fasthttp"
	_ "go.uber.org/automaxprocs"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"

	"pkg/gracefulShutdown"
	"pkg/http/fiber"
	"pkg/kafka"
	"pkg/log/model"

	_ "exchange/docs"
	"exchange/internal/config"
	analyticWriterRepository "exchange/internal/services/analyticWriter/repository"
	analyticWriterService "exchange/internal/services/analyticWriter/service"
	authService "exchange/internal/services/auth/service"
	billingEndpointHTTP "exchange/internal/services/billing/endpoint/http"
	billingNetwork "exchange/internal/services/billing/network"
	billingRepository "exchange/internal/services/billing/repository"
	billingService "exchange/internal/services/billing/service"
	currencyRepository "exchange/internal/services/currency/repository"
	currencyService "exchange/internal/services/currency/service"
	dspRepository "exchange/internal/services/dsp/repository"
	dspService "exchange/internal/services/dsp/service"
	eventRepository "exchange/internal/services/event/repository"
	eventService "exchange/internal/services/event/service"
	exchangeEndpointHTTP "exchange/internal/services/exchange/endpoint/http"
	exchangeNetwork "exchange/internal/services/exchange/network"
	exchangeRepository "exchange/internal/services/exchange/repository"
	exchangeService "exchange/internal/services/exchange/service"
	fraudScoreNetwork "exchange/internal/services/fraudScore/network"
	fraudScoreService "exchange/internal/services/fraudScore/service"
	settingsRepository "exchange/internal/services/setting/repository"
	settingsService "exchange/internal/services/setting/service"
	sspRepository "exchange/internal/services/ssp/repository"
	sspService "exchange/internal/services/ssp/service"
	"exchange/internal/services/transactor"
	userRepository "exchange/internal/services/user/repository"
	userService "exchange/internal/services/user/service"
	"exchange/migrations"
	"pkg/database/clickhouse"
	"pkg/database/geolite"
	"pkg/database/mongo"
	"pkg/database/pgsql"
	"pkg/database/redis"
	"pkg/errors"
	"pkg/log"
	"pkg/migrator"
	"pkg/panicRecover"
)

// @title EXCHANGE FiberServer Documentation
// @version @{version} build @{build}
// @description API Documentation for EXCHANGE service

//go:generate go install github.com/swaggo/swag/cmd/swag@v1.8.2
//go:generate go mod download
//go:generate swag init -o ../docs --parseInternal --parseDependency

const (
	version = "@{version}"
	build   = "@{build}"
	commit  = "@{commit}"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), err)
	}
}

func run() error {

	// Перехватываем возможную панику
	defer panicRecover.PanicRecover(func(err error) {
		log.Fatal(context.Background(), err)
	})

	// Основной контекст приложения
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Парсим флаги
	logFormat := flag.String("log-format", string(log.JSONFormat), "text - Human readable string\njson - JSON format")
	envMode := flag.String("env-mode", "local", "Environment mode for log label: test, prod")
	flag.Parse()

	var logHandlers []log.Handler
	switch *logFormat {
	case "text":
		logHandlers = append(logHandlers, log.NewConsoleHandler(os.Stdout, log.LevelDebug))
	case "json":
		logHandlers = append(logHandlers, log.NewJSONHandler(os.Stdout, log.LevelDebug))
	}

	// Загружаем конфиг
	log.Info(ctx, "loading environment variables")
	conf := config.Load()

	// Конфигурируем все синглтон сервисы
	if err := config.InitSingletones(conf); err != nil {
		return err
	}

	// Получаем имя тачки
	hostname, err := os.Hostname()
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	// Инициализируем логгер
	log.Init(
		model.SystemInfo{
			Hostname: hostname,
			Version:  version,
			Build:    build,
			Commit:   commit,
			Env:      *envMode,
		},
		logHandlers...,
	)

	// Получаем кафку для работы сервиса
	log.Info(ctx, "init kafka producer")
	producerKafka, err := kafka.NewAsyncProducer(conf.Kafka)
	if err != nil {
		return err
	}

	go func() {
		for err := range producerKafka.Errors() {
			var kafkaErr *sarama.ProducerError
			if errors.As(err, &kafkaErr) && kafkaErr != nil && kafkaErr.Msg != nil {

				var key, value []byte
				if kafkaErr.Msg.Key != nil {
					key, _ = kafkaErr.Msg.Key.Encode()
				}
				if kafkaErr.Msg.Value != nil {
					value, _ = kafkaErr.Msg.Value.Encode()
				}

				log.Error(ctx, errors.InternalServer.Wrap(kafkaErr), log.ParamsOption(
					"topic", kafkaErr.Msg.Topic,
					"key", string(key),
					"value", string(value),
					"metadata", fmt.Sprintf("%v", kafkaErr.Msg.Metadata),
					"timestamp", kafkaErr.Msg.Timestamp.String(),
				))
			} else {
				log.Error(ctx, errors.InternalServer.Wrap(err))
			}
		}
	}()

	log.Info(ctx, "create kafka topics")
	err = createKafkaTopics(conf)
	if err != nil {
		return err
	}

	// Подключаемся к базе MongoDB
	log.Info(ctx, "connection to mongo")
	mongo, err := mongo.NewClientMongo(conf.Mongo, conf.ServiceName)
	if err != nil {
		return err
	}

	// Подключаемся к базе данных Postgresql
	log.Info(ctx, "connection to pgsql")
	pgsql, err := pgsql.NewClientPgsql(conf.Pgsql)
	if err != nil {
		return err
	}
	defer pgsql.Close()

	// Подключаемся к базе данных Clickhouse
	log.Info(ctx, "connection to clickhouse")
	clickhouse, err := clickhouse.NewClientClickhouse(conf.Clickhouse)
	if err != nil {
		return err
	}
	defer clickhouse.Close()

	// Запускаем миграции в pgsql
	log.Info(ctx, "run pgsql migrations")
	pgsqlMigrator, err := migrator.NewMigrator(pgsql,
		migrator.MigratorConfig{
			EmbedMigrations: migrations.EmbedMigrationsPgsql,
			Dialect:         migrator.DialectPostgres,
			Dir:             "pgsql",
		},
	)
	if err != nil {
		return err
	}
	if err = pgsqlMigrator.Up(ctx); err != nil {
		return err
	}

	// Запускаем миграции в clickhouse
	log.Info(ctx, "run clickhouse migrations")
	clickhouseMigrator, err := migrator.NewMigrator(clickhouse,
		migrator.MigratorConfig{
			EmbedMigrations: migrations.EmbedMigrationsClickhouse,
			Dialect:         migrator.DialectClickHouse,
			Dir:             "clickhouse",
		},
	)
	if err != nil {
		return err
	}
	if err = clickhouseMigrator.Up(ctx); err != nil {
		return err
	}

	// Подключаемся к редису
	log.Info(ctx, "connection to redis")
	redisADMs, err := redis.NewClientRedis(conf.Redis, 0)
	if err != nil {
		return err
	}
	redisUnusedBids, err := redis.NewClientRedis(conf.Redis, 1)
	if err != nil {
		return err
	}

	// Группа для ожидания готовности всех зависимостей
	dependenciesErrWg := &errgroup.Group{}

	// Получаем подключение к базе данных geoLite
	log.Info(ctx, "start connection to geoLite in goroutine")

	// Аллоцируем переменную для хранения базы геолокации, чтобы потом по магии указателей
	// после парсинга бд она стала доступна в сервисах
	geoLite := new(geolite.Reader)
	defer geoLite.Close()

	// Получаем базу данных геолокаций в горутине
	dependenciesErrWg.Go(func() error {

		// Открываем и парсим файл с базой данных
		_geoLite, err := geolite.NewClientGeoLite(conf.GeoLiteFilePath)
		if err != nil {
			return err
		}

		// Копируем данные из объекта _geoLite в указатель geoLite, который уже передан в нужные сервисы,
		// соответственно в сервисах по заранее аллоцированному переданному указателю полявится база данных
		*geoLite = *_geoLite
		return nil
	})

	httpClient := &fasthttp.Client{
		Name:                          "",
		NoDefaultUserAgentHeader:      false,
		Dial:                          nil,
		DialDualStack:                 false,
		TLSConfig:                     nil,
		MaxConnsPerHost:               0,
		MaxIdleConnDuration:           0,
		MaxConnDuration:               0,
		MaxIdemponentCallAttempts:     0,
		ReadBufferSize:                0,
		WriteBufferSize:               0,
		ReadTimeout:                   0,
		WriteTimeout:                  0,
		MaxResponseBodySize:           0,
		DisableHeaderNamesNormalizing: false,
		DisablePathNormalizing:        false,
		MaxConnWaitTimeout:            0,
		RetryIf:                       nil,
		ConnPoolStrategy:              0,
		StreamResponseBody:            false,
		ConfigureClient:               nil,
	}

	// Инициализируем репозитории
	exchangeRepository := exchangeRepository.NewExchangeRepository(pgsql, mongo, geoLite, redisADMs, redisUnusedBids)
	sspRepository := sspRepository.NewSSPRepository(pgsql)
	dspRepository := dspRepository.NewDSPRepository(pgsql)
	settingsRepository := settingsRepository.NewSettingsRepository(pgsql)
	currencyRepository := currencyRepository.NewCurrencyRepository(pgsql)
	billingRepository := billingRepository.NewBillingRepository(producerKafka)
	userRepository := userRepository.NewUserRepository(pgsql)
	eventRepository := eventRepository.NewEventRepository(producerKafka, mongo)
	analyticWriterRepository := analyticWriterRepository.NewAnalyticWriterRepository(producerKafka)
	transactor := transactor.NewTransactor(pgsql)

	// Инициализируем сетевые слои
	exchangeNetwork := exchangeNetwork.NewExchangeNetwork(httpClient)
	billingNetwork := billingNetwork.NewBillingNetwork(httpClient)
	fraudScoreNetwork := fraudScoreNetwork.NewFraudScoreNetwork(httpClient, conf.FraudScoreKey)

	// Инициализируем сервисы
	fraudScoreService := fraudScoreService.NewFraudscoreService(fraudScoreNetwork)
	sspService := sspService.NewSSPService(sspRepository, transactor)
	dspService := dspService.NewDSPService(dspRepository, transactor)
	currencyService := currencyService.NewCurrencyService(currencyRepository)
	settingsService := settingsService.NewSettingsService(settingsRepository)
	eventService := eventService.NewEventService(eventRepository, exchangeRepository)
	analyticWriterService := analyticWriterService.NewAnalyticWriterService(analyticWriterRepository)
	exchangeService := exchangeService.NewExchangeService(
		exchangeRepository,
		sspService,
		dspService,
		currencyService,
		settingsService,
		exchangeNetwork,
		eventService,
		fraudScoreService,
		analyticWriterService,
	)
	billingService := billingService.NewBillingService(dspService, billingRepository, exchangeRepository, billingNetwork, analyticWriterService)
	userService, err := userService.NewUserService(userRepository, false)
	if err != nil {
		return err
	}
	_ = authService.NewAuthService(userService)

	// Переменная для индикации готовности сервиса к приему трафика
	var isReady = make(chan struct{})

	// Создаем errgroup для управления горутинами запуска серверов
	serversErrWg, ctx := errgroup.WithContext(ctx)

	// Создаем горутину на открытие трафика
	// TODO: Возможно, стоит придумать/поискать мидлвару, которое позволит блокировать любые запросы, если мы не готовы
	// принимать трафик на сервис (оно понятно, что кубер сам распоряжается трафиком, но дополнительная перестраховка не помешает)
	serversErrWg.Go(func() error {

		// Ждем, пока все связанные зависимости подтянутся и будут готовы к работе
		if err := dependenciesErrWg.Wait(); err != nil {
			return err
		}

		// Закрываем канал и тем самым отвечаем 200 на запросы готовности
		close(isReady)

		return nil
	})

	// Настраиваем и получаем HTTP-сервер
	httpServer, err := fiber.GetDefaultServer(conf.ServiceName, conf.FiberServer, isReady)
	if err != nil {
		return err
	}

	// Инициализируем эндпоинты
	httpServer.Get("/openapi/*", swaggerHandler.HandlerDefault)                   // GET /openapi/*
	httpServer.Get("/version", fiber.NewVersionHandler(version, build, hostname)) // GET /version

	exchangeEndpointHTTP.MountExchangeEndpoint(httpServer, exchangeService) // ANY /rtb/*
	billingEndpointHTTP.MountBillingEndpoint(httpServer, billingService)    // ANY /billing/*

	// Создаем горутину на запуск HTTP-сервера
	serversErrWg.Go(func() error {
		log.Info(ctx, fmt.Sprintf("HTTP server is listening :%v", conf.FiberServer.Port))
		if err = httpServer.Listen(net.JoinHostPort("", conf.FiberServer.Port)); err != nil {
			return errors.InternalServer.Wrap(err)
		}
		return nil
	})

	// Добавляем в группу горутину для graceful shutdown
	gracefulShutdown.AddGracefulShutdownErrGroup(serversErrWg, ctx, httpServer, nil)

	return serversErrWg.Wait()
}

// TODO: Вынести куда-нибудь
const (
	partitions        = 10
	replicationFactor = 1
)

func createKafkaTopics(conf *config.Config) error {

	topics := []string{
		conf.Queue.Topic.Event.ExchangeBidResponsesToSSP,
		conf.Queue.Topic.Event.SSPBidRequestsToExchange,
		conf.Queue.Topic.Event.SSPBidRequestsToExchange,
		conf.Queue.Topic.OldAnalytic.SSPWins,
		conf.Queue.Topic.OldAnalytic.DSPWins,
		conf.Queue.Topic.Analytic.SSPToExchangeRequests,
		conf.Queue.Topic.Analytic.ExchangeToDSPRequests,
		conf.Queue.Topic.Analytic.DSPToExchangeResponses,
		conf.Queue.Topic.Analytic.ExchangeToSSPResponses,
		conf.Queue.Topic.Analytic.SSPBillings,
		conf.Queue.Topic.Analytic.DSPBillings,
	}

	createTopicRequests := make([]kafka.CreateTopicRequest, 0, len(topics))

	for _, topic := range topics {
		createTopicRequests = append(createTopicRequests, kafka.CreateTopicRequest{
			Topic:             topic,
			Partitions:        partitions,
			ReplicationFactor: replicationFactor,
		})
	}

	if err := kafka.CreateTopics(conf.Kafka, createTopicRequests...); err != nil {
		return errors.InternalServer.Wrap(err)
	}
	return nil
}
