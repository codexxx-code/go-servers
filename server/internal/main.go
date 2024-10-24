package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/sashabaranov/go-openai"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/sync/errgroup"

	"pkg/database/postgresql"
	"pkg/http/middleware"
	"pkg/http/router"
	"pkg/http/server"
	"pkg/log"
	"pkg/log/model"
	"pkg/migrator"
	"pkg/panicRecover"
	"pkg/stackTrace"
	"server/internal/config"
	_ "server/internal/docs"
	forecastEndpoint "server/internal/services/forecast/endpoint"
	forecastRepository "server/internal/services/forecast/repository"
	forecastService "server/internal/services/forecast/service"
	"server/internal/services/generator/ai/chatGPT"
	generatorEndpoint "server/internal/services/generator/endpoint"
	generatorRepository "server/internal/services/generator/repository"
	generatorService "server/internal/services/generator/service"
	"server/internal/services/scheduler"
	tgBotService "server/internal/services/tgBot/service"
	"server/migrations"
)

// @title Forecast Server Documentation
// @version @{version} (build @{build})
// @description API Documentation for Coin
// @contact.name Ilia Ivanov
// @contact.email bonavii@icloud.com
// @contact.url

// @securityDefinitions.apikey AuthJWT
// @in header
// @name Authorization
// @description JWT-токен авторизации

//go:generate go install github.com/swaggo/swag/cmd/swag@v1.8.2
//go:generate go mod download
//go:generate swag init -o docs --parseInternal --parseDependency

const version = "@{version}"
const build = "@{build}"

func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), err)
	}
}

func run() error {

	// Основной контекст приложения
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Перехватываем возможную панику
	defer panicRecover.PanicRecover(func(err error) {
		log.Fatal(ctx, err)
	})

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

	// Получаем имя хоста
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	// Инициализируем логгер
	log.Init(
		model.SystemInfo{
			Hostname: hostname,
			Version:  version,
			Build:    build,
			Env:      *envMode,
		},
		logHandlers...,
	)

	// Получаем конфиг
	log.Info(ctx, "Получаем конфиг")
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	// Инициализируем все синглтоны
	log.Info(ctx, "Инициализируем синглтоны")
	if err = initSingletons(cfg); err != nil {
		return err
	}

	// Подключаемся к базе данных
	log.Info(ctx, "Подключаемся к БД")
	postrgreSQL, err := postgresql.NewClientSQL(cfg.Repository, cfg.DBName)
	if err != nil {
		return err
	}
	defer postrgreSQL.Close()

	// Запускаем миграции в базе данных
	// TODO: Подумать, как откатывать миграции при ошибках
	log.Info(ctx, "Запускаем миграции")
	postgreSQLMigrator := migrator.NewMigrator(
		postrgreSQL,
		migrator.MigratorConfig{
			EmbedMigrations: migrations.EmbedMigrationsPostgreSQL,
			Dir:             "pgsql",
		},
	)
	if err = postgreSQLMigrator.Up(ctx); err != nil {
		return err
	}

	// Инициализируем телеграм бота
	log.Info(ctx, "Инициализируем Telegram-бота")
	tgBotService, err := tgBotService.NewTgBotService(cfg.Telegram.Token, cfg.Telegram.ChatID, cfg.Telegram.Enabled)
	if err != nil {
		return err
	}
	if cfg.Telegram.Enabled {
		defer tgBotService.Bot.Close()
	}

	// Инициализируем клиента ChatGPT
	log.Info(ctx, "Инициализируем OpenAI ChatGPT")
	openai := openai.NewClient(cfg.ChatGPTApiKey)
	chatGPTService := chatGPT.NewChatGPTService(openai)

	// Регистрируем репозитории
	generatorRepository := generatorRepository.NewGeneratorRepository(postrgreSQL)
	forecastRepository := forecastRepository.NewForecastRepository(postrgreSQL)

	// Регистрируем сервисы
	forecastService := forecastService.NewForecastService(forecastRepository)
	generatorService := generatorService.NewGeneratorService(chatGPTService, generatorRepository, forecastService, tgBotService)
	scheduler := scheduler.NewScheduler(generatorService)

	log.Info(ctx, "Запускаем планировщик")
	if err := scheduler.Start(); err != nil {
		return err
	}

	r := router.NewRouter()
	forecastEndpoint.MountForecastEndpoints(r, forecastService)
	generatorEndpoint.MountGeneratorEndpoints(r, generatorService)
	r.Mount("/swagger", httpSwagger.WrapHandler)

	server, err := server.GetDefaultServer(cfg.HTTP, r)
	if err != nil {
		return err
	}

	// Создаем wait группу
	eg, ctx := errgroup.WithContext(ctx)

	// Запускаем HTTP-сервер
	eg.Go(func() error { return server.Serve(ctx) })

	// Запускаем горутину, ожидающую завершение контекста
	eg.Go(func() error {

		// Если контекст завершился, значит процесс убили
		<-ctx.Done()

		// Плавно завершаем работу сервера
		server.Shutdown(ctx)

		return nil
	})

	// Ждем завершения контекста или ошибок в горутинах
	return eg.Wait()

}

func initSingletons(cfg config.Config) error {

	stackTrace.Init(cfg.ServiceName)

	if err := middleware.Init(cfg.ServiceName); err != nil {
		return err
	}

	return nil
}
