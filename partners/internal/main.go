package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/sync/errgroup"

	"partners/internal/config"
	_ "partners/internal/docs"
	ebayEndpoint "partners/internal/services/ebay/endpoint"
	ebayNetwork "partners/internal/services/ebay/network"
	ebayService "partners/internal/services/ebay/service"

	"pkg/http/middleware"
	"pkg/http/router"
	"pkg/http/server"
	"pkg/log"
	"pkg/log/model"
	"pkg/panicRecover"
	"pkg/stackTrace"
)

// @title Partners Server Documentation
// @version @{version} (build @{build}) (commit @{commit})
// @description API Documentation for adaptor for partners
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
			Commit:   commit,
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

	// Регистрируем сетевые уровни
	ebayNetwork := ebayNetwork.NewEbayNetwork(cfg.Ebay)

	// Регистрируем сервисы
	ebayService := ebayService.NewEbayService(ebayNetwork)

	r := router.NewRouter()
	ebayEndpoint.MountEbayEndpoints(r, ebayService)
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
